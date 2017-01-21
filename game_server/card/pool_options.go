package card


func WithFengCard() OptionFunc {
	return func(opt *Option) {
		opt.generater = append(opt.generater, &FengGenerater{})
	}
}

func WithJianCard() OptionFunc{
	return func(opt *Option) {
		opt.generater = append(opt.generater, &JianGenerater{})
	}
}

func WithHuaCard() OptionFunc{
	return func(opt *Option) {
		opt.generater = append(opt.generater, &HuaGenerater{})
	}
}

func WithWanCard() OptionFunc{
	return func(opt *Option) {
		opt.generater = append(opt.generater, &WanGenerater{})
	}
}

func WithTiaoCard() OptionFunc{
	return func(opt *Option) {
		opt.generater = append(opt.generater, &WanGenerater{})
	}
}

func WithTongCard() OptionFunc{
	return func(opt *Option) {
		opt.generater = append(opt.generater, &TongGenerater{})
	}
}