package functions

import "strconv"

func ParseInteger(s string) (int, error) {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return result, nil
}
