package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Robot has a unique name as defined by the pattern #AA111
type Robot struct {
	name string
}

var robots = InitialiseRobotNames()

// Name return a new name for the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(robots) == 0 { // exhausted all names and therefore start all over again
			return "", errors.New("All names have been exhausted")
		}
		randomNumIndex := randomInt(0, len(robots))
		r.name = robots[randomNumIndex]
		robots = RemoveIndex(robots, randomNumIndex)

	}
	return r.name, nil
}

// Reset resets a robot's name
func (r *Robot) Reset() {
	r.name = ""
}

// InitialiseRobotNames generates a slice with all possible robot combinations
// when we randomly select a robot name, we just choose a random number and use it as an index
func InitialiseRobotNames() []string {
	var robots []string
	for i := 65; i < 91; i++ {
		for j := 65; j < 91; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					for m := 0; m < 10; m++ {
						newRobot := string(byte(i)) + string(byte(j)) + strconv.Itoa(k) + strconv.Itoa(l) + strconv.Itoa(m)
						robots = append(robots, newRobot)
					}
				}
			}
		}
	}
	return robots
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RemoveIndex removes the element at index position specified and return a new slice
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
