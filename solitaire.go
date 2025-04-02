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

import (
	"fmt"

	"github.com/awnumar/memguard"
)

type solitaire struct {
	// The deck of cards used in the Solitaire encryption algorithm.
	deck *Deck
}

type SolitaireOption func(*solitaire) error

// WithPassphrase takes a passphrase and sets the deck.
// It is the caller's responsibility to ensure that the passphrase is valid.
// If the passphrase is nil, it returns an error. However, the passphrase may be empty.
// In this case, the deck will be set to the initial deck:
// ace to king of clubs, diamonds, hearts, and spades, followed by the two jokers.
// It is the caller's responsibility to ensure that a proper passphrase is used
// in the appropriate context.
func WithPassphrase(passphrase []byte) SolitaireOption {
	return func(s *solitaire) error {

		s.deck = &Deck{}
		copy(s.deck[:], initialDeck)

		// Set the position to 0
		for _, c := range passphrase {
			s.deck.Advance()
			s.deck.countCut(DefaultAlphabet.Index(c) + 1)
		}
		return nil
	}
}

// WithPassphraseFromLockedBuffer takes a passphrase from a memguard.LockedBuffer and sets the deck.
// It is a convenience function for use with the memguard package.
// If the passphrase is nil, it returns an error.
// If the memguard.LockedBuffer is not Alive, WithPassphraseFromLockedBuffer will panic.
// It is the caller's responsibility to ensure that the buf is valid and not empty.
// The passphrase stored within the memguard.LockedBuffer is used to set the deck.
func WithPassphraseFromLockedBuffer(buf *memguard.LockedBuffer) SolitaireOption {
	return func(s *solitaire) error {
		if buf == nil {
			return fmt.Errorf("passphrase is required")
		}

		return WithPassphrase(buf.Bytes())(s)
	}
}

// WithPassphraseFromEnclave takes a passphrase from a memguard.Enclave and sets the deck.
// It is a convenience function for use with the memguard package.
// if the passphrase is nil, it returns an error.
// If the memguard.Enclave cannot be opened, WithPassphraseFromLockedBuffer will panic.
// It is the caller's responsibility to ensure that the passphrase is valid and not empty.
// The passphrase stored within the memguard.LockedBuffer stored in the enclave
// is used to set the deck.
func WithPassphraseFromEnclave(passphrase *memguard.Enclave) SolitaireOption {
	return func(s *solitaire) error {
		if passphrase == nil {
			return fmt.Errorf("passphrase is required")
		}
		buf, err := passphrase.Open()
		if err != nil {
			memguard.SafePanic(err)
		}
		defer buf.Destroy()
		return WithPassphraseFromLockedBuffer(buf)(s)
	}
}

func New(opts ...SolitaireOption) (*solitaire, error) {
	s := &solitaire{}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	if s.deck == nil {
		return nil, fmt.Errorf("deck is required")
	}
	return s, nil
}

func (s *solitaire) Deck() []Card {
	// Return a copy of the deck
	d := make([]Card, len(s.deck))
	copy(d, s.deck[:])
	return d
}

func (s *solitaire) Encrypt(plaintext []byte) ([]byte, error) {
	// Normalize the plaintext by removing spaces and converting to uppercase.
	normalized := normalizeCleartext(padClearText(plaintext))
	keys := s.deck.generateKeyStream(len(normalized))

	// Encrypt the plaintext using the keystream.
	// The keystream is used to determine the index of the character in the matrix.
	// The character at that index is used to encrypt the plaintext.
	ct := make([]byte, len(normalized))
	for i, c := range normalized {
		n := DefaultAlphabet.Index(c)
		key := keys[i]
		idx := (n + key + 1) % len(DefaultAlphabet)
		ct[i] = DefaultAlphabet.Char(idx)
	}

	return BlocksOfFive(ct), nil
}

func (s *solitaire) Decrypt(ciphertext []byte) ([]byte, error) {
	cleaned := nonLetters.ReplaceAll(ciphertext, []byte(""))
	// Normalize the ciphertext by removing spaces and converting to uppercase.
	if len(cleaned) == 0 || len(cleaned)%5 != 0 {
		// If the ciphertext is empty or not a multiple of 5, PANIC!
		panic("ciphertext must be a non-empty multiple of 5")
	}
	// Generate the keystream
	keys := s.deck.generateKeyStream(len(cleaned))

	// Decrypt the ciphertext using the keystream.
	ct := make([]byte, len(cleaned))
	for i, c := range cleaned {
		n := DefaultAlphabet.Index(c)
		key := keys[i]
		idx := (n - key + 1) % len(DefaultAlphabet)
		if idx < 0 {
			idx += len(DefaultAlphabet)
		}
		ct[i] = DefaultAlphabet.Char(idx)
	}
	return BlocksOfFive(ct), nil
}
