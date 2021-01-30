package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func (E Employee) changeNameByValue(val string) string {
	E.name = E.name + val
	return E.name
}

func (E *Employee) changeNameByRef(val string) {
	E.name = E.name + val
}

func (E Employee) showEmployee() {
	fmt.Printf("Employee Name:%v | Employee age:%v \n", E.name, E.age)
}

func main() {
	Mabroor := Employee{"Mabroor Ahmad", 20}

	Mabroor.showEmployee()
	Mabroor.changeNameByValue("Pro")
	Mabroor.showEmployee()
	// logging
	//fmt.Println(aim)
}

/*
###### Constants declaration ######

const name string = "Mabroor ahmad"
const age int = 20
const heightInCm float64 = 170.89

*/

/*
###### Variables ######

###### simple declaration with out initialization ######
var x,y,z int
var name, fatherName string

###### simple declaration with initialization ######
var name string = "Mabroor Ahmad"
var age int = 20
var heightInCm float64 = 170.89
var cool bool = true

###### infered (shortcut way) variable declaration  ######
x := 0
y := "Mabroor Ahmad"
height := 170.89 // default type float64
cool := true
[Note it does not work outside the function]

###### Multiple variable declaration ######
var name, age, isCool = "Mabroor Ahmad", 20, true
   [
		 or infered way
	 ]
name, age, isCool := "Mabroor Ahmad", 20, true


###### Another way of multiple declaration ######

var (
	name string = "Mabroor Ahmad"
	age  int    = 20
	cool bool
	id   int
)

[
	Note this way those variable whose value
	is not initialized are automatically
	initialized as int to 0 , string to "",
	bool to false, float to 0
]



###### Get the type of the variable ######
import (
	"reflect"
)

fmt.Println(reflect.TypeOf(variable))

[ Another way of finding a type ]

fmt.Printf("Type is : %T \n", variable)

*/

/*
###### loging ######

fmt.Printf("Hello world") // Does not ends with new line
fmt.Println("Hello world") // prints whole line means ends with new line

*/

/*
###### Functions ######

func singleReturnAndCombinedTypeDeclaredParams (x, y int) int {
	return x + y
}

func multipleReturnAndIndividualTypeDeclaredParams (name string, age int) (string, int) {
	return name, age
}

func namedReturnAndIndividualTypeDeclaredParams (name string, age int) (_name string, _age int) {
	_name = name
	_age = age
	return
}

func main() {
	defer fmt.Printf("Hello\n") // After
	fmt.Printf("World") // Before
}
[ Note defer statement is executed after the surrounding function returns ]
[ But the arguments are evaluated at start ]


func main() {
	defer fmt.Printf("4")
	defer fmt.Printf("3")
	defer fmt.Printf("2")
	defer fmt.Printf("1")
	fmt.Printf("0") // Before
}
[ Note defer calls are stored in a stack so as they are called ]

[ Note like javascript functions in go can be passed as params to other functions ]

	func sum(x, y int) int {
		return x + y
	}

	func getSum(fn func(int, int) int) int {
		return fn(3, 3)
	}

	// logging
	fmt.Println(getSum(sum))

[
	Note go functions are simmilar to closures every call
  will have seprate instance and have seprate variables
	to work with
]

*/

/*
###### Swapping the variables with lambda function ######

x, y = (func(x, y int) (int,int) {
				return y, x
			})(x, y)

	or using named return

x, y = (func(x, y int) (_x,_y int) {
				 _y = x
				 _x = y
				 return
			})(x, y)

*/

/*
###### Type casting in go ######
j := 23.3 // float64
i := int(j) // int
k := float64(i) //float64

import (
	"strconv"
)
[ This is a utility package for type conversions and parsing ]

*/

/*
###### For loop ######

for i := 0; i <= 20000; i++ {
	// loging
	fmt.Println(i)
}

[ Another way of for loop]

var counter int = 1
for counter < 50 {

	// loging
	counter += counter
	fmt.Println(counter)
}

[ for loop becomes while loop ]

for true {
	fmt.Println("Acting like while loop");
}

[ If you ommit condition it will run infinitly ]

for {
		// Logging
		fmt.Println("Runnig Infinitely")
}

*/

/*
###### if else Statement ######

if x != 0 {
	fmt.Println("x is not equal to 0")
} else if x == 0 {
	fmt.Println("x is equal to 0")
} else {
	fmt.Println("I hope i will run some day")
}

[ You can run statement or declare variables before comparing ]

if x := 3; x < 23 {
		fmt.Println("I am less then 23")
	} else {
		fmt.Println("I hope i will run some day")
}

[ Variables declared in if are also available in else ]

if x := 3; x < 23 {
		fmt.Println("I am less then 23")
	} else if x > 23 {
		//x can be accessed here
		fmt.Println("I hope i will run some day")
}

*/

/*
###### Scope of variables ######

valibales are block scoped

var j int = 23 // Global

{
	var i int = 0 // local scoped

	i // known
	j // known

}

i // unknown
j // known


*/

