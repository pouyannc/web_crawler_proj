package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct{
		name string
		inputHTML string
		inputURL string
		expected []string
		expectErr bool
	}{
		{
			name: "absolute and relative URLs",
			inputHTML: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			inputURL: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name: "no href",
			inputHTML: `
<html>
	<body>
		<a>
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			inputURL: "https://blog.boot.dev",
			expected: nil,
		},
		{
			name: "invalid href URL",
			inputHTML: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			inputURL: "https://blog.boot.dev",
			expected: nil,
		},
		{
			name: "bad HTML",
			inputHTML: `
<html body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
</html body>
`,
			inputURL: "https://blog.boot.dev",	
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name: "couldn't parse base URL",
			inputHTML: `
<html body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
</html body>
`,
			inputURL: `:\\invalidURL`,
			expected: nil,
			expectErr: true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputHTML, tc.inputURL)
			if (err != nil) != tc.expectErr {
				if tc.expectErr {
					t.Errorf("Test %v - '%s' FAIL: expected error", i, tc.name)
					return
				}
				t.Errorf("Test %v - '%s' FAIL: error occured: %v", i, tc.name, err)
				return
			}
			
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}


}