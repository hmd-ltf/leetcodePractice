package main

func main() {
	nums := []int{1, 1, 3}

	unique := removeDuplicates(nums)

	print(unique)
}

func removeDuplicates(nums []int) int {
	ii := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[ii] {
			ii++
			nums[ii] = nums[i]
		}
	}

	return ii + 1
}
