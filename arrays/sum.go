package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num;
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	// lengthOfSlices := len(slices)
	// sums := make([]int, lengthOfSlices)

	// for i, slice := range slices {
	// 	sums[i] = Sum(slice)
	// }
	// return sums;

	var sums []int

	for _, slice := range slices {
		sums = append(sums, Sum(slice));
	}
	return sums;
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}