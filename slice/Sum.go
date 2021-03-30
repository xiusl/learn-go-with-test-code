package slice

func Sum(numbers [5]int) (sum int) {
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

