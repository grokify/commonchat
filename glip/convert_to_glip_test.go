package glip

import (
	"testing"
)

var tripleBacktickToCodeTests = []struct {
	v    string
	want string
}{
	{"```foobar```", "\n[code]\nfoobar\n[/code]\n"},
}

func TestTripleBacktickToCode(t *testing.T) {
	for _, tt := range tripleBacktickToCodeTests {
		got := TripleBacktickToCode(tt.v)
		if got != tt.want {
			t.Errorf("glip.TripleBacktickToCode(\"%v\") Mismatch: want[%v], got [%v]", tt.v, tt.want, got)
		}
	}
}
