package hu_checker

import (
	"testing"
	"mahjong/game_server/card"
	"github.com/bmizerany/assert"
)

func TestQF_NotIsHu(t *testing.T) {
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

	//cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	//cards.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	cards.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	qf := NewQF(&HuConfig{Name:"QF_HU", IsEnabled:true, Score:2, Desc:"全番"})
	isHu, conf := qf.IsHu(cards)
	assert.Equal(t, conf.Name, "QF_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "全番")
	assert.Equal(t, isHu, true)

}

func TestQF_IsHu(t *testing.T) {
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

	//cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	//cards.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	cards.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	cards.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	qf := NewQF(&HuConfig{Name:"QF_HU", IsEnabled:true, Score:2, Desc:"全番"})
	isHu, conf := qf.IsHu(cards)
	assert.Equal(t, conf.Name, "QF_HU")
	assert.Equal(t, conf.IsEnabled, true)
	assert.Equal(t, conf.Score, 2)
	assert.Equal(t, conf.Desc, "全番")
	assert.Equal(t, isHu, false)

}