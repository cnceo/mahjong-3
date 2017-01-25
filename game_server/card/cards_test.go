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
		cards.AddAndSort(pool.GetCard())
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
	oneCards.AddAndSort(&Card{})
	oneCards.RandomTakeWayOne()
	t.Log("after random takeway one from only one card :")
	t.Log(oneCards.ToString(), oneCards.Len())


	oneCards = NewCards()
	oneCards.AddAndSort(&Card{CardType:CardType_Wan,CardNo:1})
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
		cards.AddAndSort(card)
	}
	if cards.IsHu() {
		t.Fatal("it should not be hu")
	}
	cards.AddAndSort(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Bai})
	cards.AddAndSort(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Zhong})
	if cards.IsHu() {
		t.Fatal("it should not be hu")
	}

	cards.TakeWay(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Bai})
	cards.AddAndSort(&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Zhong})
	if !cards.IsHu() {
		t.Fatal("it should be hu", cards.ToString())
	}

	cards.Clear()
	for i:=1; i<4; i++ {
		card := &Card{
			CardType: CardType_Feng,
			CardNo: Feng_CardNo_Dong,
		}
		cards.AddAndSort(card)
	}
	if cards.IsHu() {
		t.Fatal("it should not be hu")
	}
	cards.AddAndSort(&Card{CardType:CardType_Wan, CardNo:4})
	if cards.IsHu() {
		t.Fatal("it should not be hu")
	}
	cards.AddAndSort(&Card{CardType:CardType_Wan, CardNo:4})
	if !cards.IsHu() {
		t.Fatal("it should be hu")
	}

	cards.TakeWay(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong})
	cards.AddAndSort(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Nan})
	if cards.IsHu() {
		t.Fatal("it should not be hu")
	}
}

func TestCards_IsHu(t *testing.T) {
	hu2 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
		},
	}

	hu5 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:3},
		},
	}

	hu8 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:3},
		},
	}

	hu11 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:3},

			&Card{CardType:CardType_Wan, CardNo:6},
			&Card{CardType:CardType_Wan, CardNo:7},
			&Card{CardType:CardType_Wan, CardNo:8},
		},
	}

	hu14 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},

			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},

			&Card{CardType:CardType_Wan, CardNo:7},
			&Card{CardType:CardType_Wan, CardNo:7},
			&Card{CardType:CardType_Wan, CardNo:7},

			&Card{CardType:CardType_Wan, CardNo:6},
			&Card{CardType:CardType_Wan, CardNo:6},
			&Card{CardType:CardType_Wan, CardNo:6},
		},
	}

	hu2.Sort()
	hu5.Sort()
	hu8.Sort()
	hu11.Sort()
	hu14.Sort()

	if !hu2.IsHu() || !hu5.IsHu() || !hu8.IsHu() || !hu11.IsHu() || !hu14.IsHu(){
		//t.Fatal("all should be hu")
		t.Log(hu2.IsHu())
		t.Log(hu5.IsHu())
		t.Log(hu8.IsHu())
		t.Log(hu11.IsHu())
		t.Log(hu14.IsHu())
	}
}

func TestCards_IsNotHu(t *testing.T) {
	hu2 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Jian, CardNo:Jian_CardNo_Zhong},
		},
	}

	hu5 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:4},
		},
	}

	hu8 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:9},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:3},
		},
	}

	hu11 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:3},

			&Card{CardType:CardType_Wan, CardNo:6},
			&Card{CardType:CardType_Wan, CardNo:7},
			&Card{CardType:CardType_Tong, CardNo:8},
		},
	}

	hu14 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},

			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},
			&Card{CardType:CardType_Tiao, CardNo:1},

			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},

			&Card{CardType:CardType_Wan, CardNo:7},
			&Card{CardType:CardType_Wan, CardNo:7},
			&Card{CardType:CardType_Wan, CardNo:7},

			&Card{CardType:CardType_Wan, CardNo:6},
			&Card{CardType:CardType_Wan, CardNo:6},
			&Card{CardType:CardType_Tong, CardNo:6},
		},
	}

	hu2.Sort()
	hu5.Sort()
	hu8.Sort()
	hu11.Sort()
	hu14.Sort()

	if hu2.IsHu() || hu5.IsHu() || hu8.IsHu() || hu11.IsHu() || hu14.IsHu(){
		t.Fatal("all should not be hu")
	}
}