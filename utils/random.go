package utils

import (
	cryptorand "crypto/rand"
	"io"
	"math/rand"
	"time"

	"github.com/mohamedayari1/blockchain/proto"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(cryptorand.Reader, hash)
	return hash
}

func RandomBlock() *proto.Block {
	header := &proto.Header{
		Version:   1,
		Height:    uint32(rand.Intn(1000)),
		PrevHash:  RandomHash(),
		RootHash:  RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
	}
	return &proto.Block{
		Header:       header,
		Transactions: []*proto.Transaction{},
	}
}