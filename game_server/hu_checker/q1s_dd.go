package hu_checker

import (
	"mahjong/game_server/card"
)

//清一色对对胡

type Q1SDD struct {
	config	*HuConfig
}

func NewQ1SDD(config *HuConfig) *Q1SDD {
	return &Q1SDD{
		config:	config,
	}
}

func (q1sdd *Q1SDD) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !q1sdd.config.IsEnabled {
		return false, q1sdd.config
	}

	cardsInHand := cardsGetter.GetInHandCards()
	if cardsInHand.Len() == 0  {
		//fmt.Println(1)
		return false, q1sdd.config
	}

	if cardsInHand.At(0).IsZiCard() || !cardsInHand.IsAllCardSameType() {
		//fmt.Println(2)
		return false, q1sdd.config
	}

	cardType := cardsInHand.At(0).CardType

	//不能有吃的牌
	for tmpType := card.CardType_Wan; tmpType < card.Max_CardType; tmpType++{
		chiCards := cardsGetter.GetAlreadyChiCards(tmpType)
		if chiCards != nil && chiCards.Len() > 0 {
			//fmt.Println(3)
			return false, q1sdd.config
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
			//fmt.Println(4)
			return false, q1sdd.config
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
			//fmt.Println(5)
			return false, q1sdd.config
		}
	}

	//如果全是AAA类型的牌并且能胡的牌的话，那么牌的数量应该是 (cardTypeCnt-1)*3 + 2
	cardCnt := cardsInHand.CalcDiffCardCnt()
	huCardNum := (cardCnt - 1) * 3 + 2
	//fmt.Println("huCardNum :", huCardNum, "cardsInHand.Len :", cardsInHand.Len(), "cardCnt:", cardCnt)
	if cardsInHand.Len() == huCardNum && cardsInHand.IsHu() {
		return true, q1sdd.config
	}
	return false, q1sdd.config
}
