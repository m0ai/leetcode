func removeDuplicates(nums []int) int {
  var lastNum, k int = nums[0], 1
  for i := 1; i < len(nums); i++ { 
      if nums[i] > lastNum {
        lastNum = nums[i]
        nums[k] = lastNum
        k ++
      }
      continue
  }

  return k
}