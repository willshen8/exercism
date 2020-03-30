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
type Team struct {
	team   string
	wins   int
	losses int
	draws  int
}

// Points calculates the team's tally points
func (t *Team) Points() int {
	return t.wins*3 + t.draws
}

// Tally takes an input from io.reader and outputs a team's results in table format via io.writer
func Tally(r io.Reader, w io.Writer) error {
	var teamResults []Team
	reader := bufio.NewReader(r)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if commentRegex.MatchString(string(line)) || string(line) == "" {
			continue
		}

		results := strings.Split(string(line), ";")
		if len(results) != 3 {
			return ErrInput
		}
		teamA := results[0]
		teamB := results[1]
		result := results[2]

		if Find(teamResults, teamA) == -1 {
			teamResults = append(teamResults, Team{team: teamA})
		}
		if Find(teamResults, teamB) == -1 {
			teamResults = append(teamResults, Team{team: teamB})
		}

		switch result {
		case "win":
			teamResults[Find(teamResults, teamA)].wins++
			teamResults[Find(teamResults, teamB)].losses++
		case "loss":
			teamResults[Find(teamResults, teamA)].losses++
			teamResults[Find(teamResults, teamB)].wins++
		case "draw":
			teamResults[Find(teamResults, teamA)].draws++
			teamResults[Find(teamResults, teamB)].draws++
		default:
			return ErrInput
		}
	}

	sort.Slice(teamResults, func(i, j int) bool { return teamResults[i].team < teamResults[j].team })
	sort.Slice(teamResults, func(i, j int) bool { return teamResults[i].Points() > teamResults[j].Points() })

	output := fmt.Sprintf("%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for _, result := range teamResults {
		mp := result.draws + result.losses + result.wins
		output += fmt.Sprintf("%-31v|%3v |%3v |%3v |%3v |%3v\n",
			result.team, mp, result.wins, result.draws, result.losses, result.Points())
	}

	if _, err := w.Write([]byte(output)); err != nil {
		fmt.Println(err)
	}
	return nil
}

// Find returns the index of the team name and -1 if not found
func Find(teams []Team, x string) int {
	for i, n := range teams {
		if x == n.team {
			return i
		}
	}
	return -1
}
