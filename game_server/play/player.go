package play

import (
	"mahjong/game_server/card"
	"mahjong/game_server/hu_checker"
)

const (
	MAX_MAGIC_NUM = 4
)

type Player struct {
	magicCards			[]*card.Card		//赖子牌
	playingCards 		*card.PlayingCards	//
	huChecker			[]hu_checker.Checker
}

func NewPlayer(huChecker []hu_checker.Checker) *Player {
	player :=  &Player{
		magicCards: 	make([]*card.Card, 0),
		playingCards:	card.NewPlayingCards(),
		huChecker:		huChecker,
	}
	return player
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
				return isHu, conf.Desc, conf.Score
			}
		} else {
			//TODO 支持赖子牌的胡牌计算
			for i := 0; i<magicLen; i++ {
				for _, card := range card.GetAllCards().Data() {
					player.playingCards.AddCard(card)
					isHu, conf := checker.IsHu(player.playingCards)
					if isHu {
						return isHu, conf.Desc, conf.Score
					} else {
						player.playingCards.DropCard(card)
					}
				}
			}
		}
	}

	return false, "", 0
}

func (player *Player) ToString() string{
	return player.playingCards.ToString()
}

func (player *Player) computeMagicCandidate() []*card.Cards {
	magicNum := len(player.magicCards)
	if magicNum == 0 {
		return nil
	}
	allCards := card.GetAllCards()
	allCardsLen := allCards.Len()
	for i := 0; i < allCardsLen; i++ {
		for j := 0; j < allCardsLen; j++ {
			for k := 0; k<allCardsLen ; k++ {
				for l := 0 ; l<allCardsLen ;  {

				}
			}

		}
	}
	return magicCards
}