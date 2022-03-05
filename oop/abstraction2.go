package oop

type Triangle struct {
	base, height float32
}
type Square struct {
	length float32
}
type Rectangle struct {
	length, width float32
}
type Circle struct {
	radius float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}
func (l Square) Area() float32 {
	return l.length * l.length
}
func (r Rectangle) Area() float32 {
	return r.length * r.width
}
func (c Circle) Area() float32 {
	return 3.14 * (c.radius * c.radius)
}

type Area interface {
	Area() float32
}
