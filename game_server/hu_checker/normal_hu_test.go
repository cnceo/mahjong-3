package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestNormalHu_IsHu(t *testing.T) {
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

	hu := NewQ1SDD(&HuConfig{Name:"Normal_HU", IsEnabled:true, Score:2, Desc:"平胡"})
	isHu, conf := hu.IsHu(cards)
	assert.Equal(t, conf.Name, "Normal_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "平胡")
	assert.Equal(t, isHu, true)

}
