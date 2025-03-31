package solitaire

import "fmt"

type card int

const (
	ace card = iota + 1
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

	jokerA card = 53
	jokerB card = 54
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

func (c suit) Value(card int) int {
	if card < 1 || card > 13 {
		panic("card must be between 1 and 13")
	}
	if (card+int(c))%26 == 0 {
		return 26
	}
	return (card + int(c)) % 26
}

type Card struct {
	suit suit
	card card
}

func (c Card) IsJokerA() bool {
	return c.card == jokerA
}

func (c Card) IsJokerB() bool {
	return c.card == jokerB
}

func (c Card) Suit() suit {
	return c.suit
}

func (c Card) Face() int {
	return int(c.card)
}

func (c Card) Value() int {
	//TODO: value of jokers
	if c.suit == Jokers {
		return 53
	}
	return int(c.suit) + int(c.card)
}

func (c Card) String() string {
	if c.card == jokerA {
		return "Joker A"
	}
	if c.card == jokerB {
		return "Joker B"
	}
	var name string = fmt.Sprintf("%d", c.card)
	switch c.card {
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
	{suit: Clubs, card: ace},
	{suit: Clubs, card: two},
	{suit: Clubs, card: three},
	{suit: Clubs, card: four},
	{suit: Clubs, card: five},
	{suit: Clubs, card: six},
	{suit: Clubs, card: seven},
	{suit: Clubs, card: eight},
	{suit: Clubs, card: nine},
	{suit: Clubs, card: ten},
	{suit: Clubs, card: jack},
	{suit: Clubs, card: queen},
	{suit: Clubs, card: king},
	{suit: Diamonds, card: ace},
	{suit: Diamonds, card: two},
	{suit: Diamonds, card: three},
	{suit: Diamonds, card: four},
	{suit: Diamonds, card: five},
	{suit: Diamonds, card: six},
	{suit: Diamonds, card: seven},
	{suit: Diamonds, card: eight},
	{suit: Diamonds, card: nine},
	{suit: Diamonds, card: ten},
	{suit: Diamonds, card: jack},
	{suit: Diamonds, card: queen},
	{suit: Diamonds, card: king},
	{suit: Hearts, card: ace},
	{suit: Hearts, card: two},
	{suit: Hearts, card: three},
	{suit: Hearts, card: four},
	{suit: Hearts, card: five},
	{suit: Hearts, card: six},
	{suit: Hearts, card: seven},
	{suit: Hearts, card: eight},
	{suit: Hearts, card: nine},
	{suit: Hearts, card: ten},
	{suit: Hearts, card: jack},
	{suit: Hearts, card: queen},
	{suit: Hearts, card: king},
	{suit: Spades, card: ace},
	{suit: Spades, card: two},
	{suit: Spades, card: three},
	{suit: Spades, card: four},
	{suit: Spades, card: five},
	{suit: Spades, card: six},
	{suit: Spades, card: seven},
	{suit: Spades, card: eight},
	{suit: Spades, card: nine},
	{suit: Spades, card: ten},
	{suit: Spades, card: jack},
	{suit: Spades, card: queen},
	{suit: Spades, card: king},
	{suit: Jokers, card: jokerA},
	{suit: Jokers, card: jokerB},
}
