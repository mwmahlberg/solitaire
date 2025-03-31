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
