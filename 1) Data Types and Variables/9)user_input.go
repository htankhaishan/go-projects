package main
import ("fmt")

func main() {
	/*
	// Multiple Inptus
	var name string
	var is_male bool
	fmt.Print("Enter your name & are you male?: ")
	fmt.Scanf("%s %t", &name, &is_male)
	fmt.Println(name, is_male)
	*/

	// Scanf return values (Count and Error)
	var a string
	var b int
	fmt.Print("Enter a string and an integer: ")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count : ", count)
	fmt.Println("err : ", err)
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}