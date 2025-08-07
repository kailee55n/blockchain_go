package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 16

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork { // Create a new ProofOfWork instance for the given block
	// Create a target based on the difficulty level
	// The target is a big integer that represents the maximum hash value allowed for a valid block
	// The lower the target, the more difficult it is to find a valid hash
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target} // Initialize the ProofOfWork with the block and target

	// Print the target for debugging purposes
	fmt.Printf("Target: %064x\n", pow.target)

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte { //	Prepare the data for hashing
	// The data includes the previous block's hash, the current block's transactions hash,
	// the timestamp, the target bits, and the nonce
	// This data will be hashed to find a valid proof-of-work
	// The nonce is incremented until a valid hash is found
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining a new block")
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		if math.Remainder(float64(nonce), 100000) == 0 {
			fmt.Printf("\r%x", hash)
		}
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
