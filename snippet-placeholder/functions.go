// Copyright (C) 2023 Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

package main

import (
	"os"
	"text/template"
	"time"
)

// NewFuncMap returns a map of functions for use in snippets templates.
func NewFuncMap() template.FuncMap {
	return template.FuncMap{
		"getenv": os.Getenv,
		"now":    now,
	}
}

func now(fmt string) string {
	return time.Now().Format(fmt)
}
