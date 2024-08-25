
package main


import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expectedURL      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expectedURL: "blog.boot.dev/path",
		},
		{
			name:     "URL with slash",
			inputURL: "https://blog.boot.dev/path/",
			expectedURL: "blog.boot.dev/path",
		},
		{
			name:     "HTTP url",
			inputURL: "http://blog.boot.dev/path",
			expectedURL: "blog.boot.dev/path",
		},
		{
			name:     "HTTP url with slash",
			inputURL: "http://blog.boot.dev/path/",
			expectedURL: "blog.boot.dev/path",
		},
	}

	for i, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			normalizedURLs:= NormalizeURL(testCase.inputURL)

			if normalizedURLs != testCase.expectedURL {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, testCase.name, testCase.expectedURL, normalizedURLs)
			}
		})
	}
}