package card

import (
	"mahjong/game_server/util"
	"sort"
)

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

func (cards *Cards) Clear() {
	cards.data = cards.data[0:0]
}

func (cards *Cards)Sort() {
	sort.Sort(cards)
}

func (cards *Cards) IsAA() bool {
	if cards.Len() != 2 {
		return false
	}
	return cards.data[0].SameAs(cards.data[1])
}

func (cards *Cards) IsAAA() bool  {
	if cards.Len() != 3 {
		return false
	}
	return cards.data[0].SameAs(cards.data[1]) && cards.data[0].SameAs(cards.data[2])
}

func (cards *Cards) IsABC() bool {
	if cards.Len() != 3 {
		return false
	}
	if cards.data[0].IsZiCard() || cards.data[1].IsZiCard() || cards.data[2].IsZiCard() {
		return false
	}

	if cards.data[0].CardType != cards.data[1].CardType {
		return false
	}
	if cards.data[0].CardType != cards.data[2].CardType {
		return false
	}
	if cards.data[1].CardType != cards.data[2].CardType {
		return false
	}

	return cards.data[0].CardNo + cards.data[2].CardNo == 2 * cards.data[1].CardNo
}

func (cards *Cards) Is5Card() bool {
	if cards.Len() != 5  {
		return false
	}
	var left, three, right *Cards
	left, three = cards.splitLeftOtherAndThree()
	if three.IsAAA() && left.IsAA() {
		return true
	}

	if three.IsABC() && left.IsAA() {
		return true
	}

	three, right = cards.splitThreeAndRightOther()
	if three.IsAAA() && right.IsAA() {
		return true
	}
	if three.IsABC() && right.IsAA() {
		return true
	}

	return false
}

//把牌分成左右2份：左边其它，3个右边
func (cards *Cards) splitLeftOtherAndThree() (other *Cards, three *Cards){
	length := cards.Len()
	if length <= 3 {
		return nil, nil
	}
	other = &Cards{
		data:	cards.data[0 : length-3],
	}
	three = &Cards{
		data:	cards.data[length-3:],
	}
	return
}

//把牌分成左右2份：左边3个，其它的在右边
func (cards *Cards) splitThreeAndRightOther() (three *Cards, right *Cards){
	length := cards.Len()
	if length <= 3 {
		return nil, nil
	}

	three = &Cards{
		data:	cards.data[0:3],
	}
	right = &Cards{
		data:	cards.data[3:],
	}
	return
}

func (cards *Cards) ToString() string {
	str := ""
	for _, card := range cards.data{
		str += card.Name() + ","
	}
	return str
}
