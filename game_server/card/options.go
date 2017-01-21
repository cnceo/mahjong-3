package card

type RuleOptions struct {
	//是否能吃牌
	WithChi 		bool    `json:"with_chi"`

	//是否能碰牌
	WithPeng		bool    `json:"with_peng"`

	//是否能杠牌
	WithGang		bool    `json:"with_gang"`

	//是否允许放炮
	WithFangPao		bool    `json:"with_fang_pao"`
}

