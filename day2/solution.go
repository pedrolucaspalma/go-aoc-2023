package day2

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Subset struct {
	red   int
	green int
	blue  int
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
		subsetPtr.red += int(castedNumber)
	case ball == "blue":
		subsetPtr.blue += int(castedNumber)
	case ball == "green":
		subsetPtr.green += int(castedNumber)
	}
}

type game []Subset
type words []string

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
	for _, gameString := range separatedGames {
		// all strings are expected to be ASCII chars
		if len(gameString) < 7 {
			continue
		}
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
	red:   12,
	green: 13,
	blue:  14,
}

func returnValidGamesSum(gamesPtr *[]game) int {
	var maxNumberOfTotalBalls = SolutionMaxValues.blue + SolutionMaxValues.green + SolutionMaxValues.red
	var sum int
	for gameIdx, game := range *gamesPtr {
		var greenBallsFromGame int
		var redBallsFromGame int
		var blueBallsFromGame int

		for _, round := range game {
			greenBallsFromGame += round.green
			redBallsFromGame += round.red
			blueBallsFromGame += round.blue
		}

		if greenBallsFromGame <= SolutionMaxValues.green &&
			blueBallsFromGame <= SolutionMaxValues.blue &&
			redBallsFromGame <= SolutionMaxValues.red &&
			greenBallsFromGame+blueBallsFromGame+redBallsFromGame <= maxNumberOfTotalBalls {
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
