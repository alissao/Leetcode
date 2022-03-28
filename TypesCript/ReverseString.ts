/**
 Do not return anything, modify s in-place instead.
 */
function reverseString(s: string[]): void {
  if (s.length == 1) {
    return;
  }

  let end = s.length;
  let start = 0;

  return swapItems(s, start, end);
};

function swapItems(s: string[], start: number, end: number): void {
  let startString = s[start];
  s[start] = s[end - 1];
  s[end - 1] = startString;
  if (start + 1 >= end - 1) {
    return;
  } else {
    swapItems(s, start + 1, end - 1);
  }
}