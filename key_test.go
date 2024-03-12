package main

// TestUnit by Adrian Lorenz

import (
	"testing"

	"github.com/adrian-lorenz/nox-vault/crsa"
	"github.com/stretchr/testify/assert"
)

func TestKeyGen(t *testing.T) {
	kk, cerr := crsa.GenKey()
	if cerr != nil {
		t.Error(cerr)
	}
	assert.NotEmpty(t, kk)
	assert.NotEmpty(t, kk.PublicKey)

}
