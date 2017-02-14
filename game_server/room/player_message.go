package room

import (
	"mahjong/game_server/card"
	"fmt"
)

type PlayerMsgType	int

const  (
	PlayerMsgGet			PlayerMsgType = iota + 1
	PlayerMsgDrop
	PlayerMsgChi
	PlayerMsgPeng
	PlayerMsgGang
	PlayerMsgZiMo
	PlayerMsgDianPao
	PlayerMsgEnterRoom
	PlayerMsgLeaveRoom
	PlayerMsgPlayingGameEnd
	PlayerMsgRoomClosed
)

type PlayerMessage struct {
	Type	PlayerMsgType
	Owner 	*Player
	Msg 	interface{}
}

func (msg *PlayerMessage) String() string {
	if msg == nil {
		return "{nil PlayerMessage}"
	}
	return fmt.Sprintf("{type=%v, owner=%v}", msg.Type, msg.Owner)
}

func newPlayerMsg(t PlayerMsgType, player *Player, msg interface{}) *PlayerMessage {
	return &PlayerMessage{
		Owner:	player,
		Type: t,
		Msg: msg,
	}
}

//玩家获得牌的消息
type GetCardMsg struct {
	Card *card.Card
	canZiMo bool
}
func NewPlayerGetCardMsg(msgOwner *Player, msg *GetCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgGet, msgOwner, msg)
}

//玩家出牌的消息
type DropCardMsg struct {
	WhatCard *card.Card
	canChiGroup []*card.Cards
	canPeng	bool
	canGang bool
	canDianPao bool
}
func NewPlayerDropCardMsg(msgOwner *Player, msg *DropCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgDrop, msgOwner, msg)
}

//玩家吃牌的消息
type ChiCardMsg struct {
	ChiPlayer		*Player
	FromPlayer		*Player
	WhatCard		*card.Card
	WhatGroup		*card.Cards
}
func NewPlayerChiCardMsg(msgOwner *Player, msg *ChiCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgChi, msgOwner, msg)
}

//玩家碰牌的消息
type PengCardMsg struct {
	PengPlayer		*Player
	FromPlayer		*Player
	WhatCard		*card.Card
}
func NewPlayerPengCardMsg(msgOwner *Player, msg *PengCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgPeng, msgOwner, msg)
}

//玩家杠牌的消息
type GangCardMsg struct {
	GangPlayer		*Player
	FromPlayer		*Player
	WhatCard		*card.Card
}
func NewPlayerGangCardMsg(msgOwner *Player, msg *GangCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgGang, msgOwner, msg)
}

type PlayerScore struct {
	P *Player
	Score int
}
//玩家自摸的消息
type ZiMoMsg struct {
	HuPlayer		*Player			// 胡牌的玩家
	WhatCard		*card.Card		// 胡的牌
	Desc			string			// 胡的描述
	PlayerScore 	[]*PlayerScore	// 玩家的分数
}
func NewPlayerZiMoMsg(msgOwner *Player, msg *ZiMoMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgZiMo, msgOwner, msg)
}


//玩家点炮的消息
type DianPaoMsg struct {
	HuPlayer		*Player			// 胡牌的玩家
	FromPlayer		*Player			//
	WhatCard		*card.Card		// 胡的牌
	Desc			string			// 胡的描述
	PlayerScore 	[]*PlayerScore	// 玩家的分数
}
func NewPlayerDianPaoMsg(msgOwner *Player, msg *DianPaoMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgDianPao, msgOwner, msg)
}

//玩家进入房间的消息
type EnterRoomMsg struct {
	EnterPlayer *Player
	AllPlaer 	[]*Player
}
func NewPlayerEnterRoomMsg(msgOwner *Player, msg *EnterRoomMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgEnterRoom, msgOwner, msg)
}

//玩家离开房间的消息
type LeaveRoomMsg struct {
	LeavePlayer *Player
	AllPlaer 	[]*Player
}
func NewPlayerLeaveRoomMsg(msgOwner *Player, msg *LeaveRoomMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgLeaveRoom, msgOwner, msg)
}

//一盘游戏结束的消息
type PlayingGameEndMsg struct {}
func NewPlayerPlayingGameEndMsg(msgOwner *Player, msg *PlayingGameEndMsg) *PlayerMessage{
	return newPlayerMsg(PlayerMsgPlayingGameEnd, msgOwner, msg)
}

//房间结束的消息
type RoomClosedMsg struct {}
func NewPlayerRoomClosedMsg(msgOwner *Player, msg *RoomClosedMsg) *PlayerMessage{
	return newPlayerMsg(PlayerMsgRoomClosed, msgOwner, msg)
}