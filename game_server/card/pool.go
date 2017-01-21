package card

import "mahjong/game_server/util"


type Option struct {
	generater  []PoolGenerater
}

func newOption() *Option {
	return &Option{
		generater:	make([]PoolGenerater, 0),
	}
}

type OptionFunc func(opt *Option)


type Pool struct {
	opt   *Option
	cards []*Card
}

func NewPool(opts ...OptionFunc) *Pool {
	pool := &Pool{
		opt:	newOption(),
		cards:	make([]*Card, 0),
	}
	for _, opt := range opts {
		opt(pool.opt)
	}

	pool.generate()
	return pool
}

func (pool *Pool) generate() {
	for _, generater := range pool.opt.generater {
		cards := generater.Generate()
		pool.cards = append(pool.cards, cards...)
	}
}

func (pool *Pool) GetCard() *Card {
	return pool.randomTakeWay()
}

func (pool *Pool) randomTakeWay() *Card {
	num := len(pool.cards)
	if num == 0 {
		return nil
	}
	idx := util.RandomN(num)
	card := pool.cards[idx]
	pool.cards = append(pool.cards[0:idx], pool.cards[idx+1:]...)
	return card
}
