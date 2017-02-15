package hu_checker

import "mahjong/game_server/card"

type CardsGetter interface {
	IsHu() bool
	GetCardsInHandByType(cardType int) *card.Cards
	GetAlreadyChiCards(cardType int) *card.Cards
	GetAlreadyPengCards(cardType int) *card.Cards
	GetAlreadyGangCards(cardType int) *card.Cards
}
