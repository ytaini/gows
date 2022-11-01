/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 12:37:57
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-31 12:39:16
 * @Description:
	fmt包api
*/
package gofmt

// 读取标准输入
/*
	fmt包中提供了3类读取输入的函数：
		Scan家族：从标准输入os.Stdin中读取数据，包括Scan()、Scanf()、Scanln()
		SScan家族：从字符串中读取数据，包括Sscan()、Sscanf()、Sscanln()
		Fscan家族：从io.Reader中读取数据，包括Fscan()、Fscanf()、Fscanln()

	其中：
		Scanln、Sscanln、Fscanln在遇到换行符的时候停止
		Scan、Sscan、Fscan将换行符当作空格处理
		Scanf、Sscanf、Fscanf根据给定的format格式读取，就像Printf一样

	这3家族的函数都返回读取的记录数量，并会设置报错信息，例如读取的记录数量不足、超出或者类型转换失败等。
*/
