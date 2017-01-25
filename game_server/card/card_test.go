package card

import (
	"testing"
	"github.com/bmizerany/assert"
)

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
//	if !card1.IsZiCard() || !card2.IsZiCard() || !card3.IsZiCard() {
//		t.Fatal("it should be zi card")
//	}
	assert.Equal(t, card1.IsZiCard(), true)
	assert.Equal(t, card2.IsZiCard(), true)
	assert.Equal(t, !card3.IsZiCard(), true)
}

func TestCard_IsOk(t *testing.T) {
	assert.Equal(t, (&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong}).IsOk() ,true)
	assert.Equal(t, (&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Nan}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Xi}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Bei}).IsOk(),true)

	assert.Equal(t, (&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Zhong}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Fa}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Bai}).IsOk(),true)

	assert.Equal(t, (&Card{CardType:CardType_Hua, CardNo:Hua_CardNo_Ju}).IsOk(),true)

	assert.Equal(t, (&Card{CardType:CardType_Wan, CardNo:1}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Wan, CardNo:2}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Wan, CardNo:3}).IsOk(),true)

	assert.Equal(t, (&Card{CardType:CardType_Tiao, CardNo:1}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Tiao, CardNo:2}).IsOk(),true)
	assert.Equal(t, (&Card{CardType:CardType_Tiao, CardNo:3}).IsOk(),true)

	assert.Equal(t, (&Card{CardType:CardType_Wan, CardNo:10}).IsOk(), false)
	assert.Equal(t, (&Card{CardType:CardType_Tiao, CardNo:20}).IsOk(), false)
	assert.Equal(t, (&Card{CardType:CardType_Tong, CardNo:30}).IsOk(), false)
}