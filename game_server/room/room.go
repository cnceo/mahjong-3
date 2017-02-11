package room

import (
	"mahjong/game_server/card"
	"mahjong/game_server/log"
	"time"
	"mahjong/game_server/util"
)

const (
	RoomStatusWaitAllPlayerEnter	int = iota	// 等待玩家进入房间
	RoomStatusWaitStartPlayGame				// 等待游戏开始
	RoomStatusPlayingGame					// 正在进行游戏，结束后会进入RoomStatusEndPlayGame
	RoomStatusEndPlayGame					// 游戏结束后会回到等待游戏开始状态，或者进入结束房间状态
	RoomStatusEnd							// 房间结束状态，比如东南西北风都打完了
)

type RoomObserver interface {
	OnPlayingGameEnd(room *Room)
	OnRoomClosed(room *Room)
}

type PlayerRoomOperate struct {
	isLeaveRoom bool	//是否离开房间
	player *Player
	notify chan bool
}

type Room struct {
	id				int64					//房间id
	config 			*RoomConfig				//房间配置
	players 		[]*Player				//当前房间的玩家列表
	observers		[]RoomObserver			//房间观察者，需要实现OnRoomClose，房间close的时候会通知它

	roomStatus		int						//房间当前的状态

	firstMasterPlayer *Player				//第一个做东的玩家
	lastHuPlayer	*Player					//最后一次胡牌的玩家
	playedGameCnt	int						//已经玩了的游戏的次数

	//begin playingGameData, reset when start playing game
	cardPool		*card.Pool				//洗牌池
	magicCard		*card.Card				//当前的癞子牌
	masterPlayer	*Player					//做东的玩家，打筛子的玩家
	curOperatePlayer	*Player				//获得当前操作的玩家，可能是摸牌，碰牌，杠牌，吃牌，等他出牌
	quanFeng		*QuanFeng						//当前风圈
	huResult 		*HuResult
	otherPlayerOperate  []*PlayerOperate //当有玩家出牌，其它玩家的操作队列，依据优先级高低处理：胡 > 碰/杠 > 吃
	//end playingGameData, reset when start playing game

	playerOpCh		chan *PlayerOperate		//用户操作的channel
	playerRoomOperateCh	chan *PlayerRoomOperate	//用户房间操作的channel，进或者离开，一旦开始游戏，不允许离开房间
}

func NewRoom(config *RoomConfig) *Room {
	room := &Room{
		id:				0,
		config:			config,
		players:		make([]*Player, 0),
		cardPool:		card.NewPool(),
		observers:		make([]RoomObserver, 0),
		roomStatus:		RoomStatusWaitAllPlayerEnter,
		quanFeng:		newQuanFeng(card.Feng_CardNo_Dong),
		playedGameCnt:	0,

		otherPlayerOperate:	make([]*PlayerOperate, 0),

		playerOpCh:		make(chan *PlayerOperate, 1024),
		playerRoomOperateCh:	make(chan *PlayerRoomOperate, 4),
	}

	room.firstMasterPlayer = room.selectMasterPlayer()
	room.init()
	return room
}

func (room *Room) addObserver(observer RoomObserver) {
	room.observers = append(room.observers, observer)
}

func (room *Room) delObserver(observer RoomObserver) {
	for idx, ob := range room.observers {
		if ob == observer {
			room.observers = append(room.observers[0:idx], room.observers[idx+1:]...)
			return
		}
	}
}

func (room *Room) Start() {
	go func() {
		for  {
			room.checkStatus()
		}
	}()
}

func (room *Room) Enter(player *Player, notify chan bool) {
	enter := &PlayerRoomOperate{
		isLeaveRoom: false,
		player:	player,
		notify: notify,
	}
	room.playerRoomOperateCh <- enter
}

func (room *Room) Leave(player *Player, notify chan bool) {
	leave := &PlayerRoomOperate{
		isLeaveRoom: true,
		player:	player,
		notify: notify,
	}
	room.playerRoomOperateCh <- leave
}

func (room *Room) checkStatus() {
	switch room.roomStatus {
	case RoomStatusWaitAllPlayerEnter:
		room.waitAllPlayerEnter()
	case RoomStatusWaitStartPlayGame:
		room.startPlayGame()
	case RoomStatusPlayingGame:
		room.playingGame()
	case RoomStatusEndPlayGame:
		room.endPlayGame()
	case RoomStatusEnd:
		room.close()
	}
}

