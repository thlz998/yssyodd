package app

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ans := strs[0]
	for i := range strs {
		println(i, "TTTTTTTTTTTT")
		for j := 0; j < len(ans) && j < len(strs[i]); j++ {
			println(i, j)
			println(ans[j:j+1] != strs[i][j:j+1])
			println("ans[j]", ans[j:j+1])
			println("strs[i][j]", ans[j:j+1])
			println("\n")
			if ans[j:j+1] != strs[i][j:j+1] {
				break
			}
			println(j, "===================")

			// ans = ans[0 : j+1]
			if ans == "" {
				return ans
			}
		}
	}

	return ans
}
