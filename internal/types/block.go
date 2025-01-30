package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
// Using Go structs with field tags for JSON serialization
type Block struct {
	Index        uint64    `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	Data         string    `json:"data"`
	Hash         string    `json:"hash"`
	PrevHash     string    `json:"prev_hash"`
	Nonce        uint64    `json:"nonce"`
	Difficulty   uint8     `json:"difficulty"`
}

// Demonstrates method receiver in Go
func (b *Block) CalculateHash() string {
	// Fix: Use fmt.Sprintf instead of direct string conversion
	record := fmt.Sprintf("%d%s%s%s%d",
		b.Index,
		b.Timestamp.String(),
		b.Data,
		b.PrevHash,
		b.Nonce,
	)

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Interface
type Hashable interface {
	CalculateHash() string
}
