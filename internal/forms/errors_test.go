package forms

import (
	"testing"
)

func TestErrors_Add(t *testing.T) {
	var err = errors(map[string][]string{})
	testField := "test"
	testMsg := "This is a test"
	err.Add(testField, testMsg)

	if len(err[testField]) == 0 {
		t.Errorf("does not have the %s field when it should", testField)
	}
	lenErr := len(err)
	if lenErr != 1 {
		t.Errorf("added more than 1 error field, added %d error fields instead", lenErr)
	}
}
