package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("inputDay1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sumTotal := 0

	for scanner.Scan() {
		var inputLine = scanner.Text()

		var firstFound bool = false
		first := ' '
		last := ' '
		sumLine := 0

		for i := 0; i < len(inputLine); i++ {
			val := ' '
			strLength := len(inputLine)

			switch string(inputLine[i]) {
			case "o":
				if strLength >= i+3 && "one" == string(inputLine[i:i+3]) {
					val = '1'
				}
				break
			case "t":
				if strLength >= i+3 && "two" == string(inputLine[i:i+3]) {
					val = '2'
				} else if strLength >= i+5 && "three" == string(inputLine[i:i+5]) {
					val = '3'
				}
				break
			case "f":
				if strLength >= i+4 && "four" == string(inputLine[i:i+4]) {
					val = '4'
				} else if strLength >= i+4 && "five" == string(inputLine[i:i+4]) {
					val = '5'
				}
				break
			case "s":
				if strLength >= i+3 && "six" == string(inputLine[i:i+3]) {
					val = '6'
				} else if strLength >= i+5 && "seven" == string(inputLine[i:i+5]) {
					val = '7'
				}
				break
			case "e":
				if strLength >= i+5 && "eight" == string(inputLine[i:i+5]) {
					val = '8'
				}
				break
			case "n":
				if strLength >= i+4 && "nine" == string(inputLine[i:i+4]) {
					val = '9'
				}
				break
			default:
				if isDigit(string(inputLine[i])) {
					val = rune(inputLine[i])
				}
			}

			if !firstFound && val != ' ' {
				first = val
				firstFound = true
			}
			if val != ' ' {
				last = val
			}

		}

		fmt.Println(inputLine + ": " + string(first) + string(last))

		stringLine := string(first) + string(last)
		sumLine, err := strconv.Atoi(stringLine)

		if err != nil {
			fmt.Println(err)
		}

		sumTotal += sumLine
	}

	fmt.Println(sumTotal)

	file.Close()
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
