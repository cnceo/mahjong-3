package room

import "mahjong/game_server/util"

type RoomMgr struct {
	config 	*RoomConfig
	rooms 	map[int64]*Room
}

func NewRoomMgr() *RoomMgr {
	return &RoomMgr{
		rooms:	make(map[int64]*Room, 0),
	}
}

func (mgr *RoomMgr) Init(conf string) error{
	return util.InitJsonConfig(conf, mgr.config)
}

func (mgr *RoomMgr) CreateRoom() *Room {
	room := NewRoom(mgr.config)
	mgr.rooms[room.Id] = room
	room.AddObserver(mgr)
	room.Start()
	return room
}

func (mgr *RoomMgr) OnRoomClose(room *Room) {
	delete(mgr.rooms, room.Id)
}