package calc

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](x, y T) T {
	return x + y
}

func Subtract[T Number](x, y T) T {
	return x - y
}
