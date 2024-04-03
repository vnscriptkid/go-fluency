package main

import "fmt"

func main() {
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Bill"}

	var i int
	var v string
	for i, v = range five {
		fmt.Printf("Value[%s]\tAddress[%p]  IndexAddr[%p]\n",
			v, &v, &five[i])
	}

	// Value[Annie]    Address[0x14000010040]  IndexAddr[0x1400010a000]
	// Value[Betty]    Address[0x14000010040]  IndexAddr[0x1400010a010]
	// Value[Charley]  Address[0x14000010040]  IndexAddr[0x1400010a020]
	// Value[Doug]     Address[0x14000010040]  IndexAddr[0x1400010a030]
	// Value[Bill]     Address[0x14000010040]  IndexAddr[0x1400010a040]

	// The values 0x1400010a000 and 0x1400010a010 are hexadecimal numbers
	// In hexadecimal, 0x10 is equivalent to 16 in decimal.
	// So, the distance between 0x1400010a000 and 0x1400010a010 is 16 bytes.
	// string is 2 words in size of 8 bytes each, so 16 bytes in total ()

	fmt.Println()
	for i2, v2 := range five {
		fmt.Printf("Value[%s]\tAddress[%p]  IndexAddr[%p]\n",
			v2, &v2, &five[i2])
	}

	// Value[Annie]    Address[0x1400008e0b0]  IndexAddr[0x14000096050]
	// Value[Betty]    Address[0x1400008e0d0]  IndexAddr[0x14000096060]
	// Value[Charley]  Address[0x1400008e0f0]  IndexAddr[0x14000096070]
	// Value[Doug]     Address[0x1400008e110]  IndexAddr[0x14000096080]
	// Value[Bill]     Address[0x1400008e130]  IndexAddr[0x14000096090]
}
