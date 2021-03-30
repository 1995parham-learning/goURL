package css

import (
	"fmt"
	"strings"
)

// ColonSeparatedStrings represents an array of colon separated strings.
type ColonSeparatedStrings []string

// ParseErrors contains errors of colon separated strings' parsing process.
type ParseErrors []string

func (pe ParseErrors) Error() string {
	err := ""

	for _, p := range pe {
		err += fmt.Sprintln(p)
	}

	return err
}

// ToMap creates a map from the array of colon separated strings.
func (a ColonSeparatedStrings) ToMap() (map[string]string, error) {
	query := make(map[string]string)
	errs := make([]string, 0)

	for _, q := range a {
		keyValue := strings.Split(q, ":")

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
