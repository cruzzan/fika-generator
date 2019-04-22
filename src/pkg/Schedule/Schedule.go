package Schedule

import (
	. "github.com/cruzzan/fika-generator/src/pkg/Rotation"
	"math"
)

type Schedule struct {
	RotationCount int // Number of rotations
	RotationSize int // Max number of people in rotation
	Rotations []Rotation
	MaxRotationMembership int // Max number of rotations a person should be in
}

func NewSchedule(rCount int, rSize int, pCount int) Schedule {
	return Schedule{
		RotationCount: rCount,
		RotationSize: rSize,
		MaxRotationMembership: maxRotationMembership(pCount, rSize, rCount),
	}
}

func (s Schedule) AddRotation(r Rotation) {
	s.Rotations = append(s.Rotations, r)
}

func maxRotationMembership(pCount int, rSize int, rCount int) int {
	return int(math.Ceil(float64(pCount/(rSize * rCount))))
}
