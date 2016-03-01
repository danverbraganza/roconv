/**
  Command-line tool for performing roman numeral conversions.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"text/scanner"

	"github.com/danverbraganza/roconv/romans"
)

var acceptedModes = map[string]func(string) (string, error){
	"I": romans.ToArabicString,
	"i": romans.ToArabicString,
	"1": romans.FromArabicString,
}

var mode = flag.String("mode", "1",
	"The mode to run the converter. \n"+
		"\tIf 'I', input is expected in roman numerals, and output in arabic. \n"+
		"\tIf '1', input is expected in arabic, and output in roman numerals. \n"+
		"\tDefaults to 1.")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "roconv is a commandline utility to convert between roman and arabic numerals.\n"+
			"\tInput is either read from args if present, or failing that from the standard input.\nUsage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	function, ok := acceptedModes[*mode]
	if !ok {
		flag.Usage()
		os.Exit(64) // EX_USAGE
	}

	var processOne = func(arg string) {
		result, err := function(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Encountered error converting %v. Error was %v.",
				arg, err)
			os.Exit(65) // EX_DATAERR
		}
		fmt.Println(result)
	}

	if flag.NArg() > 0 { // Read from Nargs.
		for _, arg := range flag.Args() {
			processOne(arg)
		}
	} else { // Read from stdin.
		s := scanner.Scanner{}
		s.Init(os.Stdin)
		for s.Scan() != scanner.EOF {
			processOne(s.TokenText())
		}
	}
}
