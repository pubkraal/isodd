package isodd

import "testing"

func TestIsOdd(t *testing.T) {
	tables := []struct {
		given  int
		expect bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{7, true},
		{-1, true},
		{-3, true},
		{-2, false},
		{12033482393, true},
		{10000000002, false},
		{9007199254740991, true},
	}

	for _, row := range tables {
		res := IsOdd(row.given)
		if res != row.expect {
			t.Errorf("IsOdd(%v) == %v. Expected %v", row.given, res, row.expect)
		}
	}
}
