package card

type CardInHand struct {
	ruleOpt		*RuleOptions
	ShowCard	[]*Card			//已经吃、碰、杠，显示给对手的牌
	HideCard	[]*Card			//还在手上隐藏不可见的牌
	HuaCard		[]*Card			//拿到的花牌
}

func NewCardInHand(opt RuleOptions) *CardInHand {
	return &CardInHand{
		ruleOpt: 	opt,
	}
}

//能吃什么牌
func (cardInHand *CardInHand) CanChiWhat() []*Card {
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
	return nil
}

func (cardInHand *CardInHand) CanHuWhat() []*Card {
	//todo
	return nil
}

func (cardInHand *CardInHand) Hu (card *Card) bool {
	//todo
	return true
}

//胡牌后的分数，只有胡牌才有该分数
func (cardInhand *CardInHand) Score() int {
	//todo
	return 0
}

