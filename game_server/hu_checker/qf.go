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

	cardsInHand := cardsGetter.GetInHandCards()
	for _, card := range cardsInHand.Data() {
		if !card.IsZiCard() {
			return false, qf.config
		}
	}

	//不能有吃的牌
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		chiCards := cardsGetter.GetAlreadyChiCards(cardType)
		if chiCards != nil && chiCards.Len() > 0 {
			return false, qf.config
		}
	}

	//不能有碰的非字的牌
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		if card.CardType_Jian == cardType || card.CardType_Feng == cardType {
			continue
		}
		pengCards := cardsGetter.GetAlreadyPengCards(cardType)
		//fmt.Println("chiCards", chiCards, "cardType :", cardType)
		if pengCards != nil && pengCards.Len() > 0 {
			return false, qf.config
		}
	}

	//不能有杠的非字的牌
	for cardType := card.CardType_Wan; cardType < card.Max_CardType; cardType++{
		if card.CardType_Jian == cardType || card.CardType_Feng == cardType {
			continue
		}
		gangCards := cardsGetter.GetAlreadyGangCards(cardType)
		//fmt.Println("chiCards", chiCards, "cardType :", cardType)
		if gangCards != nil && gangCards.Len() > 0 {
			return false, qf.config
		}
	}

	return cardsInHand.IsHu(), qf.config
}