func (room *Room) isRoomEnd() bool {
	if room.config.WithQuanFeng {
		if !room.quanFeng.isLastQuanFeng() {//不是最后一圈，肯定没结束
			return false
		}
		return room.computeQuanFeng().isFirstQuanFeng()
	}

	return room.playedGameCnt >= room.config.MaxPlayGameCnt
}

func (room *Room) close() {
	for _, observer := range room.observers {
		observer.OnRoomClosed(room)
	}
}

func (room *Room) checkAllPlayerEnter() {
	if len(room.players) >= room.config.NeedPlayerNum {
		room.switchStatus(RoomStatusWaitStartPlayGame)
	}
}

func (room *Room) switchStatus(status int) {
	log.Debug("room status switch,", room.roomStatus, " =>", status)
	room.roomStatus = status
}

func (room *Room) startPlayGame()  {
	room.resetPlayingGameData()
	room.switchStatus(RoomStatusPlayingGame)
}

func (room *Room) resetPlayingGameData() {
	// 重置牌池, 洗牌
	room.cardPool.ReGenerate()

	// 计算癞子牌，如果有的话
	room.computeMagicCard()

	// 初始化素有玩家
	room.initAllPlayer()

	// 选择东家
	room.masterPlayer = room.selectMasterPlayer()

	//选完东家后，计算圈风
	if room.config.WithQuanFeng {
		room.quanFeng = room.computeQuanFeng()
	}

	// 设定获得牌的玩家的索引为东家
	room.curOperatePlayer = room.masterPlayer

	//重置上一盘的结果
	room.huResult = nil

	room.otherPlayerOperate = room.otherPlayerOperate[0:0]

}

func (room *Room) playingGame() {
	// 发牌给玩家
	if !room.putFrontCardToPlayer(room.curOperatePlayer) {
		room.switchStatus(RoomStatusEndPlayGame)
	} else {
		room.waitCurPlayerOperate()
	}
}

func (room *Room) endPlayGame() {
	room.playedGameCnt++

	for _, ob := range room.observers {
		ob.OnPlayingGameEnd(room)
	}

	if room.isRoomEnd() {
		room.switchStatus(RoomStatusEnd)
	} else {
		room.switchStatus(RoomStatusWaitStartPlayGame)
	}
}

//等待玩家进入
func (room *Room) waitAllPlayerEnter() {
	/*
	waitTime := time.Duration(room.config.WaitPlayerEnterRoomTimeout)
	select {
	case <-time.After(time.Second * waitTime):
		room.switchStatus(RoomStatusEnd)		//超时没有全部玩家进入，则结束
	case enter := <-room.playerEnterCh:
		if room.addPlayer(enter.player) {	//	玩家进入成功
			enter.notify <- true
			room.checkAllPlayerEnter()
		} else{
			enter.notify <- false			//玩家进入失败
		}
	}
	*/

	breakTimerTime := time.Duration(0)
	timeout := time.Duration(room.config.WaitPlayerEnterRoomTimeout)
	for {
		timer := timeout - breakTimerTime
		select {
		case <-time.After(timer):
			room.switchStatus(RoomStatusEnd) //超时发现没有全部玩家都进入房间了，则结束
			return
		case op := <-room.playerRoomOperateCh:
			if op.isLeaveRoom { //离开房间
				room.delPlayer(op.player)
				room.delObserver(op.player)
				op.notify <- true
			} else {//进入房间
				if room.addPlayer(op.player) { //	玩家进入成功
					room.addObserver(op.player)
					op.notify <- true
					room.checkAllPlayerEnter()
				} else {
					op.notify <- false //玩家进入失败
				}
			}
		}
	}
}

//给所有玩家发初始化的13张牌
func (room *Room) initAllPlayer() {
	for _, player := range room.players {
		player.Reset()
		for num := 0; num < 13; num++ {
			room.putFrontCardToPlayer(player)
		}
	}
}

//添加玩家
func (room *Room) addPlayer(player *Player) bool {
	if room.roomStatus != RoomStatusWaitAllPlayerEnter {
		return false
	}
	room.players = append(room.players, player)
	return true
}

