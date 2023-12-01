fun fizzBuzz(start: Int, end: Int) {
    for (k in start..end) {
        val multipleOfThree = k % 3 == 0
        val multipleOfFive = k % 5 == 0

        val resultBuilder = StringBuilder()
        
        if (multipleOfThree) {
            resultBuilder.append("Fizz")
        }
        
        if (multipleOfFive) {
            resultBuilder.append("Buzz")
        }

        if (resultBuilder.isEmpty()) continue

        println(resultBuilder.toString())
    }
}

fun main() {
    fizzBuzz(1, 100)
}