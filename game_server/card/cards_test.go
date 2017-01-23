package card

import "testing"

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
	cards.Sort()
	t.Log("after sort :")
	t.Log(cards.ToString(), cards.Len())

	t.Log("after random take way one card")
	card := cards.RandomTakeWayOne()
	t.Log(cards.ToString(), cards.Len(), card.Name())

	oneCards := NewCards()
	oneCards.AppendCard(&Card{})
	oneCards.RandomTakeWayOne()
	t.Log("after random takeway one from only one card :")
	t.Log(oneCards.ToString(), oneCards.Len())


	oneCards = NewCards()
	oneCards.AppendCard(&Card{CardType:CardType_Wan,CardNo:1})
	oneCards.TakeWay(&Card{CardType:CardType_Wan,CardNo:1})
	t.Log("after takeway from only one card :")
	t.Log(oneCards.ToString(), oneCards.Len())

	left, three := cards.splitLeftOtherAndThree()
	t.Log("left :", left.ToString())
	t.Log("three :", three.ToString())

	var right *Cards
	three, right = cards.splitThreeAndRightOther()
	t.Log("three :", three.ToString())
	t.Log("right :", right.ToString())
}

func TestCards_Is5Card(t *testing.T) {
	cards := NewCards()
	for i:=1; i<=3; i++ {
		card := &Card{
			CardType: CardType_Wan,
			CardNo: i,
		}
		cards.AppendCard(card)
	}
	if cards.Is5Card() {
		t.Fatal("it should not be 5 card")
	}
	cards.AppendCard(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Bai})
	cards.AppendCard(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Zhong})
	if cards.Is5Card() {
		t.Fatal("it should not be 5 card")
	}

	cards.TakeWay(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Bai})
	cards.AppendCard(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Zhong})
	if !cards.Is5Card() {
		t.Fatal("it should be 5 card")
	}

	cards.Clear()
	for i:=1; i<4; i++ {
		card := &Card{
			CardType: CardType_Feng,
			CardNo: Feng_CardNo_Dong,
		}
		cards.AppendCard(card)
	}
	if cards.Is5Card() {
		t.Fatal("it should not be 5 card")
	}
	cards.AppendCard(&Card{CardType:CardType_Wan, CardNo:4})
	if cards.Is5Card() {
		t.Fatal("it should not be 5 card")
	}
	cards.AppendCard(&Card{CardType:CardType_Wan, CardNo:4})
	if !cards.Is5Card() {
		t.Fatal("it should be 5 card")
	}

	cards.TakeWay(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong})
	cards.AppendCard(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Nan})
	if cards.Is5Card() {
		t.Fatal("it should not be 5 card")
	}
}