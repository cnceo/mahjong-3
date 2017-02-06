package room

import (
	"mahjong/game_server/card"
	"mahjong/game_server/log"
	"time"
)

const (
	RoomStatusWaitAllPlayerEnter	int = iota	// 等待玩家进入房间
	RoomStatusWaitStartPlayGame				// 等待游戏开始
	RoomStatusPlayingGame					// 正在进行游戏，结束后会进入RoomStatusEndPlayGame
	RoomStatusEndPlayGame					// 游戏结束后会回到等待游戏开始状态，或者进入结束房间状态
	RoomStatusEnd							// 房间结束状态，比如东南西北风都打完了
)

type RoomObserver interface {
	OnRoomClose(room *Room)
}

type Room struct {
	Id				int64
	config 			*RoomConfig
	players 		[]*Player
	cardPool		*card.Pool
	observers		[]RoomObserver
	roomStatus		int
}

func NewRoom(config *RoomConfig) *Room {
	room := &Room{
		config:		config,
		players:	make([]*Player, 0),
		observers:	make([]RoomObserver, 0),
		roomStatus:	RoomStatusWaitAllPlayerEnter,
	}
	return room
}

func (room *Room) AddObserver(observer RoomObserver) {
	room.observers = append(room.observers, observer)
}

func (room *Room) AddPlayer(player *Player) bool {
	if room.roomStatus != RoomStatusWaitAllPlayerEnter {
		return false
	}
	room.players = append(room.players, player)
	return true
}

func (room *Room) Start() {
	go func() {
		var oldStatus int
		for  {
			oldStatus = room.roomStatus
			room.checkStatus()
			if oldStatus == room.roomStatus {
				time.Sleep(time.Second)
			}
		}
	}()
}

func (room *Room) checkStatus() {
	switch room.roomStatus {
	case RoomStatusWaitAllPlayerEnter:
		room.checkAllPlayerEnter()
	case RoomStatusWaitStartPlayGame:
		room.startPlayGame()
	case RoomStatusPlayingGame:
		room.checkEndPlayGame()
	case RoomStatusEndPlayGame:
		room.checkRestartPlayGameOrEndRoom()
	case RoomStatusEnd:
		room.close()
	}
}

func (room *Room) isEnd() bool {
	//todo
	return true
}

func (room *Room) close() {
	for _, player := range room.players {
		player.LeaveRoom()
	}
	for _, observer := range room.observers {
		observer.OnRoomClose(room)
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

func (room *Room) startPlayGame() {
	room.switchStatus(RoomStatusPlayingGame)
	//todo
}

func (room *Room) checkEndPlayGame() {

}

func (room *Room) checkRestartPlayGameOrEndRoom() {
	if room.isEnd() {
		room.switchStatus(RoomStatusEnd)
	} else {
		room.switchStatus(RoomStatusWaitStartPlayGame)
	}
}