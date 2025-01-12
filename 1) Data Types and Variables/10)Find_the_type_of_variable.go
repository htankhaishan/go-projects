/*
	%T format specifier
	reflect.TypeOf function from the reflect package.
*/
/*
package main
import "fmt"

func main() {
	var grades int = 42
	var message string = "Hello, World!"
	var isCheck bool = true
	var amount float64 = 42424.34

	fmt.Printf("Type of grades %v is %T\n", grades, grades)
	fmt.Printf("Type of grades %v is %T\n", message, message)
	fmt.Printf("Type of grades %v is %T\n", isCheck, isCheck)
	fmt.Printf("Type of grades %v is %T\n", amount, amount)
	
}
*/

// using reflect.TypeOf()
/*
package main
import ("fmt"; "reflect")
func main() {
	fmt.Printf("Type of grades %v \n", reflect.TypeOf(100))
	fmt.Printf("Type of message %v \n", reflect.TypeOf("priyanka"))
	fmt.Printf("Type of isCheck %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type of amount %v \n", reflect.TypeOf(true))
}
	*/

package main
import ("fmt"; "reflect")
func main() {
	var grades int = 42
	var message string = "Hello, World!"

	fmt.Printf("Type of grades %v is %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("Type of message %v is %v \n", message, reflect.TypeOf(message))
}