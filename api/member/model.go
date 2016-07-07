package member

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"code.olipicus.com/go_rest_api/api/utility/mongo"

	"gopkg.in/mgo.v2/bson"
)

var dayList = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
}

var mu sync.Mutex

//Member Model
type Member struct {
	_id  bson.ObjectId
	Name string `json:"name"`
	Pic  string `json:"pic"`
	Day  string `json:"day"`
}

//ResultDay Model
type ResultDay struct {
	Day     string   `json:"day"`
	Color   string   `json:"color"`
	Members []Member `json:"members"`
}

//RandomDay Function
func randomDay() string {
	rand.Seed(time.Now().UnixNano())
	numRandom := rand.Intn(6-1) + 1
	return dayList[numRandom]
}

func calMaxMemberInDay(allMember int, memberHasDay int, dayCount int) int {
	//log.Printf("allMember : %v , memberHasDay : %v , memberInDay : %v , dayCount : %v", allMember, memberHasDay, memberInDay, dayCount)
	//log.Printf("%v >= %v(%v/%v)  %v", memberHasDay, dayCount, allMember, dayCount, allMember/dayCount)

	//Prevent Error
	if memberHasDay >= allMember {
		return -1
	}

	avgMemberInDay := float64(allMember) / float64(dayCount)

	if memberHasDay >= dayCount*int(avgMemberInDay) {
		return int(math.Ceil(avgMemberInDay))
	}

	return int(math.Floor(avgMemberInDay))
}

func getCount(id string, day string, mgh *mongo.Helper) (int, int, int) {
	mu.Lock()
	memberCollection := mgh.GetCollecitonObj("member")
	allMember, _ := memberCollection.Find(bson.M{}).Count()
	memberHasDay, _ := memberCollection.Find(bson.M{"day": bson.M{"$exists": 1}}).Count()
	memberInDay, _ := memberCollection.Find(bson.M{"day": day}).Count()

	defer mu.Unlock()
	return allMember, memberHasDay, memberInDay
}

func getDayAvailable(id string, mgh *mongo.Helper) string {
	day := randomDay()
	allMember, memberHasDay, memberInDay := getCount(id, day, mgh)
	maxMemberInDay := calMaxMemberInDay(allMember, memberHasDay, len(dayList))

	if memberInDay < maxMemberInDay {
		return day
	}
	return getDayAvailable(id, mgh)

}

//AssignDay Function
func AssignDay(id string, mgh *mongo.Helper) Member {

	var member Member
	mgh.GetOneDataToObj("member", id, &member)

	//log.Printf("%v", objMember)

	//Don't Assign Day if already have day
	if member.Day != "" {
		return member
	}

	day := getDayAvailable(id, mgh)

	mu.Lock()
	member.Day = day
	mgh.UpdateData("member", id, member)
	mu.Unlock()

	return member

}
