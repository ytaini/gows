/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-01 16:16:32
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-01 17:16:58
 */
package interpolation

func BinarySearchNoRecur(seq []int, target int) (index int) {
	left := 0
	right := len(seq) - 1
	for left <= right {
		var mid int
		if left != right {
			mid = left + (right-left)*(target-seq[left])/(seq[right]-seq[left])
		} else {
			mid = right
		}
		if seq[mid] == target {
			return mid
		}
		if seq[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
