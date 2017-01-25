package card

type PlayingCards struct {
	cardsInHand			*Cards		//手上的牌
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

func (playingCards *PlayingCards) AddCard(card *Card) {
	playingCards.cardsInHand.AddAndSort(card)
}

//吃牌，要吃whatCard，以及吃哪个组合whatGroup
func (playingCards *PlayingCards) Chi(whatCard *Card, whatGroup *Cards) bool {
	if whatCard.IsZiCard() {//字牌不能吃
		return false
	}

	if !playingCards.canChi(whatCard, whatGroup) {
		return false
	}
	for _, card := range whatGroup.Data() {
		if card.SameAs(whatCard) {
			continue
		}
		playingCards.cardsInHand.TakeWay(card)
		playingCards.cardsAlreadyChi[whatCard.CardType].AppendCard(card)
	}
	playingCards.cardsAlreadyChi[whatCard.CardType].AddAndSort(whatCard)

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

func (playingCards *PlayingCards) initCardsSlice()[]*Cards {
	cardsSlice := make([]*Cards, Max_CardType-1)
	for idx := 0; idx < Max_CardType-1; idx++ {
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

//检查是否能吃
func (playingCards *PlayingCards) canChi(whatCard *Card, whatGroup *Cards) bool {
	groups := playingCards.cardsInHand.ComputeChiGroup(whatCard)
	for _, group := range groups {
		if group.SameAs(whatGroup) {
			return true
		}
	}
	return false
}