package join

import (
	"testing"
)

type testData struct {
	list   []string
	expect string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{list: []string{}, expect: ""},
		{list: []string{"apple"}, expect: "apple"},
		{list: []string{"apple", "orange"}, expect: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, expect: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.expect {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\" ", test.list, got, test.expect)
		}
	}
}
