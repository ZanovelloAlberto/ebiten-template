package sprite

type Map = struct {
	camera Point
}

var Pmap = Map{
	camera: Point{
		x: screenWidth/2 - PPlayer.point.x,
		y: screenHeight/2 - PPlayer.point.y,
	},
}

func GetPoint(point Point) (x float64, y float64) {
	a := point.Less(&Pmap.camera)
	return a.Get()
}

// ============= UTILS ==================

type Point struct {
	x float64
	y float64
}

func (a *Point) Less(b *Point) Point {

	return Point{
		a.x - b.x, a.y - b.y,
	}
}

func (a *Point) Get() (x float64, y float64) {
	return a.x, a.y
}
