package solitaire

import "fmt"

type solitaire struct {
	// The deck of cards used in the Solitaire encryption algorithm.
	deck *Deck
}

type SolitaireOption func(*solitaire) error

func WithPassphrase(passphrase string) SolitaireOption {
	return func(s *solitaire) error {
		// if passphrase == "" {
		// 	return fmt.Errorf("passphrase cannot be empty")
		// }
		s.deck = &Deck{}
		copy(s.deck[:], initialDeck)

		// Set the position to 0
		for _, c := range []byte(passphrase) {
			s.deck.Advance()
			s.deck.countCut(findCharIndex(c) + 1)
		}
		return nil
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

	// Format the ciphertext with a space between every 5 characters
	// and a newline after every four groups.
	formatted := make([]byte, 0, len(ct)+len(ct)/5+len(ct)/20)
	groupCount := 0
	for i, c := range ct {
		if i > 0 && i%5 == 0 {
			groupCount++
			if groupCount%4 == 0 {
				formatted = append(formatted, '\n')
			} else {
				formatted = append(formatted, ' ')
			}
		}
		formatted = append(formatted, c)
	}
	ct = formatted
	return ct, nil
}

func (s *solitaire) Decrypt(ciphertext []byte) ([]byte, error) {
	// Normalize the ciphertext by removing spaces and converting to uppercase.
	if len(ciphertext) == 0 || len(ciphertext)%5 != 0 {
		// If the ciphertext is empty or not a multiple of 5, PANIC!
		panic("ciphertext must be a non-empty multiple of 5")
	}
	// Generate the keystream
	keys := s.generateKeyStream(len(ciphertext))

	// Decrypt the ciphertext using the keystream.
	ct := make([]byte, len(ciphertext))
	for i, c := range ciphertext {
		n := findCharIndex(c)
		key := keys[i]
		idx := (n - key + 1) % len(alphabet)
		if idx < 0 {
			idx += len(alphabet)
		}
		ct[i] = findCharByIndex(idx)
	}
	return ct, nil
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
