package main

func main() {
	nums1 := []int{0}
	m := 0
	nums2 := []int{2}
	n := 1

	merge(nums1, m, nums2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for i := m + n - 1; i >= 0; i-- {
		if m >= 0 && n >= 0 {
			if nums1[m] >= nums2[n] {
				nums1[i] = nums1[m]
				m--
			} else {
				nums1[i] = nums2[n]
				n--
			}
		} else if n >= 0 {
			nums1[i] = nums2[n]
			n--
		} else if m >= 0 {
			nums1[i] = nums1[m]
			m--
		}
	}
}
