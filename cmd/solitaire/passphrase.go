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
	"errors"
	"fmt"
	"regexp"

	"github.com/alecthomas/kong"
	"github.com/awnumar/memguard"
)

var isValidPassphrase = regexp.MustCompile(`^[a-zA-Z]+$`)

type passphrase struct {
	enc *memguard.Enclave
}

func (p *passphrase) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}

	p.enc = memguard.NewEnclave(text)
	return nil
}

func (p *passphrase) Validate(ctx *kong.Context) error {
	if p.enc == nil {
		return errors.New("passphrase is required, use --passphrase. e.g. --passphrase=secret. If you really want to use an empty passphrase, use --passphrase=''")
	}
	b, err := p.enc.Open()
	if err != nil {
		return fmt.Errorf("failed to open passphrase containr: %w", err)
	}
	defer b.Destroy()
	if len(b.Bytes()) == 0 {
		ctx.Printf("WARN: passphrase is empty, this is not recommended")
	}
	if len(b.Bytes()) < 80 {
		ctx.Printf("WARN: passphrase should be at least 80 characters long")
	}
	// Check if the passphrase contains only alphanumeric characters
	if !isValidPassphrase.Match(b.Bytes()) {
		return errors.New("passphrase must contain only letters from A-Z and a-z")
	}
	return nil
}
