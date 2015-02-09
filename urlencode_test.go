package stringutil

import (
	"testing"
)

func TestUrlEncoded(t *testing.T) {
	str := "Mike Rapadas"
	expected := "Mike%20Rapadas"
	actual, err := UrlEncoded(str)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Should match actual", expected, actual)
	}
}
