/**
 * Created with IntelliJ IDEA.
 * User: hcombs
 * Date: 8/3/13
 * Time: 10:40 AM
 * To change this template use File | Settings | File Templates.
 */
package poker

import "fmt"

//
//   This routine initializes the deck.  A deck of cards is
//   simply an integer array of length 52 (no jokers).  This
//   array is populated with each card, using the following
//   scheme:
//
//   An integer is made up of four bytes.  The high-order
//   bytes are used to hold the rank bit pattern, whereas
//   the low-order bytes hold the suit/rank/prime value
//   of the card.
//
//   +--------+--------+--------+--------+
//   |xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
//   +--------+--------+--------+--------+
//
//   p = prime number of rank (deuce=2,trey=3,four=5,five=7,...,ace=41)
//   r = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
//   cdhs = suit of card
//   b = bit turned on depending on rank of card
//
func init_deck( deck *Deck ) {

	suit := CLUB // (0x8000)

	n := 0 //index into the deck array

	for i := 0; i < NUM_SUITS; i++ {
		for  j := uint(0); j < uint(NUM_CARDS) ; j++ {
			deck[n] = NewCard( CardRank(j), Suit(suit) )
			n++
		}
		suit >>= 1 //go to the 'next' suit
	}
}

func NewCard( rank CardRank, suit Suit ) Card {



	return Card(
		primes[ rank ]  |
		uint(rank) << 8 |
		uint(suit)      |
		(1 << (16+rank) ) )
}

//
// Creates a new Deck initialized to the value of
func NewDeck() *Deck {
	var d Deck
	init_deck(&d)
	return &d
}

// Return the position within the deck for a given rank & suit,
// or (-1, false) if not found.
func (d Deck) find ( c CardRank, s Suit ) ( pos int, found bool) {
	for i, card := range d {
       if ( ( card & RANK_MASK ) != 0 ) && ( (int(card) & SUIT_MASK) == int(s) ) {
		 return i, true
	   }
	}
	return -1, false
}

// return the corresponding HandRank for a given value
func HandRankFromScore( val int ) HandRank {
	switch {
		case val > 6185: return HIGH_CARD
		case val > 3325: return ONE_PAIR
		case val > 2467: return TWO_PAIR
		case val > 1609: return THREE_OF_A_KIND
		case val > 1599: return STRAIGHT
		case val > 322:  return FLUSH
		case val > 166:  return FULL_HOUSE
		case val > 10:   return FOUR_OF_A_KIND
	}
	return STRAIGHT_FLUSH
}

// String value for a Hand
func (h Hand) String() string {
	return fmt.Sprintf("%s %s %s %s %s", h[0],h[1],h[2],h[3],h[4])
}

// String value for a Card
func (c Card) String() string {
	rank := (c >> 8) & 0xF
	return fmt.Sprintf("%s%s", cardRankString[rank], Suit(c & SUIT_MASK) )
}
