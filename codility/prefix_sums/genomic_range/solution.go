package genomicrange

import u "codility.com/prefix-sums/utils"

// = [0 C CA CAG CAGC CAGCC CAGCCT CAGCCTA]
// = [0 2 21 213 2132 21322 213224 2132241]
// = [0 0100 1100 1110 1210 1310 1311 2311]

// (2,4) : CAGCC - CAG = 1210 - 1100 = 0100 = 2
// (5,5) : min = 4
// (0,6) : 2311 - 0100 = 2211 = 1

//         0 1 2 3 4 5 6 7 8 9 10
// let S = T T T C G A T T C C G
//       =[4 4 4 2 3 1 4 4 2 2 3]

//  = [0 T TT TTT TTTC TTTCG TTTCGA TTTCGAT TTTCGATT TTTCGATTC TTTCGATTCC TTTCGATTCCG]
//  = [0 4    44   444  4442 44423 444231 4442314 44423144 444231442 4442314422 44423144223]
//  = [0 0001 0002 0003 0103 0113  1113   1114 	  1115	   1215		 1315		1325]

// (2,4) : 0103 - 0002 = 0101 = 2
// (5,5) : min = 1
// (0,6) : 1113 - 0001 = 1112 = 1
// (7,9) : 1215 - 1114 = 0101 = 2

func subtract(A, B []int) []int {
	res := make([]int, 4)
	for i := range len(A) {
		res[i] = A[i] - B[i]
	}
	return res
}

func minNucleotide(P []int) int {
	for i, ncl := range P {
		if ncl > 0 {
			return i+1
		}
	}
	return 1
}

// O(N + M) solution
func Solution(S string, P, Q []int) []int {
	M := len(P)
	solution := make([]int, M)
	pref := u.PrefixSumCount(S)
	nclMap := map[string]int{"A" : 1, "C" : 2, "G" : 3, "T" : 4}

	for i := range M {
		x, y := P[i], Q[i]
		if x == y {
			solution[i] = nclMap[string(S[x])]
			continue
		}
		if x  < 1 {
			result := subtract(pref[y], []int{0,0,0,0})
			solution[i] = minNucleotide(result)
			continue
		}
		result := subtract(pref[y], pref[x-1])
		solution[i] = minNucleotide(result)		
	}
	return solution
}

// O(n*m) solution
func Solution1(S string, P, Q []int) []int {
	M := len(P)
	solution := make([]int, M)

	for i := range M {
		min := 4
		for j := P[i]; j <= Q[i]; j++ {
			if string(S[j]) == "A" {
				min = 1
				break
			}
			var ncl int
			switch string(S[j]) {
			case "C":
				ncl = 2
			case "G":
				ncl = 3
			case "T":
				ncl = 4
			}
			if ncl < min {
				min = ncl
			}
		}
		solution[i] = min
	}
	return solution
} 