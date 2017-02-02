package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestDD_IsHu(t *testing.T) {
	cards := card.NewPlayingCards()
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:2})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:2})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:2})

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})


	ddHu := NewDD(&HuConfig{Name:"DD_HU", IsEnabled:true, Score:2, Desc:"对对胡"})
	isHu, conf := ddHu.IsHu(cards)
	assert.Equal(t, conf.Name, "DD_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "对对胡")
	assert.Equal(t, isHu, true)

}
