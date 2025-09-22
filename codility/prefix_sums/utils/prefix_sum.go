package utils

func PrefixSum(A []int) []int {
	n := len(A)
	P := make([]int, n+1)
	P[0] = 0

	for i := 1; i <= n; i++ {
		P[i] = P[i-1] + A[i-1]
	}
	return P
}

func PrefixSumMod(A []int, K int) []int {
	n := len(A)
	P := make([]int, n+1)
	P[0] = 0

	for i := 1; i <= n; i++ {
		if A[i-1] % K == 0 {
			P[i] = P[i-1] + 1
			continue
		}
		P[i] = P[i-1]
	}
	return P
}


func PrefixSumCount(S string) [][]int {
	nclMap := map[string]int{"A" : 1, "C" : 2, "G" : 3, "T" : 4}
	P := make([][]int, len(S))
	for x := range P {
		P[x] = make([]int, 4)
	}

	for i, ncl := range S {
		key := nclMap[string(ncl)] - 1
		if (i > 0) {
			copy(P[i], P[i-1])
		}
		P[i][key]++
	}
	return P
}

func CountTotal(P []int, x, y int) int {
	return P[y+1] - P[x]
}