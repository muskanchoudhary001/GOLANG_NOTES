package main

import ("fmt")

type Person struct{
  name string
  age int
  job string 
  salary int
}
func main() {
	/*
		Create a Slice With The make() Function

		The make() function can also be used to create a slice.
		Syntax
		slice_name := make([]type, length, capacity)

	*/
	// slice2 := make([]int, 3, 5)
 
	// fmt.Println(slice2)
	// fmt.Println(len(slice2))
	// fmt.Println(cap(slice2))
	/*
	Access Elements of a Slice

You can access a specific slice element by referring to the index number.

In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.
	*/
 //prices := [6] int {12,34,56,6,5,89}//array

//  fmt.Println(prices[1])
//  fmt.Println(prices[0])

 //Change Elements of a Slice
//  prices[0] = 78
//  fmt.Println(prices)
 /*
 Append Elements To a Slice

You can append elements to the end of a slice using the append()function:
Syntax
slice_name = append(slice_name, element1, element2, ...)
 */
// prices1 := append(prices, 99,56,90,100,3)
// fmt.Println(prices1)
/*
Append One Slice To Another Slice

To append all the elements of one slice to another slice, use the append()function:
Syntax
slice3 = append(slice1, slice2...)
*/
// prices2 := append (prices , prices1...)
// fmt.Println(prices2)

/*
Change The Length of a Slice

Unlike arrays, it is possible to change the length of a slice.
*/
// arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
// myslice1 := arr1[1:5] // Slice array
// fmt.Printf("myslice1 = %v\n", myslice1)
// fmt.Printf("length = %d\n", len(myslice1))
// fmt.Printf("capacity = %d\n", cap(myslice1))

// myslice1 = arr1[1:3] // Change length by re-slicing the array
// fmt.Printf("myslice1 = %v\n", myslice1)
// fmt.Printf("length = %d\n", len(myslice1))
// fmt.Printf("capacity = %d\n", cap(myslice1))

/*OPERATORS
Although the + operator is often used to add together two values, it can also be used to add together a variable and a value, or a variable and another variable:

*/
// var (
//     sum1 = 100 + 50 // 150 (100 + 50)
//     sum2 = sum1 + 250 // 400 (150 + 250)
//     sum3 = sum2 + sum2 // 800 (400 + 400)
//   )
//   fmt.Println(sum3)
//BEFORE RANGE , GO ALSO COVERS CONDITIONS STATEMENTS , LOOPS ,OPERATORS THEY ALL ARE SAME AS OF IN C++ !!! SO NOT COVERED ANY EXAMPLES OF THESE

/*
The Range Keyword

The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

The range keyword is used like this:
*/
// for index, value := array|slice|map {
// 	// code to be executed for each iteration
//  }
// fruits := [3]string{"apple", "orange", "banana"}
// for idx, val := range fruits {
//    fmt.Printf("%v\t%v\n", idx, val)
// }
//Tip: To only show the value or the index, you can omit the other output using an underscore (_).

/*FUNCTION
A function is a block of statements that can be used repeatedly in a program.

A function will not execute automatically when a page loads.

A function will be executed by a call to the function.

CREATE A FUNCTION

To create (often referred to as declare) a function, do the following:

    Use the func keyword.
    Specify a name for the function, followed by parentheses ().
    Finally, add code that defines what the function should do, inside curly braces {}.

Syntax
func FunctionName() {
  // code to be executed
} 

FUNCTION CALLING
func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}

FUNCTION PARAMETERS AND ARGUMENTS

Syntax
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
} 

EXAMPLE 
func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}

Multiple Parameters

Inside the function, you can add as many parameters as you want:
EXAMPLE

 


*/
// func familyName(fname string, age int) {
// 	fmt.Println("Hello", age, "year old", fname, "Refsnes")
//   }
  
//   func main() {
// 	familyName("Liam", 3)
// 	familyName("Jenny", 14)
// 	familyName("Anja", 30)
//   }
/*Go Recursion Functions
Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

In the following example, testcount() is a function that calls itself. We use the x variable as the data, which increments with 1 (x + 1) every time we recurse. The recursion ends when the x variable equals to 11 (x == 11). 


*/
// func testcount(x int) int {
// 	if x == 11 {
// 	  return 0
// 	}
// 	fmt.Println(x)
// 	return testcount(x + 1)
//   }
  
//   func main(){
// 	testcount(1)
//   } 


  /*
  Go Structures

A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

Declare a Struct

To declare a structure in Go, use the type and struct keywords:
Syntax
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}

  */
var pers1 Person
var pers2 Person

//pers1 specification
pers1.name="bitu"
pers1.age=45
pers1.job="engineer"
pers1.salary=60000

//pers2 specification
pers2.name="nitu"
pers2.age=39
pers2.job="farmer"
pers2.salary=30445
//access and print pers1
fmt.Println("name: ",pers1.name)
fmt.Println("age: ",pers1.age)
fmt.Println("job: ",pers1.job)
fmt.Println("salary: ",pers1.salary)
//access and print pers2
fmt.Println("name: ",pers2.name)
fmt.Println("age: ",pers2.age)
fmt.Println("job: ",pers2.job)
fmt.Println("salary: ",pers2.salary)

//PASS STRUCT AS FUNCTION ARGUMENTS

printPerson(pers1)  

printPerson(pers2)

///GO MAPS 
//USED TO STORE DATA VALUES IN KER : VALUE PAIRS , EVERY ELEMENT IS A KEY VALUE PAIR
//DOESN'T ALLOW DUPLICATES , UNORDERED AND CHANGEABLE COLLECTION
//LENGTH OF A MAP IS THE NO. OF ITS ELEMENTS 
//DEFAULT VALUE OF A MAP IS NIL
//HOLD REFERENCES TO AN UNDERLYING HASH TABLE
/*Syntax
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
*/

var mapdemo = make(map[string]string)//map is empty
mapdemo["brand"] = "ford"
mapdemo["model"] = "Mustang"
mapdemo["year"] = "1964"

mapdemo1 := make(map[string]int)
mapdemo1["oslo"]=1
mapdemo1["Bergen"]=2
mapdemo1["Trondheim"]=3
mapdemo1["Stavanger"]=4

fmt.Println(mapdemo)
fmt.Println(mapdemo1)
//ACCESSING ELEMENT
fmt.Printf(mapdemo["brand"])


//CREATING AN EMPTY MAP
/*
Syntax
var a map[KeyType]ValueType
Note: The make()function is the right way to 
create an empty map. If you make an empty map in a different 
way and write to it, it will causes a runtime panic.
Allowed Key Types

The map key can be of any data type for which the 
equality operator (==) is defined. These include:

    Booleans
    Numbers
    Strings
    Arrays
    Pointers
    Structs
    Interfaces (as long as the dynamic type supports equality)

Invalid key types are:

    Slices
    Maps
    Functions

These types are invalid because the equality operator (==) is
 not defined for them.
Allowed Value Types

The map values can be any type.

Access Map Elements

You can access map elements by:
Syntax
value = map_name[key]


Update and Add Map Elements

Updating or adding an elements are done by:
Syntax
map_name[key] = value

Remove Element from Map

Removing elements is done using the delete() function.
Syntax
delete(map_name, key)

Check For Specific Elements in a Map

You can check if a certain key exists in a map using:
Syntax
val, ok :=map_name[key]

*/
}
func printPerson(pers Person) {
  fmt.Println( "Name: ",pers.name)
  fmt.Println("Age: ",pers.age)
  fmt.Println("Job: ",pers.job)
  fmt.Println("Salary: ", pers.salary)
}

