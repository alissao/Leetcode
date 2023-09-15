package Java.algorythms;

public class MarvinRareNotes {
  
  public static void main(String... args) {
    int[] notes = {1, 2, 3, 2, 2, 1, 5, 5};
    // int[] notes = {2,4,3,4,2,3,4,2,3,4,3,2,4,3,2,34,2,4,3,2,4,3,2,34,7};
    var rareNote = getRareNotes(notes);
    System.out.println(rareNote[0]);
  }

  static int[] getRareNotes(int[] notes) {

    java.util.List<Integer> notesArrayList = java.util.Arrays
        .stream(notes)
        .mapToObj(Integer::valueOf)
        .collect(java.util.stream.Collectors.toList());
    int firstPointer = 0;
    boolean removeFirstPointer = false;
    int secondPointer = notesArrayList.size() - 1;
    Integer rareNote = null;

    while (firstPointer < notesArrayList.size() - 1 && firstPointer <= secondPointer) {
      if (rareNote == null) {
        rareNote = notesArrayList.get(firstPointer);
      }
      if (rareNote == notesArrayList.get(secondPointer) && firstPointer != secondPointer) {
        notesArrayList.remove(secondPointer);
        removeFirstPointer = true;
        secondPointer--;
      }
      if (firstPointer == secondPointer && removeFirstPointer) {
        notesArrayList.remove(firstPointer);
        rareNote = null;
        secondPointer = notes.length - 1;
        removeFirstPointer = false;
      } else if (firstPointer == secondPointer && !removeFirstPointer) {
        firstPointer++;
        secondPointer = notes.length - 1;
      }
    }
    
    if (notesArrayList.size() == 1) {
      return notesArrayList.stream().mapToInt(i -> i).toArray();
    } else {
      return null;
    }
  }

}
