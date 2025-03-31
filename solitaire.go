package solitaire

import "fmt"

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
		// if passphrase == "" {
		// 	return fmt.Errorf("passphrase cannot be empty")
		// }
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
