package Errors

import "fmt"

type AgeError struct { // Custom error type
	Age int
}

func (e *AgeError) Error() string {
	return fmt.Sprintf("Возраст %d меньше допустимого!", e.Age)
}

func checkAge(age int) error {
	if age < 18 {
		return &AgeError{Age: age}
	}
	return nil
}
