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

	JokerA card = 53
	JokerB card = 54
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

type color int

const (
	Clubs    color = 0
	Diamonds color = 13
	Hearts   color = 26
	Spades   color = 39
	Jokers   color = 52
)

func (c color) String() string {
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

func (c color) Value(card int) int {
	if card < 1 || card > 13 {
		panic("card must be between 1 and 13")
	}
	if (card+int(c))%26 == 0 {
		return 26
	}
	return (card + int(c)) % 26
}

type Card struct {
	color color
	card  card
}

func (c Card) IsJokerA() bool {
	return c.card == JokerA
}

func (c Card) IsJokerB() bool {
	return c.card == JokerB
}

func (c Card) Color() color {
	return c.color
}

func (c Card) Face() int {
	return int(c.card)
}

func (c Card) Value() int {
	//TODO: value of jokers
	if c.color == Jokers {
		return 53
	}
	return int(c.color) + int(c.card)
}

func (c Card) String() string {
	if c.card == JokerA {
		return "Joker A"
	}
	if c.card == JokerB {
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
	return fmt.Sprintf("%s %s", c.color.String(), name)
}

var initialDeck = []Card{
	{color: Clubs, card: ace},
	{color: Clubs, card: two},
	{color: Clubs, card: three},
	{color: Clubs, card: four},
	{color: Clubs, card: five},
	{color: Clubs, card: six},
	{color: Clubs, card: seven},
	{color: Clubs, card: eight},
	{color: Clubs, card: nine},
	{color: Clubs, card: ten},
	{color: Clubs, card: jack},
	{color: Clubs, card: queen},
	{color: Clubs, card: king},
	{color: Diamonds, card: ace},
	{color: Diamonds, card: two},
	{color: Diamonds, card: three},
	{color: Diamonds, card: four},
	{color: Diamonds, card: five},
	{color: Diamonds, card: six},
	{color: Diamonds, card: seven},
	{color: Diamonds, card: eight},
	{color: Diamonds, card: nine},
	{color: Diamonds, card: ten},
	{color: Diamonds, card: jack},
	{color: Diamonds, card: queen},
	{color: Diamonds, card: king},
	{color: Hearts, card: ace},
	{color: Hearts, card: two},
	{color: Hearts, card: three},
	{color: Hearts, card: four},
	{color: Hearts, card: five},
	{color: Hearts, card: six},
	{color: Hearts, card: seven},
	{color: Hearts, card: eight},
	{color: Hearts, card: nine},
	{color: Hearts, card: ten},
	{color: Hearts, card: jack},
	{color: Hearts, card: queen},
	{color: Hearts, card: king},
	{color: Spades, card: ace},
	{color: Spades, card: two},
	{color: Spades, card: three},
	{color: Spades, card: four},
	{color: Spades, card: five},
	{color: Spades, card: six},
	{color: Spades, card: seven},
	{color: Spades, card: eight},
	{color: Spades, card: nine},
	{color: Spades, card: ten},
	{color: Spades, card: jack},
	{color: Spades, card: queen},
	{color: Spades, card: king},
	{color: Jokers, card: JokerA},
	{color: Jokers, card: JokerB},
}
