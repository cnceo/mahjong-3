package card

import "testing"

func TestGenerater(t *testing.T) {
	t.Log("风牌")
	allCards := NewCards()
	cards := (&FengGenerater{}).Generate()
	for _, card := range cards.Data() {
		t.Log(card.Name())
	}
	allCards.AppendCards(cards)

	t.Log("箭牌")
	cards = (&JianGenerater{}).Generate()
	for _, card := range cards.Data() {
		t.Log(card.Name())
	}
	allCards.AppendCards(cards)

	t.Log("花牌")
	cards = (&HuaGenerater{}).Generate()
	for _, card := range cards.Data() {
		t.Log(card.Name())
	}
	allCards.AppendCards(cards)

	t.Log("万子牌")
	cards = (&WanGenerater{}).Generate()
	for _, card := range cards.Data() {
		t.Log(card.Name())
	}
	allCards.AppendCards(cards)

	t.Log("条子牌")
	cards = (&TiaoGenerater{}).Generate()
	for _, card := range cards.Data() {
		t.Log(card.Name())
	}
	allCards.AppendCards(cards)

	t.Log("筒牌")
	cards = (&TongGenerater{}).Generate()
	for _, card := range cards.Data() {
		t.Log(card.Name())
	}
	allCards.AppendCards(cards)

	t.Log(allCards.Len())
	for _, card := range allCards.Data() {
		t.Log(card.Name())
	}
}
