package card

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
}

func (pool *Pool) GetCard() *Card {
	return pool.cards.RandomTakeWayOne()
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