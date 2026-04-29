// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner_test

import (
	"fmt"
	"log"

	"github.com/gravwell/gcfg/scanner"
	"github.com/gravwell/gcfg/token"
)

func ExampleScanner_Scan() {
	// src is the input that we want to tokenize.
	src := []byte(`[profile "A"]
color = blue ; Comment`)

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file, err := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Init(file, src, nil /* no error handler */, scanner.ScanComments); err != nil {
		log.Fatal(err)
	}

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit, _ := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%q\t%q\n", fset.Position(pos), tok, lit)
	}

	// output:
	// 1:1	"["	""
	// 1:2	"IDENT"	"profile"
	// 1:10	"STRING"	"\"A\""
	// 1:13	"]"	""
	// 1:14	"\n"	""
	// 2:1	"IDENT"	"color"
	// 2:7	"="	""
	// 2:9	"STRING"	"blue"
	// 2:14	"COMMENT"	"; Comment"
}
