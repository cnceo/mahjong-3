package card

type CardInHand struct {
	ruleOpt			*RuleOptions
	ShowCard		*Cards
	HideCard		*Cards
	HuaCard			*Cards
}

func NewCardInHand(opt *RuleOptions) *CardInHand {
	return &CardInHand{
		ruleOpt: 	opt,
	}
}

//吃牌组合
func (cardInHand *CardInHand) CanChiWhat(card *Card) []*Card {
	if !cardInHand.ruleOpt.WithChi {
		return nil
	}

	//todo check which can chi
	return nil
}

//吃牌
func (cardInHand *CardInHand) Chi(card *Card) bool {
	if !cardInHand.ruleOpt.WithChi {
		return false
	}

	//todo chi what

	return true
}

func (cardInHand *CardInHand) CanPengWhat() []*Card {
	//todo
	return nil
}

func (cardInHand *CardInHand) Peng(card *Card) bool {
	//todo
	return false
}

func (cardInHand *CardInHand) CanGangWhat() []*Card {
	//todo
	return nil
}

func (cardInHand *CardInHand) Gang(card *Card) bool {
	//todo
	return true
}

func (cardInHand *CardInHand) CanHuWhat() []*Card {
	//todo
	return nil
}

func (cardInHand *CardInHand) Hu (card *Card) bool {
	//todo
	return true
}

//计算胡牌后的分数，当且仅当胡牌成功后才应该调该函数计算分数
func (cardInhand *CardInHand) ComputeScore() int {
	//todo
	return 0
}

//出牌
func (cardInHand *CardInHand) TakeWay(card *Card) bool{
	return cardInHand.HideCard.TakeWay(card)
}


