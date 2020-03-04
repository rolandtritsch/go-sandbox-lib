package play

import "testing"

func TestControl(t *testing.T) {
	cases := []struct {
		name string;
		age int;
		result string;
	}{
		{"roland", 55, "Old"},
		{"karl", 40, "Young"},
	}

	for _, c := range cases {
		got := RunControl(c.name, c.age, 50)
		if got != c.result {
			t.Errorf("RunControl(%q) == %q, want %q", c.name, got, c.result)
		}
	}
}
