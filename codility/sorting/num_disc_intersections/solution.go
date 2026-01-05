package numdiscintersections

import (
	"fmt"
	// "slices"
)

func Solution(A []int) int {
	// discs1 := make([]int, 0, len(A))
	discs2 := make([][]int, 0, len(A))
	
	// for _, x := range A {
	// 	fmt.Println(x)
	// 	discs1 = discSort1(discs1, x)
	// }

	for i := range A {
		discs2 = discSort2(discs2, []int{i-A[i],i+A[i]})
	}
	fmt.Println("sorted discs:", discs2)

	// discs := make(map[int][][]int, len(A))
	// keys := make([]int, len(A))

	// for i := range A {
	// 	disc := [][]int{{i-A[i],i+A[i]}}
	// 	keys[i] = lenDisc(disc[0])

	// 	if discs[lenDisc(disc[0])] != nil {
	// 		discs[lenDisc(disc[0])] = append(discs[lenDisc(disc[0])], disc[0])
	// 		continue
	// 	}
	// 	discs[lenDisc(disc[0])] = disc
	// }
	// fmt.Println("keys:", keys)

	// slices.Sort(keys)
	// slices.Reverse(keys)
	
	intersections := 0

	// for _, x := range keys {
	// 	fmt.Println(discs[x])
	// }

	for a := range discs2 {
		minA := discs2[a][0]
		maxA := discs2[a][1]

		for b := a + 1; b < len(discs2); b++ {

			minB := discs2[b][0]
			maxB := discs2[b][1]

			if (minB >= minA && minB <= maxA) || (maxB >= minA && maxB <= maxA) {
				intersections++
			}
		}
	}
	// for a := range keys {
	// 	keyA := keys[a]
	// 	if (len(discs[keyA]) == 1) {
	// 		minA := discs[keyA][0][0]
	// 		maxA := discs[keyA][0][1]

	// 		for b := a + 1 ; b < len(keys) ; b++ {
	// 			keyB := keys[b]
	// 			if len(discs[keys[b]]) == 1 {
	// 				minB := discs[keyB][0][0]
	// 				maxB := discs[keyB][0][1]

	// 				if (minB >= minA || minB <= maxA) || (maxB >= minA || maxB <= maxA) {
	// 					intersections++
	// 				}
	// 			}
	// 			fmt.Println(discs[keyA], discs[keyB], intersections)
	// 		}
	// 	}
	// 	break
 	// }

	return intersections
}

func lenDisc(disc []int) int {
	return disc[1] - disc[0]
}

// func discSort1(prevArray []int, nextNum int) []int {
	
// 	prevArray = append(prevArray, nextNum)
// 	n := len(prevArray)

// 	if len(prevArray) > 1 {

// 		for i := n - 1; i >= 0; i-- {
// 			if i > 0 && prevArray[i] > prevArray[i - 1] {

// 				swap := prevArray[i]
// 				prevArray[i] = prevArray[i - 1]
// 				prevArray[i - 1] = swap
// 			}
// 		}
// 	}

// 	return prevArray
// }

func discSort2(prevArray [][]int, nextRange []int) [][]int {

	prevArray = append(prevArray, nextRange)
	n := len(prevArray)

	if len(prevArray) > 1 {

		for i := n - 1; i >= 0; i-- {
			if i > 0 && lenDisc(prevArray[i]) > lenDisc(prevArray[i - 1]) {
				swap := prevArray[i]
				prevArray[i] = prevArray[i - 1]
				prevArray[i - 1] = swap
			}
		}
	}

	return prevArray
}

