package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestD3Y_IsHu(t *testing.T) {
	cards := card.NewPlayingCards()
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Fa})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Fa})

	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Bai})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Bai})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Bai})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})

	cards.Gang(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Bai})
	cards.Peng(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Fa})

	d3y := NewD3Y(&HuConfig{Name:"D3Y_HU", IsEnabled:true, Score:2, Desc:"大三元"})
	isHu, conf := d3y.IsHu(cards)
	assert.Equal(t, conf.Name, "D3Y_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "大三元")
	assert.Equal(t, isHu, true)

}
