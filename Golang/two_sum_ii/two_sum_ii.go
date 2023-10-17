package twosumii

func TwoSum(numbers []int, target int) []int {
	pointer1 := 0
	pointer2 := 1
	numbersLen := len(numbers)

	for ; pointer1 < numbersLen && pointer2 < numbersLen; {
		twoSum := numbers[pointer1] + numbers[pointer2]
		if twoSum == target {
			break
		}
		pointer2++
		if pointer2 == numbersLen {
			pointer1++
			pointer2 = pointer1 + 1
		}
	}

	return []int{pointer1+1, pointer2+1}
}