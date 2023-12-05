package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const redCount = 12
const greenCount = 13
const blueCount = 14

func main() {
	file, err := os.Open("Day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sumValidGames := 0
	power := 0

	for scanner.Scan() {
		gameValid := true

		inputLine := scanner.Text()

		gameNumber := getGameNumber(inputLine)

		semicolonSeparatedGame := getSemicolonSeparated(inputLine)

		minRed := 0
		minGreen := 0
		minBlue := 0

		for i := 0; i < len(semicolonSeparatedGame); i++ {
			commaSeparatedGame := getCommaSeparated(semicolonSeparatedGame[i])

			for j := 0; j < len(commaSeparatedGame); j++ {

				colour := string(strings.Split(commaSeparatedGame[j], " ")[1])
				switch colour {
				case "red":
					valueRed := strings.Split(commaSeparatedGame[j], " ")[0]

					if isDigit(string(valueRed)) {
						number, err := strconv.Atoi(valueRed)
						if err != nil {
							fmt.Println(err)
						}

						if number > redCount {
							gameValid = false
						}

						if number > minRed {
							minRed = number
						}
					}

				case "green":
					valueRed := strings.Split(commaSeparatedGame[j], " ")[0]

					if isDigit(string(valueRed)) {
						number, err := strconv.Atoi(valueRed)
						if err != nil {
							fmt.Println(err)
						}

						if number > greenCount {
							gameValid = false
						}

						if number > minGreen {
							minGreen = number
						}
					}

				case "blue":
					valueRed := strings.Split(commaSeparatedGame[j], " ")[0]

					if isDigit(string(valueRed)) {
						number, err := strconv.Atoi(valueRed)
						if err != nil {
							fmt.Println(err)
						}

						if number > blueCount {
							gameValid = false
						}

						if number > minBlue {
							minBlue = number
						}
					}

				}
			}
		}
		if gameValid {
			sumValidGames += gameNumber
		}
		power += (minBlue * minRed * minGreen)

	}

	fmt.Println(sumValidGames)
	fmt.Println(power)

	file.Close()
}

func getGameNumber(s string) int {
	gameNumber := strings.Split((strings.Split(s, ":")[0]), " ")[1]

	if isDigit(string(gameNumber)) {
		number, err := strconv.Atoi(gameNumber)

		if err != nil {
			fmt.Println(err)
		}
		return number
	}
	return 0
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getSemicolonSeparated(s string) []string {
	semicolonSeparated := strings.Split((strings.Split(s, ":")[1]), ";")
	return semicolonSeparated
}

func getCommaSeparated(s string) []string {
	commaSeparated := strings.Split(s, ",")

	for k := 0; k < len(commaSeparated); k++ {
		commaSeparated[k] = strings.TrimSpace(commaSeparated[k])
	}
	return commaSeparated
}
