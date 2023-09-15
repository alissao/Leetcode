package Java.algorythms;

public class KyloRen {
  
  public static void main(String... args) {
    String act = "ACT";
    String cat = "CAT";
    boolean isCatAnAct = areAnagrams(act, cat);
    System.out.println("is a cat an Act? " + isCatAnAct);
    
    String tact = "TACT";
    boolean isCatAnTact = areAnagrams(tact, cat);
    System.out.println("is a cat an Tact? " + isCatAnTact);

    String kyloRen = "Kylo Ren";
    String KY = "KN Lore Y";
    boolean isKyloRenKY = areAnagrams(kyloRen, KY);
    System.out.println("is Kylo Ren an K Y? " + isKyloRenKY);

  }

  static boolean areAnagrams(String word1, String word2) {

    var treatedString1 = word1.toLowerCase().trim().replace(" ", "");
    var treatedString2 = word2.toLowerCase().trim().replace(" ", "");
    boolean isAnagram = false;

    if (treatedString1.length() != treatedString2.length()) {
      return isAnagram;
    }
    
    isAnagram = true;

    for (char charFromString1 : treatedString1.toCharArray()) {
      if (!treatedString2.contains(charFromString1 + "")) {
        isAnagram = false;
        break;
      }
      treatedString2 = treatedString2.replaceFirst(charFromString1 + "", "");
    }

    return isAnagram;
  }

}
