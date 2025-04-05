/*
 *  Copyright 2025 Markus Mahlberg <138420+mwmahlberg@users.noreply.github.com>
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package solitaire

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuitTests struct {
	suite.Suite
}

func (s *SuitTests) TestSuitString() {
	testCases := []struct {
		desc     string
		suit     suit
		expected string
	}{
		{
			desc:     "Clubs",
			suit:     clubs,
			expected: "♣",
		},
		{
			desc:     "Diamonds",
			suit:     diamonds,
			expected: "♦",
		},
		{
			desc:     "Hearts",
			suit:     hearts,
			expected: "♥",
		},
		{
			desc:     "Spades",
			suit:     spades,
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

func (s *SuitTests) TestSuitInvalid() {
	var suitInvalidSuit = suit(100)
	s.Panics(func() {
		str := suitInvalidSuit.String()
		s.Fail("Expected panic, but got: ", str)
	})
}

type RankTests struct {
	suite.Suite
}

func (s *RankTests) TestRankString() {
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

func (s *RankTests) TestRankStringNonFace() {
	for i := 3; i <= 10; i++ {
		r := rank(i)
		s.Run(r.String(), func() {
			result := r.String()
			s.Equal(fmt.Sprintf("%d", i), result)
		})
	}
}

func (s *RankTests) TestRankInvalid() {
	var rankInvalidRank = rank(100)
	s.Panics(func() {
		str := rankInvalidRank.String()
		s.Fail("Expected panic, but got: ", str)
	})
}

func (s *RankTests) TestRankShort() {
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

func (s *RankTests) TestRankShortNonFace() {
	for i := 3; i <= 10; i++ {
		r := rank(i)
		s.Run(r.String(), func() {
			result := r.Short()
			s.Equal(fmt.Sprintf("%d", i), result)
		})
	}
}

type CardTests struct {
	suite.Suite
}

func (s *CardTests) TestCardSuit() {
	c := card{rank: ace, suit: clubs}
	s.Equal(clubs, c.Suit(), "Expected suit to be Clubs")
	c = card{rank: two, suit: diamonds}
	s.Equal(diamonds, c.Suit(), "Expected suit to be Diamonds")
	c = card{rank: queen, suit: hearts}
	s.Equal(hearts, c.Suit(), "Expected suit to be Hearts")
	c = card{rank: king, suit: spades}
	s.Equal(spades, c.Suit(), "Expected suit to be Spades")
	c = card{rank: jokerA}
	s.Equal(suit(0), c.Suit(), "Expected suit to be 0 for Joker A")
}
func (s *CardTests) TestCardRank() {

	testCases := []struct {
		desc     string
		card     card
		expected rank
	}{
		{
			desc:     "Ace of Clubs",
			card:     card{rank: ace, suit: clubs},
			expected: ace,
		},
		{
			desc:     "Two of Diamonds",
			card:     card{rank: two, suit: diamonds},
			expected: two,
		},
		{
			desc:     "Queen of Hearts",
			card:     card{rank: queen, suit: hearts},
			expected: queen,
		},
		{
			desc:     "King of Spades",
			card:     card{rank: king, suit: spades},
			expected: king,
		},
		{
			desc:     "Joker A",
			card:     card{rank: jokerA},
			expected: jokerA,
		},
		{
			desc:     "Joker B",
			card:     card{rank: jokerB},
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
func (s *CardTests) TestCardValue() {
	c := card{rank: ace, suit: clubs}
	s.Equal(1, c.Value(), "Expected value to be 1 for Ace of Clubs")
	c = card{rank: two, suit: diamonds}
	s.Equal(15, c.Value(), "Expected value to be 2 for Two of Diamonds")
	c = card{rank: queen, suit: hearts}
	s.Equal(38, c.Value(), "Expected value to be 38 for Queen of Hearts")
	c = card{rank: king, suit: spades}
	s.Equal(52, c.Value(), "Expected value to be 52 for King of Spades")
	c = card{rank: jokerA, suit: 0}
	s.Equal(53, c.Value(), "Expected value to be 53 for Joker A")
	c = card{rank: jokerB, suit: 0}
	s.Equal(53, c.Value(), "Expected value to be 53 for Joker B")
}

func (s *CardTests) TestCardString() {

	testCases := []struct {
		desc     string
		card     card
		expected string
	}{
		{
			desc:     "Ace of Clubs",
			card:     card{rank: ace, suit: clubs},
			expected: "♣ A",
		},
		{
			desc:     "Two of Diamonds",
			card:     card{rank: two, suit: diamonds},
			expected: "♦ 2",
		},
		{
			desc:     "Queen of Hearts",
			card:     card{rank: queen, suit: hearts},
			expected: "♥ Q",
		},
		{
			desc:     "King of Spades",
			card:     card{rank: king, suit: spades},
			expected: "♠ K",
		},
		{
			desc:     "Joker A",
			card:     card{rank: jokerA},
			expected: "Joker A",
		},
		{
			desc:     "Joker B",
			card:     card{rank: jokerB},
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

func TestSuit(t *testing.T) {
	suite.Run(t, new(SuitTests))
}

func TestRank(t *testing.T) {
	suite.Run(t, new(RankTests))
}
func TestCard(t *testing.T) {
	suite.Run(t, new(CardTests))
}
