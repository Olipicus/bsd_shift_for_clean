package member

import "testing"

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
		t.Errorf("Max Member should be 4 but get : %v", result)
	}

	result = calMaxMemberInDay(6, 0, 5)
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

}
