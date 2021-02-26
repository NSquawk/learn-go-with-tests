package arrays

// returns sum of all integers in an array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sum := []int{}

	for _, numberArray := range numbers {
		sum = append(sum, Sum(numberArray))
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {
	sum := []int{}

	for _, numberArray := range numbers {
		if len(numberArray) == 0 {
			sum = append(sum, 0)
		} else {
			tailSum := Sum(numberArray[1:])
			sum = append(sum, tailSum)
		}
	}
	return sum
}
