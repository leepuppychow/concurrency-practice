// Exercise 8.5 --> refactor the surface drawing program (3.2) to execute concurrently

package surfaces

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Coordinate struct {
	x int
	y int
}

func GeneratePlot() string {
	coordCh := make(chan Coordinate)
	polygonCh := make(chan string)

	go generateCoordinates(coordCh)
	go generatePolygons(coordCh, polygonCh)
	return generateSVG(polygonCh)
}

func generateCoordinates(out chan<- Coordinate) {
	defer close(out)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			out <- Coordinate{x: i, y: j}
		}
	}
}

func generatePolygons(in <-chan Coordinate, out chan<- string) {
	defer close(out)
	for coord := range in {
		x, y := coord.x, coord.y
		ax, ay := corner(x+1, y)
		bx, by := corner(x, y)
		cx, cy := corner(x, y+1)
		dx, dy := corner(x+1, y+1)
		out <- fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy)
	}
}

func generateSVG(in <-chan string) string {
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for polygon := range in {
		svg += polygon
	}
	return svg + "</svg>"
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
