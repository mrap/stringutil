package stringutil

import (
	"testing"
)

func TestSubstrs(t *testing.T) {
	str := "may"
	expected := []string{"m", "ma", "may", "a", "ay", "y"}
	actual := Substrs(str, 0)
	if !sameValues(expected, actual) {
		t.Errorf("Should have same strings", expected, actual)
	}
}

func TestSubstrsUnicode(t *testing.T) {
	str := "≈ay"
	actual := Substrs(str, 0)
	expected := []string{"≈", "≈a", "≈ay", "a", "ay", "y"}
	if !sameValues(expected, actual) {
		t.Errorf("Should have same strings", expected, actual)
	}
}

func TestSubstrsLimited(t *testing.T) {
	str := "may"
	expected := []string{"ma", "may", "ay"}
	actual := Substrs(str, 2)
	if !sameValues(expected, actual) {
		t.Errorf("Should have same strings", expected, actual)
	}
}

func sameValues(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[string]bool)
	for i, _a := range a {
		m[_a] = true
		_b := b[i]
		m[_b] = true
	}
	if len(m) != len(a) {
		return false
	}

	return true
}
