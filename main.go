package main

import (
	"fmt"
)

func main() {
	var (
		i int
		j float64
		k bool
		l string
	)
	// loging
	fmt.Println(i, j, k, l)
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
	bool to false, float to 0.0
]



###### Get the type of the variable ######
import (
	"reflect"
)

fmt.Println(reflect.TypeOf(variable))

*/

/*
###### loging ######

fmt.Printf("Hello world") // Does not ends with new line
fmt.Println("Hello world") // prints whole line means ends with new line

*/

/*
###### Functions defination ######

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
