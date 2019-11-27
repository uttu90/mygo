// Structure and its initialization
package main
import "fmt"
import "math"

type Circle struct {
	x, y, r float32
}

// Method is func with a reciever
func (c *Circle) area() float32 {
	return math.Pi * c.r * c.r
}

type Point struct {
	x, y float32
}

// Note: Method with value reciever and method with pointer reciever
func (p Point) add(p1 Point) Point {
	p.x = p.x + p1.x
	p.y = p.y + p1.y
	return p
}

func (p *Point) transform(dx float32, dy float32) {
	p.x = p.x + dx
	p.y = p.y + dy
}

func main() {
	c1 := new(Circle)
	fmt.Println(c1.x, c1.y, c1.r)

	c2 := Circle{ 1, 1, 3}
	fmt.Println(c2.x, c2.y, c2.r)
	fmt.Println(c2.area())

	p1 := Point{ 1, 1 }
	p2 := Point{ 2, 4 }
	p3 := p1.add(p2)
	fmt.Println(p1.x, p1.y, p3.x, p3.y) //
	p1.transform(3, 4)
	fmt.Println(p1.x, p1.y)
}
