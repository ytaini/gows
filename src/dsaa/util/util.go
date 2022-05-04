package util

//返回两数的最大公约数
//gcd(a,b)=gcd(b,a%b)
//gcd(a,0)=a
//gcd(0,b)=b
func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}

func GCDv2(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

//返回两数最小值
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//返回两数最大值
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//返回int切片中的最大值
func MaxIntSlice(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res

}
