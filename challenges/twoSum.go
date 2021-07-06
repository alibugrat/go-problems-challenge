package challenges

func TwoSum(numbers []int, target int) []int {
	for i, number := range numbers {
		for a, n := range numbers {
			if a != i && n+number == target {
				return []int{i, a}
			}
		}
	}
	return []int{}
}
