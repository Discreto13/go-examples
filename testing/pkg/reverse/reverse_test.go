package reverse

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"aibohphobia", "aibohphobia"},
		{"Hello, 世界", "界世 ,olleH"},
	}
	for _, test := range testCases {
		if output := Reverse(test.input); output != test.expectedOutput {
			t.Errorf("incorrect output for %s: expected %s but got %s", test.input, test.expectedOutput, output)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"aibohphobia", "aibohphobia"},
		{"Hello, 世界", "界世 ,olleH"},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			_ = Reverse(test.input)
		}
	}
}
