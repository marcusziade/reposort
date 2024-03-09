package main

import (
	"fmt"
	"os"

	reposort "github.com/marcusziade/reposort/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the directory path as an argument.")
		os.Exit(1)
	}

	path := os.Args[1]
	err := reposort.SortRepositories(path)
	if err != nil {
		fmt.Printf("Error sorting repositories: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Repositories sorted into language directories.")
}
