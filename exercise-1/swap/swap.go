package swap

func Nums(nums []int) []int {
	i := 0
	j := 1

	for j < len(nums) {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i+= 2
		j+= 2
	}

	return nums
}