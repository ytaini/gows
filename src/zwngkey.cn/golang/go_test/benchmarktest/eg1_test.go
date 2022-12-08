/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 05:29:46
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-08 07:06:35
 */
package benchmarktest

import (
	"strings"
	"testing"
	"time"
)

func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
func Benchmark(b *testing.B) {
	b.StopTimer()
	time.Sleep(1 * time.Second) // 假设需要做一些耗时的无关操作
	// b.ResetTimer() // 重置计时器
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Split("沙河有啥又有何", "有")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	b.ReportAllocs()
	b.SetParallelism(2)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Split("沙河有啥又有何", "有")
		}
	})

}
