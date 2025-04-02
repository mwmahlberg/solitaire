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

type alphabet [26]byte

// defaultAlphabet is the default alphabet used in the algorithm.
// It is the English alphabet in uppercase letters.
// Theoretically, any alphabet can be used.
// However, the algorithm is designed to work with the English alphabet.
// The alphabet is used to map the letters to numbers.
// The mapping is as follows:
// A = 1, B = 2, C = 3, ..., Z = 26
var DefaultAlphabet = alphabet{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z'}

// Index returns the index of the byte in the alphabet.
// If the byte is not in the alphabet, it returns -1.
// The index is 0-based, so A = 0, B = 1, C = 2, ..., Z = 25 for the
// default alphabet.
// The index is 1-based for the algorithm, so A = 1, B = 2, C = 3, ..., Z = 26.
func (t alphabet) Index(b byte) int {
	for i, c := range t {
		if b == c {
			return i
		}
	}
	return -1
}

func (t alphabet) Char(index int) byte {
	index = index % 26
	if index == 0 {
		index = 26
	}
	return DefaultAlphabet[index-1]
}
