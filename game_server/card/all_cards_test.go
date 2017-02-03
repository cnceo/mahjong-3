package card

import (
	"testing"
	"time"
)

func TestCandidate(t *testing.T) {
	start := time.Now()
	t.Log("OneMagicCandidate")
	total := 0
	for _, cards := range OneMagicCandidate {
		total += cards.Len()
		//t.Log(cards.ToString())
	}
	t.Log("OneMagicCandidate", total)

	t.Log("TwoMagicCandidate")
	total = 0
	for _, cards := range TwoMagicCandidate {
		total += cards.Len()
		//t.Log(cards.ToString())
	}
	t.Log("TwoMagicCandidate", total)

	t.Log("ThreeMagicCandidate")
	total = 0
	for _, cards := range ThreeMagicCandidate {
		total += cards.Len()
		//t.Log(cards.ToString())
	}
	t.Log("ThreeMagicCandidate", total)

	t.Log("FourMagicCandidate")
	total = 0
	for _, cards := range FourMagicCandidate {
		total += cards.Len()
		//t.Log(cards.ToString())
	}
	t.Log("FourMagicCandidate", total)

	//time.Sleep(time.Second)
	//time.Sleep(time.Millisecond)
	//time.Sleep(time.Microsecond)
	//time.Sleep(time.Nanosecond)
	end := time.Now()
	t.Log("spent :", end.Sub(start))
}
