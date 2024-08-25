package main

import (
	"reflect"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		expectedURL string
	}{
		{
			name:        "remove scheme",
			inputURL:    "https://blog.boot.dev/path",
			expectedURL: "blog.boot.dev/path",
		},
		{
			name:        "URL with slash",
			inputURL:    "https://blog.boot.dev/path/",
			expectedURL: "blog.boot.dev/path",
		},
		{
			name:        "HTTP url",
			inputURL:    "http://blog.boot.dev/path",
			expectedURL: "blog.boot.dev/path",
		},
		{
			name:        "HTTP url with slash",
			inputURL:    "http://blog.boot.dev/path/",
			expectedURL: "blog.boot.dev/path",
		},
	}

	for i, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			normalizedURLs := NormalizeURL(testCase.inputURL)

			if normalizedURLs != testCase.expectedURL {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, testCase.name, testCase.expectedURL, normalizedURLs)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {

	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expectedArray []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
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
			expectedArray: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "Complicated URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one/two/three/four/five">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
			<a href="https://random.com">
			<span>Random</span>
		</a>
	</body>
</html>
`,
			expectedArray: []string{"https://blog.boot.dev/path/one/two/three/four/five", 
			"https://other.com/path/one", 
			"https://random.com"},
		},
	}

	for i, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			URLArray, err := GetURLsFromHTML(testCase.inputBody, testCase.inputURL)

			if reflect.DeepEqual(URLArray, testCase.expectedArray) == false || err != nil {
				t.Errorf("Test %v | FAIL: expected URlArray: %v, actual: %v", i, testCase.expectedArray, URLArray)
			}
		})
	}

}
