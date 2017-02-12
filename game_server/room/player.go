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
	playingCards 	*card.PlayingCards	//玩家手上的牌
	huChecker		[]hu_checker.Checker	//检查胡牌的checker

	gangFromPlayers []*Player			//杠了谁的牌

	observers	 []PlayerObserver
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
}

func (player *Player) AddObserver(ob PlayerObserver) {
	player.observers = append(player.observers, ob)
}

func (player *Player) AddMagicCard(card *card.Card) {
	log.Debug(player, "Player.AddMagicCard :", card.Name())
	player.playingCards.AddMagicCard(card)
}

func (player *Player) AddCard(card *card.Card) {
	log.Debug(player, "Player.AddCard :", card.Name())
	player.playingCards.AddCard(card)
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
	result := player.ZiMo()
	return result.IsHu
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
	result := player.IsHu()
	if !result.IsHu {
		player.Drop(card)
	}
	log.Debug(player, "DianPao card :", card.Name(), " result :", result)
	return result
}

//自摸
func (player *Player) ZiMo() *HuResult {
	result := player.IsHu()
	log.Debug(player, "ZiMo result :", result)
	return result
}

func (player *Player) IsHu() *HuResult {
	magicLen := player.playingCards.GetMagicCardsLen()
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
	player.notifyObserver(NewPlayingGameEndMsg(player, msg))
}

func (player *Player) OnRoomClosed() {
	log.Debug(player, "OnRoomClosed")
	player.room = nil
	player.Reset()

	msg := &RoomClosedMsg{}
	player.notifyObserver(NewRoomClosedMsg(player, msg))
}

func (player *Player) String() string{
	if player == nil {
		return "{player=nil}"
	}
	return fmt.Sprintf("{player=%v}", player.id)
}

func (player *Player) computeMagicCandidate() []*card.Cards {
	magicNum := player.playingCards.GetMagicCardsLen()
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
	default:
	//do nothing
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
		//ob.OnPlayerSuccessOperated(player, op)
		ob.OnMsg(player, msg)
	}
}

//begin player operate event handler
func (player *Player) onPlayerGet(op *PlayerOperate) {

}

func (player *Player) onPlayerDrop(op *PlayerOperate) {

}

func (player *Player) onPlayerChi(op *PlayerOperate) {

}

func (player *Player) onPlayerPeng(op *PlayerOperate) {

}

func (player *Player) onPlayerGang(op *PlayerOperate) {

}

func (player *Player) onPlayerZiMo(op *PlayerOperate) {

}

func (player *Player) onPlayerDianPao(op *PlayerOperate) {

}

func (player *Player) onPlayerEnterRoom(op *PlayerOperate) {

}

func (player *Player) onPlayerLeaveRoom(op *PlayerOperate) {

}
//end player operate event handler

