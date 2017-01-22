package card

import (
	"testing"
//	"time"
)

func TestPool(t *testing.T) {
	pool := NewPool(
		WithFengCard(),
		WithJianCard(),
		WithHuaCard(),
		WithWanCard(),
		WithTiaoCard(),
		WithTongCard(),
	)

	/*
	for _, card := range pool.cards{
		t.Log(card.Name())
	}
	*/
	//t.Log(pool.GetCardNum())
	for{
		card := pool.GetCard()
		if card == nil {
			break
		}
		//t.Log(card.Name(), pool.GetCardNum())
		//time.Sleep(time.Second)
	}
	if pool.GetCardNum() != 0 {
		t.Fatal("card num of pool should be 0")
	}
}
