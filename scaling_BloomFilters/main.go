package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
)

const (
	// lowering the false positive rate
	k    = 4       // number of hash functions
	size = 1000000 // size of the bit array
)

type BloomFilter struct {
	bitArray [size]bool
}

func (bf *BloomFilter) Add(item string) {
	hashes := HashFunctions(item)
	for _, hash := range hashes {
		position := hash % uint32(size)
		bf.bitArray[position] = true
	}
}

func (bf *BloomFilter) Exists(item string) bool {
	hashes := HashFunctions(item)
	for _, hash := range hashes {
		position := hash % uint32(size)
		if !bf.bitArray[position] {
			// If it is false, then the item is definitely not in the filter
			return false
		}
	}
	// If the loop completes without finding any false bits, the function returns true,
	// indicating that the item might be in the filter (with a certain false positive probability)
	return true
}

// The optimal value of k (the number of hash functions)
// in a Bloom filter depends on the size of the bit array (m)
// and the expected number of items to be stored in the filter (n)
func OptimalK(m int, n int) int {
	k := float64(m) / float64(n) * math.Log(2)
	return int(math.Round(k))
}

func FalsePositiveRate(n int, m int, k int) float64 {
	p := math.Pow(1-math.Exp(-float64(k)*float64(n)/float64(m)), float64(k))
	return p
}

// create k different hash values for the input item by slicing different portions of the SHA-256 hash
func HashFunctions(item string) []uint32 {
	hash := sha256.Sum256([]byte(item)) // output is 256 bits (32 bytes)
	hashes := make([]uint32, k)

	for i := 0; i < k; i++ {
		start := i * 4
		// extracts a 4-byte (32-bit) portion of the SHA-256 hash starting
		// from the index start and ending at start + 4
		// convert these 4 bytes into a uint32 value
		hashes[i] = binary.BigEndian.Uint32(hash[start : start+4])
	}

	return hashes
}

func main() {
	bloom := &BloomFilter{}

	bloom.Add("apple")
	bloom.Add("banana")
	bloom.Add("orange")

	fmt.Println("Checking if 'apple' exists:", bloom.Exists("apple"))   // true
	fmt.Println("Checking if 'banana' exists:", bloom.Exists("banana")) // true
	fmt.Println("Checking if 'grape' exists:", bloom.Exists("grape"))   // false
}
