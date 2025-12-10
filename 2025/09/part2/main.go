package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type rect struct {
	vertices []point
	area int
}

func readInput(filename string) []string {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	contents := string(file)
	contents = strings.TrimSpace(contents)

	return strings.Split(contents, "\n")
}

func getCoords(lines []string) []point {
	coords := []point{}
	for _, line := range lines {
		coordStr := strings.Split(strings.TrimSpace(line), ",")
		coord := []int{}

		for _, c := range coordStr {
			coordInt, err := strconv.Atoi(c)

			if err != nil {
				panic(err)
			}

			coord = append(coord, coordInt)
		}

		coords = append(coords, point{x: coord[0], y: coord[1]})
	}

	return coords
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateArea(pointA, pointB point) int {
	x1 := min(pointA.x, pointB.x)
	x2 := max(pointA.x, pointB.x)
	y1 := min(pointA.y, pointB.y)
	y2 := max(pointA.y, pointB.y)

	dimensions := []int{
		x2 - x1 + 1,
		y2 - y1 + 1,
	}

	res := 1

	for _, plane := range dimensions {
		res *= plane
	}

	return res
}

func computeAllAreas(points []point) []rect {
	areas := []rect{}
	for _, pointA := range points {
		for _, pointB := range points {
			area := calculateArea(pointA, pointB)

			pointC := point{
				x: pointA.x,
				y: pointB.y,
			}
			pointD := point{
				x: pointB.x,
				y: pointA.y,
			}

			rectangle := rect{
				vertices: []point{pointA, pointC, pointB, pointD},
				area: area,
			}

			if area < 0 {
				continue
			}

			areas = append(areas, rectangle)
		}
	}

	return areas
}

func cmp(a, b rect) int {
	return b.area - a.area
}

func getBoundingRect(r rect) (point, point) {
	minX := math.MaxInt
	maxX := 0
	minY := math.MaxInt
	maxY := 0

	for _, v := range r.vertices {
		if v.x > maxX {
			maxX = v.x
		}
		if v.x < minX {
			minX = v.x
		}
		if v.y > maxY {
			maxY = v.y
		}
		if v.y < minY {
			minY = v.y
		}
	}

	pMin := point{x: minX, y: minY}
	pMax := point{x: maxX, y: maxY}
	return pMin, pMax
}

func intersects(r rect, polygon []point) bool {
	n := len(polygon)

	for i := range polygon {
		p1 := polygon[i]
		p2 := polygon[(i+1) % n]

  	pMin, pMax := getBoundingRect(r)
  
  	if p1.x == p2.x {
  		x := p1.x
  
  		minY := min(p1.y, p2.y)
  		maxY := max(p1.y, p2.y)
  
  		if pMin.x < x && x < pMax.x && max(minY, pMin.y) < min(maxY, pMax.y) {
  			return true
  		}
  	} else if p1.y == p2.y {
  		y := p1.y
  
  		minX := min(p1.x, p2.x)
  		maxX := max(p1.x, p2.x)
  
  		if pMin.y < y && y < pMax.y && max(minX, pMin.x) < min(maxX, pMax.x) {
  			return true
  		}
  	}
	}

	return false
}

func getMaxRects(rects []rect, polygon []point) []rect {
	candidates := []rect{}

	for _, r := range rects {
		if !intersects(r, polygon) {
			candidates = append(candidates, r)
		}
	}

	return candidates
}

func main() {
	lines := readInput("input.txt")

	polygon := getCoords(lines)
	rects := computeAllAreas(polygon)
	slices.SortFunc(rects, cmp)

	candidates := getMaxRects(rects, polygon)

	if len(candidates) == 0 {
		fmt.Println("Couldn't find rectangles that match the criteria")
		return
	}

	fmt.Println("The rectangle with the max area is", candidates[0].area, candidates[0])
}
