// Copyright © 2018 Kent Gibson <warthog618@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package charset

import "fmt"

// Decoder provides a mapping from GSM7 byte to UTF8 rune.
type Decoder map[byte]rune

// Encoder provides a mapping from UTF8 rune to GSM7 byte.
type Encoder map[rune]byte

// NationalLanguageIdentifier indicates the character set in use, as defined in
// 3GPP TS 23.038 Section 6.2.1.2.4.
type NationalLanguageIdentifier int

const (
	// Default character set.
	Default NationalLanguageIdentifier = iota
	// Turkish character set.
	Turkish
	// Spanish character set
	Spanish
	// Portuguese character set
	Portuguese
	// Bengali character set
	Bengali
	// Gujaranti character set
	Gujaranti
	// Hindi character set
	Hindi
	// Kannada character set
	Kannada
	// Malayalam character set
	Malayalam
	// Oriya character set
	Oriya
	// Punjabi character set
	Punjabi
	// Tamil character set
	Tamil
	// Telugu character set
	Telugu
	// Urdu character set
	Urdu
)

// Display prints the character set for a given character set decoder.
func Display(m Decoder) {
	specials := map[rune]string{
		'\n':   "LF",
		'\r':   "CR",
		'\f':   "FF",
		' ':    "SP",
		0x1b:   "ESC",
		0x20ac: " €",
	}
	fmt.Printf("      ")
	for c := 0; c < 8; c++ {
		fmt.Printf("0x%d_ ", c)
	}
	fmt.Println("")
	for r := 0; r < 0x10; r++ {
		fmt.Printf("0x_%x: ", r)
		for c := 0; c < 8; c++ {
			k := byte(c*0x10 + r)
			if v, ok := m[k]; ok {
				if s, ok := specials[v]; ok {
					fmt.Printf("%3s  ", s)
				} else if v >= 0x400 {
					fmt.Printf("%04x ", v)
				} else {
					fmt.Printf("  %c  ", v)
				}
			} else {
				fmt.Printf("     ")
			}
		}
		fmt.Println()
	}
}
