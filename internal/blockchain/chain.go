package blockchain

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"mini_chain/internal/types"
)

// Blockchain structure demonstrates mutex for concurrency access
type Blockchain struct {
	chain []types.Block
	mutex sync.RWMutex
}


// NewBlockchain demonstrates a constructor for the Blockchain struct
func NewBlockchain() *Blockchain {
	// Create the genesis block
	genesisBlock := types.Block{
		Index: 0,
		Timestamp: time.Now(),
		Data: "Genesis Block",
		PrevHash: "0",
		Difficulty: 1,
	}

	genesisBlock.Hash = genesisBlock.CalculateHash()

	// Using make to initialize the blockchain slice
	return &Blockchain{
		chain: []types.Block{genesisBlock},
	}
}

// AddBlock demonstrates adding a new block to the blockchain
func (bc *Blockchain) AddBlock (data string) error {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	prevBlock := bc.chain[len(bc.chain)-1]

	newBlock := types.Block{
		Index: prevBlock.Index + 1,
		Timestamp: time.Now(),
		Data: data,
		PrevHash: prevBlock.Hash,
		Difficulty: bc.calculateDifficulty(),
	}

	// Mining with simple proof of work
	if err := bc.mineBlock(&newBlock); err != nil {
		return fmt.Errorf("failed to mine block: %w", err)
	}

	bc.chain = append(bc.chain, newBlock)
	return nil
}

// IsValid demonstrates validating the blockchain
func (bc *Blockchain) IsValid() error {
	bc.mutex.RLock()
	defer bc.mutex.RUnlock()


	// Range loop through the chain
	for i := 1; i < len(bc.chain); i++ {
		currentBlock := bc.chain[i]
		prevBlock := bc.chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return fmt.Errorf("invalid hash at block %d", i)
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return fmt.Errorf("invalid chain link at block %d", i)
		}

		// Verify the hash meets difficulty requirement
		targetPrefix := strings.Repeat("0", int(currentBlock.Difficulty))
		if !strings.HasPrefix(currentBlock.Hash, targetPrefix) {
			return fmt.Errorf("invalid proof of work at block %d", i)
		}
	}
	return nil
}

func (bc *Blockchain) PrintChain() {
    bc.mutex.RLock()
    defer bc.mutex.RUnlock()

    for _, block := range bc.chain {
        fmt.Printf("Block #%d [Hash: %s]\n", block.Index, block.Hash)
        fmt.Printf("Data: %s\n\n", block.Data)
    }
}

func (bc *Blockchain) calculateDifficulty() uint8 {
	if len(bc.chain) < 2 {
		return 1
	}

	lastBlock := bc.chain[len(bc.chain)-1]
	timeDiff := time.Since(lastBlock.Timestamp)

	if timeDiff < 5*time.Second {
		return lastBlock.Difficulty + 1
	} else if timeDiff > 10*time.Second && lastBlock.Difficulty > 1 {
		return lastBlock.Difficulty - 1
	}

	return lastBlock.Difficulty
}