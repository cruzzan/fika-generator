package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
	. "github.com/cruzzan/fika-generator/src/pkg/Rotation"
	. "github.com/cruzzan/fika-generator/src/pkg/Schedule"
	"os"
)

var (
	app = kingpin.New("fika-generator", "Generate fika schedules for your dev department")
	names = app.Flag("name", "The name of a person in the rotation. (Repeatable)").Short('n').Required().Strings()
	rotations = app.Flag("rotations", "The number of rotations to generate").Int()
	teamSize = app.Flag("team", "The number of team members in each rotation").Int()
)

func main() {
	cmdArgs := os.Args[1:]
	kingpin.MustParse(app.Parse(cmdArgs))

	rc := checkAndDefaultRotations(*rotations)
	ts := checkAndDefaultTeamSize(*teamSize)

	s := NewSchedule(rc, ts)

	workingList := make([]string, 0)
	// Create rotation
	for i := 0; i < s.RotationCount; i++ {
		rotation := NewRotation()

		for j := 0; j < s.RotationSize; j++ {
			// Refill the working list of names if it's empty
			if len(workingList) < s.RotationSize {
				for _, n := range *names{
					workingList = append(workingList, n)
				}
			}

			workingList = rotation.AddMember(workingList)
		}

		s.Rotations = append(s.Rotations, rotation)
	}


	fmt.Printf("%v\n", s.Rotations)
}

func checkAndDefaultRotations(rCount int) int {
	if rCount == int(0) {
		rCount = int(1)
		fmt.Printf("Number of rotations was not set, defaulting to %d\n", rCount)
	}

	return rCount
}

func checkAndDefaultTeamSize(ts int) int {
	if ts == int(0) {
		ts = int(1)
		fmt.Printf("Number of team members was not set, defaulting to %d\n", ts)
	}

	return ts
}
