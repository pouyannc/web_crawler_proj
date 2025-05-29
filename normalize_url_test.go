package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name string
		inputURL string
		expected string
		expectError bool
	}{
		{
			name: "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name: "remove www",
			inputURL: "https://www.blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "lowercase capital letters",
			inputURL: "https://www.BLOG.boot.dev/PATH",
			expected: "blog.boot.dev/path",
		},
		{
			name: "handle invalid URL",
			inputURL: `:\\invalid`,
			expected: "",
			expectError: true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if (err != nil) != tc.expectError {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}