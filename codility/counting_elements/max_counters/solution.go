package maxcounters

func Solution(N int, A []int) []int {
	current_max := 0
	counters := make([]int, N)
	base := 0

	for _, v := range A {
		if v - 1 > N {
			continue
		}
		if v >= 1 && v <= N {
			if counters[v-1] < base {
				counters[v-1] = base
			}

			counters[v-1]++

			if counters[v-1] > current_max {
				current_max = counters[v-1]
			}

		}
		if N != 1 && v > N {
			base = current_max
		}
	}

	for i := range counters {
		if counters[i] < base {
			counters[i] = base
		}
	}
	
	return counters
}