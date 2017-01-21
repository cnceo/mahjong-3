package card

type PoolGenerater interface {
	Generate() []*Card //生成牌池
}

//风牌生成器
type FengGenerater struct{}

func (generater *FengGenerater) Generate() []*Card {
	return generateCards(CardType_Feng, Feng_CardNo_Dong, Feng_CardNo_Bei, 4)
}

//箭牌生成器
type JianGenerater struct{}

func (generater *JianGenerater) Generate() []*Card {
	return generateCards(CardType_Jian, Jian_CardNo_Zhong, Jian_CardNo_Bai, 4)
}

//花牌生成器
type HuaGenerater struct{}

func (generater *HuaGenerater) Generate() []*Card {
	return generateCards(CardType_Hua, Hua_CardNo_Chun, Hua_CardNo_Ju, 1)
}

//万子牌生成器
type WanGenerater struct{}

func (generater *WanGenerater) Generate() []*Card {
	return generateCards(CardType_Wan, 1, 9, 4)
}

//条子牌生成器
type TiaoGenerater struct{}

func (generater *TiaoGenerater) Generate() []*Card {
	return generateCards(CardType_Tiao, 1, 9, 4)
}

//筒子牌生成器
type TongGenerater struct{}

func (generater *TongGenerater) Generate() []*Card {
	return generateCards(CardType_Tong, 1, 9, 4)
}

//从牌编号startCardNo到endCardNo，生成num个牌类型为cardType的牌
func generateCards(cardType int, startCardNo, endCardNo int, num int) []*Card {
	cards := make([]*Card, 0)

	for cardNo := startCardNo; cardNo <= endCardNo; cardNo++ {
		for i := 0; i < num; i++ {
			card := &Card{
				CardType: cardType,
				CardNo:   cardNo,
			}
			cards = append(cards, card)
		}
	}
	return cards
}
