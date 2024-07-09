// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	_ = words

	// Print the byte and rune length of the strings
	// Hint: Use len and utf8.RuneCountInString
	fmt.Printf("\n%v \n%v\n", "Lengths", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		fmt.Printf("\tByte length: %v \n", utf8.RuneCountInString(word))
		fmt.Printf("\tRune lenfth: %v \n", len([]byte(word)))
	}

	// Print the bytes of the strings in hexadecimal
	// Hint: Use % x verb
	fmt.Printf("\n%v \n%v\n", "Bytes in hex", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		fmt.Print("\tHex: 0x")
		for _, byt := range []byte(word) {
			fmt.Printf("%x ", byt)
		}
		fmt.Println()
	}

	// Print the runes of the strings in hexadecimal
	// Hint: Use % x verb
	fmt.Printf("\n%v \n%v\n", "Runes in hex", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		fmt.Printf("\tHex: 0x% x\n", word)
	}

	// Print the runes of the strings as rune literals
	// Hint: Use for range
	fmt.Printf("\n%v \n%v\n", "Rune literals", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		fmt.Print("\tRunes: ")
		for _, run := range word {
			fmt.Printf("%q ", run)
		}
		fmt.Println()
	}

	// Print the first rune and its byte size of the strings
	// Hint: Use utf8.DecodeRuneInString
	fmt.Printf("\n%v \n%v\n", "First Rune & Byte Size", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		run, len := utf8.DecodeRuneInString(word)
		fmt.Printf("\tFirst rune: %c\n", run)
		fmt.Printf("\tByte size:  %v\n", len)
	}

	// Print the last rune of the strings
	// Hint: Use utf8.DecodeLastRuneInString
	fmt.Printf("\n%v \n%v\n", "Last Rune & Byte Size", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		run, len := utf8.DecodeLastRuneInString(word)
		fmt.Printf("\tLast rune: %c\n", run)
		fmt.Printf("\tByte size: %v\n", len)
	}

	// Slice and print the first two runes of the strings
	fmt.Printf("\n%v \n%v\n", "First Two Runes of String", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		runes := []rune(word)
		fmt.Printf("\tFirst two runes: %c\n", runes[:2])
	}

	// Slice and print the last two runes of the strings
	fmt.Printf("\n%v \n%v\n", "Last Two Runes of String", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		runes := []rune(word)
		fmt.Printf("\tLast two runes: %c\n", runes[len(runes)-2:])
	}

	// Convert the string to []rune
	// Print the first and last two runes
	fmt.Printf("\n%v \n%v\n", "First 2 and Last 2 Runes in String", strings.Repeat("=", 50))
	for _, word := range words {
		fmt.Println(word)
		runes := []rune(word)
		fmt.Printf("\tFirst two runes: %c\n", runes[:2])
		fmt.Printf("\tLast two runes: %c\n", runes[len(runes)-2:])
	}
}
