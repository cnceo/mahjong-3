package card


type Card struct {
	CardType int //牌类型
	CardNo   int //牌编号
}

//card是否在other的前一位牌
//比如1万在2万的前一位牌
func (card *Card) PrevAt(other *Card) bool{
	if !card.SameTypeAs(other) {
		return false
	}

	return card.CardNo + 1 == other.CardNo
}

func (card *Card) Less(other *Card) bool {
	if other == nil || card == nil{
		return false
	}

	if card.CardType < other.CardType {
		return true
	} else if card.CardType > other.CardType {
		return false
	}

	return card.CardNo < other.CardNo
}

func (card *Card) Swap(other *Card) {
	tmp := &Card{}
	tmp.CopyFrom(card)
	card.CopyFrom(other)
	other.CopyFrom(tmp)
}

//是否同一类型的牌
func (card *Card) SameTypeAs(other *Card) bool {
	if other == nil || card == nil {
		return false
	}
	return other.CardType == card.CardType
}

func (card *Card) SameAs(other *Card) bool {
	if other == nil || card == nil {
		return false
	}
	if other.CardType != card.CardType {
		return false
	}
	if other.CardNo != card.CardNo {
		return false
	}
	return true
}

func (card *Card) CopyFrom(other *Card) {
	if other == nil || card == nil {
		return
	}
	card.CardType = other.CardType
	card.CardNo = other.CardNo
}

//是否字牌：风、箭、花 牌
func (card *Card) IsZiCard() bool {
	if card == nil {
		return false
	}
	if card.CardType == CardType_Wan {
		return false
	}

	if card.CardType == CardType_Tiao {
		return false
	}

	if card.CardType == CardType_Tong  {
		return false
	}
	return true
}

func (card *Card) Name() string {
	if card == nil {
		return ""
	}
	cardNameMap := cardNameMap()
	noNameMap, ok1 := cardNameMap[card.CardType]
	if !ok1 {
		return "unknow card type"
	}

	name, ok2 := noNameMap[card.CardNo]
	if !ok2 {
		return "unknow card no"
	}
	return name
}

//检查三张牌是不是ABC顺子牌
func IsABC(card1, card2, card3 *Card) bool {
	if card1 == nil || card2 == nil || card3 == nil {
		return false
	}

	if card1.IsZiCard() || card2.IsZiCard() || card3.IsZiCard() {
		return false
	}

	return card1.PrevAt(card2) && card2.PrevAt(card3)
}

//检查四张牌是不是ABBC的牌型
func IsABBC(card1, card2, card3, card4 *Card) bool {
	if card1 == nil || card2 == nil || card3 == nil ||
		card4 == nil{
		return false
	}

	if card1.IsZiCard() || card2.IsZiCard() || card3.IsZiCard() ||
		card4.IsZiCard() {
		return false
	}

	return card1.PrevAt(card2) && card2.SameAs(card3) && card3.PrevAt(card4)
}

//检查五张牌是不是ABBBC的牌型
func IsABBBC(card1, card2, card3, card4, card5 *Card) bool {
	if card1 == nil || card2 == nil || card3 == nil ||
		card4 == nil || card5 == nil{
		return false
	}

	if card1.IsZiCard() || card2.IsZiCard() || card3.IsZiCard() ||
		card4.IsZiCard() || card5.IsZiCard() {
		return false
	}

	return card1.PrevAt(card2) && card2.SameAs(card3) && card3.SameAs(card4) &&
		card4.PrevAt(card5)
}

//检查六张牌是不是ABBBBC的牌型
func IsABBBBC(card1, card2, card3, card4, card5, card6 *Card) bool {
	if card1 == nil || card2 == nil || card3 == nil ||
		card4 == nil || card5 == nil || card6 == nil{
		return false
	}

	if card1.IsZiCard() || card2.IsZiCard() || card3.IsZiCard() ||
		card4.IsZiCard() || card5.IsZiCard() || card6.IsZiCard() {
		return false
	}

	return card1.PrevAt(card2) && card2.SameAs(card3) && card3.SameAs(card4) &&
		card4.SameAs(card5) && card5.PrevAt(card6)
}

func cardNameMap() map[int]map[int]string {
	return map[int]map[int]string{
		CardType_Feng: {
			Feng_CardNo_Dong: "东",
			Feng_CardNo_Nan:  "南",
			Feng_CardNo_Xi:   "西",
			Feng_CardNo_Bei:  "北",
		},

		CardType_Jian: {
			Jian_CardNo_Zhong: "中",
			Jian_CardNo_Fa:    "发",
			Jian_CardNo_Bai:   "白",
		},

		CardType_Hua: {
			Hua_CardNo_Chun: "春",
			Hua_CardNo_Xia:  "夏",
			Hua_CardNo_Qiu:  "秋",
			Hua_CardNo_Dong: "冬",
			Hua_CardNo_Mei:  "梅",
			Hua_CardNo_Lan:  "兰",
			Hua_CardNo_Zhu:  "竹",
			Hua_CardNo_Ju:   "菊",
		},

		CardType_Wan: {
			1: "一万",
			2: "二万",
			3: "三万",
			4: "四万",
			5: "五万",
			6: "六万",
			7: "七万",
			8: "八万",
			9: "九万",
		},

		CardType_Tiao: {
			1: "一条",
			2: "二条",
			3: "三条",
			4: "四条",
			5: "五条",
			6: "六条",
			7: "七条",
			8: "八条",
			9: "九条",
		},

		CardType_Tong: {
			1: "一筒",
			2: "二筒",
			3: "三筒",
			4: "四筒",
			5: "五筒",
			6: "六筒",
			7: "七筒",
			8: "八筒",
			9: "九筒",
		},
	}
}
