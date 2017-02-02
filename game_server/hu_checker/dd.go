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
		//fmt.Println("chiCards", chiCards, "cardType :", cardType)
		if chiCards != nil && chiCards.Len() > 0 {
			return false, dd.config
		}
	}

	cardsInHand := cardsGetter.GetInHandCards()
	cardCnt := cardsInHand.CalcDiffCardCnt()

	//如果全是3张的牌并且能胡的牌的话，那么牌的数量应该是 (cardTypeCnt-1)*3 + 2
	huCardNum := (cardCnt - 1) * 3 + 2
	//fmt.Println("huCardNum :", huCardNum, "cardsInHand.Len :", cardsInHand.Len(), "cardCnt:", cardCnt)
	if cardsInHand.Len() == huCardNum && cardsInHand.IsHu() {
		return true, dd.config
	}
	return false, dd.config
}

