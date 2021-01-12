package fibonacci



func Fibonacci(number int) int {
	var calculated = make(map[int]int)
	if number <= 1 {
		return number
	} else if val, ok := calculated[number]; ok {
		return val
	} else {
		result := fibonacci(number-1) + fibonacci(number-2)
		calculated[number] = result
		return result
	}
}
