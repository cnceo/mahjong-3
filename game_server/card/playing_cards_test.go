package card

import (
	"testing"
	"github.com/bmizerany/assert"
)

func TestPlayingCards_Chi(t *testing.T) {
	playingCards := NewPlayingCards()
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:3})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:5})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:7})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:9})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:2})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:4})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:6})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:8})

	playingCards.AddCard(&Card{CardType:CardType_Tong, CardNo:6})
	playingCards.AddCard(&Card{CardType:CardType_Tong, CardNo:8})

	group := NewCards()
	group.AppendCard(&Card{CardType:CardType_Wan, CardNo:4})
	group.AppendCard(&Card{CardType:CardType_Wan, CardNo:5})
	group.AppendCard(&Card{CardType:CardType_Wan, CardNo:6})
	chi := playingCards.Chi(&Card{CardType:CardType_Wan, CardNo:5}, group)

	assert.Equal(t, chi, true)
	assert.Equal(t, playingCards.cardsInHand[CardType_Wan].Len(), 7)
	assert.Equal(t, playingCards.cardsAlreadyChi[CardType_Wan].SameAs(group), true)

	groupTong := NewCards()
	groupTong.AppendCard(&Card{CardType:CardType_Tong, CardNo:6})
	groupTong.AppendCard(&Card{CardType:CardType_Tong, CardNo:7})
	groupTong.AppendCard(&Card{CardType:CardType_Tong, CardNo:8})
	chiTongErr := playingCards.Chi(&Card{CardType:CardType_Tong, CardNo:8}, groupTong)

	assert.Equal(t, chiTongErr, false)
	assert.Equal(t, playingCards.cardsInHand[CardType_Tong].Len(), 2)

	chiTongOk := playingCards.Chi(&Card{CardType:CardType_Tong, CardNo:7}, groupTong)
	assert.Equal(t, chiTongOk, true)
	assert.Equal(t, playingCards.cardsInHand[CardType_Tong].Len(), 0)
	assert.Equal(t, playingCards.cardsAlreadyChi[CardType_Tong].SameAs(groupTong), true)

//	t.Log(playingCards.ToString())
}


func TestPlayingCards_Peng(t *testing.T) {
	playingCards := NewPlayingCards()
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})

	playingCards.AddCard(&Card{CardType:CardType_Tong, CardNo:6})
	playingCards.AddCard(&Card{CardType:CardType_Tong, CardNo:6})

	pengWan := playingCards.Peng(&Card{CardType:CardType_Wan, CardNo:1})
	assert.Equal(t, pengWan, true)
	assert.Equal(t, playingCards.cardsInHand[CardType_Wan].Len(), 1)
	assert.Equal(t, playingCards.cardsAlreadyPeng[CardType_Wan].Len(), 3)

	pengTong := playingCards.Peng(&Card{CardType:CardType_Tong, CardNo:6})
	assert.Equal(t, pengTong, true)
	assert.Equal(t, playingCards.cardsInHand[CardType_Tong].Len(), 0)
	assert.Equal(t, playingCards.cardsAlreadyPeng[CardType_Tong].Len(), 3)

	pengJian := playingCards.Peng(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Bai})

	assert.Equal(t, pengJian, false)
	//t.Log(playingCards.ToString())
}


func TestPlayingCards_Gang(t *testing.T) {
	playingCards := NewPlayingCards()
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})

	playingCards.AddCard(&Card{CardType:CardType_Tong, CardNo:6})
	playingCards.AddCard(&Card{CardType:CardType_Tong, CardNo:6})

	gangWan := playingCards.Gang(&Card{CardType:CardType_Wan, CardNo:1})
	assert.Equal(t, gangWan, true)
	assert.Equal(t, playingCards.cardsInHand[CardType_Wan].Len(), 0)
	assert.Equal(t, playingCards.cardsAlreadyGang[CardType_Wan].Len(), 4)

	gangTong := playingCards.Gang(&Card{CardType:CardType_Tong, CardNo:6})
	assert.Equal(t, gangTong, false)
	assert.Equal(t, playingCards.cardsInHand[CardType_Tong].Len(), 2)
	assert.Equal(t, playingCards.cardsAlreadyGang[CardType_Tong].Len(), 0)

	gangJian := playingCards.Peng(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Bai})

	assert.Equal(t, gangJian, false)
	//t.Log(playingCards.ToString())
}

func TestPlayingCards_Reset(t *testing.T) {
	playingCards := NewPlayingCards()
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})
	playingCards.AddCard(&Card{CardType:CardType_Wan, CardNo:1})

	t.Log(playingCards.GetInHandCards(CardType_Wan).ToString())
	assert.Equal(t, playingCards.GetInHandCards(CardType_Wan).Len(), 3)
	playingCards.Reset()
	assert.Equal(t, playingCards.GetInHandCards(CardType_Wan).Len(), 0)
	t.Log(playingCards.GetInHandCards(CardType_Wan).ToString())
}