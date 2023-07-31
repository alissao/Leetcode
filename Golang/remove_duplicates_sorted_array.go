package leetcode

func removeDuplicates(nums []int) int {
  count := 0

  for i, x := range nums {
    if i < len(nums) - 1 && x == nums[i+1] {
      continue
    }
    nums[count] = nums[i]
    count++
  }
  return count
}