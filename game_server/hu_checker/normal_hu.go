package hu_checker

//平胡

type NormalHu struct {
	config	*HuConfig
}

func NewNormalHu(config *HuConfig) *NormalHu {
	return &NormalHu{
		config:	config,
	}
}

func (normalHu *NormalHu) IsHu(cardsGetter CardsGetter) (bool, *HuConfig) {
	if !normalHu.config.IsEnabled {
		//fmt.Println(1)
		return false, normalHu.config
	}

	return cardsGetter.GetInHandCards().IsHu(), normalHu.config
}
