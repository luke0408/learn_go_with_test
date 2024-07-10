package main

func Sum(numbers []int) int {
	sum := 0

	// range: 인덱스와 값을 반환 (인덱스는 사용하지 않을 때 _로 무시)
	for _, number := range numbers {
		sum += number
	}

	return sum
}
