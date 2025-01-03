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
// Count  *****************************
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
	str5 := "Mr Rakesh Pradhan"
	fmt.Println(strings.Count(str5, "a"))

	//Cut
	//this function return 3 parameters before cut string after cut string and the string or the
	//part we will provide it is present or not AND it sacesensitve
	beforeCut, afterCut, found := strings.Cut(str5, "Rakesh")

	fmt.Println("after cut function", afterCut)
	fmt.Println("before cut function", beforeCut)
	fmt.Println("available", found)

	//func CutPrefix and CutSuffix  are same like cut

	suffix, found1 := strings.CutSuffix(str1, "sh")

	fmt.Println("before cut function", suffix)
	fmt.Println("available", found1)

	prefix, found2 := strings.CutPrefix(str1, "Ra")

	fmt.Println("after cut function", prefix)
	fmt.Println("available", found2)

	//func Fields
	//it convert a string into a array in basis of one or more white space
	filds := strings.Fields(str5)

	fmt.Printf("type of fields is %T", str5)
	fmt.Printf("type of fields is %T", filds)

	//  INDEX

	fmt.Println(strings.Index("chicken", "ick"))

}
