package room

import (
	"mahjong/game_server/card"
	"mahjong/game_server/hu_checker"
	"mahjong/game_server/log"
	"time"
)

type HuResult struct {
	IsHu 	bool
	Desc	string
	Score	int
}

func newHuResult(isHu bool, desc string, score int) *HuResult {
	return &HuResult{
		IsHu: isHu,
		Desc: desc,
		Score: score,
	}
}

type Player struct {
	id				int64			// 玩家id
	room			*Room			// 玩家所在的房间
	playingCards 	*card.PlayingCards	//玩家手上的牌
	huChecker		[]hu_checker.Checker

	gangFromPlayers []*Player
}

func NewPlayer(huChecker []hu_checker.Checker) *Player {
	player :=  &Player{
		playingCards:	card.NewPlayingCards(),
		huChecker:		huChecker,
		gangFromPlayers: make([]*Player, 0),
	}
	return player
}

func (player *Player) Reset() {
	player.playingCards.Reset()
	player.gangFromPlayers = player.gangFromPlayers[0:0]
}

func (player *Player) AddMagicCard(card *card.Card) {
	player.playingCards.AddMagicCard(card)
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

func (player *Player) Gang(card *card.Card, fromPlayer *Player) bool {
	ok := player.playingCards.Gang(card)
	if ok {
		player.gangFromPlayers = append(player.gangFromPlayers, fromPlayer)
	}
	return ok
}

func (player *Player) Drop(card *card.Card) bool {
	return player.playingCards.DropCard(card)
}

func (player *Player) DropTail() *card.Card {
	return player.playingCards.DropTail()
}

//是否点炮，胡了别人的牌
func (player *Player) IsDianPao(card *card.Card) *HuResult {
	player.AddCard(card)
	result := player.IsHu()
	if !result.IsHu {
		player.Drop(card)
	}
	return result
}

//是否自摸
func (player *Player) IsZiMo() *HuResult {
	return player.IsHu()
}

func (player *Player) IsHu() *HuResult {
	magicLen := player.playingCards.GetMagicCardsLen()
	for _, checker := range player.huChecker {
		if magicLen == 0 {
			isHu, conf := checker.IsHu(player.playingCards)
			if isHu {
				log.Debug("checker :", checker.GetConfig().ToString(), "succ")
				return newHuResult(isHu, conf.Desc, conf.Score)
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
					return newHuResult(isHu, conf.Desc, conf.Score)
				} else {
					player.playingCards.DropCards(cards)
				}
			}
		}
		log.Debug("checker :", checker.GetConfig().ToString(), "failed")
	}

	return newHuResult(false, "", 0)
}

func (player *Player) EnterRoom(room *Room) bool{
	result := make(chan bool, 1)
	room.Enter(player, result)
	select {
	case <-time.After(time.Second * 3):
		return false
	case succ := <- result :
		if succ {
			player.room = room
		}
		return succ
	}
	return false
}

func (player *Player) OnPlayingGameEnd(room *Room) {
	//todo
}

func (player *Player) OnRoomClosed(room *Room) {
	if player.room == room {
		player.room = nil
	}
}

func (player *Player) LeaveRoom() bool{
	if player.room == nil {
		return true
	}
	result := make(chan bool, 1)
	select {
	case <-time.After(time.Second * 3):
		return true
	case succ := <- result :
		return succ
	}
	return true
}

func (player *Player) ToString() string{
	return player.playingCards.ToString()
}

func (player *Player) computeMagicCandidate() []*card.Cards {
	magicNum := player.playingCards.GetMagicCardsLen()
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