function longestCommonPrefix(strs: string[]): string {
  return returnCommonPrefix(strs, 0);
};

function returnCommonPrefix(strs: string[], commonLength: number): string {
  
  let firstWord = strs[0];
  
  if (!!firstWord) {
    let searchString = firstWord.substring(0, commonLength + 1);

    if (strs.length == 1) {
      return firstWord;
    }
    
    let filteredArray = strs.filter((word) => {
      return word.startsWith(searchString);
    })

    if (filteredArray.length == strs.length) {
      return commonLength + 1 <= firstWord.length ? returnCommonPrefix(strs, commonLength + 1) : firstWord;
    } else {
      return firstWord.substring(0, commonLength);
    }
  }
  return "";
}