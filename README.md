# GO: Remainder Sorting

implement a function that receives an array of strings and
sort them based on the following heuristics:
- the primary sort is by increasing remainder of the strings length, modulo 3
- if multiple strings have the same remainder, they should be sorted in alphabetical order

E.g: strArr=["Colorado","Utah","Wisconsin","Oregon"]
should return sortedArray =["Oregon","Wisconsin","Utah","Colorado"]

run function with 
```bash
    go run main.go
```

run test with
```bash
    go test
```