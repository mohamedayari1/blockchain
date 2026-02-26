package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGeneratePrivateKey tests the generation of a new random private key
// and verifies that the derived public key has the correct properties.
func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, prevKeyLen, len(privKey.Bytes()))
	assert.NotNil(t, privKey)

	pubKey := privKey.Public()
	assert.Equal(t, pubKeyLen, len(pubKey.Bytes()))
	assert.NotNil(t, pubKey)
}

// TestPublicKeyVerify tests that signatures can be created and verified.
// It ensures that a valid signature verifies correctly and invalid messages fail verification.
func TestPublicKeyVerify(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("test message")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("wrong message")))
}

// TestPrivateKeyFromString tests creating a private key from a hex-encoded seed string.
// It verifies that the key can be successfully created and recovered.
func TestPrivateKeyFromString(t *testing.T) {
	seed := "4a3b2c1d0e9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e1f0a9b8c7d6e5f4a3b"
	privKey, err := NewPrivateKeyFromString(seed)
	assert.NoError(t, err)
	assert.Equal(t, seed, hex.EncodeToString(privKey.Bytes()[:seedLen]))
}

// TestPublicKeyToAddress tests that an address can be derived from a public key
// and that it has the correct format and length.
func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
	assert.NotNil(t, address)
	assert.Equal(t, addressLen*2, len(address.String())) // hex encoding doubles the length

	fmt.Println("Address:", address.String())
}