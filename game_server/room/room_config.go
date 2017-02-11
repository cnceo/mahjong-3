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
	WithQuanFeng			bool        `json:"with_quan_feng"`	//是否支持圈风?
	MaxPlayGameCnt			int            `json:"max_play_game_cnt"`	//不支持圈风的时候，最大的游戏局数
	OnlyZiMo				bool        `json:"only_zi_mo"`	//是否只能自摸
}

func NewRoomConfig() *RoomConfig {
	return &RoomConfig{}
}

func (config *RoomConfig) Init(file string) error {
	return util.InitJsonConfig(file, config)
}