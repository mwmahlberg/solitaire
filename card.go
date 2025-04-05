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
	"regexp"
	"strconv"
)

type rank int

func (r *rank) String() string {
	switch *r {
	case ace:
		return "Ace"
	case two:
		fallthrough
	case three:
		fallthrough
	case four:
		fallthrough
	case five:
		fallthrough
	case six:
		fallthrough
	case seven:
		fallthrough
	case eight:
		fallthrough
	case nine:
		fallthrough
	case ten:
		return fmt.Sprintf("%d", *r)
	case jack:
		return "Jack"
	case queen:
		return "Queen"
	case king:
		return "King"
	default:
		panic("invalid rank")
	}
}

func (r *rank) Short() string {
	switch *r {
	case ace:
		return "A"
	case jack:
		return "J"
	case queen:
		return "Q"
	case king:
		return "K"
	default:
		return r.String()
	}
}

const (
	// TODO: Make this zero-based
	// While it is more intuitive to have the first card be 1, it is not
	// necessary for the algorithm.
	ace rank = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king

	jokerA rank = 53
	jokerB rank = 54
)

type suit int

const (
	clubs    suit = 0
	diamonds suit = 13
	hearts   suit = 26
	spades   suit = 39
)

// String returns the string representation of the suit.
// It returns the Unicode character for the suit:
//
//   - ♣ for Clubs
//   - ♦ for Diamonds
//   - ♥ for Hearts
//   - ♠ for Spades
func (s suit) String() string {
	switch s {
	case clubs:
		return "♣"
	case diamonds:
		return "♦"
	case hearts:
		return "♥"
	case spades:
		return "♠"
	default:
		// This case is only necessary to satisfy the compiler.
		// In practice, this should never happen because the suit
		// is always one of the defined constants.
		// If it does, it means that the code is incorrect.
		// In that case, we panic to indicate a programming error.
		panic("invalid suit")
	}
}

// Short returns the short string representation of the suit.
// It returns a letter for the suit:
//
//   - C for Clubs
//   - D for Diamonds
//   - H for Hearts
//   - S for Spades
func (s suit) Short() string {
	switch s {
	case clubs:
		return "C"
	case diamonds:
		return "D"
	case hearts:
		return "H"
	case spades:
		return "S"
	default:
		panic("invalid suit")
	}
}

// card represents a playing card in the Solitaire deck.
// A card is described as a combination of a suit and a rank.
// The rank is an integer from 1 to 13, where 1 is Ace and 13 is King.
// The suit is an integer, where 0 is Clubs, 13 is Diamonds,
// 26 is Hearts, and 39 is Spades.
// The jokers are represented by the ranks 53 (joker A, aka "slow" Joker) and 54 (joker B, aka as "fast" joker).
type card struct {
	suit suit
	rank rank
}

// IsJokerA returns true if the card is the Joker A (slow joker).
// Joker A is called the "slow" joker because it is moved one position
// when a new key is generated.
// Joker A is represented by the rank 53.
func (c card) IsJokerA() bool {
	return c.rank == jokerA
}

// IsJokerB returns true if the card is the Joker B (fast joker).
// Joker B is called the "fast" joker because it is moved two positions
// when a new key is generated.
// Joker B is represented by the rank 54.
func (c card) IsJokerB() bool {
	return c.rank == jokerB
}

// IsJoker returns true if either [card.IsJokerA] or [card.IsJokerB] is true.
func (c card) IsJoker() bool {
	return c.IsJokerA() || c.IsJokerB()
}

func (c card) Suit() suit {
	return c.suit
}

// Rank returns the rank of the card.
// The rank is an integer from 1 to 13, where 1 is Ace and 13 is King.
func (c card) Rank() rank {
	return c.rank
}

func (c card) Value() int {
	if c.IsJoker() {
		return 53
	}
	return int(c.suit) + int(c.rank)
}

func (c card) String() string {
	if c.IsJokerA() {
		return "Joker A"
	}
	if c.IsJokerB() {
		return "Joker B"
	}

	return fmt.Sprintf("%s %s", c.suit.String(), c.rank.Short())
}

