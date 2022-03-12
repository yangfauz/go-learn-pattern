package factory

import "go-learn-pattern/design-pattern/prototype"

type Factory struct{}

func (f Factory) CopyRobot(r *prototype.Robot) *prototype.Robot {
	x := r.Clone()
	return x.(*prototype.Robot)
}
