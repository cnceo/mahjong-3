package hu_checker

type Checker interface {
	IsHu(cardsGetter CardsGetter) (succ bool, config *HuConfig)
}

