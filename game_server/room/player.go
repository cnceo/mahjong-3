package room

import (
	"mahjong/game_server/card"
	"mahjong/game_server/hu_checker"
	"mahjong/game_server/log"
)

type Player struct {
	id				int64			// 玩家id
	room			*Room			// 玩家所在的房间
	magicCards		[]*card.Card	// 玩家手上的赖子牌
	playingCards 	*card.PlayingCards	//
	huChecker		[]hu_checker.Checker
}

func NewPlayer(huChecker []hu_checker.Checker) *Player {
	player :=  &Player{
		magicCards: 	make([]*card.Card, 0),
		playingCards:	card.NewPlayingCards(),
		huChecker:		huChecker,
	}
	return player
}

func (player *Player) ResetCards() {
	player.magicCards = player.magicCards[0:0]
	player.playingCards.Reset()
}

func (player *Player) AddMagicCard(card *card.Card) {
	player.magicCards = append(player.magicCards, card)
}

func (player *Player) AddCard(card *card.Card) {
	player.playingCards.AddCard(card)
}

/*	计算指定的牌可以吃牌的组合
*/
func (player *Player) ComputeChiGroup(card *card.Card) []*card.Cards {
	return player.playingCards.ComputeChiGroup(card)
}

func (player *Player) Chi(whatCard *card.Card, whatGroup *card.Cards) bool {
	return player.playingCards.Chi(whatCard, whatGroup)
}

func (player *Player) Peng(card *card.Card) bool {
	return player.playingCards.Peng(card)
}

func (player *Player) Gang(card *card.Card) bool {
	return player.playingCards.Gang(card)
}

func (player *Player) Drop(card *card.Card) bool {
	return player.playingCards.DropCard(card)
}

func (player *Player) IsHu() (isHu bool, desc string, score int) {
	magicLen := len(player.magicCards)
	for _, checker := range player.huChecker {
		if magicLen == 0 {
			isHu, conf := checker.IsHu(player.playingCards)
			if isHu {
				log.Debug("checker :", checker.GetConfig().ToString(), "succ")
				return isHu, conf.Desc, conf.Score
			}
		} else {
			//支持赖子牌的胡牌计算, 暴力穷举法，把赖子牌的所有候选集一个个试，胜在够简单
			candidate := player.computeMagicCandidate()
			tryCnt := 0
			for _, cards := range candidate {
				tryCnt++
				player.playingCards.AddCards(cards)
				isHu, conf := checker.IsHu(player.playingCards)
				if isHu {
					log.Debug("checker :", checker.GetConfig().ToString(), "succ, tryMagicCnt :", tryCnt, ",cards:", cards.ToString())
					log.Debug("tryCnt :", tryCnt, ", cards :", cards.ToString())
					return isHu, conf.Desc, conf.Score
				} else {
					player.playingCards.DropCards(cards)
				}
			}
		}
		log.Debug("checker :", checker.GetConfig().ToString(), "failed")
	}

	return false, "", 0
}

func (player *Player) EnterRoom(room *Room) bool{
	result := make(chan bool, 1)
	room.Enter(player, result)
	succ := <-result
	if succ {
		player.room = room
	}
	return succ
}

func (player *Player) LeaveRoom() {
	player.room = nil
}

func (player *Player) ToString() string{
	return player.playingCards.ToString()
}

func (player *Player) computeMagicCandidate() []*card.Cards {
	magicNum := len(player.magicCards)
	switch magicNum {
	case 0:
		return nil
	case 1:
		return card.OneMagicCandidate
	case 2:
		return card.TwoMagicCandidate
	case 3:
		return card.ThreeMagicCandidate
	case 4:
		return card.FourMagicCandidate
	default:
		return nil
	}
	return nil
}

//别的玩家打出一张牌时触发
func (player *Player) OnOtherDropCard(other *Player, card *card.Card) {
	//TODO
}

//摸到一张牌时触发
func (player *Player) OnGetCard(card *card.Card) {
	//TODO
}

func (player *Player) DebugAllChecker() {
	for _, checker := range player.huChecker {
		log.Debug("player's checker :", checker.GetConfig())
	}
}