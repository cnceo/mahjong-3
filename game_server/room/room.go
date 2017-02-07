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

type PlayerEnterRoom struct {
	player *Player
	notify chan bool
}

type Room struct {
	Id				int64
	config 			*RoomConfig
	players 		[]*Player
	cardPool		*card.Pool
	observers		[]RoomObserver
	roomStatus		int
	magicCard		*card.Card

	playerOpCh		chan *PlayerOperate
	playerEnterCh	chan *PlayerEnterRoom
}

func NewRoom(config *RoomConfig) *Room {
	room := &Room{
		Id:				0,
		config:			config,
		players:		make([]*Player, 0),
		cardPool:		card.NewPool(),
		observers:		make([]RoomObserver, 0),
		roomStatus:		RoomStatusWaitAllPlayerEnter,

		playerOpCh:		make(chan *PlayerOperate, 1024),
		playerEnterCh:	make(chan *PlayerEnterRoom, 4),
	}

	room.init()
	return room
}

func (room *Room) AddObserver(observer RoomObserver) {
	room.observers = append(room.observers, observer)
}

func (room *Room) Start() {
	go func() {
		for  {
			room.checkStatus()
		}
	}()
}

func (room *Room) Enter(player *Player, notify chan bool) {
	enter := &PlayerEnterRoom{
		player:	player,
		notify: notify,
	}
	room.playerEnterCh <- enter
}

func (room *Room) checkStatus() {
	switch room.roomStatus {
	case RoomStatusWaitAllPlayerEnter:
		room.waitPlayerEnter()
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
	//step.1 重置牌池
	room.cardPool.ReGenerate()

	//step.2 发牌
	room.switchStatus(RoomStatusPlayingGame)
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

//等待玩家进入
func (room *Room) waitPlayerEnter() {
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
}

func (room *Room) dispenseToAllPlayer() {
	for num := 0; num < 13; num++ {
		for _, player := range room.players {
			player.AddCard(room.cardPool.GetCard())
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

	if config.HasMagicCard {
		//todo
	}
}