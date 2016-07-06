package member

import (
	"log"
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
	Name string `json:"name"`
	Pic  string `json:"pic"`
	Day  string `json:"day"`
}

//Members Model
type Members struct {
	Members []Member
}

//RandomDay Function
func RandomDay() string {
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
	resultCollection := mgh.GetCollecitonObj("result")
	memberCollection := mgh.GetCollecitonObj("member")
	allMember, _ := memberCollection.Find(bson.M{}).Count()
	memberHasDay, _ := resultCollection.Find(bson.M{"day": bson.M{"$exists": 1}}).Count()
	memberInDay, _ := resultCollection.Find(bson.M{"day": day}).Count()

	defer mu.Unlock()
	return allMember, memberHasDay, memberInDay
}

func getDayAvailable(id string, mgh *mongo.Helper) string {
	day := RandomDay()
	allMember, memberHasDay, memberInDay := getCount(id, day, mgh)
	maxMemberInDay := calMaxMemberInDay(allMember, memberHasDay, len(dayList))

	if memberInDay < maxMemberInDay {
		return day
	}
	return getDayAvailable(id, mgh)

}

//AssignDay Function
func AssignDay(id string, mgh *mongo.Helper) Member {

	var objMember Member
	mgh.GetOneDataToObj(Handler.Collection, id, &objMember)
	//log.Printf("%v", objMember)

	//Don't Assign Day if already have day
	if objMember.Day != "" {
		return objMember
	}

	resultCollection := mgh.GetCollecitonObj("result")

	day := getDayAvailable(id, mgh)

	mu.Lock()
	objMember.Day = day
	cM, _ := resultCollection.Find(bson.M{"day": "Monday"}).Count()
	cT, _ := resultCollection.Find(bson.M{"day": "Tuesday"}).Count()
	cW, _ := resultCollection.Find(bson.M{"day": "Wednesday"}).Count()
	cTh, _ := resultCollection.Find(bson.M{"day": "Thursday"}).Count()
	cF, _ := resultCollection.Find(bson.M{"day": "Friday"}).Count()
	log.Printf("%v %v %v %v %v", cM, cT, cW, cTh, cF)

	//mgh.UpdateData(Handler.Collection, id, objMember)
	mgh.InsertData("result", objMember)
	mu.Unlock()

	return objMember

}
