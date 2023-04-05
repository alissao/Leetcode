package Java;

import java.util.Scanner;

import Java.algorythms.ClockAngle;

public class Runner {

  public static void main(String... args) {
    Scanner scanner = new Scanner(System.in);
    System.out.println("Enter hours: ");
    int hours = scanner.nextInt();
    System.out.println("Enter minutes: ");
    int minutes = scanner.nextInt();
    ClockAngle clockAngleCalculator = new ClockAngle();
    int angulo = clockAngleCalculator.retornaAnguloRelogio(hours, minutes);
    System.out.println(angulo);
    scanner.close();
  }

}