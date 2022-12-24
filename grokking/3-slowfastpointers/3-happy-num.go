package grokking

func nextNum(x int) int {
	result := 0

	for x > 0 {
		d := x % 10
		x = x / 10

		result += d * d
	}

	return result
}

func findHappyNum(num int) bool {
	// 23 -> 13 -> 10 -> 1
	//       		f
	//        s

	slow := num
	fast := num

	for {
		fast = nextNum(fast)
		fast = nextNum(fast)
		slow = nextNum(slow)

		if fast == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}
