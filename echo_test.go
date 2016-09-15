package echo

import "testing"

func TestEcho(t *testing.T) {
  cases := [] struct {
      in, want string
  } { 
      { "Hello", "Hello"}, 
      { "World!", "World!"},
  }

  for _, c := range cases {
    if got := Echo(c.in); got != c.want {
	t.Errorf("Echo(%q) == %q, want %q", c.in, got, c.want)
    }
  }

}
