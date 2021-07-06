package challenges

import "math"

func Reverse(x int) int {
	dc := getDigitsCount(x)
	var result int = 0
	for i := 0; i < dc; i++ {
		multiple := int(math.Pow(float64(10), float64(dc-i)))
		digit := int(math.Pow(float64(10), float64(i)))
		number := (x % (digit * 10)) / digit
		result += number * multiple
	}
	result /= 10
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	} else {
		return result
	}
}
func getDigitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}
