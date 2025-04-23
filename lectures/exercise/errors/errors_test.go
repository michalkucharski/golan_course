package timeparse

import "testing"

func TestConvStrToTime(t *testing.T) {

	table := []struct {
		time string
		ok   bool
	}{
		{"12:45:68", false},
		{"12:46:01", true},
		{"bad", false},
	}

	for _, data := range table {
		_, err := convStrToTime(data.time)

		if data.ok && err != nil {
			t.Errorf("%v:, %v:, error should be null", data.time, err)
		}
	}
}
