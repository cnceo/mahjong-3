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

//是否是一样的牌组
func (cards *Cards) SameAs(other *Cards) bool {
	if cards == nil || other == nil {
		return false
	}

	length := other.Len()
	if cards.Len() != length {
		return false
	}

	for idx := 0; idx < length; idx++ {
		if !cards.At(idx).SameAs(other.At(idx)) {
			return false
		}
	}
	return true
}

func (cards *Cards) IsHu() bool  {
	switch cards.Len() {
	case 2:
		return cards.isHu2()
	case 5:
		return cards.isHu5()
	case 8:
		return cards.isHu8()
	case 11:
		return cards.isHu11()
	case 14:
		return cards.isHu14()
	default:
		return false
	}
	return false
}

//检查是否能吃
func (cards *Cards) canChi(whatCard *Card, whatGroup *Cards) bool {
	if whatCard.IsZiCard() {
		return false
	}
	groups := cards.computeChiGroup(whatCard)
	for _, group := range groups {
		if group.SameAs(whatGroup) {
			return true
		}
	}
	return false
}

//检查是否能碰
func (cards *Cards) canPeng(whatCard *Card) bool  {
	return cards.calcSameCardNum(whatCard) >= 2
}

//检查是否能杠
func (cards *Cards) canGang(whatCard *Card) bool {
	return cards.calcSameCardNum(whatCard) >= 3
}

//计算与指定牌一样的牌的数量
func (cards *Cards) calcSameCardNum(whatCard *Card) int {
	num := 0
	for _, card := range cards.data {
		if card.SameAs(whatCard) {
			num++
		}
	}
	return num
}

/*	计算指定的牌可以吃牌的组合
*	假设要吃的牌为C，则需要检查是否存在如下组合：
*	ABC、BCD、BCCD、BCCCD、BCCCD、CDE
*	如果存在AB,则添加组合ABC
*	如果存在BD/BCD/BCCD/BCCCD, 则添加组合BCD
*	如果存在DE,则添加组合CDE
*/
func (cards *Cards) computeChiGroup(card *Card) []*Cards {
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
		//if AB组合，加上card后相当于ABC
		if IsABC(cards.At(idx), cards.At(idx+1), card) {
			tmp := NewCards()
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(cards.At(idx+1))
			tmp.AppendCard(card)
			cardsSlice = append(cardsSlice, tmp)
		}

		//if BD组合，加上card后相当于BCD
		if IsABC(cards.At(idx), card, cards.At(idx+1))  {
			tmp := NewCards()
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(card)
			tmp.AppendCard(cards.At(idx+1))
			cardsSlice = append(cardsSlice, tmp)
		}

		//if DE组合，加上card后相当于CDE
		if IsABC(card, cards.At(idx), cards.At(idx+1)) {
			tmp := NewCards()
			tmp.AppendCard(card)
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(cards.At(idx+1))
			cardsSlice = append(cardsSlice, tmp)
		}

		//if BCD 组合，加上card后相当于BCCD
		if IsABBC(cards.At(idx), card, cards.At(idx+1), cards.At(idx+2)) {
			tmp := NewCards()
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(card)
			tmp.AppendCard(cards.At(idx+2))
			cardsSlice = append(cardsSlice, tmp)
		}

		//if BCCD 组合，加上card后相当于BCCCD
		if IsABBBC(cards.At(idx), card, cards.At(idx+1), cards.At(idx+2), cards.At(idx+3)) {
			tmp := NewCards()
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(card)
			tmp.AppendCard(cards.At(idx+3))
			cardsSlice = append(cardsSlice, tmp)
		}

		//if BCCCD 组合，加上card后相当于BCCCCD
		if IsABBBBC(cards.At(idx), card, cards.At(idx+1), cards.At(idx+2), cards.At(idx+3), cards.At(idx+4)) {
			tmp := NewCards()
			tmp.AppendCard(cards.At(idx))
			tmp.AppendCard(card)
			tmp.AppendCard(cards.At(idx+4))
			cardsSlice = append(cardsSlice, tmp)
		}
	}
	return cardsSlice
}


func (cards *Cards) ToString() string {
	str := ""
	for _, card := range cards.data{
		str += card.Name() + ","
	}
	return str
}

//把牌分成左右2份：[:idx], [idx+1:]
func (cards *Cards) Split(idx int) (left, right *Cards){
	length := cards.Len()
	if length <= idx {
		return cards, nil
	}
	left = &Cards{
		data:	cards.data[0 : length-idx],
	}
	right = &Cards{
		data:	cards.data[length-idx:],
	}
	return left, right
}

