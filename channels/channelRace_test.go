package channels

import (
	"testing"
)

func TestChannelRace(t *testing.T) {
	
	var res []string
	 winnerMap := make(map[string]int)
	for i := 0; i < 20; i++ {
		ans := WhoWins()
		res = append(res, ans)
	}
	for _, winner := range res {
		winnerMap[winner]++
	}
	for _, v := range winnerMap {
		if v == 20 {
			 t.Errorf("Expected at least the other side to win at least once")
		}
	}

}