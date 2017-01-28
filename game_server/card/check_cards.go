package card

func IsAA(card1, card2 *Card) bool {
	return card1.SameAs(card2)
}

func IsAAA(card1, card2, card3 *Card) bool {
	return IsAA(card1, card2) && IsAA(card2, card3)
}

func IsAAAA(card1, card2, card3, card4 *Card) bool {
	return IsAA(card1, card2) && IsAAA(card2, card3, card4)
}

//检查三张牌是不是ABC顺子牌
func IsABC(card1, card2, card3 *Card) bool {
	return card1.PrevAt(card2) && card2.PrevAt(card3)
}

//检查四张牌是不是ABBC的牌型
func IsABBC(card1, card2, card3, card4 *Card) bool {
	return IsAA(card2, card3) && IsABC(card1, card2, card4)
}

//检查五张牌是不是ABBBC的牌型
func IsABBBC(card1, card2, card3, card4, card5 *Card) bool {
	return IsAAA(card2, card3, card4) && IsABC(card1, card2, card5)
}

//检查六张牌是不是ABBBBC的牌型
func IsABBBBC(card1, card2, card3, card4, card5, card6 *Card) bool {
	return IsAAAA(card2, card3, card4, card5) && IsABC(card1, card2, card6)
}

//3连对
func IsAABBCC(card1, card2, card3, card4, card5, card6 *Card) bool {
	return IsABC(card1, card3, card5) &&
		IsAA(card1, card2) && IsAA(card3, card4) && IsAA(card5, card6)
}

//3连高压
func IsAAABBBCCC(card1, card2, card3, card4, card5,
card6, card7, card8, card9 *Card) bool {
	return IsABC(card1, card4, card7) && IsAAA(card1, card2, card3) &&
	IsAAA(card4, card5, card6) && IsAAA(card7, card8, card9)
}

//3连刻
func IsAAAABBBBCCCC(card1, card2, card3, card4,
card5, card6, card7, card8,
card9, card10, card11, card12 *Card) bool {
	return IsABC(card1, card5, card9) && IsAAAA(card1, card2, card3, card4) &&
	IsAAAA(card5, card6, card7, card8) && IsAAAA(card9, card10, card11, card12)
}

//6连对
func IsAABBCCDDEEFF(card1, card2, card3, card4, card5, card6,
card7, card8, card9, card10, card11, card12 *Card) bool {
	return IsAA(card1, card2) && IsAA(card3, card4) && IsAA(card5, card6) &&
	IsAA(card7, card8) && IsAA(card9, card10) && IsAA(card11, card12) &&
	IsABC(card1, card3, card5) && IsABC(card5, card7, card9) && IsABC(card7, card9, card11)
}

//3张牌是否OK, ABC/AAA格式为OK
func Is3CardsOk(cards ...*Card) bool {
	if len(cards) != 3 {
		return false
	}
	return IsAAA(cards[0], cards[1], cards[2]) || IsABC(cards[0], cards[1], cards[2])
}

//6张牌是否OK, 3连对/Is3CardsOk * 2/A + BBBB + C
func Is6CardsOk(cards ...*Card) bool {
	if len(cards) != 6 {
		return false
	}
	//左右两边都是Is3CardOk
	if Is3CardsOk(cards[0:3]...) && Is3CardsOk(cards[3:6]...) {
		return true
	}

	//3连对
	if IsAABBCC(cards[0], cards[1], cards[2], cards[3], cards[4], cards[5]) {
		return true
	}

	//中间4个相同，但是和左右各一个组成顺子
	if IsAAAA(cards[1], cards[2], cards[3], cards[4]) &&
		IsABC(cards[0], cards[1], cards[5]){
		return true
	}
	return false
}

//9张牌是否OK
func Is9CardsOk(cards ...*Card) bool {
	if len(cards) != 9 {
		return false
	}

	// 3 + 6
	if Is3CardsOk(cards[0:3]...) &&
		Is6CardsOk(cards[3:9]...) {
		return true
	}

	//6 + 3
	if Is6CardsOk(cards[0:6]...) &&
		Is3CardsOk(cards[6:9]...){
		return true
	}
	return false
}

//12张牌是否OK
func Is12CardsOk(cards ...*Card) bool {
	if len(cards) != 12 {
		return false
	}

	//3 + 9
	if Is3CardsOk(cards[0:3]...) &&
		Is9CardsOk(cards[3:12]...) {
		return true
	}

	//9 + 3
	if Is9CardsOk(cards[0:9]...) &&
		Is3CardsOk(cards[9:12]...){
		return true
	}
	return false
}