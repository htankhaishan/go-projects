// Type Casting
// Integer to Float
/*
package main

import "fmt"

func main() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f \n", f)
}
*/

// Float to Integer
/*
package main

import "fmt"

func main() {
	var f float64 = 45.89
	var i int = int(f)
	fmt.Printf("%v \n", i)
}
*/

// strconv package
/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 4
	var s string = strconv.Itoa(i)
	fmt.Printf("%q \n", s)
	fmt.Printf("%T \n", s)
}
*/

// String to Int
/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T", err, err)
}
*/

// Error for Atoi

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "200abc"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T", err, err)
}
