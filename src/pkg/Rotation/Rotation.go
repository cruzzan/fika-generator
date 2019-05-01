package Rotation

import (
	"math/rand"
	"time"
)

type Rotation struct {
	Members []string
}

func NewRotation() Rotation {
	return Rotation{}
}

func (r *Rotation) AddMember(names []string) []string {
	for {
		n, pos := randomMember(names)

		if !r.hasMember(n) {
			r.Members = append(r.Members, n)

			//Remove the newly added member from possible members
			names = append(names[:pos], names[pos+1:]...)

			break
		}
	}
	return names
}

func (r Rotation) hasMember(n string) bool {
	for _, m := range r.Members {
		if m == n {
			return true
		}
	}
	return false
}

func randomMember(names []string) (string, int) {
	// Seed the rand generator
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	pos := rnd.Intn(len(names))

	return names[pos], pos
}
