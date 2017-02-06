package room

import (
	"testing"
	"mahjong/game_server/hu_checker"
	"github.com/bmizerany/assert"
	"mahjong/game_server/card"
)

func TestNewPlayerHuY9(t *testing.T) {
	factory := hu_checker.FactoryInst()
	err := factory.Init("../hu_checker/hu_config.json")
	assert.Equal(t, err, nil)

	player := NewPlayer(factory.GetAllChecker())
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:9})

	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	//player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})

	player.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	//player.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	player.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:1})
	//player.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:1})
	player.AddCard(&card.Card{CardType:card.CardType_Tong, CardNo:1})

	player.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	player.AddMagicCard(&card.Card{CardType:card.CardType_Tiao, CardNo:2})
	player.AddMagicCard(&card.Card{CardType:card.CardType_Tiao, CardNo:2})
	player.AddMagicCard(&card.Card{CardType:card.CardType_Tiao, CardNo:2})
	//player.AddMagicCard(&card.Card{CardType:card.CardType_Tiao, CardNo:2})


	isHu, desc, score := player.IsHu()
	t.Log(isHu, desc, score)
	//assert.Equal(t, score, 24)
	assert.Equal(t, desc, "幺九")
	assert.Equal(t, isHu, true)
}

func TestNewPlayerHuQ1SDD(t *testing.T) {
	factory := hu_checker.FactoryInst()
	err := factory.Init("../hu_checker/hu_config.json")
	assert.Equal(t, err, nil)

	player := NewPlayer(factory.GetAllChecker())
	//player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:1})
	//player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:1})
	//player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:1})

	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:2})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:2})

	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:3})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:3})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:3})

	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:4})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:4})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:4})

	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:5})
	player.AddCard(&card.Card{CardType:card.CardType_Tiao, CardNo:5})

	player.Peng(&card.Card{CardType:card.CardType_Tiao, CardNo:2})
	player.Gang(&card.Card{CardType:card.CardType_Tiao, CardNo:3})

	player.AddMagicCard(&card.Card{CardType:card.CardType_Tong, CardNo:2})
	player.AddMagicCard(&card.Card{CardType:card.CardType_Tong, CardNo:2})
	player.AddMagicCard(&card.Card{CardType:card.CardType_Tong, CardNo:2})
	//player.AddMagicCard(&card.Card{CardType:card.CardType_Tiao, CardNo:2})


	isHu, desc, score := player.IsHu()
	t.Log(isHu, desc, score)
	//assert.Equal(t, score, 26)
	assert.Equal(t, desc, "清一色对对胡")
	assert.Equal(t, isHu, true)
}

func TestNewPlayerHuD4X(t *testing.T) {
	factory := hu_checker.FactoryInst()
	err := factory.Init("../hu_checker/hu_config.json")
	assert.Equal(t, err, nil)

	player := NewPlayer(factory.GetAllChecker())
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})

	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})

	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})

	//player.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	//player.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	player.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	player.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	player.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.Peng(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	isHu, desc, score := player.IsHu()
	t.Log(isHu, desc, score)
	//assert.Equal(t, score, 26)
	assert.Equal(t, desc, "大四喜")
	assert.Equal(t, isHu, true)
}

func TestNewPlayerHuZ1S(t *testing.T) {
	factory := hu_checker.FactoryInst()
	err := factory.Init("../hu_checker/hu_config.json")
	assert.Equal(t, err, nil)

	player := NewPlayer(factory.GetAllChecker())
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Dong})

	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Nan})

	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})

	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})
	player.AddCard(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Bei})

	//player.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	//player.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})

	player.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})
	player.AddCard(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	player.Gang(&card.Card{CardType:card.CardType_Feng, CardNo:card.Feng_CardNo_Xi})
	player.Peng(&card.Card{CardType:card.CardType_Jian, CardNo:card.Jian_CardNo_Zhong})

	isHu, desc, score := player.IsHu()
	t.Log(isHu, desc, score)
	//assert.Equal(t, score, 26)
	assert.Equal(t, desc, "字一色")
	assert.Equal(t, isHu, true)
}