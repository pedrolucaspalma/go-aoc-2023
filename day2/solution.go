package day2

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Subset struct {
	Red   int
	Green int
	Blue  int
}

func (subsetPtr *Subset) setResultFromBallString(ballString string) {
	var ball string
	var numberString string

	for _, c := range ballString {
		if string(c) == " " {
			continue
		}

		if _, err := strconv.ParseInt(string(c), 10, 64); err == nil {
			numberString += string(c)
			continue
		}

		ball += string(c)
	}

	castedNumber, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		return
	}

	switch {
	case ball == "red":
		subsetPtr.Red += int(castedNumber)
	case ball == "blue":
		subsetPtr.Blue += int(castedNumber)
	case ball == "green":
		subsetPtr.Green += int(castedNumber)
	}
}

type game []Subset

func Solution(input string, isTest bool) (int, error) {
	text, err := getInput(input, isTest)
	if err != nil {
		return 0, err
	}

	games := getGames(text)
	sum := returnValidGamesSum(&games)
	return sum, nil
}

func getGames(text string) []game {
	var gameStrings = splitTextIntoGameStrings(text)
	var games []game

	for _, gameString := range gameStrings {
		games = append(games, getGameFromGameString(gameString))
	}

	return games
}

func splitTextIntoGameStrings(text string) []string {
	var separatedGames = strings.Split(text, "\n")
	var stringsToReturn []string
	for _, gameString := range separatedGames[:len(separatedGames)-1] {
		// all strings are expected to be ASCII chars
		separatedGameString := strings.Split(gameString, ":")
		gameStringWithoutInitialPart := separatedGameString[1]
		stringsToReturn = append(stringsToReturn, gameStringWithoutInitialPart)
	}
	return stringsToReturn
}

func getGameFromGameString(gameString string) game {
	var roundsStrings = strings.Split(gameString, ";")
	var game game

	for _, roundString := range roundsStrings {
		var ballStrings []string = strings.Split(roundString, ",")

		var roundResult Subset
		for _, ballString := range ballStrings {
			roundResult.setResultFromBallString(ballString)
		}
		game = append(game, roundResult)
	}
	return game
}

var SolutionMaxValues = Subset{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func returnValidGamesSum(gamesPtr *[]game) int {
	var sum int
	for gameIdx, game := range *gamesPtr {
		var allRoundsAreValid = true
		for _, round := range game {
			if !(round.Green <= SolutionMaxValues.Green &&
				round.Blue <= SolutionMaxValues.Blue &&
				round.Red <= SolutionMaxValues.Red) {
				allRoundsAreValid = false
				break
			}
		}
		if allRoundsAreValid {
			sum += gameIdx + 1
		}
	}
	return sum
}

func getInput(input string, isTest bool) (string, error) {
	var basePath string

	if isTest == true {
		basePath = "./"
	} else {
		basePath = "./day2/"
	}

	var filePath = basePath + "input-" + input + ".md"

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return string(fileBytes), err
	}

	return string(fileBytes), nil
}
