package room

import "mahjong/game_server/card"

const (
	PlayerOperateChi 	int = iota + 1
	PlayerOperatePeng
	PlayerOperateGang
	PlayerOperateHu
)

type PlayerOperateChiData struct {
	Card 	*card.Card
	Group	*card.Cards
}

type PlayerOperate struct {
	Op			int
	Data		interface{}
	Operator	*Player				//操作者
	FromWho		*Player				//card来源于谁，比如杠，碰，吃，胡，来源于哪位对手或者自己
}

func NewPlayerOperateChi(card *card.Card, group *card.Cards, operator, from *Player) *PlayerOperate {
	return newPlayerOperate(
		PlayerOperateChi,
		&PlayerOperateChiData{
			Card:	card,
			Group:	group,
		},
		operator,
		from,
	)
}

func NewPlayerOperatePeng(card *card.Card, operator, from *Player) *PlayerOperate {
	return newPlayerOperate(PlayerOperatePeng, card, operator, from)
}

func NewPlayerOperateGang(card *card.Card, operator, from *Player) *PlayerOperate {
	return newPlayerOperate(PlayerOperateGang, card, operator, from)
}

func NewPlayerOperateHu(card *card.Card, operator, from *Player) *PlayerOperate {
	return newPlayerOperate(PlayerOperateHu, card, operator, from)
}

func newPlayerOperate(op int, data interface{}, operator, from *Player) *PlayerOperate{
	return &PlayerOperate{
		Operator: operator,
		FromWho: from,
		Op:	op,
		Data: data,
	}
}
