package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestY9_IsHu(t *testing.T) {
	cards := card.NewPlayingCards()
	cards.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})
	cards.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})
	cards.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})

	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	cards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:1})

	cards.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	y9 := NewY9(&HuConfig{Name:"Y9_HU", IsEnabled:true, Score:2, Desc:"幺九"})
	isHu, conf := y9.IsHu(cards)
	assert.Equal(t, conf.Name, "Y9_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "幺九")
	assert.Equal(t, isHu, true)
}

func TestY9_NotIsHu(t *testing.T) {
	cards := card.NewPlayingCards()
	cards.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})
	cards.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})
	cards.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})

	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	cards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:1})
	cards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:2})
	cards.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:3})

	cards.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	y9 := NewY9(&HuConfig{Name:"Y9_HU", IsEnabled:true, Score:2, Desc:"幺九"})
	isHu, conf := y9.IsHu(cards)
	assert.Equal(t, conf.Name, "Y9_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "幺九")
	assert.Equal(t, isHu, false)
}
