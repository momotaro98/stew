package stew_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/momotaro98/stew"
)

const (
	errMsg = "test error"
)

func f() error {
	return errors.New(errMsg)
}

func sampleFunc() error {
	err := f()
	if err != nil {
		return stew.Wrap(err)
	}
	return nil
}

func TestWrap(t *testing.T) {
	// Act
	err := sampleFunc()
	// Assert
	if !strings.HasPrefix(err.Error(), errMsg) {
		t.Errorf("got: %s, expected: %s", err.Error(), errMsg)
	}
	if !strings.HasSuffix(err.Error(), "stew_test.sampleFunc\n") {
		t.Errorf("got: %s, expected: %s", err.Error(), "stew_test.sampleFunc")
	}
}
