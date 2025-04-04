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
//
// This package provides an implementation of the Solitaire algorithm, to be used both as
// a library and as a command line tool.
//
//
// [encryption algorithm created by Bruce Schneier in 1999]: https://schneier.com/academic/solitaire

package solitaire
