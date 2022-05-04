/*
 * @Author: zwngkey
 * @Date: 2022-04-26 15:45:38
 * @LastEditTime: 2022-04-30 22:09:24
 * @Description:
 		1.defer的执行时机
		2.defer语句的注意点
*/
package godefer

//在 for 中使用 defer 产生的问题一般比较隐晦，在特定场景下就很致命，先看下面这个例子：
func Eg21() {
	/*
		for _, file := range files {
			if f, err = os.Open(file); err != nil {
				return
			}
			// 这是错误的方式，当循环结束时文件没有关闭
			// 循环结尾处的 defer 没有执行，所以文件一直没有关闭
			// why?
			// defer 语句 在 return后才执行.不会在for循环结束就执行.

			defer f.Close()
			// 对文件进行操作
			f.Process(data)
		}
	*/

	// 更好的做法是：不用defer

	/*
		for _, file := range files {
			if f, err = os.Open(file); err != nil {
				return
			}
			// 对文件进行操作
			f.Process(data)
			// 关闭文件
			f.Close()
		}
	*/

	//或者	将循环中的操作提取到一个函数中.
	/*
		func main() {
			for _, file := range files {
				write(file)
			}
		}

		func write(file string) {
		   ...
		   if f, err = os.Open(file); err != nil {
		      return
		   }
		   defer f.Close()
		   // 对文件进行操作
		   f.Process(data)
		}
	*/

}
