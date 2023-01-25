package slice

import (
	"fmt"
	"strconv"
)

func doubleEach(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	// Access by idx
	fmt.Println("Num at index 2: ", slice[2])

	// Len of array
	fmt.Println("Len of slice: ", len(slice))

	// Capacity of array
	// cap === len
	fmt.Println("Capacity of slice: ", cap(slice))

	slice = append(slice, 6)

	// cap !== len
	fmt.Println("Capacity of slice: ", cap(slice))

	// Loop through slice
	for i := 0; i < len(slice); i++ {
		fmt.Println("Ele at idx "+strconv.Itoa(i)+" is: ", slice[i])
	}

	doubleEach(slice) // Pass by ref by nature
	fmt.Println("== Slice has been doubled each")

	// Loop through slice
	for i := 0; i < len(slice); i++ {
		fmt.Println("Ele at idx "+strconv.Itoa(i)+" is: ", slice[i])
	}

}
