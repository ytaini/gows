package dp

func ClimbStairs(n int) int {
	//f[i]表示爬到第i个台阶的方法数.
	f := make([]int, 46)

	f[0] = 1
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
