package greedy

import "math"

func jumpGame(nums []int) bool {
	// Your code will replace the placeholder return statement.

	// [ 3,  3,  4,  _,  _... ]
	//   i [i+1 i+2 i+3]
	//           ^

	// farthestIdxCanGo = 0
	// curPosIdx = 0

	farthestIdxCanGo := 0

	curPosIdx := 0

	for {
		potentialFarthest := curPosIdx + nums[curPosIdx]

		farthestIdxCanGo = int(math.Max(float64(potentialFarthest), float64(farthestIdxCanGo)))

		if farthestIdxCanGo >= len(nums)-1 {
			return true
		}

		if curPosIdx == farthestIdxCanGo {
			return false
		}

		curPosIdx++
	}

}
