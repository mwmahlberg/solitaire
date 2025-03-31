package solitaire

import "fmt"

type rank int

const (
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

var alphabet = [26]byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z'}

func findCharIndex(b byte) int {
	for i, c := range alphabet {
		if b == c {
			return i
		}
	}
	return -1
}

func findCharByIndex(i int) byte {
	i = i % 26
	if i == 0 {
		i = 26
	}
	return alphabet[i-1]
}

type suit int

const (
	Clubs    suit = 0
	Diamonds suit = 13
	Hearts   suit = 26
	Spades   suit = 39
	Jokers   suit = 52
)

func (c suit) String() string {
	switch c {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	case Jokers:
		return "J"
	default:
		return ""
	}
}

func (c suit) Value(rank int) int {
	if rank < 1 || rank > 13 {
		panic("card must be between 1 and 13")
	}
	if (rank+int(c))%26 == 0 {
		return 26
	}
	return (rank + int(c)) % 26
}

type Card struct {
	suit suit
	rank rank
}

func (c Card) IsJokerA() bool {
	return c.rank == jokerA
}

func (c Card) IsJokerB() bool {
	return c.rank == jokerB
}

func (c Card) Suit() suit {
	return c.suit
}

func (c Card) Face() int {
	return int(c.rank)
}

func (c Card) Value() int {
	//TODO: value of jokers
	if c.suit == Jokers {
		return 53
	}
	return int(c.suit) + int(c.rank)
}

func (c Card) String() string {
	if c.rank == jokerA {
		return "Joker A"
	}
	if c.rank == jokerB {
		return "Joker B"
	}
	var name string = fmt.Sprintf("%d", c.rank)
	switch c.rank {
	case king:
		name = "K"
	case queen:
		name = "Q"
	case jack:
		name = "J"
	}
	return fmt.Sprintf("%s %s", c.suit.String(), name)
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
	{suit: Jokers, rank: jokerA},
	{suit: Jokers, rank: jokerB},
}
