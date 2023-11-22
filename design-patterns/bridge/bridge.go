package bridge

/*
the bridge pattern is similar to the adapter pattern, in that it allows
interfacing between two otherwise non-compatible components.  However it is used
when creating things that do not already exist, rather than an adapter used on
a legacy component.
*/

// as the name implies, we are building a "bridge" between components

// here is an example of the problem we are trying to solve

type Shape struct {
	x, y, colour int
}

type Circle struct {
	x, y, colour int
	radius       float32
}

type Rectangle struct {
	x, y, width, height, colour int
}

func BadBlueRectangle(x, y, width, height int) Rectangle {
	return Rectangle{
		x, y, width, height, 1,
	}
}

func BadRedRectangle(x, y, width, height int) Rectangle {
	return Rectangle{
		x, y, width, height, 2,
	}
}

func BadBlueCircle(x, y int, radius float32) Circle {
	return Circle{
		x, y, 1, radius,
	}
}

func BadGreenRectangle(x, y int, radius float32) Circle {
	return Circle{
		x, y, 2, radius,
	}
}

/*
the issue is a lot of recreating types for each instantiation of a struct, we
can use the bridge pattern to simplify the interface greatly.
*/

func GoodRectangle(x, y, width, height, colour int) Rectangle {
	return Rectangle{
		x, y, width, height, colour,
	}
}

func GoodCircle(x, y, colour int, radius float32) Circle {
	return Circle{
		x, y, colour, radius,
	}
}

/*
this is a very simplified example, but I think it shows a very easy to
write object creation paradigm to follow
*/
