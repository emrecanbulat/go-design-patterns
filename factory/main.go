package main

import (
	"fmt"

	"go-design-patterns/factory/movie"
)

func main() {
	titanic := movie.NewTitanic()
	johnWick := movie.NewJohnWick()
	printMovie(titanic)
	printMovie(johnWick)
}

func printMovie(m movie.IMovie) {
	fmt.Printf("Title : %s", m.GetTitle())
	fmt.Println()
	fmt.Printf("Runtime : %d", m.GetRuntime())
	fmt.Println()
}
