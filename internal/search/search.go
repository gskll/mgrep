package search

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Dir(path, target string) error {
	var wg sync.WaitGroup

	counter := 0
	visit := func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if dir.IsDir() {
			return nil
		}
		wg.Add(1)
		go func() error {
			count, err := searchFile(&wg, path, target)
			if err != nil {
				return nil
			}
			counter += count
			return nil
		}()
		return nil
	}

	err := filepath.WalkDir(path, visit)
	if err != nil {
		return err
	}
	wg.Wait()
	fmt.Printf("Found %d matches for %s in dir %s\n", counter, target, path)

	return nil
}

func searchFile(wg *sync.WaitGroup, path, target string) (int, error) {
	defer wg.Done()

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return 0, fmt.Errorf("Error opening file %s: %w", path, err)
	}
	scanner := bufio.NewScanner(file)

	line := 0
	counter := 0
	for scanner.Scan() {
		line++

		if strings.Contains(scanner.Text(), target) {
			counter++
			fmt.Printf("%d - %s: %s\n", line, path, strings.TrimSpace(scanner.Text()))
		}

	}
	return counter, nil
}
