package passingcars

import u "codility.com/prefix-sums/utils"

func Solution(A []int) int {
	n := len(A)
	pref := u.PrefixSum(A)
	total := 0
	multiplier := 0

	for k := 1; k <= n; k++ {
		if total > 1000000000 {
			return - 1
		}
		if pref[k] == pref[k-1] {
			multiplier++
			continue
		}
		total += multiplier
	}

	return total
}
