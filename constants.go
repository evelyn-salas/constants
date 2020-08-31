package main

import (
	"fmt"
)

//a varible defined outside of main can be made updated later
//this is not the case for every type
const a int16 = 27

//iota pt 1//
//enumerated constants (continues in main)
const (
	i = iota
	j = iota
	k = iota
	//same result from:
	//_ = iota //_ is a stand in for a throwaway variable,so if we
	//i			//didn't care about the zero value
	//j
	//k //comment out rest to test!
)

//resetting iota
//iota is scoped to a constant block. i2 will be equal to zero
const (
	i2 = iota
)

//using iota to shift values
const (
	_  = iota //ignore first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

//using iota to create a byte
const (
	isAdmin = 1 << iota
	isHq
	canSeeFin

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	//How to define a constatnt//
	//if you capitalize the first letter, the constant will
	//be exported
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	//declaration issues
	//you cannot set your constant equal to something that needs
	// to be solved/executed. This won't work:
	//const myConst float64 = math.Sin(1.57)
	//fmt.Printf("%v, %T\n", myConst, myConst) //comment out rest to test

	//more declaration examples
	const a int = 14
	const b string = "pashi"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	//math using constants
	var e int = 27 //you can add a constant and a variable if same type
	fmt.Printf("%v, %T\n", a+e, a+e)

	//*exception*
	//if you do not specify the type
	//const a = 42
	//and you this to something that does have a specific type
	//var b int16 = 27
	//go will assume the types are the same
	//fmt.Printf("%v, %T\n", a+b, a+b) //comment out rest to test

	//iota pt 2//
	//enumerated constants
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", k, k)
	//so we see here that we get an increasing output of values from 0 onwards
	//(this is a shorter way of creating loops using i as a counter in java)

	//resetting iota
	fmt.Printf("%v %T\n", i2, i2)

	//using iota to shift values
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB) //quick way to figure out file size :D

	//using iota to create a byte
	var roles byte = isAdmin | canSeeFin | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHq&roles == isHq)

}
