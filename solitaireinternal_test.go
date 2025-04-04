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

	"github.com/awnumar/memguard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SolitaireSuite struct {
	suite.Suite
}

func (s *SolitaireSuite) TestPassphraseFromLockedBuffer() {
	// Create a locked buffer with the passphrase
	buf := memguard.NewBufferFromBytes([]byte("CRYPTONOMICON"))
	f := WithPassphraseFromLockedBuffer(buf)
	solitaireInstance := &Solitaire{}
	err := f(solitaireInstance)
	assert.NoError(s.T(), err, "Expected no error with valid buffer")
}

func (s *SolitaireSuite) TestPassphraseFromLockedBufferWithNilBuffer() {
	// Create a nil buffer
	var buf *memguard.LockedBuffer
	f := WithPassphraseFromLockedBuffer(buf)
	solitaireInstance := &Solitaire{}
	err := f(solitaireInstance)
	assert.Error(s.T(), err, "Expected error with nil buffer")
	assert.Equal(s.T(), ErrNilBuffer, err, "Expected ErrNilBuffer error")
}

func (s *SolitaireSuite) TestPassphraseFromDestroyedLockedBuffer() {
	// Create a destroyed buffer
	buf := memguard.NewBufferFromBytes([]byte("CRYPTONOMICON"))
	buf.Destroy()
	f := WithPassphraseFromLockedBuffer(buf)
	solitaireInstance := &Solitaire{}
	err := f(solitaireInstance)
	assert.Error(s.T(), err, "Expected error with destroyed buffer")
	assert.Equal(s.T(), ErrBufferNotAlive, err, "Expected ErrBufferNotAlive error")
}

func (s *SolitaireSuite) TestWithPassphraseFromEnclave() {
	passphrase := memguard.NewEnclave([]byte("CRYPTONOMICON"))
	passphrase.Open()
	f := WithPassphraseFromEnclave(passphrase)
	solitaireInstance := &Solitaire{}
	err := f(solitaireInstance)
	assert.NoError(s.T(), err, "Expected no error when passphrase is valid")
}

func (s *SolitaireSuite) TestWithPassphraseFromNilEnclave() {
	f := WithPassphraseFromEnclave(nil)
	solitaireInstance := &Solitaire{}
	err := f(solitaireInstance)
	assert.Error(s.T(), err, "Expected error when passphrase is nil")
}

func (s *SolitaireSuite) TestWithPassphraseFromInvalidEnclave() {
	passphrase := &memguard.Enclave{}
	f := WithPassphraseFromEnclave(passphrase)
	solitaireInstance := &Solitaire{}

	s.Panics(func() {
		f(solitaireInstance)
	}, "Expected panic when enclave is invalid")
}

func TestSolitaireInternal(t *testing.T) {
	suite.Run(t, new(SolitaireSuite))
}
