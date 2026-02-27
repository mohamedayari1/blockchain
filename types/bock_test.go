package types

import (
	"testing"

	"github.com/mohamedayari1/blockchain/proto"
	"github.com/mohamedayari1/blockchain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestHashBlock_Length verifies the hash is always exactly 32 bytes (SHA-256).
func TestHashBlock_Length(t *testing.T) {
	block := utils.RandomBlock()
	hash := HashBlock(block)
	assert.Len(t, hash, 32, "SHA-256 hash must be 32 bytes")
}

// TestHashBlock_Deterministic verifies the same block always produces the same hash.
func TestHashBlock_Deterministic(t *testing.T) {
	block := utils.RandomBlock()
	hash1 := HashBlock(block)
	hash2 := HashBlock(block)
	assert.Equal(t, hash1, hash2, "same block must produce identical hashes")
}

// TestHashBlock_Unique verifies two independently-generated blocks produce different hashes.
func TestHashBlock_Unique(t *testing.T) {
	block1 := utils.RandomBlock()
	block2 := utils.RandomBlock()
	hash1 := HashBlock(block1)
	hash2 := HashBlock(block2)
	// This could theoretically collide but is astronomically unlikely with random 32-byte fields.
	assert.NotEqual(t, hash1, hash2, "different blocks should produce different hashes")
}

// TestHashBlock_SensitiveToHeaderChanges verifies that mutating any header field changes the hash.
func TestHashBlock_SensitiveToHeaderChanges(t *testing.T) {
	block := utils.RandomBlock()
	originalHash := HashBlock(block)

	// Mutate version
	block.Header.Version++
	assert.NotEqual(t, originalHash, HashBlock(block), "changing Version must change hash")
	block.Header.Version--

	// Mutate height
	block.Header.Height++
	assert.NotEqual(t, originalHash, HashBlock(block), "changing Height must change hash")
	block.Header.Height--

	// Mutate prevHash
	block.Header.PrevHash = utils.RandomHash()
	assert.NotEqual(t, originalHash, HashBlock(block), "changing PrevHash must change hash")
}

// TestHashBlock_EmptyBlock verifies that even an empty block can be hashed without panic.
func TestHashBlock_EmptyBlock(t *testing.T) {
	block := &proto.Block{}
	require.NotPanics(t, func() {
		hash := HashBlock(block)
		assert.Len(t, hash, 32)
	})
}

// TestHashBlock_BlockWithTransactions verifies adding transactions changes the hash.
func TestHashBlock_BlockWithTransactions(t *testing.T) {
	block := utils.RandomBlock()
	hashBefore := HashBlock(block)

	// Add a transaction
	block.Transactions = append(block.Transactions, &proto.Transaction{})
	hashAfter := HashBlock(block)

	assert.NotEqual(t, hashBefore, hashAfter, "adding a transaction must change the block hash")
}