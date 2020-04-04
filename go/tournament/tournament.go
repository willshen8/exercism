package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Result records a team's records in the tournament
type Result struct {
	matches int
	wins    int
	losses  int
	draws   int
	points  int
}

// TeamPoints is a struct used to help with sorting by both team and points
type TeamPoints struct {
	team   string
	points int
}

// Tally takes an input from io.reader and outputs a team's results in table format via io.writer
func Tally(r io.Reader, w io.Writer) error {
	teamResults := make(map[string]Result)
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
		teamA := results[0]
		teamB := results[1]
		result := results[2]

		a, b := teamResults[teamA], teamResults[teamB]
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

	teamPoints := make([]TeamPoints, 0, len(teamResults))
	for team, result := range teamResults {
		teamPoints = append(teamPoints, TeamPoints{team: team, points: result.points})
	}

	sort.Slice(teamPoints, func(i, j int) bool {
		if teamPoints[i].points != teamPoints[j].points {
			return teamPoints[i].points > teamPoints[j].points
		}
		return teamPoints[i].team < teamPoints[j].team
	})

	fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range teamPoints {
		result := teamResults[v.team]
		fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n",
			v.team, result.matches, result.wins, result.draws, result.losses, result.points)
	}
	return nil
}
