package main

import "fmt"

func main() {

	var kata = "selamat malam"
	var length = len([]rune(kata))
	var i = 0

	for {
		fmt.Println(kata[i : i+1])

		i++

		if i == length {
			break
		}
	}

}
