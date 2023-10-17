for i in range(100):
  multiple_of_three = i % 3 == 0
  multiple_of_five = i % 5 == 0

  result = ''

  if (multiple_of_three):
    result += 'Fizz'
  
  if (multiple_of_five):
    result += 'Buzz'

  print(result if len(result) > 0 else i)