package romans

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Values maps each string in a roman numeral to its numeric value.
var Values = map[byte]uint32{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// These constants are the building blocks for the main algorithm, and are used
// as indexes into Symbols.
const (
	I = iota
	V
	X
)

// OVERLINE are the combining characters for the overline.
var OVERLINE = "\xCC\x85"

// Symbols maps each position with the symbols it may have.
var Symbols = []string{
	"I",
	"V",
	"X",
	"L",
	"C",
	"D",
	"M",
	OVERLINE + "V", // Combining character overline
	OVERLINE + "X", // is reversed because we reverse all the
	OVERLINE + "L", // symbols later.
	OVERLINE + "C",
	OVERLINE + "D",
	OVERLINE + "M",
	"", // 2 string padding for slicing until M̅
	"",
}

// GetSymbols returns the 3 symbols that could possibly be used to represent a
// number
func GetSymbols(i int) []string {
	return Symbols[i*2 : i*2+3]
}

// FromArabic converts an arabic integer into a roman numeral string.
// All conversions are digitwise.
func FromArabic(n uint32) (string, error) {
	var reversedOutput bytes.Buffer

	// Perform conversion by taking every digit, and mapping it.
	for i, pos, nPos := 0, uint32(1), uint32(10); pos <= n; i, pos, nPos = i+1, nPos, 10*nPos {
		digit := (n % nPos) / pos
		n -= n % pos

		if i > 6 { // No more characters available.
			syms := GetSymbols(i - 1) // Go back to the previous level
			for j := uint32(0); j < n/pos; j++ {
				// Write 10 for every one.
				for k := 0; k < 10; k++ {
					// Tally ho!
					reversedOutput.WriteString(syms[I])
				}
			}
			break
		}

		syms := GetSymbols(i)
		switch {
		case digit == 9:
			reversedOutput.WriteString(syms[X])
			reversedOutput.WriteString(syms[I])
		case digit == 4:
			reversedOutput.WriteString(syms[V])
			reversedOutput.WriteString(syms[I])
		default:
			if digit >= 5 {
				reversedOutput.WriteString(syms[V])
			}
			for j := uint32(0); j < digit%5; j++ {
				reversedOutput.WriteString(syms[I])
			}
		}
	}

	// Now reverse it!
	reversed := reversedOutput.String()
	return reverseByRune(reversed), nil
}

// Reverse a string by runes, not bytes.
func reverseByRune(s string) string {
	temp := []rune(s)
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return string(temp)
}

// ToArabic converts a roman string to an arabic integer.
func ToArabic(r string) (uint32, error) {
	r = strings.ToUpper(r)

	// Perform conversion by iterating backwards over letters.
	// Keep track of the max letter seen so far. If we see a smaller one,
	// decrement result by the current value, else increment.
	var max, total uint32
	for i := len(r) - 1; i >= 0; i-- {
		var multiplier uint32 = 1

		// Check if the overbar was passed.
		if r[i] == OVERLINE[1] && r[i-1] == OVERLINE[0] {
			multiplier = 1000
			i = i - 2 // Step over the overbar bytes
		}

		val, ok := Values[r[i]]
		if !ok {
			return 0, fmt.Errorf("Unexpected character: %s", r[i])
		}
		val *= multiplier

		if val < max {
			total -= val
		} else {
			total += val
			max = val
		}
	}
	return total, nil
}

// FromArabicString is a convenience method that parses the arabic
// numeral from a string and then calls FromArabic.
func FromArabicString(n string) (string, error) {
	val, err := strconv.ParseUint(n, 10, 32)
	if err != nil {
		return "", err
	}
	return FromArabic(uint32(val))
}

// ToArabicString is a convenience method that calls toArabic, and
// then converts the return result into a string
func ToArabicString(r string) (string, error) {
	val, err := ToArabic(r)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(uint64(val), 10), nil
}
