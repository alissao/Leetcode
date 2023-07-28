func removeElement(nums []int, val int) int {
  insertOn := len(nums) - 1
  diffElCount := 0

  for i := insertOn; i >= 0; i-- {
    if nums[i] == val {
      el := nums[insertOn]
      nums[insertOn] = nums[i]
      nums[i] = el
      insertOn--
    } else {
      diffElCount++
    }
  }
  return diffElCount
}