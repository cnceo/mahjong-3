package room

import "mahjong/game_server/card"

const (
	PlayerOperateDrop		int = iota + 1
	PlayerOperateChi
	PlayerOperatePeng
	PlayerOperateGang
	PlayerOperateZiMo
	PlayerOperateDianPao
)

type PlayerOperateChiData struct {
	Card 	*card.Card
	Group	*card.Cards
}

type PlayerOperate struct {//玩家操作
	Op			int
	Data		interface{}
	Operator	*Player				//操作者
}

func NewPlayerOperateDrop(card *card.Card, operator *Player) *PlayerOperate {
	return newPlayerOperate(
		PlayerOperateChi,
		card,
		operator,
	)
}

func NewPlayerOperateChi(card *card.Card, group *card.Cards, operator *Player) *PlayerOperate {
	return newPlayerOperate(
		PlayerOperateChi,
		&PlayerOperateChiData{
			Card:	card,
			Group:	group,
		},
		operator,
	)
}

func NewPlayerOperatePeng(card *card.Card, operator *Player) *PlayerOperate {
	return newPlayerOperate(PlayerOperatePeng, card, operator)
}

func NewPlayerOperateGang(card *card.Card, operator *Player) *PlayerOperate {
	return newPlayerOperate(PlayerOperateGang, card, operator)
}

func NewPlayerOperateDianPao(card *card.Card, operator *Player) *PlayerOperate {
	return newPlayerOperate(PlayerOperateDianPao, card, operator)
}

func NewPlayerOperateZiMo(operator *Player) *PlayerOperate {
	return newPlayerOperate(PlayerOperateZiMo, nil, operator)
}

func newPlayerOperate(op int, data interface{}, operator *Player) *PlayerOperate{
	return &PlayerOperate{
		Op:	op,
		Data: data,
		Operator: operator,
	}
}