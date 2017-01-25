package player

import "mahjong/game_server/card"



type Player struct {
	CardsInHand			[]*card.Cards		//手上的牌
	CardsAlreadyChi		[]*card.Cards		//已经吃了的牌
	CardsAlreadyPeng	[]*card.Cards		//已经碰了的牌
	CardsAlreadyGang	[]*card.Cards		//已经杠了的牌
}

func NewPlayer() *Player {
	player :=  &Player{
	}
	player.CardsInHand = player.initCardsSlice()
	player.CardsAlreadyChi = player.initCardsSlice()
	player.CardsAlreadyPeng = player.initCardsSlice()
	player.CardsAlreadyGang = player.initCardsSlice()
	return player
}

func (player *Player) initCardsSlice()[]*card.Cards {
	cardsSlice := make([]*card.Cards, card.Max_CardType)
	for idx := 0; idx < card.Max_CardType; idx++ {
		cardsSlice[idx] = card.NewCards()
	}
	return cardsSlice
}

func (player *Player) AddCard(card *card.Card) {
	player.CardsInHand[card.CardType].AddAndSort(card)
}

func (player *Player) Chi(card *card.Card) bool {
	player.CardsAlreadyChi[card.CardType].AppendCard(card)
	return true
}

func (player *Player) Peng(card *card.Card) bool {
	player.CardsAlreadyPeng[card.CardType].AppendCard(card)
	return true
}

func (player *Player) Gang(card *card.Card) bool {
	player.CardsAlreadyGang[card.CardType].AppendCard(card)
	return true
}

func (player *Player) cardsSliceToString(cardsSlice[]*card.Cards) string{
	str := ""
	for _, cards := range cardsSlice{
		str += cards.ToString()
	}
	return str
}

func (player *Player) ToString() string{
	str := ""
	str += "cardsInHand:\n" + player.cardsSliceToString(player.CardsInHand) + "\n"
	str += "cardsAlreadyChi:\n" + player.cardsSliceToString(player.CardsAlreadyChi) + "\n"
	str += "cardsAlreadyPeng:\n" + player.cardsSliceToString(player.CardsAlreadyPeng) + "\n"
	str += "cardsAlreadyGang:\n" + player.cardsSliceToString(player.CardsAlreadyGang) + "\n"
	return str
}