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

func TestKeyStream(t *testing.T) {
	suite.Run(t, new(KeyStreamSuite))
}
