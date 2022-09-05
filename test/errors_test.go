package test

import (
	"testing"
)

type MockStruct struct {
	MockName string `json:"mockName" validate:"required"`
}

func TestValidationRequiredValues(t *testing.T) {
	t.Logf("testing required values to make sure that you can't pass empty values")
}

func TestValidationEmail(t *testing.T) {
	t.Logf("testing email validation to make sure only valid emails can get sent")
}

func TestValidationOther(t *testing.T) {
}

func TestCheckError(t *testing.T) {
	t.Fail()
}
