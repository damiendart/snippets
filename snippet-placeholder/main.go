// Copyright (C) 2023 Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func usage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s [filename]\n\n", os.Args[0])
	fmt.Fprintln(w, "If no filename is specified, the snippet will be read from standard input.")
}

func main() {
	var buf strings.Builder
	var err error
	var input *os.File
	var tmpl *template.Template

	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" {
			usage(os.Stdout)
			os.Exit(0)
		}
	}

	if len(os.Args) > 2 {
		fmt.Fprintf(
			os.Stderr,
			"Error: Incorrect number of positional arguments given (1 required, %d given)\n\n",
			len(os.Args)-1,
		)
		usage(os.Stderr)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1, os.Args[1] == "-":
		input = os.Stdin
	default:
		input, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer input.Close()
	}

	_, err = io.Copy(&buf, input)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = template.New("").
		Option("missingkey=error").
		Funcs(NewFuncMap()).
		Parse(buf.String())
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
