package main

import "fmt"

func minWindow(str1 string, str2 string) string {
	fmt.Println("str1, str2", str1, str2)

	startIndex1 := 0
	startIndex2 := 0
	minSubstring := ""
	minLen := len(str1)
	end := 1

	for startIndex1 < len(str1) {
		fmt.Println("startIndex1 start", startIndex1)

		if startIndex2 == len(str2)-1 && str1[startIndex1] == str2[startIndex2] {
			end = startIndex1 + 1

			fmt.Println("	found end!!! startIndex1, startIndex2", end, startIndex1, startIndex2)

			// go back to the beginning
			startIndex1--
			startIndex2--
			fmt.Println("	found end!!! startIndex1, startIndex2", end, startIndex1, startIndex2)

			for {
				fmt.Println("	loop startIndex1, startIndex2", startIndex1, startIndex2)

				if startIndex2 == 0 {
					fmt.Println("Shortest??? end, startIndex1, startIndex2 ", end, startIndex1, startIndex2)

					for ; startIndex1 >= 0; startIndex1-- {
						fmt.Println(" forfor end, startIndex1, startIndex2 ", end, startIndex1, startIndex2)

						if str1[startIndex1] == str2[startIndex2] {
							if sstr := str1[startIndex1:end]; len(sstr) < minLen || len(sstr) == len(str1) {
								fmt.Println("min sstr", sstr)

								minSubstring = sstr
								minLen = len(sstr)
							}
							startIndex1 = end
							startIndex2 = 0
							fmt.Println("!!11111111111111", startIndex1, startIndex2)

							goto label
						}
					}

				}

				if str1[startIndex1] == str2[startIndex2] {
					startIndex2--
				}
				startIndex1--
			}
		}
		if str1[startIndex1] == str2[startIndex2] {
			fmt.Println("str1[startIndex1] == str2[startIndex2]", startIndex1, startIndex2, str1[startIndex1], str2[startIndex2])

			startIndex2++
		}
		startIndex1++

		fmt.Println("--- endloop startIndex1, startIndex2", startIndex1, startIndex2)

	label:
		// fmt.Println("--- got to label!!! endloop startIndex1, startIndex2", startIndex1, startIndex2)

	}
	fmt.Println()

	return minSubstring
}
