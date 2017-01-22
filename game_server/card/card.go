package card

import (
	"sort"
	"mahjong/game_server/util"
)

type Card struct {
	CardType int //牌类型
	CardNo   int //牌编号
}

func (card *Card) SameAs(other *Card) bool {
	if other == nil {
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
	card.CardType = other.CardType
	card.CardNo = other.CardNo
}

func (card *Card) Name() string {
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

type Cards struct {
	data 	[]*Card
}

func NewCards() *Cards{
	return &Cards{
		data :	make([]*Card, 0),
	}
}

func (cards *Cards) Data() []*Card {
	return cards.data
}

func (cards *Cards) Len() int {
	return len(cards.data)
}

func (cards *Cards) Less(i, j int) bool {
	if cards.data[i].CardType < cards.data[j].CardType {
		return true
	} else if cards.data[i].CardType > cards.data[j].CardType {
		return false
	}

	if cards.data[i].CardNo < cards.data[j].CardNo {
		return true
	}
	return false
}

func (cards *Cards) Swap(i, j int) {
	swap := cards.data[i]
	cards.data[i] = cards.data[j]
	cards.data[j] = swap
}

func (cards *Cards) AppendCard(card *Card) {
	cards.data = append(cards.data, card)
}

func (cards *Cards) AppendCards(other *Cards) {
	cards.data = append(cards.data, other.data...)
}

func (cards *Cards) TakeWay(drop *Card) bool {
	for idx, card := range cards.data {
		if card.SameAs(drop) {
			cards.data = append(cards.data[0:idx], cards.data[idx+1:]...)
			return true
		}
	}
	return false
}

func (cards *Cards) RandomTakeWayOne() *Card {
	length := cards.Len()
	if length == 0 {
		return nil
	}
	idx := util.RandomN(length)
	card := cards.data[idx]
	cards.data = append(cards.data[0:idx], cards.data[idx+1:]...)
	return card
}

func (cards *Cards)Sort() {
	sort.Sort(cards)
}

func (cards *Cards) ToString() string {
	str := ""
	for _, card := range cards.data{
		str += card.Name() + ","
	}
	return str
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
