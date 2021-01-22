package sort

func bubbleSort(sliceToSort []int) []int {
	for index := 0; index < len(sliceToSort); index++ {
		for tmpIndex := index + 1; tmpIndex < len(sliceToSort); tmpIndex++ {
			if sliceToSort[index] > sliceToSort[tmpIndex] {
				sliceToSort[index], sliceToSort[tmpIndex] = sliceToSort[tmpIndex], sliceToSort[index]
			}
		}
	}
	return sliceToSort
}
func insertSort(sliceToSort []int) []int {
	for index := 1; index < len(sliceToSort); index++ {
		compareValue := sliceToSort[index]
		tmpIndex := index
		for ; tmpIndex >= 1 && sliceToSort[tmpIndex-1] > compareValue; tmpIndex-- {
			sliceToSort[tmpIndex] = sliceToSort[tmpIndex-1]
		}
		sliceToSort[tmpIndex] = compareValue
	}
	return sliceToSort
}
