package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
	"math"
	"math/rand"
	"os"
	"time"
)

var (
	app = kingpin.New("fika-generator", "Generate fika schedules for your dev department")
	names = app.Flag("name", "The name of a person in the rotation. (Repeatable)").Short('n').Required().Strings()
	rotations = app.Flag("rotations", "The number of rotations to generate").Int()
	teamSize = app.Flag("team", "The number of team members in each rotation").Int()
)

type rotation struct {
	members []string
} 

type schedule struct {
	rotationCount int // Number of rotations
	rotationSize int // Max number of people in rotation
	rotations []rotation
	maxRotationMembership int // Max number of rotations a person should be in
}


func main() {
	cmdArgs := os.Args[1:]
	kingpin.MustParse(app.Parse(cmdArgs))

	if *rotations == int(0) {
		*rotations = int(1)
		fmt.Printf("Number of rotations was not set, defaulting to %d\n", *rotations)
	}

	if *teamSize == int(0) {
		*teamSize = int(1)
		fmt.Printf("Number of team members was not set, defaulting to %d\n", *teamSize)
	}

	s := schedule{
		rotationCount: *rotations,
		rotationSize: *teamSize,
		maxRotationMembership: int(math.Ceil(float64(len(*names)/(*teamSize * *rotations)))),
	}

	// Create rotation
	for i := 0; i < s.rotationCount; i++ {

		possibleMembers := narrowChoices(s, *names)
		rotation := rotation{}
		fmt.Println(possibleMembers)

		for j := 0; j < s.rotationSize; j++ {
			// Seed the rand generator
			src := rand.NewSource(time.Now().UnixNano())
			rnd := rand.New(src)

			pos := rnd.Intn(len(possibleMembers))

			rotation.members = append(rotation.members, possibleMembers[pos])

			//Remove the newly added member from possible members
			possibleMembers = append(possibleMembers[:pos], possibleMembers[pos+1:]...)
			fmt.Println(possibleMembers)
		}

		s.rotations = append(s.rotations, rotation)
	}


	fmt.Printf("%v\n", s.rotations)
}

func narrowChoices(s schedule, n []string) []string {
	s = schedule{}
	var names = make([]string, 0)
	for _, name := range n {
		names = append(names, name)
	}
	return names
}
