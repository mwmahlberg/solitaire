package solitaire_test

import (
	"fmt"

	"github.com/mwmahlberg/solitaire"
)

// This shows how to create a new solitaire instance without
// a passphrase. This will use the default deck order to
// generate the key stream.

func Example() {
	// Create a new solitaire instance without a passphrase for
	// keying the deck. This will use the default deck order to
	// generate the key stream. Only use this for tests!
	s, err := solitaire.New()
	if err != nil {
		panic(err)
	}
	ciphertext, err := s.Encrypt([]byte("AAAAAAAAAA"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", ciphertext)
	// Output: EXKYI ZSGEH
}
