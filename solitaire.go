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
 *  SPDX-License-Identifier: Apache-2.0
 *
 */

package solitaire

import (
	"fmt"
	"regexp"

	"github.com/awnumar/memguard"
)

// Passphrase related errors.
// Errors returned by file systems can be tested against these errors
// using errors.Is.
var (
	// returned when the [memguard.LockedBuffer] is nil
	ErrNilBuffer = fmt.Errorf("buffer is nil")

	// returned when the [memguard.LockedBuffer] is not alive
	ErrBufferNotAlive = fmt.Errorf("buffer is not alive")

	// returned when the passphrase is nil
	// Note however that the passphrase may be empty.
	ErrPassphraseIsNil = fmt.Errorf("passphrase is required")

	// returned when the passphrase contains non-letter characters
	ErrInvalidPassphrase = fmt.Errorf("passphrase contains non-letter characters")
)

var (
	uppercaseAscii = regexp.MustCompile("^[[:upper:]]*$")
)

type Solitaire struct {
	// The deck of cards used in the Solitaire encryption algorithm.
	deck     *deck
	alphabet alphabet
}

type SolitaireOption func(*Solitaire) error

// WithPassphrase takes a passphrase and sets the deck.
// It is the caller's responsibility to ensure that the passphrase is valid.
// If the passphrase is nil, WithPassphrase will panic with [ErrPassphraseIsNil].
//
// However, the passphrase may be empty.
// In this case, the deck will be set to the initial deck:
// ace to king of clubs, diamonds, hearts, and spades, followed by the two jokers.
//
// It is the caller's responsibility to ensure that a proper passphrase is used
// in the appropriate context. For simple tests, an empty passphrase might even
// be required to trace the examples given in [Bruce Schneier's original blog post].
//
// However, it is highly recommended to use a proper passphrase to ensure the
// security of the encryption. The passphrase should be a string of uppercase
// ASCII letters (A-Z) and should not contain any spaces or other characters.
//
// [Bruce Schneier's original blog post]: https://www.schneier.com/academic/solitaire/
func WithPassphrase(passphrase []byte) SolitaireOption {
	return func(s *Solitaire) error {

		if passphrase == nil {
			memguard.SafePanic(ErrPassphraseIsNil)
		}
		if !uppercaseAscii.Match(passphrase) {
			return ErrInvalidPassphrase
		}
		s.deck = &deck{}
		copy(s.deck[:], initialDeck)

		// Set the position to 0
		for _, c := range passphrase {
			s.deck.advance()
			s.deck.countCut(s.alphabet.Index(c) + 1)
		}
		return nil
	}
}

// WithPassphraseFromLockedBuffer takes a passphrase from a [memguard.LockedBuffer] and sets the deck.
// It is a convenience function for use with the [memguard] package.
// If the passphrase is nil, it returns an error.
// If the memguard.LockedBuffer is not Alive, WithPassphraseFromLockedBuffer will panic.
// It is the caller's responsibility to ensure that the buf is valid and not empty.
// The passphrase stored within the memguard.LockedBuffer is used to set the deck.
//
// Returns [ErrNilBuffer] if the buffer is nil or [ErrBufferNotAlive] if the buffer is not alive.
func WithPassphraseFromLockedBuffer(buf *memguard.LockedBuffer) SolitaireOption {
	return func(s *Solitaire) error {

		if buf == nil {
			return ErrNilBuffer
		}
		if !buf.IsAlive() {
			return ErrBufferNotAlive
		}

		return WithPassphrase(buf.Bytes())(s)
	}
}

// WithPassphraseFromEnclave takes a passphrase from a [memguard.Enclave] and sets the deck.
// It is a convenience function for use with the memguard package.
// if the passphrase is nil, it returns an error.
// If the [memguard.Enclave] cannot be opened, WithPassphraseFromLockedBuffer will panic.
// It is the caller's responsibility to ensure that the passphrase is valid and not empty.
// The passphrase stored within the [memguard.LockedBuffer] stored in the enclave
// is used to set the deck.
func WithPassphraseFromEnclave(passphrase *memguard.Enclave) SolitaireOption {
	return func(s *Solitaire) error {
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

// New creates a new [Solitaire] instance.
// It takes a variadic number of [SolitaireOption] functions that modify the
// [Solitaire] instance.
// The options are applied in the order they are provided.
// If no options are provided, the default deck is used.
func New(opts ...SolitaireOption) (*Solitaire, error) {
	s := &Solitaire{
		alphabet: defaultAlphabet,
	}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	if s.deck == nil {
		s.deck = &deck{}
		copy(s.deck[:], initialDeck)
	}
	return s, nil
}

// Deck returns the current state of the deck.
// If export is true, the deck is returned in a format suitable for
// import.
func (s *Solitaire) Deck(export bool) []string {
	// Return a copy of the deck
	d := make([]string, len(s.deck))
	for i, c := range s.deck {
		if export {
			d[i] = c.Short()
			continue
		}
		d[i] = c.String()
	}
	return d
}

func (s *Solitaire) SetDeck(d []string) error {
	s.deck = new(deck)
	if len(d) != 54 {
		return fmt.Errorf("deck must be 54 cards, got %d", len(d))
	}
	var err error
	for i, str := range d {
		if s.deck[i], err = cardFromString(str); err != nil {
			return fmt.Errorf("invalid card %s: %w", str, err)
		}
	}
	return nil
}

// Encrypt encrypts the given plaintext using the Solitaire algorithm.
// It takes a byte slice as input and returns a byte slice as output.
//

func (s *Solitaire) Encrypt(plaintext []byte) ([]byte, error) {
	// Normalize the plaintext by removing spaces and converting to uppercase.
	normalized := padClearText(normalizeCleartext(plaintext))

	// The character at that index is used to encrypt the plaintext.
	ct := make([]byte, len(normalized))
	for i, c := range normalized {
		n := s.alphabet.Index(c)
		key := s.deck.key()
		idx := (n + key + 1) % len(s.alphabet)
		ct[i] = s.alphabet.Char(idx)
	}

	return blocksOfFive(ct), nil
}

// Decrypt decrypts the given ciphertext using the Solitaire algorithm.
// It takes the ciphertext as a byte slice and returns the cleartext as a byte slice.
//
// The ciphertext is normalized by removing all non-letter characters before decryption.
func (s *Solitaire) Decrypt(ciphertext []byte) ([]byte, error) {
	cleaned := nonLetters.ReplaceAll(ciphertext, []byte(""))
	// Normalize the ciphertext by removing spaces and converting to uppercase.
	if len(cleaned) == 0 || len(cleaned)%5 != 0 {
		// If the ciphertext is empty or not a multiple of 5, PANIC!
		panic("ciphertext must be a non-empty multiple of 5")
	}

	// Decrypt the ciphertext using the keystream.
	ct := make([]byte, len(cleaned))
	for i, c := range cleaned {
		n := s.alphabet.Index(c)
		key := s.deck.key()
		idx := (n - key + 1) % len(s.alphabet)
		if idx < 0 {
			idx += len(s.alphabet)
		}
		ct[i] = s.alphabet.Char(idx)
	}
	return blocksOfFive(ct), nil
}
