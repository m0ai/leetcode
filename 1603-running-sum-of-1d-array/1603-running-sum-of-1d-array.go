func runningSum(nums []int) []int {
    for i, n := range nums {
        if i == 0 {
            continue
        }
        nums[i] = nums[i-1] + n
    }
    return nums
}