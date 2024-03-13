package main

import "fmt"

// Public Variables Decalaration
const LoginToken = "Loajhjkfhsdjfsknkajsdfjhsajfhj"

func main() {
	var username string = "Rased"
	fmt.Println(username)
	fmt.Printf("Variable of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable of type %T \n", smallVal)

	var smallFloat float32 = 255.45677
	fmt.Println(smallFloat)
	fmt.Printf("Variable of type %T \n", smallFloat)

	var bigFloat float64 = 255.456774567876
	fmt.Println(bigFloat)
	fmt.Printf("Variable of type %T \n", bigFloat)

	// Default values and some aliases
	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("Variable of type %T \n", anotherVariable)

	var anotherString string 
	fmt.Println(anotherString)
	fmt.Printf("Variable of type %T \n", anotherString)

	// Implicit Type
	var website = "raseprogrammer.com"
	fmt.Println(website) // this one is not changeable

	// No var Style
	// NO var operater only allowed to write in method or into the function
	numberOfUser := 3300000 // this one is not changeable
	fmt.Println(numberOfUser)

	// Call public variables here
	fmt.Println(LoginToken)
	fmt.Printf("Variable of type %T \n", LoginToken)
}