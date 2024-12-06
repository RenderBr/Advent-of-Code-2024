package main

type Coordinate struct {
	Character   string
	GuardPassed bool
}

func (c *Coordinate) IsGuard() bool {
	return c.Character == "^" || c.Character == ">" || c.Character == "<" || c.Character == "v"
}

func (c *Coordinate) GetDirection() Direction {
	switch c.Character {
	case "^":
		return NORTH
	case ">":
		return EAST
	case "v":
		return SOUTH
	case "<":
		return WEST
	default:
		return NOT_GUARD
	}
}

func (c *Coordinate) MoveTo(currentX int, currentY int, newX int, newY int, coords *[][]Coordinate) {
	(*coords)[newY][newX].Character = c.Character
	(*coords)[currentY][currentX].Character = "."

	(*coords)[newY][newX].GuardPassed = true
}

func (c *Coordinate) Turn() {
	currentDirection := c.GetDirection()

	// Update the new direction in a cyclic manner
	newDirection := (int(currentDirection) + 1) % 4

	// Update the character
	c.Character = GetDirectionalCharacter(newDirection)
}

func (c *Coordinate) IsObstacle() bool {
	return c.Character == "#"
}
