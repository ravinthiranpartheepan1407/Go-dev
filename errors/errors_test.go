package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		eval bool
	}{
		{"14:07:00", true},
		{"1:4:7", true},
		{"1:-4:7", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.eval && err != nil {
			t.Errorf("%v: %v error should be nil", data.time, err)
		}
	}
}
