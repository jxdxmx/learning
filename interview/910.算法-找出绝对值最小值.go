package interview

// 2022-07-10

func FindMin_910() int {
	var nums = []int{-1, 2, 3}
	if len(nums) == 1 {
		return nums[0]
	}
	var i = 0
	for nums[i] < 0 {
		i++
	}
	if i == 0 {
		i++
	}
	if nums[i]+nums[i-1] > 0 {
		return nums[i-1]
	} else {
		return nums[i]
	}
}