func (room *Room) delPlayer(player *Player)  {
	for idx, p := range room.players {
		if p == player {
			room.players = append(room.players[0:idx], room.players[idx+1:]...)
			return
		}
	}
}

//初始化cardool
func (room *Room) init() {
	config := room.config
	if config.WithFengCard {
		room.cardPool.AddFengGenerater()
	}
	if config.WithJianCard {
		room.cardPool.AddJianGenerater()
	}
	if config.WithHuaCard {
		room.cardPool.AddHuaGenerater()
	}
	if config.WithWanCard {
		room.cardPool.AddWanGenerater()
	}
	if config.WithTiaoCard {
		room.cardPool.AddTiaoGenerater()
	}
	if config.WithTongCard {
		room.cardPool.AddTongGenerater()
	}
}

//计算癞子牌
func (room *Room) computeMagicCard() {
	if !room.config.HasMagicCard {
		return
	}
	cardIdx := room.config.NeedPlayerNum * 13
	card := room.cardPool.At(cardIdx)
	room.magicCard = card.Next()
}

//是否癞子牌
func (room *Room) isMagicCard(card *card.Card) bool {
	if !room.config.HasMagicCard {
		return false
	}
	return card.SameAs(room.magicCard)
}

//发牌给指定玩家
func (room *Room) putFrontCardToPlayer(player *Player) bool {
	card := room.cardPool.PopFront()
	if card == nil {
		return false
	}
	if room.isMagicCard(card) {
		player.AddMagicCard(card)
	} else {
		player.AddCard(card)
	}
	return true
}

func (room *Room) putTailCardToPlayer(player *Player) bool {
	card := room.cardPool.PopTail()
	if card == nil {
		return false
	}
	if room.isMagicCard(card) {
		player.AddMagicCard(card)
	} else {
		player.AddCard(card)
	}
	return true
}

//选择东家
func (room *Room) selectMasterPlayer() *Player {
	if room.playedGameCnt == 0 { //第一盘，随机一个做东
		idx := util.RandomN(len(room.players))
		return room.players[idx]
	}

	if room.lastHuPlayer == nil {//流局，上一盘没有人胡牌
		return room.masterPlayer
	}

	if !room.config.WithQuanFeng { //不支持圈风，那就谁胡谁做东
		return room.lastHuPlayer
	}

	//支持圈风，那就如果他胡了就继续做东，否则他的下一个玩家做东
	if room.masterPlayer == room.lastHuPlayer { //上一次做东的人最后一次胡牌了，继续他做东
		return room.masterPlayer
	}
	return room.nextPlayer(room.masterPlayer)
}

//等待获得牌的玩家操作, 胡？杠？出牌？如果没有任何操作，超时的话，自动帮他出一张牌
func (room *Room) waitCurPlayerOperate() {
	breakTimerTime := time.Duration(0)
	timeout := time.Duration(room.config.WaitPlayerOperateTimeout)
	for  {
		timer := timeout - breakTimerTime
		select {
		case <-time.After(timer):
			//超时没有操作，自动帮他出一张牌
			room.curOperatePlayer.DropTail()
		case op := <-room.playerOpCh:
			if op.Operator != room.curOperatePlayer {
				//不是当前玩家的操作，直接无视
				return
			}
			if op.Op == PlayerOperateChi || op.Op == PlayerOperatePeng || op.Op == PlayerOperateDianPao {
				//当前玩家不可能吃牌、碰牌、点炮胡别人
				return
			}
			room.dealPlayerOperate(op)
		}
	}
}

//当玩家出牌后，等待其他玩家操作
func (room *Room) waitOtherPlayerOperateAfterDrop() {
	breakTimerTime := time.Duration(0)
	timeout := time.Duration(room.config.WaitPlayerOperateTimeout)
	for  {
		timer := timeout - breakTimerTime
		select {
		case <- time.After(timer):
			//超时没有其它玩家有任何操作, 设置下一个操作者，继续
			room.curOperatePlayer = room.nextPlayer(room.curOperatePlayer)
		case op := <-room.playerOpCh :
			if op.Operator == room.curOperatePlayer {//操作者不可能是出牌者，直接无视
				return
			}
			room.dealPlayerOperate(op)
		}
	}
}

