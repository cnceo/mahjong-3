package hu_checker

import "mahjong/game_server/card"

//全番，全部都是番牌：东南西北、中发白

type QF struct {
	config	*HuConfig
}

func NewQF(config *HuConfig) *QF {
	return &QF{
		config:	config,
	}
}

func (qf *QF) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !qf.config.IsEnabled {
		return false, qf.config
	}

	totalCardNum := cardsGetter.GetInHandCards(card.CardType_Feng).Len()
	totalCardNum += cardsGetter.GetInHandCards(card.CardType_Jian).Len()

	totalCardNum += cardsGetter.GetAlreadyPengCards(card.CardType_Feng).Len()
	totalCardNum += cardsGetter.GetAlreadyPengCards(card.CardType_Jian).Len()

	totalCardNum += cardsGetter.GetAlreadyGangCards(card.CardType_Feng).Len()/4*3
	totalCardNum += cardsGetter.GetAlreadyGangCards(card.CardType_Jian).Len()/4*3

	if totalCardNum != 14 {//一定还有其他非番牌的牌
		return false, qf.config
	}
	return cardsGetter.IsHu(), qf.config
}
