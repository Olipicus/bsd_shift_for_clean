package memberimp

import (
	"log"
	"testing"

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
	memberService := MemberService{State: "develop"}
	if _, err := memberService.GetResultByDay("Monday"); err != nil {
		t.Errorf("Got Error ")
	}

}

func TestAssignDay(t *testing.T) {
	done := make(chan bool)
	srv := MemberService{State: "develop"}
	mgh := srv.getMongoHelper()
	defer mgh.Close()

	callFunc := func(id string) {
		obj := assignDay(id, &mgh)
		done <- true
		log.Printf("%v", obj)
	}
	go callFunc("586235273b848801bf4630a7")
	go callFunc("586235393b848801bf4630a8")
	go callFunc("586235443b848801bf4630a9")
	go callFunc("586235533b848801bf4630aa")
	go callFunc("586235653b848801bf4630ab")
	go callFunc("5862356b3b848801bf4630ac")
	go callFunc("5862356d3b848801bf4630ad")
	go callFunc("5862356e3b848801bf4630ae")
	go callFunc("586235733b848801bf4630af")
	go callFunc("5862357a3b848801bf4630b0")
	go callFunc("586235853b848801bf4630b1")
	go callFunc("586235863b848801bf4630b2")
	go callFunc("586235963b848801bf4630b4")
	go callFunc("5862359b3b848801bf4630b5")
	go callFunc("586235b43b848801bf4630b6")
	go callFunc("586235b73b848801bf4630b7")
	go callFunc("586236153b848801bf4630b8")
	go callFunc("5862366d3b848801bf4630b9")
	go callFunc("5862367b3b848801bf4630ba")
	go callFunc("586236b13b848801bf4630bb")
	go callFunc("5888421febb6dae792ffdfd1")
	go callFunc("5921b9198a67c9071e1d1fe9")

	resultCollection := mgh.GetCollecitonObj("member")

	for i := 1; i <= 22; i++ {
		<-done
	}

	dayList := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	for _, strDay := range dayList {
		count, _ := resultCollection.Find(bson.M{"day": strDay}).Count()

		if count > 5 {
			t.Errorf("Result > 4 [%v : %v]", strDay, count)
		}
	}

}
