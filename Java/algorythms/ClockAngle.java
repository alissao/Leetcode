package Java.algorythms;

public class ClockAngle {
  
  // private final int minuteAngle = 6;
  // private final int hourAngle = 30;

  public int retornaAnguloRelogio(int hora, int minuto) {
    final int minuteAngle = 6;
    final int hourAngle = 30;
    int analogHour = hora > 11 ? hora - 12 : hora;
    int startingAngle = analogHour == 0 ? 360 : analogHour * hourAngle;
    int finishingAngle = minuto == 0 ? 360 : minuto * minuteAngle;
    int preResult = (360 - startingAngle) - (360 - finishingAngle);
    return preResult < 0 ? 360 + preResult : preResult;
  }


}
