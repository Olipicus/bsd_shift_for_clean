package member

import (
	"testing"
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
		} else {
			t.Logf("Test ok, got result: '%s'", day)
		}
	}

}
