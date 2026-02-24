package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
	"encoding/hex"
)


const (
	prevKeyLen = 64
	pubKeyLen = 32
	seedLen = 32
	addressLen = 20
)

type PrivateKey struct  {
	key ed25519.PrivateKey
}


type PublicKey struct  {
	key ed25519.PublicKey
}

func (p *PublicKey) Adress() Adress {
	return Adress{
		value: p.key[len(p.key) - addressLen:],
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
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

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, msg),
	}
}

func (p *PrivateKey) Public() *PublicKey {
	pubKey := p.key[prevKeyLen - pubKeyLen:]
	return &PublicKey{key: ed25519.PublicKey(pubKey)}
}


type Signature struct {
	value []byte
}


func (s *Signature) Bytes() []byte {
	return s.value
}

func (s *Signature) Verify(pk *PublicKey, msg []byte) bool {
	return ed25519.Verify(pk.key, msg, s.value)
}

type Adress struct {
	value []byte
} 


func (a Adress) Bytes() []byte {
	return a.value
}


func (a Adress) String() string {
	return hex.EncodeToString(a.value)
}



