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

func TestErrors_Get(t *testing.T) {
	var err = errors(map[string][]string{})
	testField := "test"
	testMsg := "This is a test"
	err[testField] = append(err[testField], testMsg)

	testGet := err.Get(testField)
	if testGet != testMsg {
		t.Errorf("has \"%s\" as returned message when it should be \"%s\"", testGet, testMsg)
	}
	testGet = err.Get("non-exist")
	if testGet != "" {
		t.Errorf("has %s when it should be empty for non-existent field", testGet)
	}

}
