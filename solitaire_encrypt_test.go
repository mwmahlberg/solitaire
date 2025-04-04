package solitaire_test

import (
	"errors"
	"fmt"

	"github.com/mwmahlberg/solitaire"
)

func Example_encrypt() {
	// Create a new solitaire instance with the passphrase "CRYPTONOMICON"
	s, err := solitaire.New(solitaire.WithPassphrase([]byte("CRYPTONOMICON")))
	switch err {
	case solitaire.ErrPassphraseIsNil:
		// Handle the case when the passphrase is nil. A passphrase should
		// never be nil, but may be empty, that is []byte("").
		//
		// If you want to use the default deck order, you can instead use the
		// solitaire.New() function without any options.

		fallthrough
	case solitaire.ErrInvalidPassphrase:
		// Handle the case when the passphrase contains a non-ASCII character.
		// The passphrase should be a valid ASCII string. The reason for this
		// is that the algorithm uses the ASCII values of the characters to
		// generate the deck order. If the passphrase contains non-ASCII
		// characters, the algorithm may now work as expected. The usage of
		// non-ASCII characters is not tested and hence not supported for now.

		panic(errors.Join(errors.New("cannot create solitaire instance"), err))
	}
	ciphertext, err := s.Encrypt([]byte("SOLITAIRE"))
	if err != nil {
		panic(errors.Join(errors.New("cannot encrypt"), err))
	}
	fmt.Printf("%s", ciphertext)
	// Output: KIRAK SFJAN
}
