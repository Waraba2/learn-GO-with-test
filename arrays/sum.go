package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
  var sums []int
  for _, number := range numbersToSum {
    if len(number) == 0 {
      sums = append(sums, 0)
    } else {
    tail := number[1:]
    sums = append(sums, Sum(tail))
    }
  }

  return sums
}
