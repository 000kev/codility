package countdiv

import u "codility.com/prefix-sums/utils"


// O(B - A) time complexity
func Solution1(A int, B int, K int) int {
	length := B - A + 1

	arr := make([]int, length)
	for i := A; i <= B; i++ {
		arr[i-A] = i
	}
	
	pref := u.PrefixSumMod(arr, K)
	
	return u.CountTotal(pref, 0, length-1)
}

// O(1) time complexity
func Solution2(A, B, K int) int {
	if A == 0 {
		return B/K + 1
	}
	return B/K - (A-1)/K
}