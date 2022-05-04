/*
 * @Author: zwngkey
 * @Date: 2022-05-02 15:20:37
 * @LastEditTime: 2022-05-02 16:02:37
 * @Description: 为什么 Go 标准库中有些函数只有签名，没有函数体？
 */
package goother

import (
	_ "unsafe"
)

//没有函数体,编译没有报错.
func F1()

/*
	首先，函数肯定得有实现，没有函数体，一定是在其他某个地方。Go 中一般有两种形式:
		1.函数签名使用Go，然后通过该包中的汇编文件来实现它
			比如，在标准库 sync/atomic 包中的函数基本只有函数签名.比如：atomic.StoreInt32
				func StoreInt32(addr *int32, val int32)

				atomic目录下有一个文件：asm.s，它提供了具体的实现，即通过汇编来实现：
					TEXT ·StoreInt32(SB),NOSPLIT,$0
					    JMP    runtime∕internal∕atomic·Store(SB)

					具体的实现，在 runtime∕internal/atomic/ 文件夹中

		2.通过 //go:linkname 指令来实现
			比如，在标准库 time 包中的 Sleep 函数：
				func Sleep(d Duration)

			按照 Go 源码的风格，这时候一般需要去 runtime 包中找。我们会找到 time.go，其中有一个函数：
			//go:linkname timeSleep time.Sleep
			func timeSleep(ns int64) {
			    ...
			}
			这就是我们要找的 time.Sleep 的实现。

			指令格式:
			//go:linkname 函数名 包名.函数名

			注意:
				使用 //go:linkname，必须导入 unsafe 包，所以，有时候会见到：import _ "unsafe" 这样的代码。
*/
