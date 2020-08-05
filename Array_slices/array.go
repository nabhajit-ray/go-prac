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

	//length := len(numbersToAdd)
	//sums := make([]int, length)
	var sums []int
	for _, number := range numbersToAdd {
		sums = append(sums, SumArray(number))
	}

	return sums
}

func SumAllTails(numbersToAdd ...[]int) []int {

	var sums []int
	for _, number := range numbersToAdd {
		if len(number) == 0 {
			//sums = append(sums, 0)
			//sums = sums
		} else {

			number = number[len(number)-1:]
			sums = append(sums, SumArray(number))
		}
	}

	return sums
}
