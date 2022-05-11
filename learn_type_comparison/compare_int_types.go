package main

import (
	"fmt"
	"os"
	"strconv"
)


func DoCompareGeneric [T int | int64 | int32, U int | int64 | int32] (i T, j U) bool {
	return int64(i) == int64(j)
}

func DoCompareInt(i, j int) bool {
	return i == j
}


func ShowCompare [T int | int64 | int32, U int | int64 | int32] (cmp func(T, U) bool, i T, j U) {
	var result string
	if cmp(i, j) {
		result = "=="
	} else {
		result = "!="
	}

	fmt.Printf("%v (%T) %v %v (%T)\n", i, i, result, j, j)
}

func main() {
	arg_one, _ := strconv.Atoi(os.Args[1])
	arg_two, _ := strconv.Atoi(os.Args[2])
	i := int(arg_one)
	j := int64(arg_two)
	// var i32 int32
		ShowCompare(DoCompareInt, i, i)

	cmpGeneric := make(DoCompareGeneric, int, int64)

	ShowCompare(cmpGeneric, i, j)
	// DoCompare(i, i64)
	// DoCompare(i, i32)
}