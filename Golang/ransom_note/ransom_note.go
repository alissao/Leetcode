package ransomnote

func CanConstruct(ransomNote string, magazine string) bool {
	requireds := map[rune]int{}
	requiredsCount := 0

	for _, run := range ransomNote {
		requireds[run]++
		requiredsCount++
	}

	for _, run := range magazine {
		if count, ok := requireds[run]; ok && count > 0 {
			requireds[run]--
			requiredsCount--
		}
	}

	return requiredsCount == 0
}
