package common

import (
	"bytes"
	"fmt"
	"testing"
)

func AssertInt(t *testing.T, expect, actual int, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%d', got: '%d'", s, expect, actual)
	return false
}

func AssertByte(t *testing.T, expect, actual byte, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%0X', got: '%0X'", s, expect, actual)
	return false
}

func AssertBytes(t *testing.T, expect, actual []byte, format string, a ...interface{}) bool {
	if bytes.Equal(expect, actual) {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%d', got: '%d'", s, expect, actual)
	return false
}

func AssertPos(t *testing.T, expect, actual Pos, format string, a ...interface{}) bool {
	if expect == actual {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%v', got: '%v'", s, expect, actual)
	return false
}