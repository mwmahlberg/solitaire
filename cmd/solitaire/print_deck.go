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
	"strings"

	"github.com/awnumar/memguard"
	"github.com/mwmahlberg/solitaire"
)

type PrintDeck struct {
	Export bool `kong:"help='print the deck as a sequence suitable for importing'"`
}

func (p *PrintDeck) Run() error {
	s, err := solitaire.New(solitaire.WithPassphraseFromEnclave(cfg.Passphrase.enc))
	if err != nil {
		memguard.SafePanic(err)
	}
	if p.Export {
		deck := make([]string, len(s.Deck()))
		for i, c := range s.Deck() {
			deck[i] = c.Short()
		}
		fmt.Println(strings.Join(deck, ","))
		return nil
	}

	for i, c := range s.Deck() {
		fmt.Printf("%2d: %s\n", i+1, c.String())
	}
	return nil
}
