package room

import (
	"mahjong/game_server/card"
	"mahjong/game_server/hu_checker"
	"mahjong/game_server/log"
	"time"
	"mahjong/game_server/util"
	"fmt"
)

type HuResult struct {
	IsHu 	bool
	IsZiMo	bool
	Desc	string
	Score	int
}

func (result *HuResult) String() string {
	if result == nil {
		return "{ishu=nil, desc=nil, score=nil}"
	}
	return fmt.Sprintf("{ishu=%v, desc=%s, score=%d}", result.IsHu, result.Desc, result.Score)
}

func newHuResult(isHu bool, desc string, score int) *HuResult {
	return &HuResult{
		IsHu: isHu,
		Desc: desc,
		Score: score,
	}
}

type PlayerObserver interface {
	OnMsg(player *Player, msg *PlayerMessage)
}

type Player struct {
	id				uint64			// 玩家id
	room			*Room			// 玩家所在的房间
	lastGetCard		*card.Card	//	最后获得的牌
	playingCards 	*card.PlayingCards	//玩家手上的牌
	huChecker		[]hu_checker.Checker	//检查胡牌的checker

	gangFromPlayers []*Player			//杠了谁的牌
	gangByPlayers   []*Player			//被谁杠了

	result	*HuResult
	observers	 []PlayerObserver

	operateCh		chan *PlayerOperate
}

func NewPlayer(huChecker []hu_checker.Checker) *Player {
	player :=  &Player{
		id:		util.UniqueId(),
		playingCards:	card.NewPlayingCards(),
		huChecker:		huChecker,
		gangFromPlayers: make([]*Player, 0),
		observers:	make([]PlayerObserver, 0),
	}
	return player
}

func (player *Player) GetId() uint64 {
	return player.id
}

func (player *Player) Reset() {
	log.Debug(player,"Player.Reset")
	player.playingCards.Reset()
	player.gangFromPlayers = player.gangFromPlayers[0:0]
	player.lastGetCard = nil
	player.result = nil
}

func (player *Player) AddObserver(ob PlayerObserver) {
	player.observers = append(player.observers, ob)
}

func (player *Player) AddMagicCard(card *card.Card) {
	log.Debug(player, "Player.AddMagicCard :", card.Name())
	player.playingCards.AddMagicCard(card)
	player.lastGetCard = card
}

func (player *Player) AddCard(card *card.Card) {
	log.Debug(player, "Player.AddCard :", card.Name())
	player.playingCards.AddCard(card)
	player.lastGetCard = card
}

/*	计算指定的牌可以吃牌的组合
*/
func (player *Player) ComputeChiGroup(card *card.Card) []*card.Cards {
	return player.playingCards.ComputeChiGroup(card)
}

func (player *Player) CanPeng(card *card.Card) bool {
	return player.playingCards.CanPeng(card)
}

func (player *Player) CanGang(card *card.Card) bool {
	return player.playingCards.CanGang(card)
}

func (player *Player) CanDianPao(card *card.Card) bool {
	player.playingCards.AddCard(card)
	result := player.IsHu()
	player.playingCards.DropCard(card)
	return result.IsHu
}

func (player *Player) CanZiMo() bool{
	return player.IsHu().IsHu
}

func (player *Player) OperateChi(whatCard *card.Card, whatGroup *card.Cards) bool {
	log.Debug(player, "OperateChi, card :", whatCard.Name(), " group :", whatGroup.ToString())
	result := make(chan bool, 1)
	data := &PlayerOperateChiData{
		Card: whatCard,
		Group: whatGroup,
	}
	op := NewPlayerOperateChi(player, result, data)
	player.room.PlayerOperate(op)
	return player.waitResult(result)
}

func (player *Player) OperatePeng(card *card.Card) bool {
	log.Debug(player, "OperatePeng card :", card.Name())
	result := make(chan bool, 1)
	data := &PlayerOperatePengData{
		Card: card,
	}
	op := NewPlayerOperatePeng(player, result, data)
	player.room.PlayerOperate(op)
	return player.waitResult(result)
}

func (player *Player) OperateGang(card *card.Card) bool {
	log.Debug(player, "OperateGang card :", card.Name())
	result := make(chan bool, 1)
	data := &PlayerOperateGangData{
		Card: card,
	}
	op := NewPlayerOperateGang(player, result, data)
	player.room.PlayerOperate(op)
	return player.waitResult(result)
}

func (player *Player) OperateDrop(card *card.Card) bool {
	log.Debug(player, "OperateDrop card :", card.Name())
	result := make(chan bool, 1)
	data := &PlayerOperateDropData{
		Card: card,
	}
	op := NewPlayerOperateDrop(player, result, data)
	player.room.PlayerOperate(op)
	return player.waitResult(result)
}

