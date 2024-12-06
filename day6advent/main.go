package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var coordinates [][]Coordinate
	size, totalObstacles := load(&coordinates)

	loops := 0
	for y := 0; y < len(coordinates); y++ {
		for x := 0; x < len(coordinates); x++ {
			if coordinates[y][x].IsObstacle() || coordinates[y][x].IsGuard() {
				continue
			}

			copied := copyArray(&coordinates)

			copied[y][x].Character = "#"

			if isLoop(&copied, size, totalObstacles) {
				loops++
			}
		}
	}

	fmt.Printf("Total loops: %d", loops)
}

func copyArray(coordinates *[][]Coordinate) [][]Coordinate {
	copyCoordinates := make([][]Coordinate, len(*coordinates))

	for i, inner := range *coordinates {
		copyCoordinates[i] = make([]Coordinate, len(inner))
		copy(copyCoordinates[i], inner)
	}

	return copyCoordinates
}

func isLoop(coordinates *[][]Coordinate, size int, totalObstacles int) bool {
	loop := false
	if !startGuard(coordinates, size, totalObstacles) {
		loop = true
	}

	moves := 0
	for y := 0; y < len(*coordinates); y++ {
		for x := 0; x < len((*coordinates)[y]); x++ {
			if (*coordinates)[y][x].GuardPassed {
				moves++
			}
		}
	}

	fmt.Printf("Moves made: %d, was stuck in loop? %t \n", moves, loop)
	return loop
}

func startGuard(coordinate *[][]Coordinate, size int, totalObstacles int) bool {
	coord := *coordinate
	counter := 0
	guardCompleted := false

	for y := 0; y < len(coord); y++ {

		for x := 0; x < len(coord[y]); x++ {
			if guardCompleted == true {
				break
			}

			if coord[y][x].IsGuard() {
				// found guard, where should it go?
				coord[y][x].GuardPassed = true
				guardX := x
				guardY := y

				for {
					keepMoving, newX, newY := moveNextGuard(guardX, guardY, coordinate)

					guardX, guardY = newX, newY

					if !keepMoving {
						break
					}

					if counter > size {
						break
					}

					counter++
				}

				guardCompleted = true
			}

		}

		if guardCompleted == true {
			break
		}
	}

	if counter > size {
		return false
	} else {
		return true
	}
}

func moveNextGuard(x int, y int, coordinate *[][]Coordinate) (keepMoving bool, newX int, newY int) {
	coord := *coordinate
	guard := &coord[y][x]

	// get direction guard is facing
	direction := guard.GetDirection()

	// is there something in front of guard?
	if direction.IsObstacleInFront(x, y, coordinate) {
		// turn appropriately
		guard.Turn()
		return true, x, y
	} else {
		dx, dy := direction.CalculateDelta()

		if isDeltaOutOfBounds(x, y, dx, dy, coordinate) {
			// win
			return false, x, y
		}

		guard.MoveTo(x, y, x+dx, dy+y, coordinate)

		return true, x + dx, dy + y
	}
}

func isDeltaOutOfBounds(x int, y int, dx int, dy int, coordinate *[][]Coordinate) bool {
	coord := *coordinate

	if y+dy < 0 || y+dy >= len(coord) {
		return true
	}

	if x+dx < 0 || x+dx >= len(coord[0]) {
		return true
	}

	return false
}

func load(coordinates *[][]Coordinate) (int, int) {
	filePath := "data.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0
	totalX := 0
	totalObstacles := 0
	for scanner.Scan() {
		lineText := scanner.Text()

		for len(*coordinates) <= y {
			*coordinates = append(*coordinates, []Coordinate{})
		}

		for x := 0; x < len(lineText); x++ {
			for len((*coordinates)[y]) <= x {
				(*coordinates)[y] = append((*coordinates)[y], Coordinate{})
			}

			(*coordinates)[y][x] = Coordinate{
				Character: string(lineText[x]),
			}

			if string(lineText[x]) == "#" {
				totalObstacles++
			}

			totalX++
		}

		y++
	}

	return totalX + y, totalObstacles
}
