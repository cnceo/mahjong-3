package card

type PlayingCards struct {
	cardsInHand			*Cards			//手上的牌
	cardsDiffType		[]*Cards		//手上的牌，分类多存一份
	cardsAlreadyChi		[]*Cards		//已经吃了的牌
	cardsAlreadyPeng	[]*Cards		//已经碰了的牌
	cardsAlreadyGang	[]*Cards		//已经杠了的牌
}

func NewPlayingCards() *PlayingCards {
	cards :=  &PlayingCards{
	}
	cards.cardsInHand = NewCards()
	cards.cardsAlreadyChi = cards.initCardsSlice()
	cards.cardsAlreadyPeng = cards.initCardsSlice()
	cards.cardsAlreadyGang = cards.initCardsSlice()
	return cards
}

//增加一张牌
func (playingCards *PlayingCards) AddCard(card *Card) {
	playingCards.cardsInHand.AddAndSort(card)
}

//丢弃一张牌
func (playingCards *PlayingCards) DropCard(card *Card) bool {
	succ := playingCards.cardsInHand.TakeWay(card)
	playingCards.cardsInHand.Sort()
	return succ
}

//吃牌，要吃whatCard，以及吃哪个组合whatGroup
func (playingCards *PlayingCards) Chi(whatCard *Card, whatGroup *Cards) bool {
	if !playingCards.canChi(whatCard, whatGroup) {
		return false
	}

	for _, card := range whatGroup.Data() {//移动除了whatCard以外的card到cardsAlreadyChi
		if card.SameAs(whatCard) {
			continue
		}
		playingCards.cardsInHand.TakeWay(card)
		playingCards.cardsAlreadyChi[whatCard.CardType].AppendCard(card)
	}

	//最后把whatCard加入cardsAlreadyChi
	playingCards.cardsAlreadyChi[whatCard.CardType].AddAndSort(whatCard)

	return true
}

//碰牌
func (playingCards *PlayingCards) Peng(whatCard *Card) bool {
	if !playingCards.canPeng(whatCard) {
		return false
	}

	playingCards.cardsInHand.TakeWay(whatCard)
	playingCards.cardsInHand.TakeWay(whatCard)
	playingCards.cardsAlreadyPeng[whatCard.CardType].AppendCard(whatCard)
	playingCards.cardsAlreadyPeng[whatCard.CardType].AppendCard(whatCard)
	playingCards.cardsAlreadyPeng[whatCard.CardType].AddAndSort(whatCard)
	return true
}

//杠牌
func (playingCards *PlayingCards) Gang(whatCard *Card) bool {
	if !playingCards.canGang(whatCard) {
		return false
	}

	playingCards.cardsInHand.TakeWay(whatCard)
	playingCards.cardsInHand.TakeWay(whatCard)
	playingCards.cardsInHand.TakeWay(whatCard)
	playingCards.cardsAlreadyGang[whatCard.CardType].AppendCard(whatCard)
	playingCards.cardsAlreadyGang[whatCard.CardType].AppendCard(whatCard)
	playingCards.cardsAlreadyGang[whatCard.CardType].AppendCard(whatCard)
	playingCards.cardsAlreadyGang[whatCard.CardType].AddAndSort(whatCard)
	return true
}

func (playingCards *PlayingCards) ToString() string{
	str := ""
	str += "cardsInHand:\n" + playingCards.cardsInHand.ToString() + "\n"
	str += "cardsAlreadyChi:\n" + playingCards.cardsSliceToString(playingCards.cardsAlreadyChi)
	str += "cardsAlreadyPeng:\n" + playingCards.cardsSliceToString(playingCards.cardsAlreadyPeng)
	str += "cardsAlreadyGang:\n" + playingCards.cardsSliceToString(playingCards.cardsAlreadyGang)
	return str
}

/*	计算指定的牌可以吃牌的组合
*/
func (playingCards *PlayingCards) ComputeChiGroup(card *Card) []*Cards {
	return playingCards.cardsInHand.computeChiGroup(card)
}

//检查是否能吃
func (playingCards *PlayingCards) canChi(whatCard *Card, whatGroup *Cards) bool {
	return playingCards.cardsInHand.canChi(whatCard, whatGroup)
}

//检查是否能碰
func (playingCards *PlayingCards) canPeng(whatCard *Card) bool  {
	return playingCards.cardsInHand.canPeng(whatCard)
}

//检查是否能杠
func (playingCards *PlayingCards) canGang(whatCard *Card) bool {
	return playingCards.cardsInHand.canGang(whatCard)
}


//初始化cards
func (playingCards *PlayingCards) initCardsSlice()[]*Cards {
	cardsSlice := make([]*Cards, Max_CardType)
	for idx := 0; idx < Max_CardType; idx++ {
		cardsSlice[idx] = NewCards()
	}
	return cardsSlice
}

func (playingCards *PlayingCards) cardsSliceToString(cardsSlice []*Cards) string{
	str := ""
	for _, cards := range cardsSlice{
		str += cards.ToString() + "\n"
	}
	return str
}

func (playingCards *PlayingCards) GetInHandCards() *Cards{
	return playingCards.cardsInHand
}

func (playingCards *PlayingCards) GetAlreadyChiCards(cardType int) *Cards{
	if cardType < 0 || cardType > Max_CardType {
		return nil
	}
	return playingCards.cardsAlreadyChi[cardType]
}

func (playingCards *PlayingCards) GetAlreadyPengCards(cardType int) *Cards{
	if cardType < 0 || cardType > Max_CardType {
		return nil
	}
	return playingCards.cardsAlreadyPeng[cardType]
}

func (playingCards *PlayingCards) GetAlreadyGangCards(cardType int) *Cards{
	if cardType < 0 || cardType > Max_CardType {
		return nil
	}
	return playingCards.cardsAlreadyGang[cardType]
}
