package main

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
	NOT_GUARD
)

func GetDirectionalCharacter(num int) string {
	switch num {
	case int(NORTH):
		return "^"
	case int(EAST):
		return ">"
	case int(SOUTH):
		return "v"
	case int(WEST):
		return "<"
	default:
		return "."
	}
}

func (d *Direction) CalculateDelta() (deltaX int, deltaY int) {
	// get bounds
	switch *d {
	case NORTH:
		{
			deltaY = -1
			break
		}
	case EAST:
		{
			deltaX = 1
			break
		}
	case SOUTH:
		{
			deltaY = 1
			break
		}
	case WEST:
		{
			deltaX = -1
			break
		}
	default:
		return 0, 0
	}

	return deltaX, deltaY
}

func (d *Direction) IsObstacleInFront(x int, y int, coords *[][]Coordinate) bool {
	coord := *coords
	dx, dy := d.CalculateDelta()

	// ensure bounds is not outside
	if isDeltaOutOfBounds(x, y, dx, dy, coords) {
		return false
	}

	return coord[y+dy][x+dx].IsObstacle()
}
