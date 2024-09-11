package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	l1 := m - 1
	l2 := n - 1
	insertAt := len(nums1) - 1

	if l1 < 0 {
		for l2 >= 0 && insertAt >= 0 {
			nums1[insertAt] = nums2[l2]
			insertAt--
			l2--
		}
		l1 = 0
	}

	for l1 >= 0 && l2 >= 0 {
		if nums1[l1] <= nums2[l2] {
			nums1[insertAt] = nums2[l2]
			insertAt--
			l2--
		} else {
			nums1[insertAt] = nums1[l1]
			insertAt--
			l1--
		}
	}

	for l2 >= 0 && insertAt >= 0 {
		nums1[insertAt] = nums2[l2]
		insertAt--
		l2--
	}
}
