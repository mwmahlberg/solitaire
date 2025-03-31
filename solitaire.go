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

func WithPassphrase(passphrase []byte) SolitaireOption {
	return func(s *solitaire) error {

		s.deck = &Deck{}
		copy(s.deck[:], initialDeck)

		// Set the position to 0
		for _, c := range passphrase {
			s.deck.Advance()
			s.deck.countCut(findCharIndex(c) + 1)
		}
		return nil
	}
}

func WithPassphraseFromLockedBuffer(buf *memguard.LockedBuffer) SolitaireOption {
	return func(s *solitaire) error {
		if buf == nil {
			return fmt.Errorf("passphrase is required")
		}

		return WithPassphrase(buf.Bytes())(s)
	}
}

func WithPassphraseFromEnclave(passphrase *memguard.Enclave) SolitaireOption {
	return func(s *solitaire) error {
		if passphrase == nil {
			return fmt.Errorf("passphrase is required")
		}
		buf, err := passphrase.Open()
		if err != nil {
			return err
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
	keys := s.generateKeyStream(len(normalized))

	// Encrypt the plaintext using the keystream.
	// The keystream is used to determine the index of the character in the matrix.
	// The character at that index is used to encrypt the plaintext.
	ct := make([]byte, len(normalized))
	for i, c := range normalized {
		n := findCharIndex(c)
		key := keys[i]
		idx := (n + key + 1) % len(alphabet)
		ct[i] = findCharByIndex(idx)
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
	keys := s.generateKeyStream(len(cleaned))

	// Decrypt the ciphertext using the keystream.
	ct := make([]byte, len(cleaned))
	for i, c := range cleaned {
		n := findCharIndex(c)
		key := keys[i]
		idx := (n - key + 1) % len(alphabet)
		if idx < 0 {
			idx += len(alphabet)
		}
		ct[i] = findCharByIndex(idx)
	}
	return BlocksOfFive(ct), nil
}

func (s *solitaire) generateKeyStream(length int) []int {
	// Generate the keystream by moving the jokers and cutting the deck.
	keys := make([]int, 0)
	for i := 0; len(keys) < length; i++ {
		s.deck.Advance()
		val := s.deck[s.deck[0].Value()].Value()
		if val == 53 {
			// Skip the joker
			continue
		}
		keys = append(keys, val)
	}
	return keys
}
