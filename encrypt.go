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
