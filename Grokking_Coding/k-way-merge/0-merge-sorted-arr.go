package kwaymerge

func mergeSorted(nums1 []int, m int, nums2 []int, n int) []int {
	// 1, 2, 3, 4, 5, 6 (i)
	// ^$
	// 2, 3, 5 (j)
	// ^

	// if nums1[i] > nums2[j] : nums[next] = nums[i], next--, i--
	// else nums[next] = nums[j], next--, j--

	// outputExpected: []int{1, 2, 3, 4, 5, 6},

	//  [6, 7, 8, 9, 10, 6, 7, 8, 9, 10]
	//^               $
	//  [1, 2, 3, 4, 5]
	//               ^

	i := m - 1
	j := n - 1
	next := len(nums1) - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[next] = nums1[i]
			i--
		} else {
			nums1[next] = nums2[j]
			j--
		}
		next--
	}

	for i < 0 && j >= 0 && next >= 0 {
		nums1[next] = nums2[j]
		next--
		j--
	}

	// Write your code here
	// Your code will replace this placeholder return statement
	return nums1
}
