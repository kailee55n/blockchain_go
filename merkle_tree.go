package main

import (
	"crypto/sha256"
)

// MerkleTree represent a Merkle tree
type MerkleTree struct {
	RootNode *MerkleNode // RootNode is the root node of the Merkle tree
}

// MerkleNode represent a Merkle tree node
type MerkleNode struct {
	Left  *MerkleNode // Left child of the node
	Right *MerkleNode // Right child of the node
	// Data is the hash of the node, if it is a leaf node, it contains the data
	// If it is an internal node, it contains the hash of the concatenation of its children's hashes
	// The hash is computed using SHA-256
	Data []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
func NewMerkleTree(data [][]byte) *MerkleTree { // Create a new Merkle tree from a sequence of data
	if len(data) == 0 {
		return &MerkleTree{nil} // Return an empty Merkle tree if no data is provided
	}
	var nodes []MerkleNode

	if len(data)%2 != 0 { //	 If the number of data items is odd, duplicate the last item
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data { // 	Iterate over the data and create leaf nodes
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}

	return &mTree
}

// NewMerkleNode creates a new Merkle tree node
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}
