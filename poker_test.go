package poker

import (
	"testing"
	//"fmt"
)

var validSuits = []struct {
	s Suit
	value uint
} {
	{ SPADE, 0x1000 },
	{ HEART, 0x2000 },
	{ DIAMOND, 0x4000 },
	{ CLUB, 0x8000 },
}


func Test_suits(t *testing.T) {

	for _,testCase := range validSuits {
		if (testCase.s != Suit(testCase.value) ) {
			t.Errorf("Expected %v got %v", testCase.value, testCase.s)
		}
	}
}

type HandFrequency [NUM_HAND_TYPES]int

// Here's the distribution we should get
//
// Straight Flush:       40
// Four of a Kind:      624
// Full House:     3744
// Flush:     5108
// Straight:    10200
// Three of a Kind:    54912
// Two Pair:   123552
// One Pair:  1098240
// High Card:  1302540

func Test_allHands( t *testing.T ) {

	var expectedNumbers = HandFrequency{0,40,624,3744,5108,10200,54912,123552,1098240,1302540 }

	var hand Hand
	deck := NewDeck();
	var frequency HandFrequency

	for a:=0; a < (CARDS_IN_DECK - 4); a++ {
		hand[0] = deck[a]
		for b:=a+1; b < (CARDS_IN_DECK - 3); b++ {
			hand[1] = deck[b]
			for c:=b+1; c < (CARDS_IN_DECK - 2); c++ {
				hand[2] = deck[c]
				for d:=c+1; d < (CARDS_IN_DECK - 1); d++ {
					hand[3] = deck[d]
					for e:=d+1; e < CARDS_IN_DECK; e++ {
						hand[4] = deck[e]
						frequency[ hand.Eval().Rank() ]++
					}
				}
			}
		}
	}

	for i,v := range expectedNumbers {
		if ( frequency[i] != v ) {
			t.Errorf("For %s, expected %d, got %d", HandRank(i), v, frequency[i])
		}
	}
}

