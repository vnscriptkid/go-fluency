package grokking

func removeDuplicates(a []int) {
	// 2, 3, 6, 9, 6, 9, 9
	//             $
	//                     ^

	// if a[cur] !== a[next]
	// swap (a, cur, next)
	// cur++,next++

	// else cur++
	//

	cur := 1
	next := 1

	for cur < len(a) {

		if a[cur] != a[next-1] {
			a[cur], a[next] = a[next], a[cur]

			next++
		}

		cur++
	}

}
