package card

var allCards []*Cards

func GetAllCards(cardType int) *Cards {
	if cardType >= Max_CardType || cardType < 0 {
		return nil
	}
	if allCards == nil {
		allCards = make([]*Cards, Max_CardType)
	}

	if allCards[cardType] == nil {
		allCards[cardType] = initAllCards(cardType)
	}
	return allCards[cardType]
}

func initAllCards(cardType int) *Cards {
	switch cardType {
	case CardType_Feng:
		return generateCards(cardType, Feng_CardNo_Dong, Feng_CardNo_Bei, 1)
	case CardType_Jian:
		return generateCards(cardType, Jian_CardNo_Zhong, Jian_CardNo_Bai, 1)
	case CardType_Wan, CardType_Tiao, CardType_Tong:
		return generateCards(cardType, 1, 9, 1)
	}
	return nil
}