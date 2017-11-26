package string

import "testing"


func Test(t *testing.T){
	var tests = []struct {
		s, expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"dave", "evad"},
		{"☹dave☀", "☀evad☹"},
		{"",""},
	}

	for _, test := range tests {
		actual := Reverse(test.s)
		if actual != test.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", test.s, actual, test.expected)
		}
	}

}


