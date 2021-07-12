package main

import (
	"fmt"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func RemainderSorting(strArr []string) []string {
	//create 3 empty slices t, u, k
	var t []string
	var u []string
	var k []string
	//if modulo of 3 push array element to slice t
	//else push array element with 1 remainder into slice u
	//else push array element with 2 remainder into slice k
	for i := range strArr {
		if len(strArr[i])%3 == 0 {
			t = append(t, strArr[i])
		} else if len(strArr[i])%3 == 1 {
			u = append(u, strArr[i])
		} else if len(strArr[i])%3 == 2 {
			k = append(k, strArr[i])
		}
	}

	c := collate.New(language.Und, collate.IgnoreCase)
	c.SortStrings(t)
	c.SortStrings(u)
	c.SortStrings(k)

	//push elements of slice u into slice t
	for j := range u {
		t = append(t, u[j])
	}
	for p := range k {
		t = append(t, k[p])
	}
	return t
}

func main() {
	arr := []string{"a", "bc", "ab", "cbd", "abc"}
	result := RemainderSorting(arr)
	fmt.Println(result)
}
