package problem0020

import "strings"

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	m := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	ss := strings.Split(s, "")
	var ts []string
	for _, si := range ss {
		if _, ok := m[si]; ok {
			ts = append(ts, m[si])
		} else if len(ts)-1 < 0 {
			return true
		} else {
			if ts[len(ts)-1] == si {
				ts = ts[:len(ts)-1]
			} else {
				return false
			}
		}
	}
	return true
}
