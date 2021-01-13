package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix, has := findPrefix(strs[0], strs[1])
	if !has {
		return prefix
	}
	for i := 2; i < len(strs); i++ {
		prefix, has = findPrefix(strs[i], prefix)
		if !has {
			return prefix
		}
	}

	return prefix
}

func findPrefix(s1, s2 string) (string, bool) {
	var index, length1, length2 = 0, len(s1), len(s2)
	if length1 == 0 || length2 == 0 {
		return "", false
	}

	for i := 0; i < length1; {
		if s1[i] == s2[i] {
			i++
			index = i
		} else {
			break
		}

		if i == length2 {
			break
		}
	}

	if index == 0 {
		return "", false
	}

	return s1[:index], true
}
