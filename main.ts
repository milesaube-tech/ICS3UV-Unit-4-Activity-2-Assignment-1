/**
 * @author 
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview A do-while program that sums integers inside and outside a given range.
 */

// variables
let start: number = 0;
let end: number = 0;
let value: number = 0;
let insideSum: number = 0;
let outsideSum: number = 0;

// get the range from the user
start = Number(prompt("Enter an integer for the start of the range: ") || "0");
end = Number(prompt("Enter an integer for the end of the range: ") || "0");

// input loop
do {
  value = Number(prompt("Enter an integer or zero (0) to end: ") || "0");

  if (value !== 0) {
    if (value >= start && value <= end) {
      insideSum = insideSum + value;
    } else {
      outsideSum = outsideSum + value;
    }
  }

} while (value !== 0);

// output
console.log("The sum of the integers inside the range is " + insideSum);
console.log("The sum of the integers outside the range is " + outsideSum);
console.log("Done.");
