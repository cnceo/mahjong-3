package hu_checker

import "mahjong/game_server/card"

//小四喜
func init()  {
	FactoryInst().register("X4X_HU",
		func(config *HuConfig) Checker {
			return NewX4X(config)
		},
	)
}

type X4X struct {
	config	*HuConfig
}

func NewX4X(config *HuConfig) *X4X {
	return &X4X{
		config:	config,
	}
}

func (x4x *X4X) GetConfig() *HuConfig {
	return x4x.config
}

func (x4x *X4X) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !x4x.config.IsEnabled {
		return false, x4x.config
	}

	inHandFengCardNum := cardsGetter.GetInHandCards(card.CardType_Feng).Len()
	pengFengCardNum := cardsGetter.GetAlreadyPengCards(card.CardType_Feng).Len()
	gangFengCardNum := cardsGetter.GetAlreadyGangCards(card.CardType_Feng).Len()/4*3

	totalFengNum := inHandFengCardNum + pengFengCardNum + gangFengCardNum

	if totalFengNum != 11 || inHandFengCardNum % 3 != 2 {//不是11张风牌，或者没有一对风牌做将
		return false, x4x.config
	}

	return cardsGetter.IsHu(), x4x.config
}
