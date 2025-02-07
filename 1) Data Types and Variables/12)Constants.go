/*
package main

import "fmt"

func main() {
	const name = "Tom"
	const age = 12
	const is_muggle = true
	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", age, age)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
}
*/

// Run-time Panic
/*
package main

import "fmt"

func main() {
	const name = "Tom"
	name = "Ben"
	fmt.Printf("%v: %T \n", name, name)
}
*/

package main

import "fmt"

const PI float64 = 3.14 //global constant

func main() {
	var radius float64 = 5.0
	var area float64
	area = PI * radius * radius
	fmt.Println("Area of the Cycle is : ", area)
}
