package room

import (
	"mahjong/game_server/card"
	"fmt"
)

type PlayerOperateType int

const (
	//被动操作，系统触发
	PlayerOperateGetInitCards	PlayerOperateType = iota + 1
	PlayerOperateGet

	//主动操作，用户触发
	PlayerOperateDrop
	PlayerOperateChi
	PlayerOperatePeng
	PlayerOperateGang
	PlayerOperateZiMo
	PlayerOperateDianPao
	PlayerOperateEnterRoom
	PlayerOperateLeaveRoom
)


type PlayerOperate struct {//玩家操作
	Op			PlayerOperateType
	Operator	*Player				//操作者
	Notify		chan bool			//	结果通知管道
	Data		interface{}
}

func (op *PlayerOperate) String() string {
	if op == nil {
		return "{operator=nil, op=nil}"
	}
	return fmt.Sprintf("{operator=%v, op=%v}", op.Operator, op.Op)
}

func newPlayerOperate(op PlayerOperateType, operator *Player, notify chan bool, data interface{}) *PlayerOperate{
	return &PlayerOperate{
		Op:	op,
		Data: data,
		Operator: operator,
		Notify:	notify,
	}
}

type PlayerOperateGetInitCardsData struct {
	CardsInHand 	[]*card.Cards
	MagicCards		*card.Cards
}
func NewPlayerOperateGetInitCards(operator *Player, notify chan bool, data *PlayerOperateGetInitCardsData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateGetInitCards, operator, notify, data)
}

type PlayerOperateChiData struct {
	Card 	*card.Card
	Group	*card.Cards
}
func NewPlayerOperateChi(operator *Player, notify chan bool, data *PlayerOperateChiData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateChi, operator, notify, data)
}

type PlayerOperateGetData struct{
	Card *card.Card
}
func NewPlayerOperateGet(operator *Player, data *PlayerOperateGetData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateGet, operator, nil, data)
}

type PlayerOperateDropData struct {
	Card *card.Card
}
func NewPlayerOperateDrop(operator *Player, notify chan bool, data *PlayerOperateDropData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateDrop, operator, notify, data)
}

type PlayerOperatePengData struct {
	Card *card.Card
}
func NewPlayerOperatePeng(operator *Player, notify chan bool, data *PlayerOperatePengData) *PlayerOperate {
	return newPlayerOperate(PlayerOperatePeng, operator, notify, data)
}

type PlayerOperateGangData struct {
	Card *card.Card
}
func NewPlayerOperateGang(operator *Player, notify chan bool, data *PlayerOperateGangData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateGang, operator, notify, data)
}

type PlayerOperateDianPaoData struct {
	Card *card.Card
}
func NewPlayerOperateDianPao(operator *Player, notify chan bool, data *PlayerOperateDianPaoData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateDianPao, operator, notify, data)
}

type PlayerOperateZiMoData struct {

}
func NewPlayerOperateZiMo(operator *Player, notify chan bool, data *PlayerOperateZiMoData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateZiMo, operator, notify, data)
}

type PlayerOperateEnterRoomData struct {

}
func NewPlayerOperateEnterRoom(operator *Player, notify chan bool, data *PlayerOperateEnterRoomData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateEnterRoom, operator, notify, data)
}

type PlayerOperateLeaveRoomData struct {

}
func NewPlayerOperateLeaveRoom(operator *Player, notify chan bool, data *PlayerOperateLeaveRoomData) *PlayerOperate {
	return newPlayerOperate(PlayerOperateLeaveRoom, operator, notify, data)
}