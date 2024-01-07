package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gskll/mgrep/internal/search"
)

func main() {
	start := time.Now()
	if len(os.Args) != 3 {
		panic("Error: 2 arguments required - \"mgrep search_string search_dir\"")
	}
	searchString, targetDir := os.Args[1], os.Args[2]
	err := search.Dir(targetDir, searchString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Search completed in %v\n", time.Since(start))
}
