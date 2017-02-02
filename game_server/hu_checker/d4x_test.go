package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestD4X_IsHu(t *testing.T) {
	cards := card.NewPlayingCards()
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	cards.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	d4x := NewD4X(&HuConfig{Name:"D4X_HU", IsEnabled:true, Score:2, Desc:"大四喜"})
	isHu, conf := d4x.IsHu(cards)
	assert.Equal(t, conf.Name, "D4X_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "大四喜")
	assert.Equal(t, isHu, true)

}
