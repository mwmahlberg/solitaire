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

import "fmt"

type Rank int

func (r *Rank) String() string {
	switch *r {
	case ace:
		return "Ace"
	case two:
		return "2"
	case three:
		return "3"
	case four:
		return "4"
	case five:
		return "5"
	case six:
		return "6"
	case seven:
		return "7"
	case eight:
		return "8"
	case nine:
		return "9"
	case ten:
		return "10"
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

func (r *Rank) Short() string {
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
	ace Rank = iota + 1
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

	jokerA Rank = 53
	jokerB Rank = 54
)

type Suit int

const (
	Clubs    Suit = 0
	Diamonds Suit = 13
	Hearts   Suit = 26
	Spades   Suit = 39
)

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
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

func (s Suit) Short() string {
	switch s {
	case Clubs:
		return "C"
	case Diamonds:
		return "D"
	case Hearts:
		return "H"
	case Spades:
		return "S"
	default:
		panic("invalid suit")
	}
}

func (s Suit) Value(rank int) int {
	if rank < 1 || rank > 13 {
		panic("card must be between 1 and 13")
	}
	if (rank+int(s))%26 == 0 {
		return 26
	}
	return (rank + int(s)) % 26
}

type Card struct {
	suit Suit
	rank Rank
}

func (c Card) IsJokerA() bool {
	return c.rank == jokerA
}

func (c Card) IsJokerB() bool {
	return c.rank == jokerB
}

func (c Card) IsJoker() bool {
	return c.IsJokerA() || c.IsJokerB()
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) Rank() Rank {
	return c.rank
}

func (c Card) Value() int {
	if c.IsJoker() {
		return 53
	}
	return int(c.suit) + int(c.rank)
}

func (c Card) String() string {
	if c.IsJokerA() {
		return "Joker A"
	}
	if c.IsJokerB() {
		return "Joker B"
	}

	return fmt.Sprintf("%s %s", c.suit.String(), c.rank.Short())
}

func (c Card) Short() string {
	if c.IsJokerA() {
		return "JA"
	}
	if c.IsJokerB() {
		return "JB"
	}

	return fmt.Sprintf("%s%s", c.suit.Short(), c.rank.Short())
}

var initialDeck = []Card{
	{suit: Clubs, rank: ace},
	{suit: Clubs, rank: two},
	{suit: Clubs, rank: three},
	{suit: Clubs, rank: four},
	{suit: Clubs, rank: five},
	{suit: Clubs, rank: six},
	{suit: Clubs, rank: seven},
	{suit: Clubs, rank: eight},
	{suit: Clubs, rank: nine},
	{suit: Clubs, rank: ten},
	{suit: Clubs, rank: jack},
	{suit: Clubs, rank: queen},
	{suit: Clubs, rank: king},
	{suit: Diamonds, rank: ace},
	{suit: Diamonds, rank: two},
	{suit: Diamonds, rank: three},
	{suit: Diamonds, rank: four},
	{suit: Diamonds, rank: five},
	{suit: Diamonds, rank: six},
	{suit: Diamonds, rank: seven},
	{suit: Diamonds, rank: eight},
	{suit: Diamonds, rank: nine},
	{suit: Diamonds, rank: ten},
	{suit: Diamonds, rank: jack},
	{suit: Diamonds, rank: queen},
	{suit: Diamonds, rank: king},
	{suit: Hearts, rank: ace},
	{suit: Hearts, rank: two},
	{suit: Hearts, rank: three},
	{suit: Hearts, rank: four},
	{suit: Hearts, rank: five},
	{suit: Hearts, rank: six},
	{suit: Hearts, rank: seven},
	{suit: Hearts, rank: eight},
	{suit: Hearts, rank: nine},
	{suit: Hearts, rank: ten},
	{suit: Hearts, rank: jack},
	{suit: Hearts, rank: queen},
	{suit: Hearts, rank: king},
	{suit: Spades, rank: ace},
	{suit: Spades, rank: two},
	{suit: Spades, rank: three},
	{suit: Spades, rank: four},
	{suit: Spades, rank: five},
	{suit: Spades, rank: six},
	{suit: Spades, rank: seven},
	{suit: Spades, rank: eight},
	{suit: Spades, rank: nine},
	{suit: Spades, rank: ten},
	{suit: Spades, rank: jack},
	{suit: Spades, rank: queen},
	{suit: Spades, rank: king},
	{rank: jokerA},
	{rank: jokerB},
}
