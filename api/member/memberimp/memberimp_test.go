package memberimp

import (
	"log"
	"testing"

	"code.olipicus.com/bsd_shift_for_clean/api/member/gen-go/member"
	"code.olipicus.com/bsd_shift_for_clean/api/utility/mongo"
	"gopkg.in/mgo.v2/bson"
)

func TestRandom(t *testing.T) {
	dayExpect := map[string]bool{
		"Monday":    true,
		"Tuesday":   true,
		"Wednesday": true,
		"Thursday":  true,
		"Friday":    true,
	}

	for i := 0; i < 100; i++ {
		day := randomDay()
		if !dayExpect[day] {
			t.Errorf("Test failed, got result: '%s'", day)
			break
		} else {
			t.Logf("Test ok, got result: '%s'", day)
		}
	}
}

func TestCalMaxMemberInDay(t *testing.T) {
	//Param : allMember int, memberHasDay int, dayCount int

	result := calMaxMemberInDay(4, 0, 5)
	if result != 1 {
		t.Errorf("Max Member should be 1 but get : %v", result)
	}

	result = calMaxMemberInDay(6, 0, 5)
	if result != 1 {
		t.Errorf("Max Member should be 1 but get : %v", result)
	}

	result = calMaxMemberInDay(6, 1, 5)
	if result != 1 {
		t.Errorf("Max Member should be 1 but get : %v", result)
	}

	result = calMaxMemberInDay(6, 2, 5)
	if result != 1 {
		t.Errorf("Max Member should be 1 but get : %v", result)
	}

	result = calMaxMemberInDay(6, 3, 5)
	if result != 1 {
		t.Errorf("Max Member should be 1 but get : %v", result)
	}
	result = calMaxMemberInDay(6, 4, 5)
	if result != 1 {
		t.Errorf("Max Member should be 1 but get : %v", result)
	}

	result = calMaxMemberInDay(6, 5, 5) //left one member
	if result != 2 {
		t.Errorf("Max Member should be 2 but get : %v", result)
	}

	result = calMaxMemberInDay(7, 6, 5) // left two member
	if result != 2 {
		t.Errorf("Max Member should be 2 but get : %v", result)
	}

	result = calMaxMemberInDay(7, 7, 5) // all member has day
	if result != -1 {
		t.Errorf("Max Member should error -1 but get : %v", result)
	}

	result = calMaxMemberInDay(6, 4, 5) // all member has day
	if result != 1 {
		t.Errorf("Max Member should error -1 but get : %v", result)
	}

}

func TestGetResultByDay(t *testing.T) {
	var memberService MemberService
	if _, err := memberService.GetResultByDay("Monday"); err != nil {
		t.Errorf("Got Error ")
	}

}

func TestAssignDay(t *testing.T) {
	done := make(chan bool)
	var mgh mongo.Helper
	mgh.Init(member.MongoAddress, member.DatabaseName)
	defer mgh.Close()

	callFunc := func(id string) {
		obj := assignDay(id, &mgh)
		done <- true
		log.Printf("%v", obj)
	}
	go callFunc("577e71c66d0a227f2293343d")
	go callFunc("577e71d16d0a227f2293343e")
	go callFunc("577e71da6d0a227f2293343f")
	go callFunc("577e71e86d0a227f22933440")
	go callFunc("577e71f16d0a227f22933441")
	go callFunc("577e7b1e6d0a227f22933444")
	go callFunc("577e7b336d0a227f22933445")
	go callFunc("577e7b3d6d0a227f22933446")
	go callFunc("577e7b486d0a227f22933447")
	go callFunc("577e7b626d0a227f22933448")
	go callFunc("577e7b626d0a227f22933449")
	go callFunc("577e7bb56d0a227f2293344a")
	go callFunc("577e7bb56d0a227f2293344b")
	go callFunc("577e7bb56d0a227f2293344c")
	go callFunc("577e7bb56d0a227f2293344d")
	go callFunc("577e7bb56d0a227f2293344e")
	go callFunc("577e7bb56d0a227f2293344f")
	go callFunc("577e7bb56d0a227f22933450")
	go callFunc("577e7bb56d0a227f22933451")

	resultCollection := mgh.GetCollecitonObj("member")

	for i := 1; i <= 19; i++ {
		<-done
	}

	dayList := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	for _, strDay := range dayList {
		count, _ := resultCollection.Find(bson.M{"day": strDay}).Count()

		if count > 4 {
			t.Errorf("Result > 4 [%v : %v]", strDay, count)
		}
	}

}
