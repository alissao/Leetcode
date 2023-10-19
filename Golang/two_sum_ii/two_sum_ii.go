package twosumii

func TwoSum(numbers []int, target int) []int {
	pointer1 := 0
	pointer2 := len(numbers) - 1

	for ; pointer1 < len(numbers) && pointer2 < len(numbers); {
		twoSum := numbers[pointer1] + numbers[pointer2]
		if twoSum > target {
			pointer2--
		} else if twoSum < target {
			pointer1++
		} else if twoSum == target {
			break
		}
	}

	return []int{pointer1+1, pointer2+1}
}