func (player *Player) OperateZiMo() bool {
	log.Debug(player, "OperateZiMo")
	result := make(chan bool, 1)
	data := &PlayerOperateZiMoData{}
	op := NewPlayerOperateZiMo(player, result, data)
	player.room.PlayerOperate(op)
	return player.waitResult(result)
}

func (player *Player) OperateDianPao(card *card.Card) bool {
	log.Debug(player, "OperateDianPao card :", card.Name())
	result := make(chan bool, 1)
	data := &PlayerOperateDianPaoData{
		Card: card,
	}
	op := NewPlayerOperateDianPao(player, result, data)
	player.room.PlayerOperate(op)
	return player.waitResult(result)
}

func (player *Player) OperateEnterRoom(room *Room) bool{
	log.Debug(player, "OperateEnterRoom room :", room)
	result := make(chan bool, 1)
	data := &PlayerOperateEnterRoomData{}
	op := NewPlayerOperateEnterRoom(player, result, data)
	player.room.PlayerOperate(op)
	ok := player.waitResult(result)
	if ok {
		player.room = room
	}

	return ok
}

func (player *Player) OperateLeaveRoom() bool{
	log.Debug(player, "OperateLeaveRoom", player.room)
	if player.room == nil {
		return true
	}

	result := make(chan bool, 1)
	data := &PlayerOperateLeaveRoomData{}
	op := NewPlayerOperateLeaveRoom(player, result, data)
	player.room.PlayerOperate(op)
	ok := player.waitResult(result)
	if ok {
		player.room = nil
	}

	return ok
}

func (player *Player) BeGangBy(gangPlayer *Player) {
	player.gangByPlayers = append(player.gangByPlayers, gangPlayer)
}

func (player *Player) Chi(whatCard *card.Card, whatGroup *card.Cards) bool {
	log.Debug(player, "Chi card :", whatCard.Name(), ", group :", whatGroup.ToString())
	return player.playingCards.Chi(whatCard, whatGroup)
}

func (player *Player) Peng(card *card.Card) bool {
	log.Debug(player, "Peng card :", card.Name())
	return player.playingCards.Peng(card)
}

func (player *Player) Gang(card *card.Card, fromPlayer *Player) bool {
	log.Debug(player, "Gang card :", card.Name())
	ok := player.playingCards.Gang(card)
	if ok {
		player.gangFromPlayers = append(player.gangFromPlayers, fromPlayer)
	}
	return ok
}

func (player *Player) Drop(card *card.Card) bool {
	log.Debug(player, "Drop card :", card.Name())
	return player.playingCards.DropCard(card)
}

func (player *Player) AutoDrop() *card.Card {
	log.Debug(player, "AutoDrop")
	return player.playingCards.DropTail()
}

//胡别人的牌
func (player *Player) DianPao(card *card.Card) *HuResult {
	player.AddCard(card)
	player.result = player.IsHu()
	player.result.IsZiMo = false
	if !player.result.IsHu {
		player.Drop(card)
	}
	log.Debug(player, "DianPao card :", card.Name(), " result :", player.result)
	return player.result
}

//自摸
func (player *Player) ZiMo() *HuResult {
	player.result = player.IsHu()
	player.result.IsZiMo = true
	log.Debug(player, "ZiMo result :", player.result)
	return player.result
}

func (player *Player) IsHu() *HuResult {
	magicLen := player.playingCards.GetMagicCards().Len()
	log.Debug(player, "IsHu magicLen :", magicLen)
	for _, checker := range player.huChecker {
		if magicLen == 0 {
			isHu, conf := checker.IsHu(player.playingCards)
			if isHu {
				log.Debug(player, "checker :", checker.GetConfig().ToString(), "succ")
				return newHuResult(isHu, conf.Desc, conf.Score)
			}
		} else {
			//支持赖子牌的胡牌计算, 暴力穷举法，把赖子牌的所有候选集一个个试，胜在够简单
			candidate := player.computeMagicCandidate()
			tryCnt := 0
			for _, cards := range candidate {
				tryCnt++
				player.playingCards.AddCards(cards)
				isHu, conf := checker.IsHu(player.playingCards)
				if isHu {
					log.Debug(player, "checker :", checker.GetConfig().ToString(), "succ, tryMagicCnt :", tryCnt, ",cards:", cards.ToString())
					log.Debug(player, "tryCnt :", tryCnt, ", cards :", cards.ToString())
					return newHuResult(isHu, conf.Desc, conf.Score)
				} else {
					player.playingCards.DropCards(cards)
				}
			}
		}
		log.Debug(player, "checker :", checker.GetConfig().ToString(), "failed")
	}

	return newHuResult(false, "", 0)
}

