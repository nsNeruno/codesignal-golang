package arrays

type Arrays struct {
}

// https://app.codesignal.com/interview-practice/task/pMvymcahZ8dY4g75q/description
func (_a Arrays) FirstDuplicate(a []int) int {
	l := len(a)

	c := make([]int, l+1)

	for i := 0; i < l; i++ {
		c[a[i]]++
		if c[a[i]] == 2 {
			return a[i]
		}
	}

	return -1
}

// https://app.codesignal.com/interview-practice/task/uX5iLwhc6L5ckSyNC/solutions
func (_a Arrays) FirstNotRepeatingCharacter(s string) string {
	mc := make(map[rune]int)
	for _, c := range s {
		mc[c]++
	}
	for _, c := range s {
		if mc[c] == 1 {
			return string(c)
		}
	}
	return "_"
}

// https://app.codesignal.com/interview-practice/task/5A8jwLGcEpTPyyjTB
func (_a Arrays) RotateImage(a [][]int) [][]int {
	l := len(a)
	r := make([][]int, l)
	for i := 0; i < l; i++ {
		r[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			r[j][l-1-i] = a[i][j]
		}
	}
	return r
}
