package go

import (
    "fmt"
    "strconv"
)

// variable can be declared globally
// we can assign as int8 to conserve memory
var x int8 = 1

func declareVariables() {
    // multiple declarations
    var i, j int = 1, 2

    // short variable declaration - not possible outside functions
    k := 3 

    // multiple short variable declaration
    a, b, c := true, 1, "hello"

    // untyped variable 
    var y = 1

    fmt.Println(x, i, j, k, a, b, c, y)
}

func declareStringsAndBytes() {
    s := "hello world"       // strings must use double quotes 
    s = `hello world`        // equivalent of """ in Python

    c := 'A'                 // ascii number for char A
    b := []byte(s)           // ascii number for each char in s
}

func convertType() {
    // example converts string to bytes
    str := "hello world"
    bytes := []byte(str)
    fmt.Println(string(bytes))
    
    // example converts a type back to []string
    type names []string
    friends := names{"Alice", "Cassandra", "Ellie"}
    fmt.Println([]string(friends))

    // example converts int to string
    // note that Aoti returns (int, error)
	num := 2000
	str = strconv.Itoa(num)
	num, _ = strconv.Atoi(str)
	fmt.Println(str, num)
}