package main

import (
	"fmt"
	"strings"
)

//there are loat of package are available in string package

// Builder
// Clone
// Compare
// Contains
// ContainsAny
// ContainsFunc
// ContainsRune
// Count
// Cut
// CutPrefix
// CutSuffix
// EqualFold
// Fields
// FieldsFunc
// HasPrefix
// HasSuffix
// Index
// IndexAny
// IndexByte
// IndexFunc
// IndexRune
// Join
// LastIndex
// LastIndexAny
// LastIndexByte
// LastIndexFunc
// Map
// NewReplacer
// Repeat
// Replace
// ReplaceAll
// Split
// SplitAfter
// SplitAfterN
// SplitN
// Title
// ToLower
// ToLowerSpecial
// ToTitle
// ToTitleSpecial
// ToUpper
// ToUpperSpecial
// ToValidUTF8
// Trim
// TrimFunc
// TrimLeft
// TrimLeftFunc
// TrimPrefix
// TrimRight
// TrimRightFunc
// TrimSpace
// TrimSuffix

//all the example are available in https://pkg.go.dev/strings

func main() {
	str1 := "Rakesh"
	str2 := "Pradhan"
	// str3 := "i am a boy"
	str4 := " name is rakesh im rakesh my name is       "

	// Builder function
	// var b strings.Builder
	// for i := 3; i >= 1; i-- {
	// 	fmt.Fprintf(&b, "%d...", i)
	// }
	// b.WriteString("ignition")
	// fmt.Println(b.String())

	// Clone

	cloneName := strings.Clone(str1)

	fmt.Println(cloneName == str1) //true
	fmt.Println(cloneName)         //rakesh

	// Compare
	fmt.Println(strings.Compare(str1, str2))     //1  bcz str1>str2
	fmt.Println(strings.Compare(str1, "Rakesh")) //0  bcz str1 == str2
	fmt.Println(strings.Compare(str1, "rakesh")) //-1 bcz str1 < rakesh

	// Contains
	fmt.Println(strings.Contains(str4, "foo"))  //false
	fmt.Println(strings.Contains(str4, "name")) //true

	//func Count ¶
	fmt.Println(strings.Count(str4, "e"))

}
