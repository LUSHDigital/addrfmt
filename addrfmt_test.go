package addrfmt_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/LUSHDigital/addrfmt"
)

func Test_Line(t *testing.T) {
	line := addrfmt.Line([2]string{"a", "b"})
	equals(t, "a", line.Type())
	equals(t, "b", line.Text())
}
func Test_Lines_Exists(t *testing.T) {
	lines := addrfmt.Lines([][2]string{
		{"a", "a"},
		{"b", "b"},
	})
	err := lines.Exists("a", "b")
	equals(t, nil, err)
	err = lines.Exists("c")
	if err == nil {
		t.Error("should be an error")
	}
}

func Test_Lines_Line(t *testing.T) {
	lines := addrfmt.Lines([][2]string{
		{"a", "a"},
		{"b", "b"},
	})
	_, err := lines.Line("a")
	equals(t, nil, err)
	_, err = lines.Line("c")
	if err == nil {
		t.Error("should be an error")
	}
	if err.Error() == "" {
		t.Error("error should have content")
	}
}

func Test_Lines_Text(t *testing.T) {
	lines := addrfmt.Lines([][2]string{
		{"a", "a"},
		{"b", "b"},
	})
	s := lines.Text("a")
	equals(t, "a", s)
	s = lines.Text("c")
	equals(t, "", s)
}

func Test_Lines_Render(t *testing.T) {
	lines := addrfmt.Lines([][2]string{
		{"a", "a"},
		{"b", "b"},
	})
	var buf bytes.Buffer
	err := lines.Render(&buf, `{{ txt "a" }}{{ txt "c" }}{{ text "b" }}`, nil)
	equals(t, nil, err)
	equals(t, "ab", buf.String())
	buf.Reset()
	err = lines.Render(&buf, `txt "a" }}{{ txt "c" }}{{ text "b"`, nil)
	if err == nil {
		t.Error("should be an error")
	}
	err = lines.Render(&buf, `{{ hello "a" }}{{ txt "c" }}{{ text "b" }}`, map[string]interface{}{
		"hello": func(s string) string {
			return fmt.Sprintf("hello %s", s)
		},
		"txt": func(s string) string {
			if s == "" {
				return "c"
			}
			return s
		},
	})
	equals(t, nil, err)
	equals(t, "hello acb", buf.String())
}

func equals(tb testing.TB, expected, actual interface{}) {
	tb.Helper()
	if expected != actual {
		tb.Fatalf("\n\texp: %#[1]v (%[1]T)\n\tgot: %#[2]v (%[2]T)\n", expected, actual)
	}
}
