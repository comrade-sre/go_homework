package searchSimple

func search(x int) (result []int) {
	stack := []int{2}
	var isSimple bool
	if x == 1 {
		result = append(result, x)
		return
	}
	result = append(result, 2)
	for i := 3; i < x; i++ {
		for _, j := range stack {
			isSimple = true
			if i%j == 0 {
				stack = append(stack, i)
				isSimple = false
				break
			}
		}
		stack = append(stack, i)
		if isSimple {
			result = append(result, i)
		}
	}
	return result
}
