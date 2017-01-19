package card

import "mahjong/game_server/util"

type Option struct {
	
}

type OptionFunc func(opt *Option)


type Pool struct {
	opt			Option
	cards		[]*Card
}

func NewPool(opts ...OptionFunc) *Pool {
	pool :=  &Pool{
	}
	for _, opt := range opts {
		opt(&pool.opt)
	}
	return pool
}

func (pool *Pool) Init() {
	
}

func (pool *Pool) GetCard() *Card {
	return pool.randomTakeWay()
}

func (pool *Pool) randomTakeWay() *Card{
	num := len(pool.cards)
	if num == 0 {
		return nil
	}
	idx := util.RandomN(num)
	card := pool.cards[idx]
	pool.cards = append(pool.cards[0:idx], pool.cards[idx+1:]...)
	return card
}