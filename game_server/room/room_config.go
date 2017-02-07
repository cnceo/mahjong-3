package room

import "mahjong/game_server/util"

type RoomConfig struct {
	NeedPlayerNum				int        `json:"need_player_num"`
	WithFengCard				bool       `json:"With_feng_card"`
	WithJianCard				bool       `json:"With_jian_card"`
	WithHuaCard				bool       `json:"With_hua_card"`
	WithWanCard				bool       `json:"With_wan_card"`
	WithTiaoCard				bool       `json:"With_tiao_card"`
	WithTongCard 			bool       `json:"With_tong_card"`
	WithChi					bool       `json:"With_chi"`
	WithPeng					bool       `json:"With_peng"`
	WithGang					bool       `json:"With_gang"`
	HasMagicCard				bool        `json:"has_magic_card"`
	WaitPlayerEnterRoomTimeout	int        `json:"wait_player_enter_room_timeout"`
	WaitPlayerOperateTimeout	int        `json:"wait_player_operate_timeout"`
}

func NewRoomConfig() *RoomConfig {
	return &RoomConfig{}
}

func (config *RoomConfig) Init(file string) error {
	return util.InitJsonConfig(file, config)
}