//等待碰、吃牌的玩家出牌，超时的话，自动帮他出一张牌
func (room *Room) waitPlayerDrop() {
	breakTimerTime := time.Duration(0)
	timeout := time.Duration(room.config.WaitPlayerOperateTimeout)
	for  {
		timer := timeout - breakTimerTime
		select {
		case <- time.After(timer):
			room.curOperatePlayer.DropTail()
			return
		case op := <-room.playerOpCh :
			if room.curOperatePlayer != op.Operator {
				return
			}
			if op.Op != PlayerOperateDrop {
				return
			}
			room.dealPlayerOperate(op)
		}
	}
}

//取指定玩家的下一个玩家
func (room *Room) nextPlayer(player *Player) *Player {
	idx := 0
	for i, p := range room.players {
		if p == player {
			idx = i
			break
		}
	}
	if idx == len(room.players) - 1 {
		return room.players[0]
	}
	return room.players[idx+1]
}

//处理玩家操作
func (room *Room) dealPlayerOperate(op *PlayerOperate) {
	switch op.Op {
	case PlayerOperateDianPao :
		if room.config.OnlyZiMo {
			//只能自摸，则不允许非摸牌者胡
			return
		}
		if card, ok := op.Data.(*card.Card); ok {
			room.huResult = op.Operator.IsDianPao(card)
			if room.huResult.IsHu {
				room.lastHuPlayer = op.Operator
				room.switchStatus(RoomStatusEndPlayGame)
			}
		}

	case PlayerOperateZiMo :
		if op.Operator != room.curOperatePlayer {
			return
		}
		room.huResult = op.Operator.IsZiMo()
		if room.huResult.IsHu {
			room.lastHuPlayer = op.Operator
			room.switchStatus(RoomStatusEndPlayGame)
		}

	case PlayerOperateDrop:
		if card, ok := op.Data.(*card.Card); ok {
			if op.Operator != room.curOperatePlayer {
				return
			}
			if op.Operator.Drop(card) {//出牌
				//通知所有其他玩家
				for _, other := range room.players {
					if other == op.Operator {
						continue
					}
					other.OnOtherDropCard(op.Operator, card)
				}
				//出牌后等待其他玩家操作
				room.waitOtherPlayerOperateAfterDrop()
			}
		}

	case PlayerOperateChi:
		if !room.config.WithChi {
			return
		}
		nextPlayer := room.nextPlayer(room.curOperatePlayer)
		if nextPlayer != op.Operator {
			return
		}
		//先加入等待队列，然后选一个优先级最高的操作来处理， todo
		//room.otherPlayerOperate = append(room.otherPlayerOperate, op)
		if chiData, ok := op.Data.(*PlayerOperateChiData); ok {
			if op.Operator.Chi(chiData.Card, chiData.Group) {
				//吃成功了，设定当前玩家为杠牌者，并等待他出牌
				room.curOperatePlayer = op.Operator
				room.waitPlayerDrop()
			}
		}

	case PlayerOperatePeng:
		if !room.config.WithPeng {
			return
		}
		if card, ok := op.Data.(*card.Card); ok {
			if op.Operator.Peng(card) {
				//碰成功了，设定当前玩家为杠牌者，并等待他出牌
				room.curOperatePlayer = op.Operator
				room.waitPlayerDrop()
			}
		}

	case PlayerOperateGang :
		if !room.config.WithGang {
			return
		}
		if room.cardPool.GetCardNum() <= len(room.players) {//牌数少于人数，不允许杠牌了
			return
		}
		if card, ok := op.Data.(*card.Card); ok {
			if op.Operator.Gang(card, room.curOperatePlayer) {
				//杠成功了，设定当前玩家为杠牌者
				room.curOperatePlayer = op.Operator
			}
		}
	default:
		//do nothing
	}
}

//计算圈风
func (room *Room) computeQuanFeng() *QuanFeng{
	if !room.config.WithQuanFeng {
		return room.quanFeng
	}

	if room.lastHuPlayer == nil {//上一盘没有人胡, 圈风不变
		return room.quanFeng
	}

	if room.lastHuPlayer == room.masterPlayer {//上一次胡牌的玩家是东家，圈风不变
		return room.quanFeng
	}

	if room.masterPlayer != room.firstMasterPlayer {//没有回到起点，还是同一个圈风
		return room.quanFeng
	}
	return room.quanFeng.next()
}