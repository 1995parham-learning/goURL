package css_test

import (
	"testing"

	"github.com/elahe-dastan/goURL/internal/css"
)

func TestColonSeparatedStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     []string
		colon     string
		separator string
		expected  map[string]string
	}{
		{
			name:      "simple",
			input:     []string{"hello:world", "salam:donya"},
			separator: "",
			colon:     ":",
			expected:  map[string]string{"hello": "world", "salam": "donya"},
		},
		{
			name:      "with separator",
			input:     []string{"hello:world,salam:donya"},
			separator: ",",
			colon:     ":",
			expected:  map[string]string{"hello": "world", "salam": "donya"},
		},
		{
			name:      "with separator and colon",
			input:     []string{"hello=world&salam=donya"},
			separator: "&",
			colon:     "=",
			expected:  map[string]string{"hello": "world", "salam": "donya"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := css.NewWithOptions(test.input, test.colon, test.separator).ToMap()
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
