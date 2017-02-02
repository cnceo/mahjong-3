package hu_checker

import "mahjong/game_server/card"

//大四喜

type D4X struct {
	config	*HuConfig
}

func NewD4X(config *HuConfig) *D4X {
	return &D4X{
		config:	config,
	}
}

func (d4x *D4X) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !d4x.config.IsEnabled {
		return false, d4x.config
	}

	cardsInHand := cardsGetter.GetInHandCards()
	inHandFengCardNum := cardsInHand.CalcCardCntAsSameCardType(card.CardType_Feng)
	pengFengCardNum := cardsGetter.GetAlreadyPengCards(card.CardType_Feng).Len()
	gangFengCardNum := cardsGetter.GetAlreadyGangCards(card.CardType_Feng).Len()

	totalDiffFengCard := inHandFengCardNum/3 + pengFengCardNum/3 + gangFengCardNum/4

	//4种风牌的碰或杠
	if totalDiffFengCard == 4 && cardsInHand.IsHu() {
		return true, d4x.config
	}
	return false, d4x.config
}
