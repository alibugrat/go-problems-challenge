package challenges

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return float64(nums[l/2])
	} else {
		sum := nums[l/2-1] + nums[l/2]
		return float64(sum) / 2
	}
}
