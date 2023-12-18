package css_test

import (
	"testing"

	"github.com/1995parham-learning/gourl/internal/css"
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
			input:     []string{"hello:world", "1:2"},
			separator: "",
			colon:     ":",
			expected:  map[string]string{"hello": "world", "1": "2"},
		},
		{
			name:      "with separator",
			input:     []string{"hello:world,1:2"},
			separator: ",",
			colon:     ":",
			expected:  map[string]string{"hello": "world", "1": "2"},
		},
		{
			name:      "with separator and colon",
			input:     []string{"hello=world&1=2"},
			separator: "&",
			colon:     "=",
			expected:  map[string]string{"hello": "world", "1": "2"},
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
