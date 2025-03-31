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
