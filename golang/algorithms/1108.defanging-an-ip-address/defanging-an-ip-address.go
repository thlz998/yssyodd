package problem1108

import "strings"

func defangIPaddr(address string) string {
	//a := strings.Split(address, ".")
	//return strings.Join(a, "[.]")
	return strings.ReplaceAll(address, ".", "[.]")
}