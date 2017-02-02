package play

import (
	"testing"
	"github.com/bmizerany/assert"
	"mahjong/game_server/card"
)

func TestPlayingCards_Chi(t *testing.T) {
	playingCards := NewPlayingCards()
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:5})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:7})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:9})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:2})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:4})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:6})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:8})

	playingCards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:6})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:8})

	group := card.NewCards()
	group.AppendCard(&card.Card{CardType:card.CardType_Wan, CardNo:4})
	group.AppendCard(&card.Card{CardType:card.CardType_Wan, CardNo:5})
	group.AppendCard(&card.Card{CardType:card.CardType_Wan, CardNo:6})
	chi := playingCards.Chi(&card.Card{CardType:card.CardType_Wan, CardNo:5}, group)

	assert.Equal(t, chi, true)
	assert.Equal(t, playingCards.cardsInHand.Len(), 9)
	assert.Equal(t, playingCards.cardsAlreadyChi[card.CardType_Wan].SameAs(group), true)

	groupTong := card.NewCards()
	groupTong.AppendCard(&card.Card{CardType:card.CardType_Tong, CardNo:6})
	groupTong.AppendCard(&card.Card{CardType:card.CardType_Tong, CardNo:7})
	groupTong.AppendCard(&card.Card{CardType:card.CardType_Tong, CardNo:8})
	chiTongErr := playingCards.Chi(&card.Card{CardType:card.CardType_Tong, CardNo:8}, groupTong)

	assert.Equal(t, chiTongErr, false)
	assert.Equal(t, playingCards.cardsInHand.Len(), 9)

	chiTongOk := playingCards.Chi(&card.Card{CardType:card.CardType_Tong, CardNo:7}, groupTong)
	assert.Equal(t, chiTongOk, true)
	assert.Equal(t, playingCards.cardsInHand.Len(), 7)
	assert.Equal(t, playingCards.cardsAlreadyChi[card.CardType_Tong].SameAs(groupTong), true)

//	t.Log(playingCards.ToString())
}


func TestPlayingCards_Peng(t *testing.T) {
	playingCards := NewPlayingCards()
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	playingCards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:6})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:6})

	pengWan := playingCards.Peng(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	assert.Equal(t, pengWan, true)
	assert.Equal(t, playingCards.cardsInHand.Len(), 3)
	assert.Equal(t, playingCards.cardsAlreadyPeng[card.CardType_Wan].Len(), 3)

	pengTong := playingCards.Peng(&card.Card{CardType:card.CardType_Tong, CardNo:6})
	assert.Equal(t, pengTong, true)
	assert.Equal(t, playingCards.cardsInHand.Len(), 1)
	assert.Equal(t, playingCards.cardsAlreadyPeng[card.CardType_Tong].Len(), 3)

	pengJian := playingCards.Peng(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Bai})

	assert.Equal(t, pengJian, false)
	//t.Log(playingCards.ToString())
}


func TestPlayingCards_Gang(t *testing.T) {
	playingCards := NewPlayingCards()
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	playingCards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:6})
	playingCards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:6})

	gangWan := playingCards.Gang(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	assert.Equal(t, gangWan, true)
	assert.Equal(t, playingCards.cardsInHand.Len(), 2)
	assert.Equal(t, playingCards.cardsAlreadyGang[card.CardType_Wan].Len(), 4)

	gangTong := playingCards.Gang(&card.Card{CardType:card.CardType_Tong, CardNo:6})
	assert.Equal(t, gangTong, false)
	assert.Equal(t, playingCards.cardsInHand.Len(), 2)
	assert.Equal(t, playingCards.cardsAlreadyGang[card.CardType_Tong].Len(), 0)

	gangJian := playingCards.Peng(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Bai})

	assert.Equal(t, gangJian, false)
	//t.Log(playingCards.ToString())
}