package card

import "testing"

func TestCard(t *testing.T) {
	t.Log(CardType_Feng)
	t.Log(CardType_Jian)
	t.Log(CardType_Hua)
	t.Log(CardType_Wan)
	t.Log(CardType_Tiao)
	t.Log(CardType_Tong)
	data := []int{1}
	t.Log(data)
	data = append(data[0:0], data[1:]...)
	t.Log(data)
}
