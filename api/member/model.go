package member

import (
	"math/rand"
	"time"
)

//Member Model
type Member struct {
	Name string
	Pic  string
}

//Members Model
type Members struct {
	Members []Member
}

//RandomDay Function
func RandomDay() string {
	rand.Seed(time.Now().UnixNano())
	numRandom := rand.Intn(6-1) + 1

	switch numRandom {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	}

	return ""
}
