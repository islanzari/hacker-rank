package main

import (
	"fmt"
)

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[i] < ar[j] {
				k1 := ar[i] + ar[j]
				if k1%k == 0 {
					count++

				}

			}

		}
	}
	return count
}

func main() {
	var ar = []int32{43, 95, 51, 55, 40, 86, 65, 81,
		51, 20, 47, 50, 65, 53, 23, 78, 75,
		75, 47, 73, 25, 27, 14, 8, 26, 58, 95,
		28, 3, 23, 48, 69, 26, 3, 73, 52, 34, 7,
		40, 33, 56, 98, 71, 29, 70, 71, 28, 12,
		18, 49, 19, 25, 2, 18, 15, 41, 51, 42, 46, 19, 98, 56, 54, 98, 72, 25, 16, 49, 34, 99, 48, 93, 64, 44, 50, 91, 44, 17, 63, 27, 3, 65,
		75, 19, 68, 30, 43, 37, 72, 54, 82, 92, 37, 52, 72, 62, 3, 88, 82, 71}
	fmt.Println(divisibleSumPairs(100, 22, ar))
}
