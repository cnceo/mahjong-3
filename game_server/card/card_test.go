package card

import "testing"

func TestCard(t *testing.T) {
}

func TestSort(t *testing.T) {
	pool := NewPool(
		WithFengCard(),
		WithJianCard(),
		WithHuaCard(),
		WithWanCard(),
		WithTiaoCard(),
		WithTongCard(),
	)

	cards := NewCards()
	for i:=0; i<13; i++ {
		cards.AppendCard(pool.GetCard())
	}
	t.Log("before sort :")
	t.Log(cards.ToString(), cards.Len())
	Sort(cards)
	t.Log("after sort :")
	t.Log(cards.ToString(), cards.Len())

	t.Log("after random take way one card")
	card := cards.RandomTakeWayOne()
	t.Log(cards.ToString(), cards.Len(), card.Name())

	cards = NewCards()
	cards.AppendCard(&Card{})
	cards.RandomTakeWayOne()
	t.Log("after random takeway one from only one card :")
	t.Log(cards.ToString(), cards.Len())

	cards = NewCards()
	cards.AppendCard(&Card{CardType:CardType_Wan,CardNo:1})
	cards.TakeWay(&Card{CardType:CardType_Wan,CardNo:1})
	t.Log("after takeway from only one card :")
	t.Log(cards.ToString(), cards.Len())
}