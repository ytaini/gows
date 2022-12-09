/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 04:42:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 05:25:14
 */
package strategy

// 抽象策略接口
type Comparator[T any] interface {
	Compare(T, T) int
}
