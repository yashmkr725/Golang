package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")
	languages := make(map[string]string)
	// adding values
	languages["Js"] = "JavaScript"
	languages["Ts"] = "TypeScript"
	languages["Py"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["Js"])

	//deletion
	delete(languages,"Py") // you can use this in slices also
	fmt.Println(languages)

	//loops are intresting in golang
	for key,value := range languages{
		fmt.Printf("the keys is %v and the value is %v \n",key,value)
	}


}
