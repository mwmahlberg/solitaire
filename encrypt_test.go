package solitaire

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EncryptionSuite struct {
	suite.Suite
}

func (s *EncryptionSuite) TestNormalize() {
	testCases := []struct {
		desc     string
		input    []byte
		expected []byte
	}{
		{
			desc:     "uppercase, lowercase, non-alphabetic characters",
			input:    []byte("Hello, World! 123"),
			expected: []byte("HELLOWORLD"),
		},
		{
			desc:     "German uppercase umlauts",
			input:    []byte{196, 214, 220}, // ÄÖÜ
			expected: []byte("AEOEUE"),
		},
		{
			desc:     "German lowercase umlauts and sharp S",
			input:    []byte{228, 246, 252, 223}, // äöüß
			expected: []byte("AEOEUESS"),
		},
		{
			desc:     "Empty input",
			input:    []byte(""),
			expected: []byte(""),
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			normalized := normalizeCleartext(tC.input)
			if string(normalized) != string(tC.expected) {
				s.Failf(tC.desc, "Expected %s, got %s", tC.expected, normalized)
			}
		})
	}
}

func (s *EncryptionSuite) TestPadding() {
	testCases := []struct {
		desc     string
		input    []byte
		expected []byte
	}{
		{
			desc:     "no padding needed",
			input:    []byte("HELLOWORLD"),
			expected: []byte("HELLOWORLD"),
		},
		{
			desc:     "padding needed",
			input:    []byte("HI"),
			expected: []byte("HIXXX"),
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			padded := padClearText(tC.input)
			if string(padded) != string(tC.expected) {
				s.Failf(tC.desc, "Expected %s, got %s", tC.expected, padded)
			}
		})
	}
}

func TestEncryptionSuite(t *testing.T) {
	suite.Run(t, new(EncryptionSuite))
}
