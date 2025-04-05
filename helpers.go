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
	"regexp"
)

const aUmlaut = 228
const oUmlaut = 246
const uUmlaut = 252
const capitalAumlaut = 196
const capitalOumlaut = 214
const capitalUumlaut = 220
const sharpS = 223

var mappings = map[byte][]byte{
	'Ä': {'A', 'E'},
	'Ö': {'O', 'E'},
	'Ü': {'U', 'E'},
	'ä': {'A', 'E'},
	'ö': {'O', 'E'},
	'ü': {'U', 'E'},
	'ß': {'S', 'S'},
}

func normalizeCleartext(plaintext []byte) []byte {

	// Normalize the plaintext by removing non-alphabetic characters
	// and converting to uppercase.
	normalized := make([]byte, 0, len(plaintext))
	for _, b := range plaintext {
		switch {
		case b == aUmlaut || b == oUmlaut || b == uUmlaut:
			fallthrough
		case b == capitalAumlaut || b == capitalOumlaut || b == capitalUumlaut:
			fallthrough
		case b == sharpS:
			// Handle umlauts and sharp S
			// Convert umlaut into their base letters with an appended e, uppercased
			// e.g. ä -> AE, ö -> OE, ü -> UE
			// Convert capital umlaut into their base letters with an appended e, uppercased
			// e.g. Ä -> AE, Ö -> OE, Ü -> UE
			// Convert sharp S into SS, uppercased
			// e.g. ß -> SS
			// Append the corresponding mapping
			if mapping, exists := mappings[b]; exists {
				normalized = append(normalized, mapping...)
			}
		case b >= 'A' && b <= 'Z':
			// Uppercase letters are added directly
			normalized = append(normalized, b)
		case b >= 'a' && b <= 'z':
			// Lowercase letters are converted to uppercase
			normalized = append(normalized, b-'a'+'A')
		default:
			// Ignore non-alphabetic characters
			// e.g. punctuation, numbers, etc.
		}
	}
	return normalized
}

var nonLetters = regexp.MustCompile(`[^\p{L}]+`)

func padClearText(plaintext []byte) []byte {
	// Pad the plaintext to a multiple of 5
	trimmed := nonLetters.ReplaceAll(plaintext, []byte(""))
	padLength := (5 - len(trimmed)%5) % 5
	if padLength == 0 {
		return trimmed
	}
	padded := make([]byte, len(trimmed)+padLength)
	copy(padded, trimmed)
	for i := len(trimmed); i < len(padded); i++ {
		padded[i] = 'X' // Padding character
	}
	return padded
}

func blocksOfFive(s []byte) []byte {

	// This is a bit of a hack, but it works.
	// Format the input with a space between every 5 characters
	// and a newline after every four groups.
	formatted := make([]byte, 0, len(s)+len(s)/5+len(s)/20)
	groupCount := 0
	for i, c := range s {
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
	return formatted
}
