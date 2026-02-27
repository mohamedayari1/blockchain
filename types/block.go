package types

import (
	"crypto/sha256"

	"github.com/mohamedayari1/blockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func HashBlock(block *proto.Block) []byte {
	data, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(data)
	return hash[:]
}