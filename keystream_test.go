package solitaire

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type KeyStreamSuite struct {
	suite.Suite
}

func (s *KeyStreamSuite) TestSuitString() {
	testCases := []struct {
		desc     string
		suit     suit
		expected string
	}{
		{
			desc:     "Clubs",
			suit:     Clubs,
			expected: "♣",
		},
		{
			desc:     "Diamonds",
			suit:     Diamonds,
			expected: "♦",
		},
		{
			desc:     "Hearts",
			suit:     Hearts,
			expected: "♥",
		},
		{
			desc:     "Spades",
			suit:     Spades,
			expected: "♠",
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			result := tC.suit.String()
			s.Equal(tC.expected, result)
		})
	}

}

func (s *KeyStreamSuite) TestSuitInvalid() {
	var suitInvalidSuit = suit(100)
	s.Panics(func() {
		str := suitInvalidSuit.String()
		s.Fail("Expected panic, but got: ", str)
	})
}

func (s *KeyStreamSuite) TestSuitValue() {
	testCases := []struct {
		desc    string
		suit    suit
		summand int
	}{
		{
			desc:    "Clubs",
			suit:    Clubs,
			summand: 0,
		},
		{
			desc:    "Diamonds",
			suit:    Diamonds,
			summand: 13,
		},
		{
			desc:    "Hearts",
			suit:    Hearts,
			summand: 26,
		},
		{
			desc:    "Spades",
			suit:    Spades,
			summand: 39,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			for i := 1; i <= 13; i++ {
				result := tC.suit.Value(i)
				expected := (i + int(tC.suit)) % 26
				if expected == 0 {
					expected = 26
				}
				s.Equal(expected, result)
			}
		})
	}
}

func (s *KeyStreamSuite) TestSuitValueInvalid() {
	s.Panics(func() {
		Clubs.Value(14)
	})
	s.Panics(func() {
		Clubs.Value(-1)
	})
}

func (s *KeyStreamSuite) TestRankString() {
	testCases := []struct {
		desc     string
		rank     rank
		expected string
	}{
		{
			desc:     "Ace",
			rank:     ace,
			expected: "Ace",
		},
		{
			desc:     "Two",
			rank:     two,
			expected: "2",
		},
		{
			desc:     "Jack",
			rank:     jack,
			expected: "Jack",
		},
		{
			desc:     "Queen",
			rank:     queen,
			expected: "Queen",
		},
		{
			desc:     "King",
			rank:     king,
			expected: "King",
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			result := tC.rank.String()
			s.Equal(tC.expected, result)
		})
	}
}

func (s *KeyStreamSuite) TestRankStringNonFace() {
	for i := 3; i <= 10; i++ {
		r := rank(i)
		s.Run(r.String(), func() {
			result := r.String()
			s.Equal(fmt.Sprintf("%d", i), result)
		})
	}
}

func (s *KeyStreamSuite) TestRankInvalid() {
	var rankInvalidRank = rank(100)
	s.Panics(func() {
		str := rankInvalidRank.String()
		s.Fail("Expected panic, but got: ", str)
	})
}

func (s *KeyStreamSuite) TestRankShort() {
	testCases := []struct {
		desc     string
		rank     rank
		expected string
	}{
		{
			desc:     "Ace",
			rank:     ace,
			expected: "A",
		},
		{
			desc:     "Two",
			rank:     two,
			expected: "2",
		},
		{
			desc:     "Jack",
			rank:     jack,
			expected: "J",
		},
		{
			desc:     "Queen",
			rank:     queen,
			expected: "Q",
		},
		{
			desc:     "King",
			rank:     king,
			expected: "K",
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			result := tC.rank.Short()
			s.Equal(tC.expected, result)
		})
	}
}

func (s *KeyStreamSuite) TestRankShortNonFace() {
	for i := 3; i <= 10; i++ {
		r := rank(i)
		s.Run(r.String(), func() {
			result := r.Short()
			s.Equal(fmt.Sprintf("%d", i), result)
		})
	}
}

