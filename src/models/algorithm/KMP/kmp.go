package algorithm

func Matchstr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	prefixFunc := getPrefix(needle)
	pat_ptr, tar_ptr := 0, 0
	for tar_ptr < len(haystack) {
		if haystack[tar_ptr] == needle[pat_ptr] {
			tar_ptr += 1
			pat_ptr += 1
		} else {
			if pat_ptr == 0 {
				tar_ptr += 1
			} else {
				pat_ptr = prefixFunc[pat_ptr-1]
			}
		}
		if pat_ptr == len(needle) {
			return tar_ptr - pat_ptr
		}
	}
	return -1
}

func getPrefix(pattern string) (prefixFunc []int) {
	pre_ptr, itr_ptr := 0, 1
	prefixFunc = append(prefixFunc, 0)
	for itr_ptr < len(pattern) {
		if pattern[pre_ptr] == pattern[itr_ptr] {
			pre_ptr += 1
			itr_ptr += 1
			prefixFunc = append(prefixFunc, pre_ptr)
		} else {
			if pre_ptr == 0 {
				prefixFunc = append(prefixFunc, 0)
				itr_ptr += 1
			} else {
				pre_ptr = prefixFunc[pre_ptr-1]
			}
		}
	}
	return
}
