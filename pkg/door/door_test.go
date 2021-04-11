package door

import "testing"

func TestNew(t *testing.T) {
	s := `{{{}}}}`
	f := New("{}")
	isEnd := f.IsClose(s)
	t.Log(isEnd)
}
