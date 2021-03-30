package css_test

import (
	"testing"

	"github.com/elahe-dastan/goURL/internal/css"
)

func TestColonSeparatedStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []string
		expected map[string]string
	}{
		{
			name:     "simple",
			input:    []string{"hello:world", "salam:donya"},
			expected: map[string]string{"hello": "world", "salam": "donya"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := css.ColonSeparatedStrings(test.input).ToMap()
			if err != nil {
				t.Fatalf("ToMap(%s) has error: %s", test.input, err)
			}

			for expectedKey, expectedValue := range test.expected {
				if result[expectedKey] != expectedValue {
					t.Fatalf("ToMap(%s)[%s] has %s instead of %s", test.input, expectedKey, result[expectedKey], expectedValue)
				}
			}
		})
	}
}
