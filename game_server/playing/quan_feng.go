package playing

import "mahjong/game_server/card"

type QuanFeng struct {
	quan	int
}

func newQuanFeng(quan int) *QuanFeng {
	return &QuanFeng{
		quan:	quan,
	}
}

func (qf *QuanFeng) next() *QuanFeng {
	if qf.quan == card.Feng_CardNo_Bei {
		return newQuanFeng(card.Feng_CardNo_Dong)
	}
	return newQuanFeng(qf.quan+1)
}

func (qf *QuanFeng) isLastQuanFeng() bool {
	return qf.quan == card.Feng_CardNo_Bei
}

func (qf *QuanFeng) isFirstQuanFeng() bool {
	return qf.quan == card.Feng_CardNo_Dong
}

func (qf *QuanFeng) sameAs(other *QuanFeng) bool {
	return qf.quan == other.quan
}