package blockchain

import (
	"fmt"
	"time"
	"strings"
	"mini_chain/internal/types"
)


// Constants
const (
	miningTimeout = 3 * time.Second
	targetPrefix = "0"
)

// mineBlock demonstrates mining a new block
func (bc *Blockchain) mineBlock(block *types.Block) error {
	targetPrefix := strings.Repeat("0", int(block.Difficulty))

	timeout := time.After(30 * time.Second)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("mining timed out")
		default:
			block.Nonce++
			hash := block.CalculateHash()

			if strings.HasPrefix(hash, targetPrefix) {
				block.Hash = hash
				return nil
			}
		}
	}
}