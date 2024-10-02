// package main

// import "testing"

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{
// 			name:     "Valid data file",
// 			path:     "data.txt",
// 			expected: "Linear Regression Line: y = ",
// 		},
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "Valid data file",
			filePath: "data.txt",
			expected: "Linear Regression Line: y = ",
		},
		{
			name:     "Empty data file",
			filePath: "empty_data.txt",
			expected: "provide data, file is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "main.go", tt.filePath)

			output, err := cmd.CombinedOutput()
			if err != nil {
				if !contains(string(output), tt.expected) {
					t.Errorf("Expected output to contain %q, got %q", tt.expected, string(output))
				}
				return
			}

			if !contains(string(output), tt.expected) {
				t.Errorf("Expected output to contain %q, got %q", tt.expected, string(output))
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
