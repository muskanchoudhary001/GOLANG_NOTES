package main

import (
	"fmt"
)

// var e int
// var f int = 90
// var g = "heeelmooo"
// const A int = 1
// const B  = 18
// const (
//   C int = 1
//   D = 3.14
//   E = "Hi!"
// )

func main() {
	//  //VARIABLE DECLARATION IN GO
	// 	fmt.Println("Hello World! I'm ready to conquer the world")
	// 	var integer int = 45
	// 	var stringvari string = "muskan jhoothi hai !!!"
	// 	var floatvari float32 = 23.34

	// 	fmt.Println(integer)
	// 	fmt.Println(stringvari)
	// 	fmt.Println(floatvari)

	// 	var a string
	// 	var b int
	// 	var c bool

	// 	a = "it is what it is "
	// 	b = 69
	// 	c = true

	// 	fmt.Println(a)
	// 	fmt.Println(b)
	// 	fmt.Println(c)

	// 	e = 78
	// 	fmt.Println(e)
	// 	fmt.Println(f)
	// 	fmt.Println(g)

	// // MULTIPLE VARIABLE DECLARATION
	// var w,x,y,z int = 1,2,3,4

	// fmt.Println(w)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(w, x, y, z)

	// // //If the type keyword is not specified, you can declare different types of variables in the same line:
	// // //NOTE:If you use the type keyword, it is only possible to declare one type of variable per line.

	// // var r, s = 6, "Hello"
	// // into,strngo := 10 ,"world"

	// // fmt.Println(r)
	// // fmt.Println(s)
	// // fmt.Println(into)
	// // fmt.Println(strngo)

	// //go variable declaration in block
	// var (
	//   u int
	//   i int = 1
	//   o string = "hello"
	// )
	// fmt.Println(u)
	// fmt.Println(i)
	// fmt.Println(o)

	// //Constants
	// //Note: The value of a constant must be assigned when you declare it.

	// const PI float32 =3.14

	// fmt.Println(PI)

	// //TYPED CONSTANTS
	// //Typed constants are declared with a defined type:
	// fmt.Println(A)

	// //UNTYPES CONSTANTS
	// //Untyped constants are declared without a type:
	// fmt.Println(B)

	// // Multiple Constants Declaration

	// // Multiple constants can be grouped together into a block for readability:
	// fmt.Println(C)
	// fmt.Println(D)
	// fmt.Println(E)

	/* GO ARRAYS
	    DECLARING AN ARRAY
	   1. With the var keyword:
	   var array_name = [length]datatype{values} // here length is defined

	   or

	   var array_name = [...]datatype{values} // here length is inferred

	   2. With the := sign:
	    array_name := [length]datatype{values} // here length is defined

	   or

	   array_name := [...]datatype{values} // here length is inferred
	*/

	// var array1 = [5] int {1,2,3,4,5}
	//  array2 := [3] int {4,8,9}

	//  fmt.Println(array1)
	//  fmt.Println(array2)

	//  var array3 = [3] string {"hyyy!!","hieee!!","heellloooo!!!"}
	//  fmt.Println(array3)
	//  fmt.Println(array3[1])

	//  array3[2] = "hyyy there !!!"
	//  fmt.Println(array3)

	// var ar1 = [4] int {}
	// var ar2 = [4] int {1,2}
	// var ar3 = [4] int {1,2,3,4}

	//Initialize Only Specific Elements
	//var ar4 = [4] int {1:34,2:67}

	// fmt.Println(ar1)
	// fmt.Println(ar2)
	// fmt.Println(ar3)
	//fmt.Println(ar4)

	//Length of array
	//fmt.Println(len(ar4))

	/*
	   GO SLICES
	   Slices are similar to arrays, but are more powerful and flexible.

	   Like arrays, slices are also used to store multiple values of the same type in a single variable.

	   However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	   In Go, there are several ways to create a slice:

	       Using the []datatype{values} format
	       Create a slice from an array
	       Using the make() function
	   SYNTEX :slice_name := []datatype{values}
	*/
	//  var myslice1 = []int {}
	//  fmt.Println(len(myslice1))
	//  fmt.Println(cap(myslice1))
	//  fmt.Println(myslice1)

	//  myslice2 := [5] string {"tinu" ," bitu" , "nitu" , "mohit","happy"}
	//  fmt.Println(len(myslice2))
	//  fmt.Println(cap(myslice2))
	//  fmt.Println(myslice2)

	/*
	   Create a Slice From an Array

	   You can create a slice by slicing an array:
	   Syntax
	   var myarray = [length]datatype{values} // An array
	   myslice := myarray[start:end] // A slice made from the array

	*/
	var arr1 = [5]int{90, 45, 56, 32, 43}
	myslice3 := arr1[1:3]

	fmt.Println(arr1)
	fmt.Println(myslice3)
	fmt.Println(len(myslice3), " ", cap(myslice3))

}
