package lc04

// FindMedianSortedArrays1
// 根据中位数的定义，当 m+n 是奇数时，中位数是两个有序数组中的第 (m+n)/2 个元素，
// 当 m+n 是偶数时，中位数是两个有序数组中的第 (m+n)/2 个元素和第 (m+n)/2+1 个元素的平均值。
// 因此，这道题可以转化成寻找两个有序数组中的第 k 小的数，其中 kk 为 (m+n)/2 或 (m+n)/2+1。
func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKth(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2, totalLength/2-1
		return float64(getKth(nums1, nums2, midIndex1)+getKth(nums1, nums2, midIndex2+1)) / 2
	}
}
func getKth(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return Min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := Min(index1+half, len(nums1)) - 1
		newIndex2 := Min(index2+half, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// FindMedianSortedArrays 通过插入排序思想,合并两个数组(归并也行)
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	tmps := make([]int, l1+l2)
	if l1 == 0 {
		copy(tmps, nums2)
	} else {
		copy(tmps, nums1)
	}
	n := l1
	for j := 0; j < l2; j++ {
		i := n - 1
		for ; i >= 0 && nums2[j] < tmps[i]; i-- {
			tmps[i+1] = tmps[i]
		}
		tmps[i+1] = nums2[j]
		n++
	}
	// 当l1+l2为奇数时,(l1+l2)/2 = (l1+l2-1)/2,
	//	(tmps[(l1+l2)/2] + tmps[(l1+l2-1)/2])/2 相当于两个相同的数字相加再除以2，还是其本身
	// 当l1+l2为偶数时,计算tmps的中位数也是用 (tmps[(l1+l2)/2] + tmps[(l1+l2-1)/2])/2
	return float64(tmps[(l1+l2)/2]+tmps[(l1+l2-1)/2]) / 2
}
