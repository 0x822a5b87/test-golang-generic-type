package generics

import (
	"testing"
)

func TestNewGenderValue(t *testing.T) {
	v := NewGenderValue[Male]("male")
	if v.GetGender() != "male" {
		t.Error()
	}
}
