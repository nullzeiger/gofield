// Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
)

// minefield with a size 10x10
const Rowcols = 10
// Number of mines
const Mines = 5
// Map symbol
const Mapchar = "󱔐"
// Mine symbol
const Minechar = "󰷚"
// Flag symbol
const Flagchar = ""
// Hero symbol
const Herochar = ""

// Arrive data structure
type Arrive struct {
	char string
	x, y int
}

// Hero data structure
type Hero struct {
	char string
	x, y int
}

// Mine data structure
type Mine struct {
	char string
	x, y int
}

// Create new map
func CreateMap(a Arrive) [][]string {
	gameMap := make([][]string, Rowcols)

	for i := 0; i < Rowcols; i++ {
		for j := 0; j < Rowcols; j++ {
			gameMap[i] = append(gameMap[i], Mapchar)
		}
	}

	gameMap[a.y][a.x] = a.char

	return gameMap
}

// Create 5 random mine
func CreateMines() [Mines]Mine {
	var m [Mines]Mine
	for i := 0; i < len(m); i++ {
		m[i] = Mine{char: Minechar, y: rand.IntN(Rowcols), x: rand.IntN(Rowcols)}
	}

	return m
}

// Create random arrive
func CreateArrive() Arrive {
	return Arrive{char: Flagchar, y: rand.IntN(Rowcols), x: rand.IntN(Rowcols)}
}

// Clear console and print map
func PrintMap(m [][]string, h Hero) {
	fmt.Printf("\x1B[2J\x1B[H")

	fmt.Println("Command to move(hjkl) and exit(q)")

	m[h.y][h.x] = h.char

	for _, i := range m {
		for _, j := range i {
			fmt.Print(j)
		}

		fmt.Println()
	}

	m[h.y][h.x] = Mapchar
}

func main() {
	hero := Hero{Herochar, 0, 0}
	mines := CreateMines()
	arrive := CreateArrive()
	newMap := CreateMap(arrive)

	PrintMap(newMap, hero)

	loop := true

	// Game loop
	for loop {
		// Read Input
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'q':
			loop = false
			break
		case 'l':
			hero.x += 1
			if hero.x > 9 {
				hero.x = 9
			}

			PrintMap(newMap, hero)
		case 'h':
			hero.x -= 1
			if hero.x < 0 {
				hero.x = 0
			}

			PrintMap(newMap, hero)
		case 'k':
			hero.y += 1
			if hero.y > 9 {
				hero.y = 9
			}

			PrintMap(newMap, hero)
		case 'j':
			hero.y -= 1
			if hero.y < 0 {
				hero.y = 0
			}

			PrintMap(newMap, hero)
		}

		if hero.x == arrive.x && hero.y == arrive.y {
			fmt.Println("Win!!")
			loop = false
			break
		}

		for i := 0; i < Mines; i++ {
			if hero.x == mines[i].x && hero.y == mines[i].y {
				hero.char = mines[i].char
				PrintMap(newMap, hero)
				fmt.Println("Game Over!!")
				loop = false
				break
			}
		}

	}
}