/*
###### Switch Statement in go ######

switch runtime.GOOS {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("OSX")
	default:
		fmt.Println("any")
	}

[ Note there is no need of break statement ]
---------

fmt.Println("When's Saturday?")
today := time.Now().Weekday()
switch time.Saturday {
case today + 0:
	fmt.Println("Today.")
case today + 1:
	fmt.Println("Tomorrow.")
case today + 2:
	fmt.Println("In two days.")
default:
	fmt.Println("Too far away.")
}

[ Note switch variables are not constant ]
---------

t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}

[ Note switch without variable is true a clean way of writing long if else chains ]
*/

/*
###### Pointers ######

var i int = 24  // simple variable
var p *int = &i // pointer to i variable
fmt.Println(*p) // Logging value at pointer
*p = 234        // changing value at pointer
fmt.Println(i)  // Logging changed value(change done through pointer) of i
j := 234				// Declaring an other variable
p = &j          // Assigning p pointer new adress of j variable AKA referencing
fmt.Println(*p) // Logging value(representing value of j) at pointer
*p = 234        // changing value(changing value of j) at pointer AKA dereferencing
fmt.Println(j)  // Logging changed value(change done through pointer) of j

x := 24.243
ptr := &x       // shortcut way of creating a pointer

*/

/*
###### Structs ######

type Point struct {
	X int
	Y int
}

p1 := Point{42, 325}           // Creating instance
or
var p1 Point = Point{42, 325}  // Creating instance
or
var p1 Point = Point{X: 2}     // Specific assignment initialization urordered
[ unspecified fields will be initialized in a default way ]
or
var p1 Point = Point{}         // Default initialization

p1.X                           // Accessing fields

pointer := &p1                 // Pointer to struct
pointer.Y = 23                 // Acessing fiels from pointer

// This is a slice created using struct
persons := []struct {
	name string
	age  int
}{
	{name: "Mabroor Ahmad", age: 20},
	{name: "Ghazi Hassan", age: 21},
}

*/

/*
###### Arrays ######
var alphabet [5]string                                      // Declaring, 5 is the length of array
var alphabets [5]string = [5]string{"a","b","c","d","e"}    // Declarative initialization
alphabets := [5]string{"a","b","c","d","e"}                 // Shortcut way
alphabets[0]                                                // Accessing values

// Iterate over elements of array
	for idx, val := range arr {
		fmt.Printf("x:%v v:%v\n", idx, val)
	}

	or

	for i := range pow {
		fmt.Printf("output: %d\n", i)
	}

	or if you dont want idx

	for _, val := range pow {
		fmt.Printf("output: %d\n", val)
	}

*/

/*
###### Slice ######

aSlice := alphabets[0:3]									// Creating a slice from an array
aSlice[0] = "xxx"											// Changing the value in the slice

[
    Note slice is just a reference to the array. Modifing it
    will modify the value of the main array so be carefull
]

s := []int{235,24,36,324,36}								// Direct initialization
s = s[:0]													// Slice the slice to give it zero length.
s = s[:4]													// Extend its length.
s = s[2:]													// Drop its first two values.
aSlice := alphabets[0:3]									// [i,j] i start_idx , j end_idx

-------Dynamic arrays
// make function is used to create dynamic arrays
// make([]type,len,cap)

b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

// 2d Array using slice

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	or

	matrix := [][]int{
		{21, 23, 52},
		{28, 22, 53},
		{35, 74, 98},
	}

-------Adding element to a slice
// element is added using append function

	append(Slice,element) // returns slice with added element

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board = append(board,
		[]string{"_", "_", "_"},
	)

*/

/*
###### Map ######
These are like javascript objects which contain set of key value pairs

	// Type
	type Point struct {
		x, y float64
	}

	// Creating map
	aim := map[string]Point{
		"mabroor": {23, 64},
		"ghazi":   {634, 234},
	}

	// Another method
	aim := make(map[string]Point)
	aim["mabroor"] = Point{x: 23, y: 324}
	aim["ghazi"] = Point{x: 21, y: 4}

	// logging
	fmt.Println(aim["mabroor"])


	// Check if element is there
	elem, ok := aim["mabroor"]

	// ok is true if element is there
	if ok {
		fmt.Println("Element is Present:", elem)
	} else {
		fmt.Println("Element is not found!")
	}

*/

/*
###### Methods ######

	Methods in go as go does not have classes but you can
	create methods for structs

	type Point struct {
		x, y float64
	}

	// assigning a method to Point
	func (p Point) render() {
		fmt.Printf("x:%v,y:%v", p.x, p.y)
	}

	// Initializing
	aim := Point{92.241, 523.234}

	// Calling the method
	aim.render()

	// This is a regular functional way of doing the same thing as above
	func renderPoint(p Point) {
		fmt.Printf("x:%v,y:%v", p.x, p.y)
	}

	//Now we have to call it like regular function

	render(aim)

	// Try to pass values by reference as it wont duplicate things
	// And also it would make you able to change the value
	// and it is more efficent

	func (E *Employee) changeNameByRef(val string) {
		E.name = E.name + val
	}

*/
