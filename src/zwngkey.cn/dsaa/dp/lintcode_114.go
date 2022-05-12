package dp

/**
 * @param m: positive integer (1 <= m <= 100)
 * @param n: positive integer (1 <= n <= 100)
 * @return: An integer
 */
func UniquePaths(m int, n int) int {
	//f[i][j] 表示到达(i,j)点的方式有多少种.
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				//当一个点位于网格边缘时,这时到达该点的路径只有一种.
				f[i][j] = 1
			} else {
				//到达(i,j)点的方式等于到达(i,j-1),(i-1,j)两点方式的和.
				f[i][j] = f[i][j-1] + f[i-1][j]
			}
		}
	}
	return f[m-1][n-1]
}
