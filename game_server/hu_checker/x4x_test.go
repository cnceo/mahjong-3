package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestX4X_IsHu(t *testing.T) {
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

	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:2})
	cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:3})

	cards.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	x4x := NewX4X(&HuConfig{Name:"X4X_HU", IsEnabled:true, Score:2, Desc:"小四喜"})
	isHu, conf := x4x.IsHu(cards)
	assert.Equal(t, conf.Name, "X4X_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "小四喜")
	assert.Equal(t, isHu, true)

}
