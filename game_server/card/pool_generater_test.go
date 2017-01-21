package card

import "testing"

func TestGenerater(t *testing.T) {
	t.Log("风牌")
	allCards := make([]*Card, 0)
	cards := (&FengGenerater{}).Generate()
	for _, card := range cards {
		t.Log(card.Name())
	}
	allCards = append(allCards, cards...)

	t.Log("箭牌")
	cards = (&JianGenerater{}).Generate()
	for _, card := range cards {
		t.Log(card.Name())
	}
	allCards = append(allCards, cards...)

	t.Log("花牌")
	cards = (&HuaGenerater{}).Generate()
	for _, card := range cards {
		t.Log(card.Name())
	}
	allCards = append(allCards, cards...)

	t.Log("万子牌")
	cards = (&WanGenerater{}).Generate()
	for _, card := range cards {
		t.Log(card.Name())
	}
	allCards = append(allCards, cards...)

	t.Log("条子牌")
	cards = (&TiaoGenerater{}).Generate()
	for _, card := range cards {
		t.Log(card.Name())
	}
	allCards = append(allCards, cards...)

	t.Log("筒牌")
	cards = (&TongGenerater{}).Generate()
	for _, card := range cards {
		t.Log(card.Name())
	}
	allCards = append(allCards, cards...)

	t.Log(len(allCards))
	for _, card := range allCards {
		t.Log(card.Name())
	}
}
