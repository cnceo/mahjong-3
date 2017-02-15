package hu_checker

import "mahjong/game_server/card"

//字一色/全番，全部都是番牌：东南西北、中发白
func init()  {
	FactoryInst().register("Z1S_HU",
		func(config *HuConfig) Checker {
			return NewZ1S(config)
		},
	)
}

type Z1S struct {
	config	*HuConfig
}

func NewZ1S(config *HuConfig) *Z1S {
	return &Z1S{
		config:	config,
	}
}

func (z1s *Z1S) GetConfig() *HuConfig {
	return z1s.config
}

func (z1s *Z1S) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !z1s.config.IsEnabled {
		return false, z1s.config
	}

	totalCardNum := cardsGetter.GetCardsInHandByType(card.CardType_Feng).Len()
	totalCardNum += cardsGetter.GetCardsInHandByType(card.CardType_Jian).Len()

	totalCardNum += cardsGetter.GetAlreadyPengCards(card.CardType_Feng).Len()
	totalCardNum += cardsGetter.GetAlreadyPengCards(card.CardType_Jian).Len()

	totalCardNum += cardsGetter.GetAlreadyGangCards(card.CardType_Feng).Len()/4*3
	totalCardNum += cardsGetter.GetAlreadyGangCards(card.CardType_Jian).Len()/4*3

	if totalCardNum != 14 {//一定还有其他非番牌的牌
		return false, z1s.config
	}
	return cardsGetter.IsHu(), z1s.config
}
