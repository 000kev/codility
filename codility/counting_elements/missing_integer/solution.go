package missinginteger

import "slices"

func getMax(A *[]int) int {
	maxi := 0
	for _, v := range *A {
		maxi = max(maxi, v)
	}
	return maxi
}

func Solution(A []int) int {
	N := len(A)
	infimum := 1
	maximum := getMax(&A)
	counter := make([]int, maximum)

	if N == 0 {
		return infimum
	}
	if N == 1 {
		if A[0] <= 0 {
			return infimum
		} 
		if A[0] - 1 > 0 {
			return A[0] - 1
		}
		return A[0] + 1
	}
	if maximum <= 0 {
		return infimum
	}

	for _, v := range A {
		if v > 0 {
			counter[v-1]++
		}
	}

	index := slices.Index(counter, 0)
	if index < 0 {
		if maximum < 0 {
			return infimum
		}
		return maximum + 1
	}
	return index + 1
}
