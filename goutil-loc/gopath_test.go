package goutil_loc

import (
	"testing"
)

func TestGopaths(t *testing.T) {
	t.Log(GetGopaths())
	t.Log(GetFirstGopath(false))
}
