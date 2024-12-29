package palindromelinkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    *ListNode
		expected bool
	}{
		{
			name: "list is palindrome",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "list is not palindrome",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}

	for _, test := range tests {
		got := isPalindrome(test.input)
		if got != test.expected {
			t.Errorf("got: %v, expected: %v", got, test.expected)
		}
	}
}
