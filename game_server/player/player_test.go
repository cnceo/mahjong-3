package player

import (
	"testing"
	"mahjong/game_server/card"
)

func TestNewPlayer(t *testing.T) {
	player := NewPlayer()
	player.AddCard(&card.Card{CardType:card.CardType_Wan, CardNo:1})
	player.Chi(&card.Card{CardType:card.CardType_Wan, CardNo:2})
	player.Peng(&card.Card{CardType:card.CardType_Wan, CardNo:3})
	player.Gang(&card.Card{CardType:card.CardType_Wan, CardNo:4})
	t.Log(player.ToString())
}
