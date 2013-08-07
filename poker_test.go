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


