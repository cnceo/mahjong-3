package card

var OneMagicCandidate 	[]*Cards
var TwoMagicCandidate 	[]*Cards
var ThreeMagicCandidate []*Cards
var FourMagicCandidate 	[]*Cards

func init() {
	allCards := getAllCards()
	diffCardNum := allCards.Len()
	//init OneMagicCandidate
	OneMagicCandidate = make([]*Cards, 0)
	for i := 0; i < diffCardNum; i++ {
		cards := NewCards()
		cards.AppendCard(allCards.At(i))
		OneMagicCandidate = append(OneMagicCandidate, cards)
	}

	//init TwoMagicCandidate
	TwoMagicCandidate = make([]*Cards, 0)
	for i:=0; i<diffCardNum; i++ {
		for j:=i; j<diffCardNum; j++ {
			cards := NewCards()
			cards.AppendCard(allCards.At(i))
			cards.AppendCard(allCards.At(j))
			TwoMagicCandidate = append(TwoMagicCandidate, cards)
		}
	}

	//init ThreeMagicCandidate
	ThreeMagicCandidate = make([]*Cards, 0)
	for i:=0; i<diffCardNum; i++ {
		for j:=i; j<diffCardNum; j++ {
			for k:=j; k<diffCardNum; k++ {
				cards := NewCards()
				cards.AppendCard(allCards.At(i))
				cards.AppendCard(allCards.At(j))
				cards.AppendCard(allCards.At(k))
				ThreeMagicCandidate = append(ThreeMagicCandidate, cards)
			}
		}
	}

	//init FourMagicCandidate
	FourMagicCandidate = make([]*Cards, 0)
	for i:=0; i<diffCardNum; i++ {
		for j:=i; j<diffCardNum; j++ {
			for k:=j; k<diffCardNum; k++ {
				for l:=k; l<diffCardNum; l++ {
					cards := NewCards()
					cards.AppendCard(allCards.At(i))
					cards.AppendCard(allCards.At(j))
					cards.AppendCard(allCards.At(k))
					cards.AppendCard(allCards.At(l))
					FourMagicCandidate = append(FourMagicCandidate, cards)
				}
			}
		}
	}

}

func getAllCards() *Cards {
	cards := NewCards()
	for cardType := CardType_Wan; cardType < Max_CardType; cardType++ {
		cards.AppendCards(genAllCards(cardType))
	}
	return cards
}

func genAllCards(cardType int) *Cards {
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