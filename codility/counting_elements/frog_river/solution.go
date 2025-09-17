package frogriver

func Solution(X int, A []int) int {
	count := X
	count_array := make([]int, X + 1)
	for i := range len(A) {
		count_array[A[i]] += 1
		if count_array[A[i]] == 1 {
			count -= 1
		}
		if count == 0 {
			return i
		}
	}

	if count > 0 {
		return -1
	}
	return 0
}
