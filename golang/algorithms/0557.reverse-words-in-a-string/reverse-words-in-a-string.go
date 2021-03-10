package problem0557

import "strings"
//
//func reverseWords(s string) string {
//	ss := strings.Split(s, " ")
//	for i, s := range ss {
//		ss[i] = revers(s)
//	}
//	return strings.Join(ss, " ")
//}
//
//func revers(s string) string {
//	bytes := []byte(s)
//	i, j := 0, len(bytes)-1
//	for i < j {
//		bytes[i], bytes[j] = bytes[j], bytes[i]
//		i++
//		j--
//	}
//	return string(bytes)
//}

func reverseWords(s string) string {
	tr := strings.Split(s, " ")
	for i, ss := range tr{
		tr[i] = revers(ss)
	}
	return strings.Join(tr, " ")
}

// 反转每个字符
func revers(s string) string {
	tr := strings.Split(s, "")
	var tm []string
	for _, ss := range tr{
		tm = append([]string{ss},tm...)
	}
	return strings.Join(tm, "")
}