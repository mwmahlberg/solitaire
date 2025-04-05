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

// Solitaire is an [encryption algorithm created by Bruce Schneier in 1999].
// It is a simple algorithm that uses a deck of cards to generate a keystream.
// The keystream is then combined with the plaintext to produce the ciphertext.
// // This package provides an implementation of the Solitaire algorithm, to be used both as
// a library and as a command line tool. You can find the latter under [github.com/mwmahlberg/solitaire/cmd/solitaire].
//
// # Disclaimer
//
// Neither the package nor the command line tool are intended to be used for production use, but rather as a
// demonstration of the algorithm.
// The implementation is not optimized for speed or memory usage, but rather for
// readability and simplicity.
//
// # The algorithm
//
// The algorithm works by using a deck of cards to generate a keystream.
// First, the deck is optionally brought into a known state using a passphrase.
// Then, all non-letter (spaces, colons, etc) characters
// are removed from the plaintext, and the remaining characters are converted to uppercase
// and padded with uppercase Xs to make its length a multiple of 5. The result is the cleartext.
// Each character in the cleartext has a value: A=1, B=2, C=3, ..., Z=26.
//
// For each character in the cleartext, a new key in the form of a card is generated from the deck.
// Each card has a value, the sum of their suits value and their rank (Ace=1,2,3,...,King=13).
//   - Clubs: 0
//   - Diamonds: 13
//   - Hearts: 26
//   - Spades: 39
//
// (Neither of the jokers are used for encrypting, they are just important for the key generation algorithm.)
//
// For example, the value of the 2 of Hearts is 26 + 2 = 28, the value of the ace of Spades is 39 + 1 = 40, and so on.
//
// For each character in the cleartext, the a key is generated from the deck, its value is added to the value of the character, and the result is taken modulo 26.
// Say the cleartext character is A (value 1) and the generated key is the Ace of Spades (value 40).
// The result is 1 + 40 = 41, and 41 mod 26 = 15. The ciphertext character is the character with value 15, which is O.
// Note that generating a key modifies the deck, as Solitaire was designed to be a output-feedback mode stream cipher.
// This basically means that for every key generated, the deck is modified, and the next key will be different.
//
// [encryption algorithm created by Bruce Schneier in 1999]: https://schneier.com/academic/solitaire
package solitaire
