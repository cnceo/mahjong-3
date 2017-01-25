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
	assert.Equal(t, playingCards.cardsInHand.Len(), 9)
	assert.Equal(t, playingCards.cardsAlreadyChi[CardType_Wan].SameAs(group), true)

	groupTong := NewCards()
	groupTong.AppendCard(&Card{CardType:CardType_Tong, CardNo:6})
	groupTong.AppendCard(&Card{CardType:CardType_Tong, CardNo:7})
	groupTong.AppendCard(&Card{CardType:CardType_Tong, CardNo:8})
	chiTongErr := playingCards.Chi(&Card{CardType:CardType_Tong, CardNo:8}, groupTong)

	assert.Equal(t, chiTongErr, false)
	assert.Equal(t, playingCards.cardsInHand.Len(), 9)

	chiTongOk := playingCards.Chi(&Card{CardType:CardType_Tong, CardNo:7}, groupTong)
	assert.Equal(t, chiTongOk, true)
	assert.Equal(t, playingCards.cardsInHand.Len(), 7)
	assert.Equal(t, playingCards.cardsAlreadyChi[CardType_Tong].SameAs(groupTong), true)

	t.Log(playingCards.ToString())
}

