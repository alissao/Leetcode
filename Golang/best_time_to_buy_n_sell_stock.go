package leetcode

func maxProfit(prices []int) int {
  var lowest, highest, lowestIndex, highestProfit = -1, -1, 0, 0

  for x, el := range prices {
    if (el < lowest || lowest == -1) {
      if (highest != -1 && highest - lowest > highestProfit) {
        highestProfit = highest - lowest 
      }
      lowest = el
      lowestIndex = x
      highest = -1
    } else if (el > highest && el > lowest && x > lowestIndex) {
      highest = el
      if (highest - lowest > highestProfit) {
        highestProfit = highest - lowest
      }
    }
  }

  return highestProfit
}