	/**
	* author:
	* version: 1.0.0
	* date: 2025-12-09
	* fileoverview: A while-loop program that sums integers inside and outside a given range.
	*/

	package main

	import ("fmt")

	func main() {
		// variables
		var start int
		var end int
		var value int
		var insideSum int = 0
		var outsideSum int = 0

		// get the range from the user
		fmt.Print("Enter an integer for the start of the range: ")
		fmt.Scan(&start)

		fmt.Print("Enter an integer for the end of the range: ")
		fmt.Scan(&end)

		// input loop (do-while equivalent)
		for {
			fmt.Print("Enter an integer or zero (0) to end: ")
			fmt.Scan(&value)

			if value != 0 {
				if value >= start && value <= end {
					insideSum = insideSum + value
				} else {
					outsideSum = outsideSum + value
				}
			} else {
				break
			}
		}

		// output
		fmt.Println("The sum of the integers inside the range is", insideSum)
		fmt.Println("The sum of the integers outside the range is", outsideSum)
		fmt.Println("Done.")
	}