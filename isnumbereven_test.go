package isnumbereven

import "testing"

func TestIsEven(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
		{-5, false},
	}

	for _, c := range cases {
		result := IsEven(c.input)
		if result != c.expected {
			t.Errorf("IsEven(%d) == %v, expected %v", c.input, result, c.expected)
		}
	}
}
