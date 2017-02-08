package hu_checker

import (
	"mahjong/game_server/card"
	//"mahjong/game_server/log"
)

//清一色对对胡
func init()  {
	FactoryInst().register("Q1SDD_HU",
		func(config *HuConfig) Checker {
			return NewQ1SDD(config)
		},
	)
}

type Q1SDD struct {
	config	*HuConfig
}

func NewQ1SDD(config *HuConfig) *Q1SDD {
	return &Q1SDD{
		config:	config,
	}
}
func (q1sdd *Q1SDD) GetConfig() *HuConfig {
	return q1sdd.config
}

func (q1sdd *Q1SDD) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !q1sdd.config.IsEnabled {
		//log.Debug("q1sdd is not enabled", q1sdd.GetConfig())
		return false, q1sdd.config
	}

	cardType := 0
	cardTypeCnt := 0
	for tmpType := card.CardType_Wan; tmpType < card.Max_CardType; tmpType++{
		cardsInHand := cardsGetter.GetInHandCards(tmpType)
		if cardsInHand != nil && cardsInHand.Len() > 0 {
			if cardsInHand.At(0).IsZiCard() {//清一色对对胡不能有字牌
				//log.Debug("card is zi,", cardsInHand.At(0).Name())
				return false, q1sdd.config
			}
			cardType = tmpType
			cardTypeCnt++
			if cardTypeCnt > 1 {//清一色对对胡不能有大于1种以上的牌
				//log.Debug("card type cnt :", cardType)
				return false, q1sdd.config
			}
		}
	}

	inHandCardNum := cardsGetter.GetInHandCards(cardType).Len()
	pengCardNum := cardsGetter.GetAlreadyPengCards(cardType).Len()
	gangCardNum := cardsGetter.GetAlreadyGangCards(cardType).Len()/4*3
	totalCardNum := inHandCardNum + pengCardNum + gangCardNum
	if totalCardNum != 14 {//不足14张肯定不是清一色
		//log.Debug("q1sdd totalCardNum:", totalCardNum)
		return false, q1sdd.config
	}

	//如果全是AAA类型的牌并且能胡的牌的话，那么牌的数量应该是 (cardTypeCnt-1)*3 + 2
	cardCnt := cardsGetter.GetInHandCards(cardType).CalcDiffCardCnt()
	huCardNum := (cardCnt - 1) * 3 + 2
	if cardsGetter.GetInHandCards(cardType).Len() != huCardNum {//不相等的话手上该类型的牌肯定不是AAA类型和将
		//log.Debug("card in hand is not hu card num, huCardNum :", huCardNum)
		return false, q1sdd.config
	}

	return cardsGetter.IsHu(), q1sdd.config
}
