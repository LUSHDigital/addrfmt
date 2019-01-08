package addrfmt

import (
	"fmt"
	"io"
	"text/template"
)

// LineMissingError happens when the address line was expected but wasn't in the set
type LineMissingError struct {
	t string
}

func (e LineMissingError) Error() string {
	return fmt.Sprintf("require address line to be present: %q", e.t)
}

// Line represents an address line as a tuple of a type and text value
type Line [2]string

// Type returns the type value
func (l Line) Type() string {
	return l[0]
}

// Text returns the text value
func (l Line) Text() string {
	return l[1]
}

// Lines represents multiple instances of Line
type Lines [][2]string

// Text returns the text of a line type
func (ls Lines) Text(t string) (s string) {
	l, err := ls.Line(t)
	if err != nil {
		return s
	}
	return l.Text()
}

// Line checks if an address line type exists and returns it
func (ls Lines) Line(t string) (Line, error) {
	for i, l := range ls {
		if ls[i][0] == t {
			return l, nil
		}
	}
	return Line{}, LineMissingError{t}
}

// Exists checks if all the provided address line type exists
func (ls Lines) Exists(ts ...string) error {
	for _, t := range ts {
		if _, err := ls.Line(t); err != nil {
			return err
		}
	}
	return nil
}

// Render address lines to a template
func (ls Lines) Render(wr io.Writer, t string, fm template.FuncMap) error {
	tmpl, err := ls.Template().Funcs(fm).Parse(t)
	if err != nil {
		return err
	}
	return tmpl.Execute(wr, nil)
}

// Template returns a new empty template with the basic rendering functions
func (ls Lines) Template() *template.Template {
	return template.New("").Funcs(map[string]interface{}{
		"txt":  ls.Text,
		"text": ls.Text,
	})
}
