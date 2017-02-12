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
type PlayerGetCardMsg struct {
	Card *card.Card
	canZiMo bool
}
func NewPlayerGetCardMsg(msgOwner *Player, msg *PlayerGetCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgGet, msgOwner, msg)
}

//玩家出牌的消息
type PlayerDropCardMsg struct {
	WhatCard *card.Card
	canChiGroup []*card.Cards
	canPeng	bool
	canGang bool
	canDianPao bool
}
func NewPlayerDropCardMsg(msgOwner *Player, msg *PlayerDropCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgDrop, msgOwner, msg)
}

//玩家吃牌的消息
type PlayerChiCardMsg struct {
	ChiPlayer		*Player
	FromPlayer		*Player
	WhatCard		*card.Card
	WhatGroup		*card.Cards
}
func NewPlayerChiCardMsg(msgOwner *Player, msg *PlayerChiCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgChi, msgOwner, msg)
}

//玩家碰牌的消息
type PlayerPengCardMsg struct {
	PengPlayer		*Player
	FromPlayer		*Player
	WhatCard		*card.Card
}
func NewPlayerPengCardMsg(msgOwner *Player, msg *PlayerPengCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgPeng, msgOwner, msg)
}

//玩家杠牌的消息
type PlayerGangCardMsg struct {
	GangPlayer		*Player
	FromPlayer		*Player
	WhatCard		*card.Card
}
func NewPlayerGangCardMsg(msgOwner *Player, msg *PlayerGangCardMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgGang, msgOwner, msg)
}

//玩家自摸的消息
type PlayerZiMoMsg struct {
	HuPlayer		*Player			// 胡牌的玩家
	WhatCard		*card.Card		// 点炮的牌
	Result			*HuResult		// 胡牌的结果
	MyScore			int				// 我的分数
}
func NewPlayerZiMoMsg(msgOwner *Player, msg *PlayerZiMoMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgZiMo, msgOwner, msg)
}


//玩家点炮的消息
type PlayerDianPaoMsg struct {
	DropCardPlayer	*Player			// 出牌的玩家
	HuPlayer		*Player			// 胡牌的玩家
	WhatCard		*card.Card		// 点炮的牌
	Result			*HuResult		// 胡牌的结果
	MyScore			int				// 我的分数
}
func NewPlayerDianPaoMsg(msgOwner *Player, msg *PlayerDianPaoMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgDianPao, msgOwner, msg)
}

//玩家进入房间的消息
type PlayerEnterRoomMsg struct {
	EnterPlayer *Player
}
func NewPlayerEnterRoomMsg(msgOwner *Player, msg *PlayerEnterRoomMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgEnterRoom, msgOwner, msg)
}

//玩家离开房间的消息
type PlayerLeaveRoomMsg struct {
	LeavePlayer *Player
}
func NewPlayerLeaveRoomMsg(msgOwner *Player, msg *PlayerLeaveRoomMsg) *PlayerMessage {
	return newPlayerMsg(PlayerMsgLeaveRoom, msgOwner, msg)
}

//一盘游戏结束的消息
type PlayingGameEndMsg struct {}
func NewPlayingGameEndMsg(msgOwner *Player, msg *PlayingGameEndMsg) *PlayerMessage{
	return newPlayerMsg(PlayerMsgPlayingGameEnd, msgOwner, msg)
}

//房间结束的消息
type RoomClosedMsg struct {}
func NewRoomClosedMsg(msgOwner *Player, msg *RoomClosedMsg) *PlayerMessage{
	return newPlayerMsg(PlayerMsgRoomClosed, msgOwner, msg)
}