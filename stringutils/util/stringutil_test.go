package stringutils 

import "testing"

func TestReverseString(t *testing.T){
	testCases := []struct {
		input  string
		expected int
	}{
		{"prevet", 6},
        {"", 0},
        {"12345", 5},
        {"こんにちは", 5},
		{"BIZ_BIRGEMIZ", 12},
	}
	for _, tc := range testCases{

		t.Run(tc.input, func(t *testing.T){

			actual := Counting(tc.input)
			if actual != tc.expected{
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}