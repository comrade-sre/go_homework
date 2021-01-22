package fibonacci

func Fibonacci(number int) int {
	var calculated = make(map[int]int)
	if number <= 1 {
		return number
	} else if val, ok := calculated[number]; ok {
		return val
	} else {
		result := Fibonacci(number-1) + Fibonacci(number-2)
		calculated[number] = result
		return result
	}
}
