package solitaire_test

import (
	"fmt"

	"github.com/mwmahlberg/solitaire"
)

// This shows how to decrypt a ciphertext using the solitaire
// package. The passphrase is used to key the deck into the same
// initial configuration as used before encryption. The
// ciphertext is then decrypted:
//  1. A key is generated from the deck.
//  2. The value of the key is subtracted from the value of the
//     ciphertext character.
//  3. The result is taken modulo 26.
//  4. The result is the value of the plaintext character from the set of
//     1=A, 2=B, ..., 26=Z.
//
// This is done for each character in the ciphertext.
// The result is the cleartext, the plaintext
// stripped of all non-letter characters, converted to uppercase and
// padded to a multiple of 5 characters with X.This means that the decrypted
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
