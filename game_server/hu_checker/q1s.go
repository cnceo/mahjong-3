package hu_checker

import (
	"mahjong/game_server/card"
)

//清1色
func init()  {
	FactoryInst().register("Q1S_HU",
		func(config *HuConfig) Checker {
			return NewQ1S(config)
		},
	)
}

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

	cardType := 0
	cardTypeCnt := 0
	for tmpType := card.CardType_Wan; tmpType < card.Max_CardType; tmpType++{
		cardsInHand := cardsGetter.GetInHandCards(tmpType)
		if cardsInHand != nil && cardsInHand.Len() > 0 {
			if cardsInHand.At(0).IsZiCard() {//清一色不能有字牌
				return false, q1s.config
			}
			cardType = tmpType
			cardTypeCnt++
			if cardTypeCnt > 1 {//清一色不能有大于1种以上的牌
				return false, q1s.config
			}
		}
	}
/*
	//不能有吃非不同类型的牌
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
*/
	inHandCardNum := cardsGetter.GetInHandCards(cardType).Len()
	chiCardNum := cardsGetter.GetAlreadyChiCards(cardType).Len()
	pengCardNum := cardsGetter.GetAlreadyPengCards(cardType).Len()
	gangCardNum := cardsGetter.GetAlreadyGangCards(cardType).Len()/4*3
	totalCardNum := inHandCardNum + chiCardNum + pengCardNum + gangCardNum
	if totalCardNum != 14 {//肯定不是清一色
		return false, q1s.config
	}

	return cardsGetter.IsHu(), q1s.config
}
