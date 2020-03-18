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

var maxRobots = 26 * 26 * 10 * 10 * 10
var usedRobotNames = make(map[string]bool)

// Name return a new name for the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = NewName()
		for usedRobotNames[r.name] {
			r.name = NewName()
			if len(usedRobotNames) == maxRobots {
				return "", errors.New("All names have been exhausted")
			}
		}
		usedRobotNames[r.name] = true
	}
	return r.name, nil
}

// Reset resets a robot's name
func (r *Robot) Reset() {
	r.name = ""
}

// NewName generate a new robot name
func NewName() string {
	return randomString(2) + fmt.Sprintf("%03d", (randomInt(0, 1000)))
}

// randomInt generate a random number between min and max
func randomInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}

// Generate a random string of A-Z chars with length specified
func randomString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(randomInt(65, 91))
	}
	return string(bytes)
}
