package member

import (
	"log"
	"math"
	"math/rand"
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

func calMaxMemberInDay(allMember int, memberHasDay int, memberInDay int, dayCount int) int {
	//log.Printf("allMember : %v , memberHasDay : %v , memberInDay : %v , dayCount : %v", allMember, memberHasDay, memberInDay, dayCount)
	//log.Printf("%v >= %v(%v/%v)  %v", memberHasDay, dayCount, allMember, dayCount, allMember/dayCount)

	avgMemberInDay := float64(allMember) / float64(dayCount)
	if allMember <= dayCount {
		return allMember
	} else if memberHasDay >= dayCount*int(avgMemberInDay) {
		return int(math.Ceil(avgMemberInDay))
	} else {
		return int(math.Floor(avgMemberInDay))
	}
}

//AssignDay Function
func AssignDay(id string, day string) Member {
	var mgh mongo.Helper
	mgh.Init(Handler.MongoAddress, Handler.DBName)

	defer mgh.Close()

	var objMember Member

	mgh.GetOneDataToObj(Handler.Collection, id, &objMember)
	log.Printf("%v", objMember)

	mgoCollection := mgh.GetCollecitonObj(collection)

	allMember, _ := mgoCollection.Find(bson.M{}).Count()
	memberHasDay, _ := mgoCollection.Find(bson.M{"day": bson.M{"$exists": 1}}).Count()
	memberInDay, _ := mgoCollection.Find(bson.M{"day": day}).Count()

	maxMemberInDay := calMaxMemberInDay(allMember, memberHasDay, memberInDay, len(dayList))

	//log.Printf("maxPersonInDay : %v, countPersonInDay : %v", maxPersonInDay, countPersonInDay)
	if memberInDay < maxMemberInDay {
		//Assign Day to Member
		objMember.Day = day

	} else {
		//Random Again
		objMember = AssignDay(id, RandomDay())

	}
	return objMember

}
