package testing

func StrArrContains(sa []string, s1 string) bool {
	for _, s2 := range sa {
		if s2 == s1 {
			return true
		}
	}
	return false
}

func ElementsMatching(sa []string) func(a []string) bool {
	return func(a []string) bool {
		for _, m := range sa {
			if !StrArrContains(a, m) {
				return false
			}
		}
		return true
	}
}
