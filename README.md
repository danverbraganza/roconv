Roman Numerals Conversion Library in Go
=======================================

Author: Danver Braganza

If you need an implementation of roman numerals, and don't want to
copy-paste the dozens of lines it would take to reimplement it
yourself, this library is for you.

As long as the strings passed to the functions are in utf-8, this
library is able to parse Roman Numerals containing characters with
overbars, such as

   V̅, L̅, C̅ and M̅

If you cannot accept these characters in your output, you may not call
the FromArabic functions with values greater than or equal to 4000
(M̅V).

Command line use
----------------

**NB:** I'm using the ! key for my prompt. It signifies a command to be entered while letting you copy paste from screen to terminal.

    ! go build roconv

Then:

    ! ./roconv -h
    roconv is a commandline utility to convert between roman and arabic numerals.
	   Input is either read from args if present, or failing that from the standard input.
    Usage:
      -mode="1": The mode to run the converter.
      If 'I', input is expected in roman numerals, and output in arabic.
      If '1', input is expected in arabic, and output in roman numerals.
      Defaults to 1.

Suddenly:

    ! ./roconv 3 4 6 1999
    III
    IV
    VI
    MCMXCIX

    ! ./roconv -mode=I
    ! MCMXCIX
    1999
    ! L̅
    50000

Testing
-------

To run tests, run

    ! test/test.sh
    If you see no output from diff, all the tests passed.

You can also run the tests in go with

    ! go test github.com/danverbraganza/go-roconv/romans
