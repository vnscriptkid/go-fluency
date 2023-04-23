package main

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strconv"
	"strings"
)

const replicationFactor = 1000

type node struct {
	nodeID string
	hash   uint32
}

type HashCircle []node

func (hc HashCircle) Len() int           { return len(hc) }
func (hc HashCircle) Less(i, j int) bool { return hc[i].hash < hc[j].hash }
func (hc HashCircle) Swap(i, j int)      { hc[i], hc[j] = hc[j], hc[i] }

func NewHashCircle() *HashCircle {
	hc := &HashCircle{}
	return hc
}

// FNV-1a 32-bit hash function, hash space size would be 2^32
// 0 -> (2^32 - 1) ~ 4.3 billion
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func addNode(hc *HashCircle, nodeID string) {
	for i := 0; i < replicationFactor; i++ {
		vNode := nodeID + ":" + strconv.Itoa(i)
		h := hash(vNode)
		*hc = append(*hc, node{nodeID: vNode, hash: h})
	}
	sort.Sort(hc)
}

func removeNode(hc *HashCircle, nodeID string) {
	newHC := &HashCircle{}
	for _, n := range *hc {
		if !strings.HasPrefix(n.nodeID, nodeID+":") {
			*newHC = append(*newHC, n)
		}
	}
	*hc = *newHC
	sort.Sort(hc) // Resort the hash circle by hash value after removing the node
}

func findNodeForKey(hc *HashCircle, key string) string {
	h := hash(key)
	// Find the first virtual node
	// with a hash value greater than or equal to the key's hash value
	i := sort.Search(hc.Len(), func(i int) bool { return (*hc)[i].hash >= h })

	// If not found, then i == hc.Len()
	if i == hc.Len() {
		i = 0
	}
	return (*hc)[i].nodeID
}

func main() {
	hashCircle := NewHashCircle()

	// Add nodes
	addNode(hashCircle, "node1")
	addNode(hashCircle, "node2")
	addNode(hashCircle, "node3")

	// Find nodes for keys
	fmt.Println("node for key1:", findNodeForKey(hashCircle, "key1"))
	fmt.Println("node for key2:", findNodeForKey(hashCircle, "key2"))
	fmt.Println("node for key3:", findNodeForKey(hashCircle, "key3"))
	fmt.Println("node for key4:", findNodeForKey(hashCircle, "key4"))
	fmt.Println("node for key5:", findNodeForKey(hashCircle, "key5"))
	fmt.Println("node for key6:", findNodeForKey(hashCircle, "key6"))
	fmt.Println("node for key7:", findNodeForKey(hashCircle, "key7"))
	fmt.Println("node for key8:", findNodeForKey(hashCircle, "key8"))

	removeNode(hashCircle, "node2")

	fmt.Println("node for key1 (after removing node2):", findNodeForKey(hashCircle, "key1"))
	fmt.Println("node for key2 (after removing node2):", findNodeForKey(hashCircle, "key2"))
	fmt.Println("node for key3 (after removing node2):", findNodeForKey(hashCircle, "key3"))
	fmt.Println("node for key4 (after removing node2):", findNodeForKey(hashCircle, "key4"))
	fmt.Println("node for key5 (after removing node2):", findNodeForKey(hashCircle, "key5"))
	fmt.Println("node for key6 (after removing node2):", findNodeForKey(hashCircle, "key6"))
	fmt.Println("node for key7 (after removing node2):", findNodeForKey(hashCircle, "key7"))
}