func (player *Player) OnPlayingGameEnd() {
	log.Debug(player, "OnPlayingGameEnd")
	msg := &PlayingGameEndMsg{}
	player.notifyObserver(NewPlayerPlayingGameEndMsg(player, msg))
}

func (player *Player) OnRoomClosed() {
	log.Debug(player, "OnRoomClosed")
	player.room = nil
	player.Reset()

	msg := &RoomClosedMsg{}
	player.notifyObserver(NewPlayerRoomClosedMsg(player, msg))
}

func (player *Player) String() string{
	if player == nil {
		return "{player=nil}"
	}
	return fmt.Sprintf("{player=%v}", player.id)
}

func (player *Player) computeMagicCandidate() []*card.Cards {
	magicNum := player.playingCards.GetMagicCards().Len()
	switch magicNum {
	case 0:
		return nil
	case 1:
		return card.OneMagicCandidate
	case 2:
		return card.TwoMagicCandidate
	case 3:
		return card.ThreeMagicCandidate
	case 4:
		return card.FourMagicCandidate
	default:
		return nil
	}
	return nil
}

//玩家成功操作的通知
func (player *Player) OnPlayerSuccessOperated(op *PlayerOperate) {
	log.Debug(player, "OnPlayerSuccessOperated", op)
	switch op.Op {
	case PlayerOperateGetInitCards:

	case PlayerOperateGet:
		player.onPlayerGet(op)
	case PlayerOperateDrop:
		player.onPlayerDrop(op)
	case PlayerOperateChi:
		player.onPlayerChi(op)
	case PlayerOperatePeng:
		player.onPlayerPeng(op)
	case PlayerOperateGang :
		player.onPlayerGang(op)
	case PlayerOperateZiMo :
		player.onPlayerZiMo(op)
	case PlayerOperateDianPao :
		player.onPlayerDianPao(op)
	case PlayerOperateEnterRoom:
		player.onPlayerEnterRoom(op)
	case PlayerOperateLeaveRoom:
		player.onPlayerLeaveRoom(op)
	}
}

func (player *Player) DebugAllChecker() {
	for _, checker := range player.huChecker {
		log.Debug(player, "player's checker :", checker.GetConfig())
	}
}

func (player *Player) GetRoom() *Room {
	return player.room
}

func (player *Player) waitResult(resultCh chan bool) bool{
	log.Debug(player, "Player.waitResult")
	select {
	case <- time.After(time.Second * 10):
		log.Debug(player, "Player.waitResult timeout")
		return false
	case result := <- resultCh:
		log.Debug(player, "Player.waitResult result :", result)
		return result
	}
	return false
}

func (player *Player) notifyObserver(msg *PlayerMessage) {
	log.Debug(player, "notifyObserverMsg", msg)
	for _, ob := range player.observers {
		ob.OnMsg(player, msg)
	}
}

func (player *Player) calcScore(huPlayer *Player) int {
	if huPlayer.result == nil {//做个保护措施
		return 0
	}

	total := 0

	//计算一个基础得分
	if player == huPlayer {//胡牌者是正分
		total = huPlayer.result.Score
	} else {//其它玩家的分数取决于胡牌者是否自摸
		if huPlayer.result.IsZiMo {
			total = -huPlayer.result.Score
		} else {//不是自摸，只计算放炮的玩家的分数
			if player == player.room.getPrevOperator() {
				total = -huPlayer.result.Score
			}
		}
	}

	//计算杠别人的分数
	for _, gangFrom := range player.gangFromPlayers {
		if gangFrom == player {
			total += 2
		} else {
			total += 1
		}
	}

	total -= len(player.gangByPlayers) * 1	//计算放杠的分数
	return total
}

//begin player operate event handler
func (player *Player) onPlyaerGetInitCards(op *PlayerOperate) {
	if data, ok := op.Data.(*PlayerOperateGetInitCardsData); ok {
		msg := &GetInitCardsMsg{
			CardsInHand: data.CardsInHand,
			MagicCards: data.MagicCards,
		}
		player.notifyObserver(NewPlayerGetInitCardsMsg(player, msg))
	}
}

