package hu_checker

import "mahjong/game_server/card"

//小四喜

type X4X struct {
	config	*HuConfig
}

func NewX4X(config *HuConfig) *X4X {
	return &X4X{
		config:	config,
	}
}

func (x4x *X4X) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !x4x.config.IsEnabled {
		return false, x4x.config
	}

	cardsInHand := cardsGetter.GetInHandCards()
	inHandFengCardNum := cardsInHand.CalcCardCntAsSameCardType(card.CardType_Feng)
	pengFengCardNum := cardsGetter.GetAlreadyPengCards(card.CardType_Feng).Len()
	gangFengCardNum := cardsGetter.GetAlreadyGangCards(card.CardType_Feng).Len()

	totalDiffFengCard := inHandFengCardNum/3 + pengFengCardNum/3 + gangFengCardNum/4

	//3个风牌的碰或杠，1对风牌做眼
	if totalDiffFengCard == 3 && inHandFengCardNum % 3 == 2 && cardsInHand.IsHu() {
		return true, x4x.config
	}
	return false, x4x.config
}