func (s *KeyStreamSuite) TestDefaultAlphabetNotPresent() {
	idx := alphabet.Index(byte('@'))
	s.Equal(-1, idx, "Expected -1 for non-alphabet character")
}

func (s *KeyStreamSuite) TestCardSuit() {
	c := Card{rank: ace, suit: Clubs}
	s.Equal(Clubs, c.Suit(), "Expected suit to be Clubs")
	c = Card{rank: two, suit: Diamonds}
	s.Equal(Diamonds, c.Suit(), "Expected suit to be Diamonds")
	c = Card{rank: queen, suit: Hearts}
	s.Equal(Hearts, c.Suit(), "Expected suit to be Hearts")
	c = Card{rank: king, suit: Spades}
	s.Equal(Spades, c.Suit(), "Expected suit to be Spades")
	c = Card{rank: jokerA}
	s.Equal(suit(0), c.Suit(), "Expected suit to be 0 for Joker A")
}
func (s *KeyStreamSuite) TestCardRank() {

	testCases := []struct {
		desc     string
		card     Card
		expected rank
	}{
		{
			desc:     "Ace of Clubs",
			card:     Card{rank: ace, suit: Clubs},
			expected: ace,
		},
		{
			desc:     "Two of Diamonds",
			card:     Card{rank: two, suit: Diamonds},
			expected: two,
		},
		{
			desc:     "Queen of Hearts",
			card:     Card{rank: queen, suit: Hearts},
			expected: queen,
		},
		{
			desc:     "King of Spades",
			card:     Card{rank: king, suit: Spades},
			expected: king,
		},
		{
			desc:     "Joker A",
			card:     Card{rank: jokerA},
			expected: jokerA,
		},
		{
			desc:     "Joker B",
			card:     Card{rank: jokerB},
			expected: jokerB,
		},
	}
	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			result := tC.card.Rank()
			s.Equal(tC.expected, result)
		})
	}

}
func (s *KeyStreamSuite) TestCardValue() {
	c := Card{rank: ace, suit: Clubs}
	s.Equal(1, c.Value(), "Expected value to be 1 for Ace of Clubs")
	c = Card{rank: two, suit: Diamonds}
	s.Equal(15, c.Value(), "Expected value to be 2 for Two of Diamonds")
	c = Card{rank: queen, suit: Hearts}
	s.Equal(38, c.Value(), "Expected value to be 38 for Queen of Hearts")
	c = Card{rank: king, suit: Spades}
	s.Equal(52, c.Value(), "Expected value to be 52 for King of Spades")
	c = Card{rank: jokerA, suit: 0}
	s.Equal(53, c.Value(), "Expected value to be 53 for Joker A")
	c = Card{rank: jokerB, suit: 0}
	s.Equal(53, c.Value(), "Expected value to be 53 for Joker B")
}

func (s *KeyStreamSuite) TestCardString() {

	testCases := []struct {
		desc     string
		card     Card
		expected string
	}{
		{
			desc:     "Ace of Clubs",
			card:     Card{rank: ace, suit: Clubs},
			expected: "♣ A",
		},
		{
			desc:     "Two of Diamonds",
			card:     Card{rank: two, suit: Diamonds},
			expected: "♦ 2",
		},
		{
			desc:     "Queen of Hearts",
			card:     Card{rank: queen, suit: Hearts},
			expected: "♥ Q",
		},
		{
			desc:     "King of Spades",
			card:     Card{rank: king, suit: Spades},
			expected: "♠ K",
		},
		{
			desc:     "Joker A",
			card:     Card{rank: jokerA},
			expected: "Joker A",
		},
		{
			desc:     "Joker B",
			card:     Card{rank: jokerB},
			expected: "Joker B",
		},
	}
	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			result := tC.card.String()
			s.Equal(tC.expected, result)
		})
	}
}

func TestKeyStream(t *testing.T) {
	suite.Run(t, new(KeyStreamSuite))
}
