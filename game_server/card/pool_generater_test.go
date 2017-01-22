package card

import "testing"

func TestGenerater(t *testing.T) {
	//t.Log("风牌")
	allCards := NewCards()
	cards := (&FengGenerater{}).Generate()
	for _, card := range cards.Data() {
		//t.Log(card.Name())
		if card.CardType != CardType_Feng {
			t.Fatal("card type should be feng")
		}
	}
	if cards.Len() != 16 {
		t.Fatal("card num should be 16")
	}
	allCards.AppendCards(cards)

	//t.Log("箭牌")
	cards = (&JianGenerater{}).Generate()
	for _, card := range cards.Data() {
		//t.Log(card.Name())
		if card.CardType != CardType_Jian {
			t.Fatal("card type should be jian")
		}
	}
	if cards.Len() != 12 {
		t.Fatal("card num should be 16")
	}
	allCards.AppendCards(cards)

	//t.Log("花牌")
	cards = (&HuaGenerater{}).Generate()
	for _, card := range cards.Data() {
		//t.Log(card.Name())
		if card.CardType != CardType_Hua {
			t.Fatal("card type should be hua")
		}
	}
	if cards.Len() != 8 {
		t.Fatal("card num should be 8")
	}
	allCards.AppendCards(cards)

	//t.Log("万子牌")
	cards = (&WanGenerater{}).Generate()
	for _, card := range cards.Data() {
		//t.Log(card.Name())
		if card.CardType != CardType_Wan {
			t.Fatal("card type should be wan")
		}
	}
	if cards.Len() != 36 {
		t.Fatal("card num should be 36")
	}
	allCards.AppendCards(cards)

	//t.Log("条子牌")
	cards = (&TiaoGenerater{}).Generate()
	for _, card := range cards.Data() {
		//t.Log(card.Name())
		if card.CardType != CardType_Tiao {
			t.Fatal("card type should be wan")
		}
	}
	if cards.Len() != 36 {
		t.Fatal("card num should be 36")
	}
	allCards.AppendCards(cards)

	//t.Log("筒牌")
	cards = (&TongGenerater{}).Generate()
	for _, card := range cards.Data() {
		//t.Log(card.Name())
		if card.CardType != CardType_Tong {
			t.Fatal("card type should be wan")
		}
	}
	if cards.Len() != 36 {
		t.Fatal("card num should be 36")
	}
	allCards.AppendCards(cards)

	if allCards.Len() != 144 {
		t.Fatal("card num should be 144")
	}
	/*
	t.Log(allCards.Len())
	for _, card := range allCards.Data() {
		t.Log(card.Name())
	}
	*/
}
