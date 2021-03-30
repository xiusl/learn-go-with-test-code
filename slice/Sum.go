package slice

func Sum(numbers []int) (sum int) {
	sum = 0

	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	// 使用 range
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

