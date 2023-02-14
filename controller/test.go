package main

import "fmt"

type Data struct {
	City, Province, Country string
}

func main() {

	data := Data{
		"Kulon Progo",
		"Yogyakarta",
		"Indonesia",
	}

	fmt.Println(data)

}
