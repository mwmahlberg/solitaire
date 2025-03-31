package main

import (
	"fmt"
	"strings"

	"github.com/awnumar/memguard"
	"github.com/mwmahlberg/solitaire"
)

type PrintDeck struct {
	Export bool `kong:"help='print the deck as a sequence suitable for importing'"`
}

func (p *PrintDeck) Run(pp passphrase) error {
	s, err := solitaire.New(solitaire.WithPassphraseFromEnclave(cfg.Passphrase.enc))
	if err != nil {
		memguard.SafePanic(err)
	}
	if p.Export {
		deck := make([]string, len(s.Deck()))
		for i, c := range s.Deck() {
			if c.IsJokerA() {
				deck[i] = "JA"
				continue
			}
			if c.IsJokerB() {
				deck[i] = "JB"
				continue
			}
			var color string
			switch c.Suit() {
			case solitaire.Clubs:
				color = "C"
			case solitaire.Diamonds:
				color = "D"
			case solitaire.Hearts:
				color = "H"
			case solitaire.Spades:
				color = "S"
			}
			deck[i] = fmt.Sprintf("%s%d", color, c.Rank())
		}
		fmt.Println(strings.Join(deck, ","))
		return nil
	}

	for i, c := range s.Deck() {
		fmt.Printf("%2d: %s\n", i+1, c.String())
	}
	return nil
}
