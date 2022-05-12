package dp

func JumpGame(inp []int) bool {
	//inp:2,3,1,1,4

	n := len(inp)
	//f[i]表示能否跳到i位置.
	f := make([]bool, n)

	//初始化
	f[0] = true

	for j := 1; j < n; j++ {
		//先假设不能跳到j位置
		f[j] = false
		for i := 0; i < j; i++ {
			//判断i位置能否跳到,其中 i < j.
			//也就是判断j前面的各个位置中是否有能跳到的位置.
			//并通过i和j的距离与在i位置能跳跃的距离进行比较.
			//如何inp[i]>= j - i.那么表示j位置也可以跳到.
			if f[i] && i+inp[i] >= j {
				f[j] = true
				break
			}
		}
	}

	return f[n-1]
}
