[![GoDoc](https://godoc.org/github.com/LUSHDigital/addrfmt?status.svg)](https://godoc.org/github.com/LUSHDigital/addrfmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/LUSHDigital/addrfmt)](https://goreportcard.com/report/github.com/LUSHDigital/addrfmt)
[![Build Status](https://travis-ci.org/LUSHDigital/addrfmt.svg?branch=master)](https://travis-ci.org/LUSHDigital/addrfmt)

# addrfmt
This is a small library and framework for validating and rendering addresses which only relies on the standard library.

## Usage
You can get the runnable code in the examples folder.

```go
package main

import (
	"os"
	"strings"

	"github.com/LUSHDigital/addrfmt"
)

var tmpl = `{{ fmt_receipient "RECIPIENT" }}
{{ txt "ADDRESS_LINE" }}
{{ txt "ADMINISTRATIVE_AREA" }}{{ if txt "POST_CODE" }}, {{ txt "POST_CODE" }}{{ end }}`

func main() {
	lines := addrfmt.Lines([][2]string{
		{"RECIPIENT", "John Doe"},
		{"ADDRESS_LINE", "Anon Ln. 42"},
		{"POST_CODE", "666 420"},
		{"ADMINISTRATIVE_AREA", "The Internet"},
	})
	err := valid(lines)
	if err != nil {
		panic(err)
	}
	funcs := map[string]interface{}{
		"fmt_receipient": func(s string) string {
			return strings.ToUpper(lines.Text(s))
		},
	}
	lines.Render(os.Stdout, tmpl, funcs)
}

var required = []string{
	"RECIPIENT",
	"ADDRESS_LINE",
	"ADMINISTRATIVE_AREA",
}

func valid(lines addrfmt.Lines) error {
	if err := lines.Exists(required...); err != nil {
		return err
	}
	line, err := lines.Line("POST_CODE")
	if err != nil {
		return err
	}
	if strings.HasSuffix(line.Text(), "666") {
		// Do some other magic validation
	}
	return nil
}
```