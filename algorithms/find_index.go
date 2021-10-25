package algorithms

// find an index of an element
func findIndex(list []int, num int) (index int, exist bool) {
	if len(list) == 0 {
		return 0, false
	}

	for i, v := range list {
		if v == num {
			return i, true
		}
	}

	return 0, false
}
