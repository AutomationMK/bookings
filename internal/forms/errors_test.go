package forms

import "testing"

func TestAdd(t *testing.T) {
	var err = errors(map[string][]string{})
	testField := "test"
	testMsg := "This is a test"
	err.Add(testField, testMsg)

	getTest := err.Get(testField)
	if getTest == "" {
		t.Errorf("does not have the %s field when it should", testField)
	}
}
