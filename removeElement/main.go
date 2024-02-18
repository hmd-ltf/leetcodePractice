package main

func main() {
	nums := []int{3, 1, 1, 3, 1}

	unique := removeElement(nums, 1)

	print(unique)
}

func removeElement(nums []int, val int) int {
	ii := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ii] = nums[i]
			ii++
		}
	}

	return ii
}
