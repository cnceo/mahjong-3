package card

import (
	"mahjong/game_server/util"
	"sort"
)

type Cards struct {
	data 	[]*Card
}

//创建一个Cards对象
func NewCards() *Cards{
	return &Cards{
		data :	make([]*Card, 0),
	}
}

//从指定的cardSlice创建一个Cards对象
func NewCardsFrom(cardSlice []*Card) *Cards{
	return &Cards{
		data: cardSlice,
	}
}

//获取cards的数据
func (cards *Cards) Data() []*Card {
	return cards.data
}

//获取第idx个牌
func (cards *Cards) At(idx int) *Card {
	if idx >= cards.Len() {
		return nil
	}
	return cards.data[idx]
}

//cards的长度，牌数
func (cards *Cards) Len() int {
	return len(cards.data)
}

//比较指定索引对应的两个牌的大小
func (cards *Cards) Less(i, j int) bool {
	return cards.At(i).Less(cards.At(j))
}

//交换索引为，j的两个数据
func (cards *Cards) Swap(i, j int) {
	length := cards.Len()
	if i >= length || j >= length {
		return
	}
	swap := cards.At(i)
	cards.data[i] = cards.At(j)
	cards.data[j] = swap
}

//追加一张牌
func (cards *Cards) AppendCard(card *Card) {
	if card == nil {
		return
	}
	cards.data = append(cards.data, card)
}

//增加一张牌并排序
func (cards *Cards) AddAndSort(card *Card){
	if card == nil {
		return
	}
	cards.AppendCard(card)
	cards.Sort()
}

//追加一个cards对象
func (cards *Cards) AppendCards(other *Cards) {
	cards.data = append(cards.data, other.data...)
}

//取走一张指定的牌，并返回成功或者失败
func (cards *Cards) TakeWay(drop *Card) bool {
	if drop == nil {
		return true
	}
	for idx, card := range cards.data {
		if card.SameAs(drop) {
			cards.data = append(cards.data[0:idx], cards.data[idx+1:]...)
			return true
		}
	}
	return false
}

//随机取走一张牌
func (cards *Cards) RandomTakeWayOne() *Card {
	length := cards.Len()
	if length == 0 {
		return nil
	}
	idx := util.RandomN(length)
	card := cards.At(idx)
	cards.data = append(cards.data[0:idx], cards.data[idx+1:]...)
	return card
}

//清空牌
func (cards *Cards) Clear() {
	cards.data = cards.data[0:0]
}

//排序
func (cards *Cards)Sort() {
	sort.Sort(cards)
}

//是否一对牌
func (cards *Cards) IsAA() bool {
	if cards.Len() != 2 {
		return false
	}
	return cards.At(0).SameAs(cards.At(1))
}

//是否3张相同的牌
func (cards *Cards) IsAAA() bool  {
	if cards.Len() != 3 {
		return false
	}
	return cards.At(0).SameAs(cards.At(1)) && cards.At(0).SameAs(cards.At(2))
}

//是否顺子牌
func (cards *Cards) IsABC() bool {
	if cards.Len() != 3 {
		return false
	}
	if cards.At(0).IsZiCard() || cards.At(1).IsZiCard() || cards.At(2).IsZiCard() {
		return false
	}


	if !cards.isAllCardSameType() {
		return false
	}

	return cards.At(0).PrevAt(cards.At(1)) && cards.At(1).PrevAt(cards.At(2))
}

//是否胡牌, 胡牌公式：m*AAA + n*ABC + AA
//递归检查：最左边或者最后边3张是否AAA/ABC、剩下的牌是否构成胡牌，直到最后剩下两张是否AA
func (cards *Cards) IsHu() bool {
	length := cards.Len()
	if (length-2)%3 != 0 {//not 2,5,8,11,14
		return false
	}

	if length == 2 {
		if cards.IsAA() {
			return true
		}
		return false
	}

	var left, three, right *Cards
	left, three = cards.splitLeftOtherAndThree()
	if three.IsAAA() && left.IsHu() {
		return true
	}

	if three.IsABC() && left.IsHu() {
		return true
	}

	three, right = cards.splitThreeAndRightOther()
	if three.IsAAA() && right.IsHu() {
		return true
	}
	if three.IsABC() && right.IsHu() {
		return true
	}

	return false
}

func (cards *Cards) ToString() string {
	str := ""
	for _, card := range cards.data{
		str += card.Name() + ","
	}
	return str
}

/*	检查指定的牌可以吃牌的组合
*	假设要吃的牌为C，则需要检查是否存在如下组合：
*	ABC、BCD、BCCD、BCCCD、BCCCD、CDE
*	如果存在AB,则添加组合ABC
*	如果存在BD/BCD/BCCD/BCCCD, 则添加组合BCD
*	如果存在DE,则添加组合CDE
*/
func (cards *Cards) ComputeChi(card *Card) []*Cards {
	if card.IsZiCard() {
		return nil
	}

	length := cards.Len()
	if length < 2 {
		return nil
	}
	cardsSlice := make([]*Cards, 0)

	//检查AB/BD/DE组合
	for idx := 0; idx < length-1; idx++ {
		//if AB组合
		if cards.At(idx).PrevAt(cards.At(idx+1)) && cards.At(idx+1).PrevAt(card) {
			tmp := NewCards()
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(cards.At(idx+1))
			tmp.AppendCard(card)
			cardsSlice = append(cardsSlice, tmp)
		}

		//if BD组合
		if cards.At(idx).PrevAt(card) && card.PrevAt(cards.At(idx+1)) {
			tmp := NewCards()
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(card)
			tmp.AppendCard(cards.At(idx+1))
			cardsSlice = append(cardsSlice, tmp)
		}

		//if DE组合
		if card.PrevAt(cards.At(idx)) && cards.At(idx).PrevAt(cards.At(idx+1)) {
			tmp := NewCards()
			tmp.AppendCard(card)
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(cards.At(idx+1))
			cardsSlice = append(cardsSlice, tmp)
		}

		//if BCD 组合
		//if BCCD 组合
		//if BCCCD 组合
		//todo
	}
	return cardsSlice
}

//检查指定的牌是否可以碰
func (cards *Cards) CheckPeng(card *Card) bool {
	cnt := 0
	for _, c := range cards.data {
		if c.SameAs(card) {
			cnt++
			if cnt == 2 {
				return true
			}
		}
	}
	return false
}

//检查指定的牌是否可以杠
func (cards *Cards) CheckGang(card *Card) bool {
	cnt := 0
	for _, c := range cards.data {
		if c.SameAs(card) {
			cnt++
			if cnt == 3 {
				return true
			}
		}
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

//是否所有的牌都是同一个类型
func (cards *Cards) isAllCardSameType() bool {
	length := cards.Len()
	for idx := 1; idx < length; idx++ {
		if !cards.At(0).SameTypeAs(cards.At(idx)) {
			return false
		}
	}
	return true
}