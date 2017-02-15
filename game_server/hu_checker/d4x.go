package hu_checker

import "mahjong/game_server/card"

//大四喜
func init()  {
	FactoryInst().register("D4X_HU",
		func(config *HuConfig) Checker {
			return NewD4X(config)
		},
	)
}

type D4X struct {
	config	*HuConfig
}

func NewD4X(config *HuConfig) *D4X {
	return &D4X{
		config:	config,
	}
}

func (d4x *D4X) GetConfig() *HuConfig {
	return d4x.config
}

func (d4x *D4X) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !d4x.config.IsEnabled {
		return false, d4x.config
	}

	inHandFengCardNum := cardsGetter.GetCardsInHandByType(card.CardType_Feng).Len()
	pengFengCardNum := cardsGetter.GetAlreadyPengCards(card.CardType_Feng).Len()
	gangFengCardNum := cardsGetter.GetAlreadyGangCards(card.CardType_Feng).Len()

	totalDiffFengCard := inHandFengCardNum/3 + pengFengCardNum/3 + gangFengCardNum/4

	//4种风牌的碰或杠
	if totalDiffFengCard != 4 {//没用全部风牌，肯定不是大四喜
		return false, d4x.config
	}
	return cardsGetter.IsHu(), d4x.config
}
