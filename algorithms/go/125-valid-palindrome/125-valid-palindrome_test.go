package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "hello",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "long with invalid charactor",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "empty string",
			input:    " ",
			expected: true,
		},
	}

	for _, test := range tests {
		got := isPalindrome(test.input)
		if got != test.expected {
			t.Errorf("want: %v, got: %v", test.expected, got)
		}
	}
}
