package card

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
	cards *Cards
}

func NewPool(opts ...OptionFunc) *Pool {
	pool := &Pool{
		opt:	newOption(),
		cards:	NewCards(),
	}
	for _, opt := range opts {
		opt(pool.opt)
	}

	pool.generate()
	return pool
}

func (pool *Pool) generate() {
	for _, generater := range pool.opt.generater {
		pool.cards.AppendCards(generater.Generate())
	}
}

func (pool *Pool) GetCard() *Card {
	return pool.cards.RandomTakeWayOne()
}

func (pool *Pool) GetCardNum() int {
	return pool.cards.Len()
}
