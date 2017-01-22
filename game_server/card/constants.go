package card

const(
	Max_CardType = CardType_Hua - CardType_Wan
)

const (
	CardType_Wan  int = iota//万子
	CardType_Tiao //条子
	CardType_Tong //筒子

	CardType_Feng //风牌：东、南、西、北、

	CardType_Jian //箭牌：中、发、白

	CardType_Hua //花牌：春、夏、秋、冬，梅、兰、竹、菊
)

const (
	Feng_CardNo_Dong int = iota + 1 //东
	Feng_CardNo_Nan                 //南
	Feng_CardNo_Xi                  //西
	Feng_CardNo_Bei                 //北
)

const (
	Jian_CardNo_Zhong int = iota + 1 //中
	Jian_CardNo_Fa                   //发
	Jian_CardNo_Bai                  //白
)

const (
	Hua_CardNo_Chun int = iota + 1 //春
	Hua_CardNo_Xia                 //夏
	Hua_CardNo_Qiu                 //秋
	Hua_CardNo_Dong                //冬
	Hua_CardNo_Mei                 //梅
	Hua_CardNo_Lan                 //兰
	Hua_CardNo_Zhu                 //竹
	Hua_CardNo_Ju                  //菊
)
