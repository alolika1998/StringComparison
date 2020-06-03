package main

import (
	"fmt"
)

func comparestring(str1 string, str2 string){
if str1 == str2 {
	fmt.Printf("%s and %s are equal\n", str1,str2)
	return
}
fmt.Printf("%s and %s are not equal\n", str1,str2)
}
func main() {
	string1 := "Golang"
	string2 := "Golang"
	comparestring(string1,string2)

	string3 := "C"
	string4 := "C++"
	comparestring(string3,string4)
}
