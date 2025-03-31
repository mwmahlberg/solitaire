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
