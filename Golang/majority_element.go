package leetcode

type NumCounter struct {
  Number  int
  Counter int
}

func majorityElement(nums []int) int {
  var majEl NumCounter = NumCounter{-1, 0}
  var i, numsLen int = 0, len(nums)
  var j = numsLen - 1
  var elegible = numsLen/2

  var tempNumCounter = NumCounter{nums[i], 0}

  for ;i < numsLen && j >= 0; {
    
    if tempNumCounter.Number == nums[j] {
      tempNumCounter.Counter++
    }
    if i < numsLen && j == 0 {
      if tempNumCounter.Counter > elegible && tempNumCounter.Counter > majEl.Counter {
        majEl = tempNumCounter
      }
      j = numsLen - 1
      i++
      if i < numsLen {
        tempNumCounter = NumCounter{nums[i], 0}
      }
      continue
    }
    j--
  }

  return majEl.Number
}