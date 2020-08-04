package main

func SumArray(numbers []int) int {
	var sum int
	//for i := 0; i < len(numbers); i++ {
	for _, number := range numbers {

		sum = sum + number
	}

	return sum

}

func SumAll(numbersToAdd ...[]int) []int {

	length := len(numbersToAdd)
	sums := make([]int, length)
	for i, number := range numbersToAdd {
		sums[i] = SumArray(number)
	}

	return sums
}
