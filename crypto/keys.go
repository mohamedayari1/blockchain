// Package crypto provides cryptographic functions for blockchain operations.
// It handles key generation, digital signatures, and address derivation using
// the Ed25519 elliptic curve signature scheme.
package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Constants define the sizes of cryptographic components.
const (
	// prevKeyLen is the length of Ed25519 private keys (64 bytes: 32-byte seed + 32-byte public key).
	prevKeyLen = 64
	// pubKeyLen is the length of Ed25519 public keys (32 bytes).
	pubKeyLen = 32
	// seedLen is the length of the seed used to generate private keys (32 bytes).
	seedLen = 32
	// addressLen is the length of blockchain addresses derived from public keys (20 bytes, similar to Ethereum).
	addressLen = 20
)

// PrivateKey represents a private key used for signing transactions.
// It wraps Ed25519 private key which includes both the seed and the public key.
type PrivateKey struct {
	key ed25519.PrivateKey
}

// NewPrivateKeyFromString creates a PrivateKey from a hex-encoded string.
// The string must represent exactly 32 bytes of data (64 hex characters).
// Returns an error if the string is not valid hex or the wrong length.
func NewPrivateKeyFromString(s string) (*PrivateKey, error) {
	keyBytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string: %w", err)
	}
	return NewPrivateKeyFromSeed(keyBytes)
}

// NewPrivateKeyFromSeed creates a PrivateKey from a 32-byte seed.
// In Ed25519, the seed is deterministic - the same seed always produces the same key pair.
// Returns an error if the seed length is incorrect.
func NewPrivateKeyFromSeed(seed []byte) (*PrivateKey, error) {
	if len(seed) != seedLen {
		return nil, fmt.Errorf("seed must be exactly %d bytes, got %d", seedLen, len(seed))
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}, nil
}


// PublicKey represents a public key derived from a PrivateKey.
// It is used to verify signatures created by the corresponding private key.
type PublicKey struct {
	key ed25519.PublicKey
}

// Address derives a blockchain address from the public key by taking the last 20 bytes.
// This is similar to Ethereum's address derivation (last 20 bytes of Keccak-256 hash).
func (p *PublicKey) Address() Address {
	return Address{
		value: p.key[len(p.key)-addressLen:],
	}
}

// Bytes returns the raw bytes of the public key.
func (p *PublicKey) Bytes() []byte {
	return p.key
}

// GeneratePrivateKey creates a new random private key.
// It generates a cryptographically secure random 32-byte seed using crypto/rand
// and derives the Ed25519 key pair from it.
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




// Bytes returns the raw bytes of the private key (64 bytes total).
func (p *PrivateKey) Bytes() []byte {
	return p.key
}

// Sign creates a digital signature for a message using this private key.
// The signature can be verified by anyone with the corresponding public key.
// Returns a Signature object containing the 64-byte Ed25519 signature.
func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, msg),
	}
}

// Public extracts and returns the corresponding PublicKey.
// The public key is derived from the last 32 bytes of the private key.
func (p *PrivateKey) Public() *PublicKey {
	pubKey := p.key[prevKeyLen-pubKeyLen:]
	return &PublicKey{key: ed25519.PublicKey(pubKey)}
}

// Signature represents a digital signature created by a PrivateKey.
// Ed25519 signatures are always 64 bytes.
type Signature struct {
	value []byte
}

// Bytes returns the raw bytes of the signature.
func (s *Signature) Bytes() []byte {
	return s.value
}

// Verify checks whether this signature is valid for a given public key and message.
// Returns true if the signature is valid, false otherwise.
func (s *Signature) Verify(pk *PublicKey, msg []byte) bool {
	return ed25519.Verify(pk.key, msg, s.value)
}

// Address represents a blockchain address derived from a PublicKey.
// It's a 20-byte value typically displayed as a hex string.
type Address struct {
	value []byte
}

// Bytes returns the raw bytes of the address.
func (a Address) Bytes() []byte {
	return a.value
}

// String returns the hex-encoded representation of the address.
func (a Address) String() string {
	return hex.EncodeToString(a.value)
}



