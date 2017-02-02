package hu_checker

import "mahjong/game_server/card"

//大三元

type D3Y struct {
	config	*HuConfig
}

func NewD3Y(config *HuConfig) *D3Y {
	return &D3Y{
		config:	config,
	}
}

func (d3y *D3Y) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !d3y.config.IsEnabled {
		return false, d3y.config
	}

	cardsInHand := cardsGetter.GetInHandCards()
	inHandJianCardNum := cardsInHand.CalcCardCntAsSameCardType(card.CardType_Jian)
	pengJianCardNum := cardsGetter.GetAlreadyPengCards(card.CardType_Jian).Len()
	gangJianCardNum := cardsGetter.GetAlreadyGangCards(card.CardType_Jian).Len()

	totalDiffJianCard := inHandJianCardNum/3 + pengJianCardNum/3 + gangJianCardNum/4

	if totalDiffJianCard == 3 && cardsInHand.IsHu() {
		return true, d3y.config
	}
	return false, d3y.config
}
