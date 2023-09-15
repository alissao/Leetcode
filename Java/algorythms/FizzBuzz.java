package Java.algorythms;

public class FizzBuzz {

  public static void main(String... args) {
    printFizzBuzz();
  }

  static void printFizzBuzz() {
    for (int i = 1; i <= 100; i++) {
      boolean multipleOfThree = i % 3 == 0;
      boolean multipleOfFive = i % 5 == 0;
      StringBuilder resultSB = new StringBuilder("");

      if (multipleOfThree) { resultSB.append("Fizz"); };

      if (multipleOfFive) { resultSB.append("Buzz"); };

      System.out.println(resultSB.length() > 0 ? resultSB.toString() : i);
    }
  }
  
}
