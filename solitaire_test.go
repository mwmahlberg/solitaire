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
	"strings"
	"testing"

	"github.com/mwmahlberg/solitaire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	setCRYPTONOMICON = "C7,C8,C9,D3,CQ,CK,DA,D2,SK,H4,D7,D8,D9,D10,DJ,DQ,DK,D4,C2,H5,H6,C5,H9,H10,HJ,HQ,H7,S2,S3,S4,S5,S6,S7,H8,SQ,JA,H2,S10,C6,D5,D6,HK,SA,S8,C10,CJ,HA,SJ,JB,H3,C3,C4,CA,S9"
)

type SolitaireTestSuite struct {
	suite.Suite
}

func (s *SolitaireTestSuite) TestSetDeck() {
	solitaire, err := solitaire.New()
	s.NoError(err, "Failed to create new solitaire instance")
	s.NotNil(solitaire, "Solitaire instance should not be nil")
	err = solitaire.SetDeck(strings.Split(setCRYPTONOMICON, ","))
	s.NoError(err, "Failed to set deck")
	c, err := solitaire.Encrypt([]byte("SOLITAIRE"))
	s.NoError(err, "Failed to encrypt plaintext")
	s.NotNil(c, "Ciphertext should not be nil")
	expected := "KIRAK SFJAN"
	s.Equal(expected, string(c), "Ciphertext does not match expected value")
}

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
