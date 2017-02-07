package card

import "testing"

func TestSort(t *testing.T) {
	pool := NewPool()

	pool.AddFengGenerater()
	pool.AddJianGenerater()
	pool.AddHuaGenerater()
	pool.AddWanGenerater()
	pool.AddTiaoGenerater()
	pool.AddTongGenerater()

	pool.ReGenerate()

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
		t.Fatal("all should be hu")
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

func TestCards_ComputeChi(t *testing.T) {
	notChiGroup := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong},
			&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Nan},
		},
	}
	groups := notChiGroup.computeChiGroup(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Xi})
	if len(groups) != 0 {
		t.Fatal("should not has group")
	}

	chiGroup := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Wan, CardNo:1},
			&Card{CardType:CardType_Wan, CardNo:2},
			&Card{CardType:CardType_Wan, CardNo:3},
			&Card{CardType:CardType_Wan, CardNo:4},
			&Card{CardType:CardType_Wan, CardNo:5},
			&Card{CardType:CardType_Wan, CardNo:5},
			&Card{CardType:CardType_Wan, CardNo:5},
			&Card{CardType:CardType_Wan, CardNo:6},
			&Card{CardType:CardType_Wan, CardNo:7},
			&Card{CardType:CardType_Wan, CardNo:8},
			&Card{CardType:CardType_Wan, CardNo:9},
		},
	}

	for cardNo:=1; cardNo<10; cardNo++{
		groups := chiGroup.computeChiGroup(&Card{CardType:CardType_Wan, CardNo:cardNo})
		length := len(groups)
		if (cardNo == 1 || cardNo == 9) && length != 1 {
			t.Fatal("should be 1 group")
		}

		if (cardNo == 2 || cardNo == 8) && length != 2 {
			t.Fatal("should be 2 group")
		}

		if (cardNo >= 3 && cardNo <= 7) && length != 3 {
			t.Fatal("should be 3 group")
		}

		if len(groups) == 0 {
			t.Fatal("should has group")
		}
	}

	nilGroups := chiGroup.computeChiGroup(&Card{CardType:CardType_Tong, CardNo:5})
	if len(nilGroups) != 0 {
		t.Fatal("it should not has group")
	}

}

func TestCards_SameAs(t *testing.T) {
	cards1 := NewCards()
	cards2 := NewCards()
	cards1.AppendCard(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong})
	cards2.AppendCard(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong})
	if !cards1.SameAs(cards2) {
		t.Fatal("should be same as")
	}

	cards2.AppendCard(&Card{CardType:CardType_Feng, CardNo:Feng_CardNo_Dong})
	if cards1.SameAs(cards2) {
		t.Fatal("should not be same as")
	}
}


func TestCards_IsHuSpe(t *testing.T) {

	//胡：23334
	// ABBBC => AA + {ABC}
	hu5 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:3},

			&Card{CardType:CardType_Tiao, CardNo:4},
			&Card{CardType:CardType_Tiao, CardNo:3},
			&Card{CardType:CardType_Tiao, CardNo:3},
		},
	}

	//胡 : 22223344、22333344、22334444、22334455、
	//  :  22223444、22233344
	//  AAAABBCC/AABBBBCC/AABBCCCC/AABBCCDD  => AA + {AABBCC}
	hu8 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},

			&Card{CardType:CardType_Tiao, CardNo:3},
			&Card{CardType:CardType_Tiao, CardNo:3},

			&Card{CardType:CardType_Tiao, CardNo:4},
			&Card{CardType:CardType_Tiao, CardNo:4},
		},
	}

	//胡: {11222333444 常规算法能检查出来}、12222333444、12223333444、12223334444{不能胡}、
	//    22223334445{不能胡}、22233334445、22233344445、{22233344455、22233344456}
	//   AAA/ABC + {Hu8}
	hu11 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Tiao, CardNo:1},

			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},

			&Card{CardType:CardType_Tiao, CardNo:3},
			&Card{CardType:CardType_Tiao, CardNo:3},
			&Card{CardType:CardType_Tiao, CardNo:3},

			&Card{CardType:CardType_Wan, CardNo:4},
			&Card{CardType:CardType_Wan, CardNo:4},
			&Card{CardType:CardType_Wan, CardNo:4},
		},
	}

	//12222333344445
	hu14 := &Cards{
		data: []*Card{
			&Card{CardType:CardType_Tiao, CardNo:1},

			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},
			&Card{CardType:CardType_Tiao, CardNo:2},

			&Card{CardType:CardType_Tiao, CardNo:3},
			&Card{CardType:CardType_Tiao, CardNo:3},
			&Card{CardType:CardType_Tiao, CardNo:3},
			&Card{CardType:CardType_Tiao, CardNo:3},

			&Card{CardType:CardType_Tiao, CardNo:4},
			&Card{CardType:CardType_Tiao, CardNo:4},
			&Card{CardType:CardType_Tiao, CardNo:4},
			&Card{CardType:CardType_Tiao, CardNo:4},

			&Card{CardType:CardType_Tiao, CardNo:5},
		},
	}


	hu5.Sort()
	hu8.Sort()
	hu11.Sort()
	hu14.Sort()

	if  !hu5.IsHu() || !hu8.IsHu() || !hu11.IsHu() || !hu14.IsHu(){
		t.Fatal("all should be hu")
	}
}
