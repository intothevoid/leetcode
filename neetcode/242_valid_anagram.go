package neetcode

func IsAnagram(s string, t string) bool {
	smap := map[rune]int{}
	tmap := map[rune]int{}

	scnt := 0
	for _, val := range s {
		smap[val] = scnt + 1
	}

	tcnt := 0
	for _, val := range t {
		tmap[val] = tcnt + 1
	}

	if len(smap) == len(tmap) {
		for k := range smap {
			if smap[k] == tmap[k] {
				continue
			}
			return false
		}
		return true
	}
	return false
}
