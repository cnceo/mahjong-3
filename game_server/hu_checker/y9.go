package hu_checker

import "mahjong/game_server/card"

//幺九胡：只能有番牌或者1、9

type Y9 struct {
	config	*HuConfig
}

func NewY9(config *HuConfig) *Y9 {
	return &Y9{
		config:	config,
	}
}

func (y9 *Y9) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !y9.config.IsEnabled {
		return false, y9.config
	}

	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++ {
		cardsInHand := cardsGetter.GetInHandCards(cardType)
		for _, card := range cardsInHand.Data() {
			if !y9.isY9Card(card){
				return false, y9.config
			}
		}
	}

	//不能有吃的牌
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		chiCards := cardsGetter.GetAlreadyChiCards(cardType)
		if chiCards != nil && chiCards.Len() > 0 {
			return false, y9.config
		}
	}

	//不能有碰的非幺九的牌
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		if card.CardType_Jian == cardType || card.CardType_Feng == cardType {
			continue
		}
		pengCards := cardsGetter.GetAlreadyPengCards(cardType)
		for _, card := range pengCards.Data() {
			if !y9.isY9Card(card) {
				return false, y9.config
			}
		}
	}

	//不能有杠的非幺九的牌
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		if card.CardType_Jian == cardType || card.CardType_Feng == cardType {
			continue
		}
		gangCards := cardsGetter.GetAlreadyGangCards(cardType)
		for _, card := range gangCards.Data() {
			if !y9.isY9Card(card) {
				return false, y9.config
			}
		}
	}

	return cardsGetter.IsHu(), y9.config
}

func (y9 *Y9) isY9Card(card* card.Card) bool {
	if card.IsZiCard() {
		return true
	}

	if card.CardNo == 1 || card.CardNo == 9 {
		return true
	}
	return false
}