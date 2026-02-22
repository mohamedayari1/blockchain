package crypto

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ed25519/go-ed25519"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.NotNil(t, privKey)
	assert.Equal(t, prevKeyLen, len(privKey.Bytes()))


func TestPublicKeyVerify(t *testing.T) {
	privKey := ed25519.NewKeyFromSeed(make([]byte, 32))
	pubKey := privKey.Public().(ed25519.PublicKey)
	msg := []byte("test message")
	sig := ed25519.Sign(privKey, msg)

	if !pubKey.Verify(msg, sig) {
		t.Errorf("PublicKey verification failed")
	}
}
