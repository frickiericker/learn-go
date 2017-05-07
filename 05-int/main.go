package main

import (
	"fmt"
)

func main() {
	a := int(10)
	b := int8(100)
	c := int16(1000)
	d := int32(10000)
	e := int64(100000)

	// Requires explicit conversions (fail if variables are const)
	_ = a + int(b)
	_ = b + int8(c)
	_ = int32(e) + d

	// Defined behavior https://golang.org/ref/spec#Conversions
	if int8(c) == -24 {
		fmt.Println("Yes")
	}

	// Unsigned
	u := uint32(0x12345678)
	v := u & uint32(d)
	fmt.Printf("%08x\n", v)
}
