package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Match struct {
	matchResults [10]string
}

const (
	winPoints  = 3
	drawPoints = 1
	losePoints = 0
	numOfTeams = 2
)

func convToInt(match string) (int, int, error) {
	teams := strings.Split(match, ":")
	if len(teams) != numOfTeams {
		return 0, 0, errors.New("invalid match result format")
	}

	teamX, err := strconv.Atoi(teams[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting team X score: %w", err)
	}

	teamY, err := strconv.Atoi(teams[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting team Y score: %w", err)
	}

	return teamX, teamY, nil

}

func calculateTeamPoints(teamX int, teamY int) int {
	if teamX > teamY {
		return winPoints
	}

	if teamX < teamY {
		return losePoints
	}

	return drawPoints
}

func main() {
	var matchs Match = Match{matchResults: [10]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}}

	var total int = 0

	for _, matchResult := range matchs.matchResults {
		teamX, teamY, err := convToInt(matchResult)

		if err != nil {
			fmt.Println(err)
			continue // Continue with next match
		}

		total += calculateTeamPoints(teamX, teamY)
	}

	fmt.Println("Total points:", total)
}
