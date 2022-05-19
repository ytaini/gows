/*
 * @Author: zwngkey
 * @Date: 2022-05-11 22:29:02
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 20:36:34
 * @Description: Go支持的几种并发同步技术。
 */
package goconcsync

/*
	什么是并发同步？
		并发同步是指如何控制若干并发计算（在Go中，即协程），从而
		  避免在它们之间产生数据竞争的现象；
		  避免在它们无所事事的时候消耗CPU资源。

*/
