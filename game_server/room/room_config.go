package room

import "mahjong/game_server/util"

type RoomConfig struct {
	NeedPlayerNum		int        `json:"need_player_num"`
	WithoutFengCard		bool       `json:"without_feng_card"`
	WithoutJianCard		bool       `json:"without_jian_card"`
	WithoutHuaCard		bool       `json:"without_hua_card"`
	WithoutWanCard		bool       `json:"without_wan_card"`
	WithoutTiaoCard		bool       `json:"without_tiao_card"`
	WithoutTongCard 	bool       `json:"without_tong_card"`
	WithoutChi			bool       `json:"without_chi"`
	WithoutPeng			bool       `json:"without_peng"`
	WithoutGang			bool       `json:"without_gang"`
	MaxMagicCardNum		int        `json:"max_magic_card_num"`
}

func NewRoomConfig() *RoomConfig {
	return &RoomConfig{}
}

func (config *RoomConfig) Init(file string) error {
	return util.InitJsonConfig(file, config)
}