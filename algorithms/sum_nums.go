package algorithms

func sumInList(list []int) int {
	if len(list) == 0 {
		return 0
	}

	sum := 0
	for _, v := range list {
		sum += v
	}

	return sum
}
