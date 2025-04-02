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
	"github.com/alecthomas/kong"
	"github.com/awnumar/memguard"
)

var cfg struct {
	Passphrase passphrase `kong:"help='Passphrase for the de- and encryption'"`
	Encrypt    encrypt    `kong:"cmd,help='Encrypt the given cleartext'"`
	Decrypt    decryptCmd `kong:"cmd,help='Decrypt the given ciphertext'"`
	PrintDeck  PrintDeck  `kong:"cmd,help='Print the deck for a given passphrase'"`
}

func main() {
	memguard.CatchInterrupt()
	defer memguard.Purge()
	ctx := kong.Parse(&cfg, kong.Name("solitaire"), kong.Description("Solitaire encryption and decryption tool"))
	err := ctx.Run()
	if err != nil {
		memguard.SafePanic(err)
	}
}
