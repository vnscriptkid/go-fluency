package Bitwise

func extraCharacterIndex(str1 string, str2 string) int {
	xorAll := 0

	for _, c := range str1 {
		xorAll ^= int(c)
	}

	for _, c := range str2 {
		xorAll ^= int(c)
	}

	var longerStr *string = &str1

	if len(str1) < len(str2) {
		longerStr = &str2
	}

	for i, c := range *longerStr {
		if xorAll == int(c) {
			return i
		}
	}

	return -1
}
