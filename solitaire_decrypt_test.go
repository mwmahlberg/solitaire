package solitaire_test

import (
	"fmt"

	"github.com/mwmahlberg/solitaire"
)

// This shows how to decrypt a ciphertext using the solitaire
// package. The passphrase is used to key the deck into the same
// initial configuration as used before encryption. The
// ciphertext is then decrypted using the same algorithm as
// used for encryption. The result is the original plaintext.
// Note however that during encryption the plaintext is padded to a
// multiple of 5 characters with X. This means that the decrypted
// ciphertext may contain trailing X characters. The
// [solitaire.Solitaire.Decrypt] function will not remove these characters, as it has
// no way of knowing how many characters were added.
func Example_decrypt() {
	// Create a new solitaire instance with the passphrase "CRYPTONOMICON"
	s, err := solitaire.New(solitaire.WithPassphrase([]byte("CRYPTONOMICON")))
	if err != nil {
		panic(err)
	}
	cleartext, err := s.Decrypt([]byte("KIRAK SFJAN"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", cleartext)
	// Output: SOLIT AIREX

}
