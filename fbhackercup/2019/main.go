package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// dat, err := ioutil.ReadFile("leapfrog_ch__sample_input.txt")
	dat, err := os.Open("leapfrog_ch__sample_input.txt")
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)

	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	err = scanner.Err()
	check(err)

	rLine := strings.TrimSpace(txtlines[0])
	t, err := strconv.Atoi(rLine)
	check(err)

	fmt.Println("test cases amount: ", t)
	for caseNumber := 1; caseNumber <= t; caseNumber++ {
		line := strings.TrimSpace(txtlines[caseNumber])
		result := checkResult(line)
		fmt.Printf("Case # %v: %s\n", caseNumber, result)
	}
}

var regexxx = regexp.MustCompile(`([B]+)\.`)

func checkResult(input string) string {
	freeLilypads := strings.Index(input, ".")
	betaFrogsExist := strings.Index(input, "B")

	if freeLilypads == -1 || betaFrogsExist == -1 {
		return "N"
	}

	// moveBeta()

	aInd := strings.Index(input, "A")
	fmt.Println("aind: ", aInd)
	if aInd > 0 {
		input = input[aInd:]
	} else {
		return "N"
	}

	if regexxx.MatchString(input) {
		return "Y"
	}

	return checkResult(input)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
