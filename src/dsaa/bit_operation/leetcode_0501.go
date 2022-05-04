package test1

func InsertBits(N int, M int, i int, j int) int {
	for k := i; k <= j; k++ {
		N = N &^ (1 << k)
	}
	return N | (M << i)
}
func InsertBits2(N int, M int, i int, j int) int {
	left := N >> (j + 1)
	left = left << (j + 1)
	mid := M << i
	right := N & (1<<i - 1)
	return left | mid | right
}
