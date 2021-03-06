package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot has a unique name as defined by the pattern #AA111
type Robot struct {
	name string
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var maxRobots = 26 * 26 * 10 * 10 * 10
var usedRobotNames = make(map[string]bool)

// Name return a new name for the robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(usedRobotNames) == maxRobots {
		return "", errors.New("All names have been exhausted")
	}
	r.name = NewName()
	for usedRobotNames[r.name] {
		r.name = NewName()

	}
	usedRobotNames[r.name] = true
	return r.name, nil
}

// Reset resets a robot's name
func (r *Robot) Reset() {
	r.name = ""
}

// NewName generate a new robot name
func NewName() string {
	r1 := string(random.Intn(26) + 'A')
	r2 := string(random.Intn(26) + 'A')
	num := random.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}
