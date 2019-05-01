package Schedule

import (
	. "github.com/cruzzan/fika-generator/src/pkg/Rotation"
)

type Schedule struct {
	RotationCount int // Number of rotations
	RotationSize int // Max number of people in rotation
	Rotations []Rotation
}

func NewSchedule(rCount int, rSize int) Schedule {
	return Schedule{
		RotationCount: rCount,
		RotationSize: rSize,
	}
}

func (s Schedule) AddRotation(r Rotation) {
	s.Rotations = append(s.Rotations, r)
}
