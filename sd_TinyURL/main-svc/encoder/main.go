package encoder

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var mapping = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+-"

func binaryToChar(binaryString string) (string, error) {
	index, err := strconv.ParseInt(binaryString, 2, 64)
	fmt.Printf("index: %v\n", index)
	if err != nil {
		return "", err
	}
	return string(mapping[index]), nil
}

func charToBinary(character string) (string, error) {
	index := -1
	for i, char := range mapping {
		if string(char) == character {
			index = i
			break
		}
	}
	if index == -1 {
		return "", fmt.Errorf("character not found in mapping")
	}
	return fmt.Sprintf("%06s", strconv.FormatInt(int64(index), 2)), nil
}

func generateRanges() {
	var partitionSize uint64 = 1000000000 // 1 billion
	var maxNum uint64 = uint64(math.Pow(2, 48) - 1)

	var i uint64
	for i = 0; i <= maxNum; i += partitionSize {
		end := i + partitionSize - 1

		if end > maxNum {
			end = maxNum
		}

		fmt.Printf("Range: [%d, %d]\n", i, end)

		if end == maxNum-1 {
			break
		}
	}
}

//  [280475000000000, 280475999999999]
//  [280476000000000, 280476999999999]
//  [280477000000000, 280477999999999]
//  [280478000000000, 280478999999999]

func Encode(int64Num int64) string {
	// step 1: convert to binary string
	binStr := strconv.FormatInt(int64Num, 2)
	// step 2: split into 6-bit chunks, if chunk is less than 6-bit, pad with 0
	encoded := ""

	fmt.Printf("binStr: %v\n", binStr)
	for i := len(binStr); i > 0; i -= 6 {
		startIdx := 0

		if i-6 > 0 {
			startIdx = i - 6
		}

		chunk := binStr[startIdx:i]

		if len(chunk) < 6 {
			chunk = strings.Repeat("0", 6-len(chunk)) + chunk
		}

		fmt.Printf("chunk: %v\n", chunk)
		// step 3: map each chunk to a character (0-9, a-z, A-Z)
		char, err := binaryToChar(chunk)
		if err != nil {
			log.Fatal("Error:", err)
		}

		// step 4: join all the characters together
		encoded = char + encoded
	}

	return encoded
}
