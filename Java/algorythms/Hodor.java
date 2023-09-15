package Java.algorythms;

public class Hodor {
  
  public static void main(String... args) {
    
    int hodorStrength = 50;
    int[] doorsStrength = {10,25,5};
    int[] zombiesInflux = {2,3,1};

    var whatDoor = holdTheDoor(hodorStrength, doorsStrength, zombiesInflux);
    System.out.println(whatDoor);
  }

  static int holdTheDoor(int hodorStrength, int[] doorsStrength, int[] zombiesInflux) {

    int easiestDoorToHold = -1;
    int longestTimeHoldingDoor = -1;
    
    for (int i = 0; i < doorsStrength.length; i++) {
      int totalStrength = hodorStrength + doorsStrength[i];
      int timeHoldingDoor = totalStrength / zombiesInflux[i];
      if (timeHoldingDoor > longestTimeHoldingDoor) {
        longestTimeHoldingDoor = timeHoldingDoor;
        easiestDoorToHold = i;
      }
    }

    return easiestDoorToHold;
  }

}
