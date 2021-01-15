package main

import (
	"fmt"
)

func main() {
	x := 234.25
	ptr := &x

	// Logging
	fmt.Printf("Pointer is :%f \n", *ptr)
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
