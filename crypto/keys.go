package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)


const (
	prevKeyLen = 64
	pubKeyLen = 32
	seedLen = 32
)

type PrivateKey struct  {
	key ed25519.PrivateKey
}


type PublicKey struct  {
	key ed25519.PublicKey
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic("Failed to generate random seed: " + err.Error())
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}


}




func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

func (p *PrivateKey) Public() *PublicKey {
	pubKey := p.key[prevKeyLen - pubKeyLen:]
	return &PublicKey{key: pubKey,}
}








