package main

func main() {
	height := []int{2, 3, 5, 1, 8}

	maxArea := maxArea(height)

	print(maxArea)
}

func maxArea(height []int) int {
	maxA := 0
	left, right := 0, len(height)-1

	for left < right {
		length := right - left
		area := 0

		if height[left] < height[right] {
			area = length * height[left]
			left++
		} else {
			area = length * height[right]
			right--
		}

		if area > maxA {
			maxA = area
		}
	}

	return maxA
}
