package issubsequence

func IsSubsequence(s string, t string) bool {
	if (len(s) == 0) { return true }

	isSubsequence := false
	pointerS := 0
	pointerT := 0

	for pointerS < len(s) && pointerT < len(t) {
		if s[pointerS] == t[pointerT] {
			if len(s)-1 == pointerS {
				isSubsequence = true
			}
			pointerS++
		}
		pointerT++
	}

	return isSubsequence
}
