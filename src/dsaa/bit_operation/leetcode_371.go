package test1

func GetSum(a, b int) (c int) {
	if b == 0 {
		return a
	} else {
		c = GetSum(a^b, (a&b)<<1)
		return c
	}
}
