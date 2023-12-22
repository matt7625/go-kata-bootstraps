package kata


type point struct {
	x int
	y int
}

func newPoint(x int, y int) *point {
	return &point{x,y}
}

func ManhattanDistance(a *point, b *point) int {
	return Abs(a.x - b.x) + Abs(a.y - b.y)
}


func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}