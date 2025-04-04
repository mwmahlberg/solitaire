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

package solitaire_test

import (
	"errors"
	"fmt"

	"github.com/awnumar/memguard"
	"github.com/mwmahlberg/solitaire"
)

func ExampleWithPassphraseFromLockedBuffer() {
	// Create a new memguard.LockedBuffer
	lockedBuffer := memguard.NewBufferFromBytes([]byte("CRYPTONOMICON"))

	s, err := solitaire.New(solitaire.WithPassphraseFromLockedBuffer(lockedBuffer))

	switch err {
	case solitaire.ErrNilBuffer:
		// Handle a nil buffer
		// For this example, we will just panic
		fallthrough
	case solitaire.ErrBufferNotAlive:
		// Handle the case when the buffer was destroyed prior to
		// calling the New function. For this example, we will just panic.
		fallthrough
	case solitaire.ErrPassphraseIsNil:
		// Handle the case when the passphrase is nil. A passphrase should
		// never be nil, but may be empty, that is []byte("").
		//
		// If you do not want to use a passphrase keying the deck,
		// you can instead use solitaire.New() without any options.
		fallthrough
	case solitaire.ErrInvalidPassphrase:
		// Handle the case when the passphrase contains a non-ASCII character.
		// For this example, we will just panic.
		panic(errors.Join(errors.New("cannot create solitaire instance"), err))
	}
	ciphertext, err := s.Encrypt([]byte("SOLITAIRE"))
	if err != nil {
		panic(errors.Join(errors.New("cannot encrypt"), err))
	}
	fmt.Printf("%s", ciphertext)
	// Output: KIRAK SFJAN
}
