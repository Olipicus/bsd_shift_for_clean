package memberimp

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"code.olipicus.com/bsd_shift_for_clean/api/member/gen-go/member"
	"code.olipicus.com/bsd_shift_for_clean/api/utility/mongo"
)

var (
	dayList = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
	}
	mu sync.Mutex
)

//MemberService : Implement
type MemberService struct {
	State string
}

func (srv MemberService) getMongoHelper() (mgh mongo.Helper) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)

	config := make(map[string]map[string]string)
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	mgh.Init(config[srv.State]["mongo_address"], config[srv.State]["db_name"])
	return
}

//GetResultByDay : implement
func (srv MemberService) GetResultByDay(day string) (result *member.ResultDay, err error) {
	mgh := srv.getMongoHelper()
	defer mgh.Close()

	var resultDay member.ResultDay
	result = &resultDay

	result.Day = day
	resultCollection := mgh.GetCollecitonObj("member")
	err = resultCollection.Find(bson.M{"day": day}).All(&result.Members)

	return
}

//GetNotAssign : implement
func (srv MemberService) GetNotAssign() (listMember []*member.Member, err error) {
	mgh := srv.getMongoHelper()
	defer mgh.Close()

	resultCollection := mgh.GetCollecitonObj("member")
	err = resultCollection.Find(bson.M{"day": bson.M{"$exists": false}}).All(&listMember)

	return
}

//AssignDay : Implement
func (srv MemberService) AssignDay(id string) (listMember []*member.Member, err error) {
	mgh := srv.getMongoHelper()
	defer mgh.Close()

	var objMember member.Member

	err = mgh.GetOneDataToObj("member", id, &objMember)

	if err == mgo.ErrNotFound {
		return
	}

	objResult := assignDay(id, &mgh)

	resultCollection := mgh.GetCollecitonObj("member")
	err = resultCollection.Find(bson.M{"day": objResult.Day}).All(&listMember)

	return

}

//GetMember : Implement
func (srv MemberService) GetMember(id string) (objMember *member.Member, err error) {
	mgh := srv.getMongoHelper()
	defer mgh.Close()

	err = mgh.GetOneDataToObj("member", id, &objMember)

	return
}

//GetResults : Implement
func (srv MemberService) GetResults() (listResult []*member.ResultDay, err error) {
	mgh := srv.getMongoHelper()
	defer mgh.Close()

	dayColor := map[int]string{
		1: "yellow",
		2: "pink",
		3: "green",
		4: "orange",
		5: "blue",
	}

	for i := 1; i <= len(dayList); i++ {
		var result member.ResultDay
		result.Day = dayList[i]
		result.Color = dayColor[i]
		resultCollection := mgh.GetCollecitonObj("member")
		err = resultCollection.Find(bson.M{"day": dayList[i]}).All(&result.Members)

		listResult = append(listResult, &result)

	}
	return
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

func assignDay(id string, mgh *mongo.Helper) (member member.Member) {

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
