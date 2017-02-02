package play

import (
	"mahjong/game_server/hu_checker"
	"mahjong/game_server/card"
)

type PlayingCards struct {
	magicCards			[]*card.Card			//可变的牌
	cardsInHand			*card.Cards			//手上的牌
	cardsAlreadyChi		[]*card.Cards		//已经吃了的牌
	cardsAlreadyPeng	[]*card.Cards		//已经碰了的牌
	cardsAlreadyGang	[]*card.Cards		//已经杠了的牌
}

func NewPlayingCards() *PlayingCards {
	cards :=  &PlayingCards{
	}
	cards.magicCards = make([]*card.Card, 0)
	cards.cardsInHand = card.NewCards()
	cards.cardsAlreadyChi = cards.initCardsSlice()
	cards.cardsAlreadyPeng = cards.initCardsSlice()
	cards.cardsAlreadyGang = cards.initCardsSlice()
	return cards
}

//增加一张牌
func (playingCards *PlayingCards) AddCard(card *card.Card) {
	playingCards.cardsInHand.AddAndSort(card)
}

//丢弃一张牌
func (playingCards *PlayingCards) DropCard(card *card.Card) bool {
	succ := playingCards.cardsInHand.TakeWay(card)
	playingCards.cardsInHand.Sort()
	return succ
}

//吃牌，要吃whatCard，以及吃哪个组合whatGroup
func (playingCards *PlayingCards) Chi(whatCard *card.Card, whatGroup *card.Cards) bool {
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
func (playingCards *PlayingCards) Peng(whatCard *card.Card) bool {
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
func (playingCards *PlayingCards) Gang(whatCard *card.Card) bool {
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

//检查是否能吃
func (playingCards *PlayingCards) canChi(whatCard *card.Card, whatGroup *card.Cards) bool {
	return playingCards.cardsInHand.CanChi(whatCard, whatGroup)
}

//检查是否能碰
func (playingCards *PlayingCards) canPeng(whatCard *card.Card) bool  {
	return playingCards.cardsInHand.CanPeng(whatCard)
}

//检查是否能杠
func (playingCards *PlayingCards) canGang(whatCard *card.Card) bool {
	return playingCards.cardsInHand.CanGang(whatCard)
}


//初始化cards
func (playingCards *PlayingCards) initCardsSlice()[]*card.Cards {
	cardsSlice := make([]*card.Cards, card.Max_CardType)
	for idx := 0; idx < card.Max_CardType; idx++ {
		cardsSlice[idx] = card.NewCards()
	}
	return cardsSlice
}

func (playingCards *PlayingCards) cardsSliceToString(cardsSlice []*card.Cards) string{
	str := ""
	for _, cards := range cardsSlice{
		str += cards.ToString() + "\n"
	}
	return str
}

func (playingCards *PlayingCards) GetInHandCards() *card.Cards{
	return playingCards.cardsInHand
}

func (playingCards *PlayingCards) GetAlreadyChiCards(cardType int) *card.Cards{
	if cardType < 0 || cardType > card.Max_CardType {
		return nil
	}
	return playingCards.cardsAlreadyChi[cardType]
}

func (playingCards *PlayingCards) GetAlreadyPengCards(cardType int) *card.Cards{
	if cardType < 0 || cardType > card.Max_CardType {
		return nil
	}
	return playingCards.cardsAlreadyPeng[cardType]
}

func (playingCards *PlayingCards) GetAlreadyGangCards(cardType int) *card.Cards{
	if cardType < 0 || cardType > card.Max_CardType {
		return nil
	}
	return playingCards.cardsAlreadyGang[cardType]
}

func (playingCards *PlayingCards) GetMagicCards() []*card.Card {
	return playingCards.magicCards
}

func (playingCards *PlayingCards) IsHu (checker hu_checker.Checker) bool {
	for  {
		
	}
}
