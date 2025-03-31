package solitaire

import (
	"testing"

	"github.com/awnumar/memguard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SolitaireSuite struct {
	suite.Suite
}

func (s *SolitaireSuite) TestWithPassphraseFromEnclave() {
	passphrase := memguard.NewEnclave([]byte("CRYPTONOMICON"))
	passphrase.Open()
	f := WithPassphraseFromEnclave(passphrase)
	solitaireInstance := &solitaire{}
	err := f(solitaireInstance)
	assert.NoError(s.T(), err, "Expected no error when passphrase is valid")
}

func (s *SolitaireSuite) TestWithPassphraseFromNilEnclave() {
	f := WithPassphraseFromEnclave(nil)
	solitaireInstance := &solitaire{}
	err := f(solitaireInstance)
	assert.Error(s.T(), err, "Expected error when passphrase is nil")
}

func (s *SolitaireSuite) TestWithPassphraseFromInvalidEnclave() {
	passphrase := &memguard.Enclave{}
	f := WithPassphraseFromEnclave(passphrase)
	solitaireInstance := &solitaire{}

	s.Panics(func() {
		f(solitaireInstance)
	}, "Expected panic when enclave is invalid")
}

func TestSolitaireInternal(t *testing.T) {
	suite.Run(t, new(SolitaireSuite))
}
