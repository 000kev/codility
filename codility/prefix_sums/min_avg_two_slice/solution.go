package minavgtwoslice

func Solution(A []int) int {
	N := len(A)
	minAvg := float64(A[0] + A[1]) / 2
	minIndex := 0
	
	for i := 0; i < N - 1; i++ {
		avg2 := float64(A[i] + A[i+1]) / 2
		if avg2 < minAvg {
			minAvg = avg2
			minIndex = i
		}

		if i < N - 2 {
			avg3 := float64(A[i] + A[i+1] + A[i+2]) / 3
			if avg3 < minAvg {
				minAvg = avg3
				minIndex = i
			}
		}
	}
	return minIndex
}

// func Solution(A []int) int {
// 	N := len(A)
// 	pref := u.PrefixSum(A)
// 	minimum := 10000
// 	minIndex := 0

// 	for i := range N {
// 		for j := i + 1; j < N; j += 1 {
// 			if i == 0 {
// 				candidate := pref[j]/j + 1
// 				if candidate < minimum {
// 					minIndex = i
// 				}
// 				minimum = min(minimum, candidate)
// 				continue
// 			}
// 			candidate := (pref[j] - pref[i-1]) / (j - i + 1)
// 			if candidate < minimum {
// 				minIndex = i
// 			}
// 			minimum = min(minimum, candidate)
// 		}
// 	}
// 	return minIndex
// }
