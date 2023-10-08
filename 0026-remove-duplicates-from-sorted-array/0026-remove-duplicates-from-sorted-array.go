func removeDuplicates(nums []int) int {
  var leastNum, k int = nums[0], 1

  for i := 1; i < len(nums); i++ { 
      if nums[i] > leastNum {
        leastNum = nums[i]
        nums[k] = leastNum
        k += 1
      }
      continue
  }

  return k
}