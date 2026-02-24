package crypto

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, prevKeyLen, len(privKey.Bytes()))
	assert.NotNil(t, privKey)

	pubKey := privKey.Public()
	assert.Equal(t, pubKeyLen, len(pubKey.Bytes()))
	assert.NotNil(t, pubKey)
}

func TestPublicKeyVerify(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("test message")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("wrong message")))
}

func TestPublicKeyToAdress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	adress := pubKey.Adress()
	assert.Equal(t, addressLen, len(adress.Bytes()))
	assert.NotNil(t, adress)
	assert.Equal(t, addressLen, len(adress.String()))
}