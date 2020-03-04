package play

import "testing"

func TestRunSwitch(t *testing.T) {
	cases := []struct {
		in int;
		want string
	}{
		{0, "Zero"},
		{1, "One"},
	}

	for _, c := range cases {
		got := RunSwitch(c.in)
		if got != c.want {
			t.Errorf("RunSwitch(%q) == %q, want %q", c.in, got, c.want)
		}
		gotType := RunSwitch2()
		if gotType != "string" {
			t.Errorf("RunSwitch2(%q) == %q, want %q", c.in, gotType, "string")
		}
	}
}