//是否所有的牌都是同一个类型
func (cards *Cards) IsAllCardSameType() bool {
	length := cards.Len()
	for idx := 1; idx < length; idx++ {
		if !cards.At(0).SameTypeAs(cards.At(idx)) {
			return false
		}
	}
	return true
}

//获取不同牌的类型的数量
func (cards *Cards) CalcDiffCardCnt() int {
	has := make(map[int64]bool)
	for _, card := range cards.data {
		has[card.MakeKey()] = true
	}
	return len(has)
}

func (cards *Cards) CalcCardCntAsSameCardType(cardType int) int {
	cnt := 0
	tmp := &Card{CardType:cardType}
	for _, card := range cards.data {
		if card.SameTypeAs(tmp) {
			cnt++
		}
	}
	return cnt
}

//胡2张牌
func (cards *Cards)isHu2() bool {
	if cards.Len() != 2 {
		return false
	}
	return IsAA(cards.data[0], cards.data[1])
}

//胡5张牌
func (cards *Cards) isHu5() bool {
	if cards.Len() != 5 {
		return false
	}

	//AA+BCD, 2 + 3
	if IsAA(cards.data[0], cards.data[1]) && Is3CardsOk(cards.data[2:5]...){
		return true
	}

	//ABC + DD, 3 + 2
	if Is3CardsOk(cards.data[0:3]...) && IsAA(cards.data[3], cards.data[4]) {
		return true
	}

	//A + BBB + C
	if IsAAA(cards.data[1], cards.data[2], cards.data[3]) &&
		Is3CardsOk(cards.data[0], cards.data[1], cards.data[4]){
		return true
	}

	return false
}

//胡8张牌
func (cards *Cards) isHu8() bool {
	if cards.Len() != 8 {
		return false
	}
	//2 + 6
	if IsAA(cards.data[0], cards.data[1]) &&
		Is6CardsOk(cards.data[2:8]...) {
		return true
	}

	//6 + 2
	if Is6CardsOk(cards.data[0:6]...)&&
		IsAA(cards.data[6], cards.data[7]) {
		return true
	}

	//3 + 2 + 3
	if Is3CardsOk(cards.data[0:3]...) && IsAA(cards.data[3], cards.data[4]) &&
		Is3CardsOk(cards.data[5:8]...) {
		return true
	}
	return false
}

//胡11张牌
func (cards *Cards) isHu11() bool {
	if cards.Len() != 11 {
		return false
	}

	//最左边的两个为眼， 2 + 9
	if IsAA(cards.data[0], cards.data[1]) &&
		Is9CardsOk(cards.data[2:11]...) {
		return true
	}

	//最右边的两个为眼， 9 + 2
	if Is9CardsOk(cards.data[0:9]...) &&
		IsAA(cards.data[9], cards.data[10]) {
		return true
	}

	//中间左边两个为眼， 3 + 2 + 6
	if Is3CardsOk(cards.data[0:3]...) && IsAA(cards.data[3], cards.data[4]) &&
		Is6CardsOk(cards.data[5:11]...) {
		return true
	}

	//中间右边两个为眼， 6 + 2 + 3
	if Is6CardsOk(cards.data[0:6]...) && IsAA(cards.data[6], cards.data[7]) &&
		Is3CardsOk(cards.data[8:11]...){
		return true
	}
	return false
}

//胡14张牌
func (cards *Cards) isHu14() bool {
	if cards.Len() != 14 {
		return false
	}

	// 2 + 12
	if IsAA(cards.data[0], cards.data[1]) && Is12CardsOk(cards.data[2:14]...) {
		return true
	}

	// 3 + 2 + 9
	if Is3CardsOk(cards.data[0:3]...) && IsAA(cards.data[3], cards.data[4]) &&
		Is9CardsOk(cards.data[5:14]...) {
		return true
	}

	// 6 + 2 +6
	if Is6CardsOk(cards.data[0:6]...) && IsAA(cards.data[6], cards.data[7]) &&
		Is6CardsOk(cards.data[8:14]...) {
		return true
	}

	// 9 + 2 + 3
	if Is9CardsOk(cards.data[0:9]...) && IsAA(cards.data[9], cards.data[10]) &&
		Is3CardsOk(cards.data[11:14]...) {
		return true
	}

	// 12 + 3
	if Is12CardsOk(cards.data[0:12]...) && IsAA(cards.data[0], cards.data[1]) {
		return true
	}
	return false
}