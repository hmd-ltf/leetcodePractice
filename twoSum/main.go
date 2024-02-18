package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)

	print(res[0], res[1])
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for i, v := range nums {
		if ii, ok := mp[target-v]; ok {
			return []int{ii, i}
		} else {
			mp[v] = i
		}
	}

	return []int{0}
}
