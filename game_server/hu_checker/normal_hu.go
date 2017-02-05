package hu_checker

//平胡
//大四喜
func init()  {
	FactoryInst().register("NORMAL_HU",
		func(config *HuConfig) Checker {
			return NewNormalHu(config)
		},
	)
}

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

	return cardsGetter.IsHu(), normalHu.config
}
