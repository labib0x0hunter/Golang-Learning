package main

import (
	"fmt"
	"strconv"
	"strings"
)

var str string = "https://heelogo.go/howareyou"
var ptrPrefix string = "https://"
var ptrSuffix string = "areyou"
var ptr string = "go"
var withOutTrim = "  labib  al  "

func print(arr ...any) {
	fmt.Println(arr...)
}

func main() {

	// Functions
	found := strings.Contains(str, ptr)
	found = strings.HasPrefix(str, ptrPrefix)
	found = strings.HasSuffix(str, ptrSuffix)

	index := strings.Index(str, ptr)
	index = strings.LastIndex(str, ptr)

	upper := strings.ToUpper("abc")
	lower := strings.ToLower("ABC")

	afterTrim := strings.TrimSpace(withOutTrim)
	afterTrim = strings.Trim(withOutTrim, " ")

	split := strings.Split("a,b,c", ",")
	split = strings.Fields("a b c")

	join := strings.Join([]string{"a", "b", "c"}, "+")
	newStr := strings.ReplaceAll(str, "go", "GO")

	print(found, index, upper, lower, afterTrim, split, join, newStr)

	// Comapre two strings
	// Case sensitive
	a, b := "abx", "Abx"
	if a == b {
		fmt.Println("Equal")
	}

	// Case insensitive
	if strings.EqualFold(a, b) {
		fmt.Println("Equal")
	}

	// Efficient concatanation
	var strB strings.Builder
	strB.WriteString("Hello")
	strB.WriteString("Go")
	strB.WriteByte('+')
	strB.Write([]byte("NewWorld"))
	strB.WriteRune(rune(233))
	print(strB.String())

	// Efficient replace
	r := strings.NewReplacer(
		"<", "&lt;",
		">", "&gt;",
		"&", "&amp;",
	)
	s := "<html> & go >"
	s = r.Replace(s)
	fmt.Println(s)

	r = strings.NewReplacer(
		"{name}", "Labib",
		"{lang}", "Go",
	)
	s = r.Replace("Hello {name}, welcome to {lang}!")
	fmt.Println(s)

	// Convertion
	n, _ := strconv.Atoi("12")
	fmt.Println(n)

	m := strconv.Itoa(12)
	fmt.Println(m)

	flt, _ := strconv.ParseFloat("123.321", 32)
	flt32 := float32(flt)
	fmt.Println(flt32)

	

}
