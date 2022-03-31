package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

func Steam_CountGamesInDatabase() (*big.Int, *big.Int) { // counts games in database + counts all games, which are not playable in your region / country	a := big.NewInt(0)
	a := big.NewInt(0)
	b := big.NewInt(0)
	filepath.Walk("../steam", func(path string, info os.FileInfo, err error) error {
		// if it is a file, then count it
		if info.IsDir() {
			return nil
		}
		// get parent directory of path
		parent := filepath.Dir(path)
		if parent == "..\\steam\\83" || parent == "../steam/83" {
			b.Add(b, big.NewInt(1))
		} else {
			a.Add(a, big.NewInt(1))
		}
		return nil
	})
	return a, b
}

func Steam_PrintOutAllGames() {
	filepath.Walk("../steam", func(path string, info os.FileInfo, err error) error {
		// if it is a file, then count it
		if info.IsDir() {
			return nil
		}
		// get parent directory of path
		parent := filepath.Dir(path)
		if parent == "..\\steam\\83" || parent == "../steam/83" {
		} else {
			// set id to the file name
			id := filepath.Base(path)
			// remove the extension from id
			id = strings.TrimSuffix(id, filepath.Ext(id))

			// read file and print the first line
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			scanner.Scan()
			fmt.Println(scanner.Text(), "("+id+")")

		}
		return nil
	})
}

func main() {
	// games that are available in your region / country
	a, b := Steam_CountGamesInDatabase()
	c := big.NewInt(0).Add(a, b)
	fmt.Println("All games:", c.String())
	fmt.Println("Games available in your region / country:", a.String())
	fmt.Println("Games not available in your region / country:", b.String())
	// Steam_PrintOutAllGames() - too much games for the terminal
}
