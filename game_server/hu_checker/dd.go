package hu_checker

import (
	"mahjong/game_server/card"
)

//对对胡

type DD struct {
	config	*HuConfig
}

func NewDD(config *HuConfig) *DD {
	return &DD{
		config:	config,
	}
}

func (dd *DD) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !dd.config.IsEnabled {
		return false, dd.config
	}

	//不能有吃的牌
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		chiCards := cardsGetter.GetAlreadyChiCards(cardType)
		if chiCards != nil && chiCards.Len() > 0 {
			return false, dd.config
		}
	}

	//计算不同的牌的数量, 所有牌的数量
	diffCardCnt := 0
	totalCardCnt := 0
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		cardsInHand := cardsGetter.GetInHandCards(cardType)
		diffCardCnt += cardsInHand.CalcDiffCardCnt()
		totalCardCnt += cardsInHand.Len()
	}

	//如果是对对胡的话，则全是AAA类型的牌并且能胡的牌的话，那么牌的数量应该是 (cardTypeCnt-1)*3 + 2
	huCardNum := (diffCardCnt - 1) * 3 + 2
	//fmt.Println("huCardNum :", huCardNum, "cardsInHand.Len :", cardsInHand.Len(), "cardCnt:", cardCnt)
	if totalCardCnt != huCardNum {//肯定含有不是AAA类型的牌
		return false, dd.config
	}
	return cardsGetter.IsHu(), dd.config
}