func (c card) Short() string {
	if c.IsJokerA() {
		return "JA"
	}
	if c.IsJokerB() {
		return "JB"
	}

	return fmt.Sprintf("%s%s", c.suit.Short(), c.rank.Short())
}

var (
	cardRegex  = regexp.MustCompile(`(?i)^(?P<Suit>C|D|H|S)(?P<Rank>[1-9]|J|Q|K)|(?P<Joker>JA|JB)$`)
	suitIndex  = cardRegex.SubexpIndex("Suit")
	rankIndex  = cardRegex.SubexpIndex("Rank")
	jokerIndex = cardRegex.SubexpIndex("Joker")
)

func cardFromString(s string) (card, error) {
	var c card
	if !cardRegex.MatchString(s) {
		return c, fmt.Errorf("invalid card %s", s)
	}

	matches := cardRegex.FindStringSubmatch(s)
	if matches[jokerIndex] != "" {
		switch matches[jokerIndex] {
		case "JA":
			c.rank = jokerA
		case "JB":
			c.rank = jokerB
		default:
			return c, fmt.Errorf("invalid joker %s", matches[jokerIndex])
		}
		return c, nil
	}

	switch matches[suitIndex] {
	case "C":
		c.suit = clubs
	case "D":
		c.suit = diamonds
	case "H":
		c.suit = hearts
	case "S":
		c.suit = spades
	default:
		return c, fmt.Errorf("invalid suit %s", matches[suitIndex])
	}

	switch matches[rankIndex] {
	case "1":
		fallthrough
	case "A":
		c.rank = ace
	case "Q":
		c.rank = queen
	case "K":
		c.rank = king
	default:
		r, err := strconv.Atoi(matches[rankIndex])
		if err != nil {
			return c, fmt.Errorf("invalid rank %s", matches[rankIndex])
		}
		if r < 1 || r > 13 {
			return c, fmt.Errorf("invalid rank %s", matches[rankIndex])
		}
		c.rank = rank(r)
	}

	return c, nil
}

var initialDeck = []card{
	{suit: clubs, rank: ace},
	{suit: clubs, rank: two},
	{suit: clubs, rank: three},
	{suit: clubs, rank: four},
	{suit: clubs, rank: five},
	{suit: clubs, rank: six},
	{suit: clubs, rank: seven},
	{suit: clubs, rank: eight},
	{suit: clubs, rank: nine},
	{suit: clubs, rank: ten},
	{suit: clubs, rank: jack},
	{suit: clubs, rank: queen},
	{suit: clubs, rank: king},
	{suit: diamonds, rank: ace},
	{suit: diamonds, rank: two},
	{suit: diamonds, rank: three},
	{suit: diamonds, rank: four},
	{suit: diamonds, rank: five},
	{suit: diamonds, rank: six},
	{suit: diamonds, rank: seven},
	{suit: diamonds, rank: eight},
	{suit: diamonds, rank: nine},
	{suit: diamonds, rank: ten},
	{suit: diamonds, rank: jack},
	{suit: diamonds, rank: queen},
	{suit: diamonds, rank: king},
	{suit: hearts, rank: ace},
	{suit: hearts, rank: two},
	{suit: hearts, rank: three},
	{suit: hearts, rank: four},
	{suit: hearts, rank: five},
	{suit: hearts, rank: six},
	{suit: hearts, rank: seven},
	{suit: hearts, rank: eight},
	{suit: hearts, rank: nine},
	{suit: hearts, rank: ten},
	{suit: hearts, rank: jack},
	{suit: hearts, rank: queen},
	{suit: hearts, rank: king},
	{suit: spades, rank: ace},
	{suit: spades, rank: two},
	{suit: spades, rank: three},
	{suit: spades, rank: four},
	{suit: spades, rank: five},
	{suit: spades, rank: six},
	{suit: spades, rank: seven},
	{suit: spades, rank: eight},
	{suit: spades, rank: nine},
	{suit: spades, rank: ten},
	{suit: spades, rank: jack},
	{suit: spades, rank: queen},
	{suit: spades, rank: king},
	{rank: jokerA},
	{rank: jokerB},
}
