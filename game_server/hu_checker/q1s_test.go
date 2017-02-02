package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestQ1S_IsHu(t *testing.T) {
	cards := card.NewPlayingCards()
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:2})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:2})

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:4})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:4})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:4})

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:5})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:5})

	cards.Peng(&card.Card{CardType:card.CardType_Wan, CardNo:2})
	cards.Gang(&card.Card{CardType:card.CardType_Wan, CardNo:3})

	hu := NewQ1S(&HuConfig{Name:"Q1S_HU", IsEnabled:true, Score:2, Desc:"清一色"})
	isHu, conf := hu.IsHu(cards)
	assert.Equal(t, conf.Name, "Q1S_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "清一色")
	assert.Equal(t, isHu, true)

}
