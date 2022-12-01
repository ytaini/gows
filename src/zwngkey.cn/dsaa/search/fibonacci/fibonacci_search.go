/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-01 17:22:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-01 20:17:32
 */
package fibonacci

var max_size = 31

func fib(f []int) {
	f[0] = 0
	f[1] = 1
	for i := 2; i < max_size; i++ {
		f[i] = f[i-1] + f[i-2]
	}
}

func FibonacciSearch(seq []int, target int) int {
	sLen := len(seq)

	f := make([]int, max_size)
	fib(f) //构造一个斐波那契数组f = [0,1,1,2,3,5,8,13,...]

	k := 0

	//在斐波那契数列找一个等于略大于查找表中元素个数的数f[k]
	for sLen > f[k]-1 {
		k++
	}

	temp := make([]int, sLen)
	copy(temp, seq)

	// 要求开始表中记录的个数为某个斐波那契数小1，即len(temp)=F(k)-1;
	for i := sLen; i < f[k]-1; i++ {
		// (如果要补充元素，则补充重复最后一个元素，直到满足F[n]-1个元素)
		temp = append(temp, seq[sLen-1])
	}

	left := 0
	right := sLen - 1

	// 开始将target值与第F(k-1)位置的记录进行比较(即mid=low+F(k-1)-1)
	for left <= right {
		mid := left + f[k-1] - 1
		if target < temp[mid] {
			right = mid - 1
			// k-=1 说明范围[low,mid-1]内的元素个数为F(k-1)-1个
			k -= 1
		} else if target > temp[mid] {
			left = mid + 1
			// k-=2 说明范围[mid+1,left]内的元素个数为
			// n-(F(k-1))=F(k)-1-F(k-1)=F(k)-F(k-1)-1=F(k-2)-1个
			k -= 2
		} else {
			if mid < sLen {
				return mid
			} else {
				return sLen - 1
			}
		}
	}
	return -1
}
