package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Result records a team's records in the tournament
type team struct {
	name    string
	matches int
	wins    int
	losses  int
	draws   int
	points  int
}

// Tally takes an input from io.reader and outputs a team's results in table format via io.writer
func Tally(r io.Reader, w io.Writer) error {
	teamResults := make(map[string]team)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		results := strings.Split(line, ";")
		if len(results) != 3 {
			return fmt.Errorf("Error with reading input, received: %s", line)
		}
		teamA, teamB, result := results[0], results[1], results[2]

		a, b := teamResults[teamA], teamResults[teamB]
		a.name, b.name = teamA, teamB
		a.matches++
		b.matches++

		switch result {
		case "win":
			a.points += 3
			a.wins++
			b.losses++
		case "loss":
			a.losses++
			b.wins++
			b.points += 3
		case "draw":
			a.draws++
			a.points++
			b.draws++
			b.points++
		default:
			return fmt.Errorf("Error with game result, it should be win, loss or draw, instead got: %s", result)
		}
		teamResults[teamA], teamResults[teamB] = a, b
	}

	sortedTeamResults := make([]team, 0, len(teamResults))
	for _, result := range teamResults {
		sortedTeamResults = append(sortedTeamResults, result)
	}

	sort.Slice(sortedTeamResults, func(i, j int) bool {
		if sortedTeamResults[i].points != sortedTeamResults[j].points {
			return sortedTeamResults[i].points > sortedTeamResults[j].points
		}
		return sortedTeamResults[i].name < sortedTeamResults[j].name
	})

	fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range sortedTeamResults {
		result := teamResults[v.name]
		fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n",
			v.name, result.matches, result.wins, result.draws, result.losses, result.points)
	}
	return nil
}
