package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// ErrInput throws an error when inputs has error and can't processed
var ErrInput = errors.New("Error with inputs")
var commentRegex = regexp.MustCompile(`#`)

// Team struct records team's records
type Result struct {
	mp     int
	wins   int
	losses int
	draws  int
	points int
}

// Tally takes an input from io.reader and outputs a team's results in table format via io.writer
func Tally(r io.Reader, w io.Writer) error {
	teamResults := make(map[string]*Result)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if commentRegex.MatchString(line) || line == "" {
			continue
		}

		results := strings.Split(line, ";")
		if len(results) != 3 {
			return ErrInput
		}
		teamA := results[0]
		teamB := results[1]
		result := results[2]

		if _, ok := teamResults[teamA]; !ok {
			teamResults[teamA] = &Result{}
		}
		if _, ok := teamResults[teamB]; !ok {
			teamResults[teamB] = &Result{}
		}
		teamResults[teamA].mp++
		teamResults[teamB].mp++

		switch result {
		case "win":
			teamResults[teamA].points += 3
			teamResults[teamA].wins++
			teamResults[teamB].losses++
		case "loss":
			teamResults[teamA].losses++
			teamResults[teamB].wins++
			teamResults[teamB].points += 3
		case "draw":
			teamResults[teamA].draws++
			teamResults[teamA].points++
			teamResults[teamB].draws++
			teamResults[teamB].points++
		default:
			return ErrInput
		}
	}

	sortedTeams := SortTeams(teamResults)

	fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for _, team := range sortedTeams {
		result := teamResults[team]
		fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n",
			team, result.mp, result.wins, result.draws, result.losses, result.points)
	}
	return nil
}

// SortedTeams sort a map of teams based on points and alphabetical order
func SortTeams(teamResults map[string]*Result) []string {
	teams := make([]string, 0, len(teamResults))
	for team := range teamResults {
		teams = append(teams, team)
	}
	sort.Strings(teams)

	// using bubble sort to sort
	for sorted := false; !sorted; {
		swapped := false
		for i := 0; i < len(teamResults)-1; i++ {
			if teamResults[teams[i+1]].points > teamResults[teams[i]].points {
				teams[i], teams[i+1] = teams[i+1], teams[i]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}

	}
	return teams
}
