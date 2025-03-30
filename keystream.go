package solitaire

import (
	"fmt"

	"github.com/awnumar/memguard"
)

type card int
type joker int
type jokerPosition int

// type jokerColor int

const (
	one card = iota + 1
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
	ace
	jokerA card = 53
	jokerB card = 54
)

var matrix = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z'}

func findCharIndex(b byte) int {
	for i, c := range matrix {
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
	return matrix[i-1]
}

type color int

const (
	clubs    color = 0
	diamonds color = 13
	hearts   color = 26
	spades   color = 39
	jokers   color = 52
)

func (c color) String() string {
	switch c {
	case clubs:
		return "C"
	case diamonds:
		return "D"
	case hearts:
		return "H"
	case spades:
		return "S"
	case jokers:
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

func (c Card) Value() int {
	//TODO: value of jokers
	if c.color == jokers {
		return 53
	}
	return int(c.color) + int(c.card)
}

var initialDeck = []Card{
	{color: clubs, card: one},
	{color: clubs, card: two},
	{color: clubs, card: three},
	{color: clubs, card: four},
	{color: clubs, card: five},
	{color: clubs, card: six},
	{color: clubs, card: seven},
	{color: clubs, card: eight},
	{color: clubs, card: nine},
	{color: clubs, card: ten},
	{color: clubs, card: jack},
	{color: clubs, card: queen},
	{color: clubs, card: king},
	{color: diamonds, card: one},
	{color: diamonds, card: two},
	{color: diamonds, card: three},
	{color: diamonds, card: four},
	{color: diamonds, card: five},
	{color: diamonds, card: six},
	{color: diamonds, card: seven},
	{color: diamonds, card: eight},
	{color: diamonds, card: nine},
	{color: diamonds, card: ten},
	{color: diamonds, card: jack},
	{color: diamonds, card: queen},
	{color: diamonds, card: king},
	{color: hearts, card: one},
	{color: hearts, card: two},
	{color: hearts, card: three},
	{color: hearts, card: four},
	{color: hearts, card: five},
	{color: hearts, card: six},
	{color: hearts, card: seven},
	{color: hearts, card: eight},
	{color: hearts, card: nine},
	{color: hearts, card: ten},
	{color: hearts, card: jack},
	{color: hearts, card: queen},
	{color: hearts, card: king},
	{color: spades, card: one},
	{color: spades, card: two},
	{color: spades, card: three},
	{color: spades, card: four},
	{color: spades, card: five},
	{color: spades, card: six},
	{color: spades, card: seven},
	{color: spades, card: eight},
	{color: spades, card: nine},
	{color: spades, card: ten},
	{color: spades, card: jack},
	{color: spades, card: queen},
	{color: spades, card: king},
	{color: jokers, card: jokerA},
	{color: jokers, card: jokerB},
}

func setUp(passphrase string) {
	deck := Deck{}
	// Initialize the deck with the initial order
	for i := 0; i < len(initialDeck); i++ {
		deck.order[i] = initialDeck[i]
	}
	// // Set the position to the first card
	// deck.SetPosition(deck.FindRedJoker())
	// deck.MoveCurrent(1)
	// for n, p := range deck.order {
	// 	fmt.Printf("%d: %s%d\n", n, p.color.String(), p.Value())
	// }
	// bj := deck.FindBlackJoker()
	// fmt.Printf("Black joker is at %d\n", bj)
	// deck.SetPosition(bj)
	// // skip two cards
	// deck.MoveCurrent(2)
	// for n, p := range deck.order {
	// 	fmt.Printf("%d: %s%d\n", n, p.color.String(), p.Value())
	// }
	// deck.TripleCut()
	// deck.CountCut()
	// val := deck.order[deck.order[0].Value()].Value()
	// fmt.Printf("Value of the card at position %d is %d\n", deck.order[0].Value(), val)

	// deck.SetPosition(deck.FindRedJoker())
	// deck.MoveCurrent(1)
	// deck.SetPosition(deck.FindBlackJoker())
	// deck.MoveCurrent(2)
	// deck.TripleCut()
	// deck.CountCut()
	// val = deck.order[deck.order[0].Value()].Value()
	// fmt.Printf("Value of the card at position %d is %d\n", deck.order[0].Value(), val)

	plaintext := []byte("AAAAAAAAAA")
	// Normalize the plaintext by removing non-alphabetic characters
	// and converting to uppercase.
	normalized := normalizeCleartext(plaintext)
	keys := make([]int, 0)

	for i := 0; len(keys) < len(normalized); i++ {
		deck.SetPosition(deck.FindRedJoker())
		deck.MoveCurrent(1)
		deck.SetPosition(deck.FindBlackJoker())
		deck.MoveCurrent(2)
		deck.TripleCut()
		deck.CountCut()
		val := deck.order[deck.order[0].Value()].Value()
		if val == 53 {
			// Skip the joker
			continue
		}
		keys = append(keys, val)
	}
	ct := make([]byte, len(normalized))
	// Print the keys
	for i, c := range normalized {
		n := findCharIndex(c)
		key := keys[i]
		idx := (n + key + 1)
		ct[i] = findCharByIndex(idx)
	}

	fmt.Printf("Cleartext:\t%s\n", string(normalized))
	// Print the ciphertext
	txt := fmt.Sprintf("Ciphertext:\t%s\n", string(ct))
	fmt.Println(txt)

}

func setupWithEnclave(e *memguard.Enclave) {
	l, err := e.Open()
	if err != nil {
		panic(err)
	}
	defer l.Destroy()

	trueSetup(l.String())

}

func trueSetup(passphrase string) {
	// Initialize the deck with the initial order
	deck := Deck{}
	copy(deck.order[:], initialDeck)

	for _, c := range []byte(passphrase) {
		deck.SetPosition(deck.FindRedJoker())
		deck.MoveCurrent(1)
		deck.SetPosition(deck.FindBlackJoker())
		deck.MoveCurrent(2)
		deck.TripleCut()
		deck.CountCut()
		deck.countCut(findCharIndex(c) + 1)
	}

	clear := []byte("SOLITAIRE")
	normalized := normalizeCleartext(padClearText(clear))
	keys := make([]int, 0)
	for i := 0; len(keys) < len(normalized); i++ {
		deck.SetPosition(deck.FindRedJoker())
		deck.MoveCurrent(1)
		deck.SetPosition(deck.FindBlackJoker())
		deck.MoveCurrent(2)
		deck.TripleCut()
		deck.CountCut()
		val := deck.order[deck.order[0].Value()].Value()
		if val == 53 {
			// Skip the joker
			continue
		}
		keys = append(keys, val)
	}

	ct := make([]byte, len(normalized))
	// Print the keys
	for i, c := range normalized {
		n := findCharIndex(c)
		key := keys[i]
		idx := (n + key + 1)
		ct[i] = findCharByIndex(idx)
	}

	// Print the ciphertext
	fmt.Printf("Ciphertext: ")
	printCipherText(ct)
}
