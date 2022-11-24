/*
  - @Author: zwngkey
  - @Date: 2022-05-06 23:23:48

* @LastEditors: wzmiiiiii
* @LastEditTime: 2022-11-24 17:38:32
  - @Description:
    Go 语言 select 关键字常见的现象、数据结构以及实现原理。
*/
package gochannel

/*
	Go 语言中的 select 也能够让 Goroutine 同时等待多个 Channel 可读或者可写，
		Channel状态改变之前，select 会一直阻塞Goroutine。

	Go 语言中的 select 关键字，它能够让一个 Goroutine 同时等待多个 Channel 达到准备状态。

	select 是与 switch 相似的控制结构，与 switch 不同的是，select 中虽然也有多个 case，
		但是这些 case 中的表达式必须都是 Channel 的收/发操作。
*/

/*
	当我们在 Go 语言中使用 select 控制结构时，会遇到两个有趣的现象：
		1.select 能在 Channel 上进行非阻塞的收发操作；
		2.select 在遇到多个 Channel 同时响应时，会随机执行一种情况；


	非阻塞的收发:
		在通常情况下，select 语句会阻塞当前 Goroutine 并等待多个 Channel 中的一个达到可以收发的状态。
		  但是如果 select 控制结构中包含 default 语句，那么这个 select 语句在执行时会遇到以下两种情况：
			1.当存在可以收发的 Channel 时，直接处理该 Channel 对应的 case；
			2.当不存在可以收发的 Channel 时，执行 default 中的语句；

		只要我们稍微想一下，就会发现 Go 语言设计的这个现象很合理。
			select 的作用是同时监听多个 case 是否可以执行，如果多个 Channel 都不能执行，那么运行 default 也是理所当然的。

		非阻塞的 Channel 发送和接收操作还是很有必要的，在很多场景下我们不希望 Channel 操作阻塞当前 Goroutine，
		  只是想看看 Channel 的可读或者可写状态，如下所示：
			errCh := make(chan error, len(tasks))
			wg := sync.WaitGroup{}
			wg.Add(len(tasks))
			for i := range tasks {
				go func() {
					defer wg.Done()
					if err := tasks[i].Run(); err != nil {
						errCh <- err
					}
				}()
			}
			wg.Wait()

			select {
			case err := <-errCh:
				return err
			default:
				return nil
			}

		在上面这段代码中，我们不关心到底多少个任务执行失败了，只关心是否存在返回错误的任务，最后的 select 语句能很好地完成这个任务。


	随机执行:
		另一个使用 select 遇到的情况是同时有多个 case 就绪时，select 会选择哪个 case 执行的问题.

		select 在遇到多个 case 同时满足可读或者可写条件时会随机选择一个 case 执行其中的代码。

		如果我们按照顺序依次判断，那么后面的条件永远都会得不到执行，而随机的引入就是为了避免饥饿问题的发生。
*/
/*
	数据结构:
		select 在 Go 语言的源代码中不存在对应的结构体，但是我们使用 runtime.scase 结构体表示 select 控制结构中的 case：
			type scase struct {
				c    *hchan         // chan
				elem unsafe.Pointer // data element
			}
		因为非默认的 case 中都与 Channel 的发送和接收有关，所以 runtime.scase 结构体中也包含一个 runtime.hchan 类型的字段
			用来存储 case 中使用的 Channel。

*/
/*
	实现原理:
		https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-select/#%e7%9b%b4%e6%8e%a5%e9%98%bb%e5%a1%9e
*/
