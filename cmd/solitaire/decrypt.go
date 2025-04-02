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

package main

import (
	"fmt"

	"github.com/awnumar/memguard"
	"github.com/mwmahlberg/solitaire"
)

type decryptCmd struct {
	Ciphertext []byte `kong:"arg,type='filecontent',help='Ciphertext to be decrypted',sep=''"` //nolint:golint
}

func (p *decryptCmd) Run() error {
	s, err := solitaire.New(solitaire.WithPassphraseFromEnclave(cfg.Passphrase.enc))
	if err != nil {
		memguard.SafePanic(err)
	}
	ct, err := s.Decrypt(p.Ciphertext)
	if err != nil {
		memguard.SafePanic(err)
	}
	if ct == nil {
		memguard.SafePanic("Plaintext is nil")
	}
	// Print the plaintext
	fmt.Printf("%s\n", ct)
	return nil
}
