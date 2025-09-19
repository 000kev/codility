package permcheck

func Solution(A []int) int {
	n := len(A)
	count_array := make([]int, n+1)
	for i := range A {
		if A[i] > n {
			return 0
		}
		count_array[A[i]] += 1
		if count_array[A[i]] > 1 {
			return 0
		}
	}
	return 1
}
