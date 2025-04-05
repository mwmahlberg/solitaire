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
	"testing"

	"github.com/stretchr/testify/suite"
)

type HelperSuite struct {
	suite.Suite
}

func (s *HelperSuite) TestNormalize() {
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

func (s *HelperSuite) TestPadding() {
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

func TestHelpers(t *testing.T) {
	suite.Run(t, new(HelperSuite))
}
