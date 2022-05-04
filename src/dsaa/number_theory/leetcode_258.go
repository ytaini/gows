package nt

/*

在数学中，数根(又称位数根或数字根Digital root)是自然数的一种性质，换句话说，每个自然数都有一个数根。
	数根是将一正整数的各个位数相加（即横向相加），若加完后的值大于10的话，
	则继续将各位数进行横向相加直到其值小于十为止[1]，或是，将一数字重复做数字和，直到其值小于十为止，则所得的值为该数的数根。

树根用途:
	数根可以计算模运算的同余，对于非常大的数字的情况下可以节省很多时间。
	数字根可作为一种检验计算正确性的方法。例如，两数字的和的数根等于两数字分别的数根的和。
	另外，数根也可以用来判断数字的整除性，如果数根能被3或9整除，则原来的数也能被3或9整除。

*/

func AddDigits(num int) int {
	return (num-1)&9 + 1
}

func AddDigits2(num int) (sum int) {

	for num != 0 {
		sum += num % 10
		num /= 10
	}
	if sum > 9 {
		sum = AddDigits(sum)
	}

	return
}

func AddDigits1(num int) (sum int) {

	if num < 10 {
		return num
	}

	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return AddDigits(sum)
}
