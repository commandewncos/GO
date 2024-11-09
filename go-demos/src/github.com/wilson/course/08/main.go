package main

import "fmt"

func main() {
	// // Define map
	// emails := make(map[string]string)
	// emails["Wilson"] = "wilson@gmail.com"
	// emails["Jessica"] = "jessica@gmail.com"
	// fmt.Println(emails)
	// fmt.Println(len(emails))
	// fmt.Println(emails["Wilson"])

	// // Delete map
	// delete(emails, "Wilson")
	// fmt.Println(emails)

	informations := map[string]string{
		"Name": "Wilson", "Email": "wilson@gmail.com", "hobby": "Read book",
	}

	informations["Name"] = "Wilson Nascimento Costa"

	fmt.Println(informations["Name"])

}
