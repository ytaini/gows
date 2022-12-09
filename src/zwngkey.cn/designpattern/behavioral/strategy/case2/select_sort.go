/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 04:39:56
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 05:39:10
 */
package strategy

// 环境类
type SelectSorter[T any] struct {
	comparator Comparator[T]
}

func NewSelectSorter[T any](comparator Comparator[T]) *SelectSorter[T] {
	return &SelectSorter[T]{
		comparator: comparator,
	}
}

func (s *SelectSorter[T]) Sort(arr []T) {
	for i := 0; i < len(arr)-1; i++ {
		k := i
		for j := i + 1; j < len(arr); j++ {
			if s.comparator.Compare(arr[k], arr[j]) > 0 {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
}
