package css

import (
	"fmt"
	"strings"
)

// ColonSeparatedStrings represents an array of colon separated strings.
type ColonSeparatedStrings struct {
	colon     string
	separator string
	array     []string
}

// ParseErrors contains errors of colon separated strings' parsing process.
type ParseErrors []string

func (pe ParseErrors) Error() string {
	err := ""

	for _, p := range pe {
		err += fmt.Sprintln(p)
	}

	return err
}

func NewWithOptions(css []string, colon string, separator string) *ColonSeparatedStrings {
	return &ColonSeparatedStrings{
		array:     css,
		colon:     colon,
		separator: separator,
	}
}

func New(css []string) *ColonSeparatedStrings {
	return NewWithOptions(css, ":", "")
}

// ToMap creates a map from the array of colon separated strings.
func (a *ColonSeparatedStrings) ToMap() (map[string]string, error) {
	query := make(map[string]string)
	errs := make([]string, 0)
	array := make([]string, 0)

	for _, q := range a.array {
		if a.separator != "" {
			array = append(array, strings.Split(q, a.separator)...)
		} else {
			array = append(array, q)
		}
	}

	for _, q := range array {
		keyValue := strings.Split(q, a.colon)

		if len(keyValue) <= 1 {
			errs = append(errs, fmt.Sprintf("%s isn't a colon separated string", q))

			continue
		}

		if _, ok := query[keyValue[0]]; ok {
			errs = append(errs, fmt.Sprintf("%s is a repetead query parameter so only the last one is considered", keyValue[0]))
		}

		query[keyValue[0]] = keyValue[1]
	}

	var err error
	if len(errs) != 0 {
		err = ParseErrors(errs)
	}

	return query, err
}
