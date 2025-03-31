package main

import (
	"fmt"

	"github.com/awnumar/memguard"
	"github.com/mwmahlberg/solitaire"
)

type encrypt struct {
	Cleartext []byte `kong:"arg,type='filecontent',help='Cleartext to be encrypted',sep=''"`
}

func (p *encrypt) Run() error {
	s, err := solitaire.New(solitaire.WithPassphraseFromEnclave(cfg.Passphrase.enc))
	if err != nil {
		memguard.SafePanic(err)
	}
	ct, err := s.Encrypt([]byte(p.Cleartext))
	if err != nil {
		memguard.SafePanic(err)
	}
	if ct == nil {
		memguard.SafePanic("Ciphertext is nil")
	}
	fmt.Println(string(ct))
	return nil
}
