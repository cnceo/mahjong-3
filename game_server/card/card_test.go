package card

import "testing"

func TestCard(t *testing.T) {
	card1 := &Card{
		CardType: CardType_Wan,
		CardNo: 1,
	}
	card2 := &Card{
		CardType: CardType_Tiao,
		CardNo: 1,
	}
	card3 := &Card{
		CardType: CardType_Tong,
		CardNo: 1,
	}
	if card1.IsZiCard() || card2.IsZiCard() || card3.IsZiCard() {
		t.Fatal("it should be not zi card")
	}

	card1 = &Card{
		CardType: CardType_Feng,
		CardNo: Feng_CardNo_Bei,
	}
	card2 = &Card{
		CardType: CardType_Jian,
		CardNo: Jian_CardNo_Bai,
	}
	card3 = &Card{
		CardType: CardType_Hua,
		CardNo: Hua_CardNo_Chun,
	}
	if !card1.IsZiCard() || !card2.IsZiCard() || !card3.IsZiCard() {
		t.Fatal("it should be zi card")
	}
}
