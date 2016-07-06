package member

import (
	"log"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"code.olipicus.com/go_rest_api/api/utility/mongo"
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
		day := RandomDay()
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

func TestAssignDay(t *testing.T) {
	done := make(chan bool)
	var mgh mongo.Helper
	mgh.Init(Handler.MongoAddress, Handler.DBName)
	defer mgh.Close()

	callFunc := func(id string) {
		obj := AssignDay(id, &mgh)
		done <- true
		log.Printf("%v", obj)
	}
	go callFunc("577bc7ed8d152a6b726b8c8f")
	go callFunc("577bc7f88d152a6b726b8c90")
	go callFunc("577bc8038d152a6b726b8c91")
	go callFunc("577bc80d8d152a6b726b8c92")
	go callFunc("577bc8168d152a6b726b8c93")
	go callFunc("577bc7ed8d152a6b726b8c8f")
	go callFunc("577bc7f88d152a6b726b8c90")
	go callFunc("577bc8038d152a6b726b8c91")
	go callFunc("577bc80d8d152a6b726b8c92")
	go callFunc("577bc8168d152a6b726b8c93")
	go callFunc("577bc7ed8d152a6b726b8c8f")
	go callFunc("577bc7f88d152a6b726b8c90")
	go callFunc("577bc8038d152a6b726b8c91")
	go callFunc("577bc80d8d152a6b726b8c92")
	go callFunc("577bc8168d152a6b726b8c93")
	go callFunc("577bc7ed8d152a6b726b8c8f")
	go callFunc("577bc7f88d152a6b726b8c90")
	go callFunc("577bc8038d152a6b726b8c91")
	go callFunc("577bc80d8d152a6b726b8c92")
	go callFunc("577bc8168d152a6b726b8c93")
	go callFunc("577bc7ed8d152a6b726b8c8f")
	go callFunc("577bc7f88d152a6b726b8c90")
	go callFunc("577bc8038d152a6b726b8c91")
	go callFunc("577bc80d8d152a6b726b8c92")
	go callFunc("577bc8168d152a6b726b8c93")
	go callFunc("577bc7ed8d152a6b726b8c8f")
	go callFunc("577bc7f88d152a6b726b8c90")
	go callFunc("577bc8038d152a6b726b8c91")
	go callFunc("577bc80d8d152a6b726b8c92")
	go callFunc("577bc8168d152a6b726b8c93")

	resultCollection := mgh.GetCollecitonObj("result")

	for i := 1; i <= 30; i++ {
		<-done
	}

	dayList := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	for _, strDay := range dayList {
		count, _ := resultCollection.Find(bson.M{"day": strDay}).Count()

		if count != 6 {
			t.Errorf("Result is not 6 [%v : %v]", strDay, count)
		}
	}

}
