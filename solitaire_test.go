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
	"fmt"
	"testing"

	"github.com/mwmahlberg/solitaire"
	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	testCases := []struct {
		passphrase string
		cleartext  string
		ciphertext string
	}{
		{
			passphrase: "CRYPTONOMICON",
			cleartext:  "SOLITAIRE",
			ciphertext: "KIRAK SFJAN",
		},
		{
			passphrase: "FOO",
			cleartext:  "AAAAAAAAAAAAAAA",
			ciphertext: "ITHZU JIWGR FARMW",
		},
		{
			passphrase: "",
			cleartext:  "AAAAAAAAAA",
			ciphertext: "EXKYI ZSGEH",
		},
		{
			passphrase: "CRYPTONOMICON",
			cleartext:  "HELLO WORLD Hello WORLD HELLO WORLD",
			ciphertext: "ZYRDF OLJHT YQIZV EDSQS\nEECJE FZXRN",
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%s(%s)->%s", tC.passphrase, tC.cleartext, tC.ciphertext), func(t *testing.T) {
			s, err := solitaire.New(solitaire.WithPassphrase([]byte(tC.passphrase)))
			assert.NoError(t, err, "Failed to create new solitaire instance")
			assert.NotNil(t, s, "Solitaire instance should not be nil")
			ct, err := s.Encrypt([]byte(tC.cleartext))
			assert.NoError(t, err, "Failed to encrypt plaintext")
			assert.NotNil(t, ct, "Ciphertext should not be nil")
			assert.Equal(t, tC.ciphertext, string(ct), "Ciphertext does not match expected value")
		})
	}
}

func TestDecryption(t *testing.T) {
	testCases := []struct {
		passphrase string
		cleartext  string
		ciphertext string
	}{
		{
			passphrase: "CRYPTONOMICON",
			cleartext:  "SOLIT AIREX",
			ciphertext: "KIRAK SFJAN",
		},
		{
			passphrase: "FOO",
			cleartext:  "AAAAA AAAAA AAAAA",
			ciphertext: "ITHZU JIWGR FARMW",
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%s(%s)->%s", tC.passphrase, tC.cleartext, tC.ciphertext), func(t *testing.T) {
			s, err := solitaire.New(solitaire.WithPassphrase([]byte(tC.passphrase)))
			assert.NoError(t, err, "Failed to create new solitaire instance")
			assert.NotNil(t, s, "Solitaire instance should not be nil")
			clear, err := s.Decrypt([]byte(tC.ciphertext))
			assert.NoError(t, err, "Failed to encrypt plaintext")
			assert.NotNil(t, clear, "Ciphertext should not be nil")
			assert.Equal(t, tC.cleartext, string(clear), "Ciphertext does not match expected value")
		})
	}
}

func TestSolitaire(t *testing.T) {
	s, err := solitaire.New(solitaire.WithPassphrase([]byte("CRYPTONOMICON")))
	assert.NoError(t, err, "Failed to create new solitaire instance")
	assert.NotNil(t, s, "Solitaire instance should not be nil")
	ct, err := s.Encrypt([]byte("SOLITAIRE"))
	assert.NoError(t, err, "Failed to encrypt plaintext")
	assert.NotNil(t, ct, "Ciphertext should not be nil")
	result := []byte("KIRAK SFJAN")
	assert.Equal(t, result, ct, "Ciphertext does not match expected value")

}
