package hu_checker

import (
	"mahjong/game_server/card"
)

//清1色

type Q1S struct {
	config	*HuConfig
}

func NewQ1S(config *HuConfig) *Q1S {
	return &Q1S{
		config:	config,
	}
}

func (q1s *Q1S) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !q1s.config.IsEnabled {
		//fmt.Println(1)
		return false, q1s.config
	}

	cardsInHand := cardsGetter.GetInHandCards()
	if cardsInHand.Len() == 0  {
		//fmt.Println(2)
		return false, q1s.config
	}

	if cardsInHand.At(0).IsZiCard() || !cardsInHand.IsAllCardSameType() {
		//fmt.Println(3)
		return false, q1s.config
	}

	cardType := cardsInHand.At(0).CardType

	//不能有吃的牌
	for tmpType := card.CardType_Wan; tmpType < card.Max_CardType; tmpType++{
		if cardType == tmpType {
			continue
		}
		chiCards := cardsGetter.GetAlreadyChiCards(tmpType)
		if chiCards != nil && chiCards.Len() > 0 {
			//fmt.Println(4)
			return false, q1s.config
		}
	}

	//不能有碰非不同类型的牌
	for tmpType := card.CardType_Wan; tmpType < card.Max_CardType; tmpType++{
		if tmpType == cardType {
			continue
		}
		pengCards := cardsGetter.GetAlreadyPengCards(tmpType)
		////fmt.Println("chiCards", chiCards, "cardType :", cardType)
		if pengCards != nil && pengCards.Len() > 0 {
			//fmt.Println(5)
			return false, q1s.config
		}
	}

	//不能有杠的非不同类型的牌
	for tmpType := card.CardType_Wan; tmpType < card.Max_CardType; tmpType++{
		if tmpType == cardType {
			continue
		}
		gangCards := cardsGetter.GetAlreadyGangCards(tmpType)
		////fmt.Println("chiCards", chiCards, "cardType :", cardType)
		if gangCards != nil && gangCards.Len() > 0 {
			//fmt.Println(6)
			return false, q1s.config
		}
	}

	return cardsInHand.IsHu(), q1s.config
}
