package leetcode

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m == 0 {
			copy(nums1, nums2)
			return
	}
	
	indexM := m - 1
	indexN := n - 1
	i := m + n - 1

	for ; indexN >= 0; i-- {
		if indexM >= 0 && nums1[indexM] > nums2[indexN] {
			nums1[i] = nums1[indexM]
      indexM--
		} else {
			nums1[i] = nums2[indexN]
      indexN--
		}
	}
	
}