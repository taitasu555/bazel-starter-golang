package stringutil_test

import (
	"embed"
	"testing"

	stringutil "github.com/taitasu555/bazel-starter-golang/pkg/string"
)

//go:embed testdata/*
var testdata embed.FS

func TestString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
	}{
		{
			name: "empty",
		},
		{
			name: "single",
		},
		{
			name: "multiple",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			givenBytes, err := testdata.ReadFile("testdata/" + tt.name + "_string_input.txt")
			if err != nil {
				t.Fatal(err)
			}
			expectedBytes, err := testdata.ReadFile("testdata/" + tt.name + "_string_output.txt")
			if err != nil {
				t.Fatal(err)
			}
			given := string(givenBytes)
			expected := string(expectedBytes)

			// Exercise
			actual := stringutil.Reverse(string(given))

			// Verify
			if actual != expected {
				t.Errorf("expected: %s, actual: %s", expected, actual)
			}
		})
	}
}
