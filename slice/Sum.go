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
	//lengthOfNumbers := len(numbersToSum)
	//sums = make([]int, lengthOfNumbers)
	//
	//for i, numbers := range numbersToSum {
	//	sums[i] = Sum(numbers)
	//}
	//return

	// 使用 append 函数
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbersToSum := range numbersToSum {
		tail := numbersToSum[1:]
		sums = append(sums, Sum(tail))
	}
	return
}