func (player *Player) onPlayerGet(op *PlayerOperate) {
	if data, ok := op.Data.(*PlayerOperateGetData); ok {
		msg := &GetCardMsg{
			canZiMo : player.CanZiMo(),
		}
		if op.Operator == player {//拿到牌只告诉自己是什么牌
			msg.Card = data.Card
		}
		player.notifyObserver(NewPlayerGetCardMsg(player, msg))
	}
}

func (player *Player) onPlayerDrop(op *PlayerOperate) {
	if data, ok := op.Data.(*PlayerOperateDropData); ok {
		if op.Operator == player {//自己出牌，不用通知自己
			return
		}
		msg := &DropCardMsg{
			WhatCard : data.Card,
			canChiGroup : player.ComputeChiGroup(data.Card),
			canPeng	: player.CanPeng(data.Card),
			canGang : player.CanGang(data.Card),
			canDianPao : player.CanDianPao(data.Card),
		}
		player.notifyObserver(NewPlayerDropCardMsg(player, msg))
	}
}

func (player *Player) onPlayerChi(op *PlayerOperate) {
	if data, ok := op.Data.(*PlayerOperateChiData); ok {
		if op.Operator == player {//自己吃牌，不用通知自己
			return
		}
		msg := &ChiCardMsg{
			ChiPlayer: op.Operator,
			FromPlayer: player.room.getPrevOperator(),
			WhatCard: data.Card,
			WhatGroup: data.Group,
		}
		player.notifyObserver(NewPlayerChiCardMsg(player, msg))
	}
}

func (player *Player) onPlayerPeng(op *PlayerOperate) {
	if data, ok := op.Data.(*PlayerOperatePengData); ok {
		if op.Operator == player {//自己碰牌，不用通知自己
			return
		}
		msg := &PengCardMsg{
			PengPlayer: op.Operator,
			FromPlayer: player.room.getPrevOperator(),
			WhatCard: data.Card,
		}
		player.notifyObserver(NewPlayerPengCardMsg(player, msg))
	}
}

func (player *Player) onPlayerGang(op *PlayerOperate) {
	if data, ok := op.Data.(*PlayerOperateGangData); ok {
		if op.Operator == player {//自己杠牌，不用通知自己
			return
		}
		msg := &GangCardMsg{
			GangPlayer: op.Operator,
			FromPlayer: player.room.getPrevOperator(),
			WhatCard: data.Card,
		}
		player.notifyObserver(NewPlayerGangCardMsg(player, msg))
	}
}

func (player *Player) onPlayerZiMo(op *PlayerOperate) {
	if _, ok := op.Data.(*PlayerOperateZiMoData); ok {
		msg := &ZiMoMsg{
			HuPlayer: op.Operator,
			WhatCard: op.Operator.lastGetCard,
			Desc: op.Operator.result.Desc,
			PlayerScore: make([]*PlayerScore, 0),
		}
		for _, tmpPlayer := range player.room.players {
			score := &PlayerScore{
				P : tmpPlayer,
				Score: tmpPlayer.calcScore(op.Operator),
			}
			msg.PlayerScore = append(msg.PlayerScore, score)
		}
		player.notifyObserver(NewPlayerZiMoMsg(player, msg))
	}
}

func (player *Player) onPlayerDianPao(op *PlayerOperate) {
	if data, ok := op.Data.(*PlayerOperateDianPaoData); ok {
		msg := &DianPaoMsg{
			HuPlayer: op.Operator,
			FromPlayer: player.room.getPrevOperator(),
			WhatCard: data.Card,
			Desc: op.Operator.result.Desc,
			PlayerScore: make([]*PlayerScore, 0),
		}
		for _, tmpPlayer := range player.room.players {
			score := &PlayerScore{
				P : tmpPlayer,
				Score: tmpPlayer.calcScore(op.Operator),
			}
			msg.PlayerScore = append(msg.PlayerScore, score)
		}
		player.notifyObserver(NewPlayerDianPaoMsg(player, msg))
	}
}

func (player *Player) onPlayerEnterRoom(op *PlayerOperate) {
	if _, ok := op.Data.(*PlayerOperateEnterRoomData); ok {
		msg := &EnterRoomMsg{
			EnterPlayer : op.Operator,
			AllPlaer: player.room.players,
		}
		player.notifyObserver(NewPlayerEnterRoomMsg(player, msg))
	}
}

func (player *Player) onPlayerLeaveRoom(op *PlayerOperate) {
	if _, ok := op.Data.(*PlayerOperateEnterRoomData); ok {
		msg := &LeaveRoomMsg{
			LeavePlayer : op.Operator,
			AllPlaer: player.room.players,
		}
		player.notifyObserver(NewPlayerLeaveRoomMsg(player, msg))
	}
}
//end player operate event handler
