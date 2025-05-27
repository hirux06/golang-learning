package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"


	fmt.Println("Languages map:", languages)
	
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "PY")
	fmt.Println("Languages map: ", languages)

	for key, val := range languages{
		fmt.Printf("For key %v shorts for %v\n", key, val)
	}
}