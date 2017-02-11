package card

import (
	"mahjong/game_server/util"
	//"mahjong/game_server/log"
)

type Pool struct {
	cards *Cards
	generaters  []PoolGenerater
}

func NewPool() *Pool {
	pool := &Pool{
		cards:	NewCards(),
		generaters:	make([]PoolGenerater, 0),
	}
	return pool
}

func (pool *Pool) generate() {
	for _, generater := range pool.generaters {
		pool.cards.AppendCards(generater.Generate())
	}
}

func (pool *Pool) ReGenerate() {
	pool.cards.Clear()
	pool.generate()
	pool.shuffle()
}

//洗牌，打乱牌
func (pool *Pool) shuffle() {
	length := pool.cards.Len()
	for cnt := 0; cnt<length; cnt++ {
		i := util.RandomN(length)
		j := util.RandomN(length)
		pool.cards.Swap(i, j)
		//log.Debug("poll shuffle swap[", i, "=>", j, "]")
	}
}

func (pool *Pool) PopFront() *Card {
	return pool.cards.PopFront()
}

func (pool *Pool) PopTail() *Card{
	return pool.cards.PopTail()
}

func (pool *Pool) At(idx int) *Card {
	return pool.cards.At(idx)
}

func (pool *Pool) GetCardNum() int {
	return pool.cards.Len()
}

func (pool *Pool) AddFengGenerater() {
	pool.generaters = append(pool.generaters, &FengGenerater{})
}

func (pool *Pool) AddJianGenerater() {
	pool.generaters = append(pool.generaters, &JianGenerater{})
}

func (pool *Pool) AddHuaGenerater() {
	pool.generaters = append(pool.generaters, &HuaGenerater{})
}

func (pool *Pool) AddWanGenerater() {
	pool.generaters = append(pool.generaters, &WanGenerater{})
}

func (pool *Pool) AddTiaoGenerater() {
	pool.generaters = append(pool.generaters, &TiaoGenerater{})
}

func (pool *Pool) AddTongGenerater() {
	pool.generaters = append(pool.generaters, &TongGenerater{})
}