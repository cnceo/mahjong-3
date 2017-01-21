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
	for{
		card := pool.GetCard()
		if card == nil {
			break
		}
		t.Log(card.Name())
		//time.Sleep(time.Second)
	}
}
