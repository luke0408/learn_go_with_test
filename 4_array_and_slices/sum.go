package main

func Sum(numbers []int) int {
	sum := 0

	// range: 인덱스와 값을 반환 (인덱스는 사용하지 않을 때 _로 무시)
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	/*
		아래와 같이 작성 가능
		하지만, len()의 반환값에 따라 out of range 에러가 발생할 수 있음

		lengthOfNumbers := len(numbersToSum)
		sums := make([]int, lengthOfNumbers)
		for i, numbers := range numbersToSum {
		  sums[i] = Sum(numbers)
		}
	*/

	var sums []int
	for _, numbers := range numbersToSum {
		// append: 슬라이스에 요소 추가
		sums = append(sums, Sum(numbers))
	}

	return sums
}
