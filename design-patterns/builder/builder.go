package builder

// A builder pattern allows a step by step method-based pattern of object creation
// rather than using a constructor.  This allows for creating an object without
// a complex constructor sequence.

type Hotdog struct {
	bread   string
	ketchup bool
	mustard bool
	kraut   bool
	relish  bool
}

func (h *Hotdog) breadType(bread string) {
	h.bread = bread
}

func (h *Hotdog) addKetchup() {
	h.ketchup = true
}

func (h *Hotdog) addMustard() {
	h.mustard = true
}

func (h *Hotdog) addKraut() {
	h.kraut = true
}

func (h *Hotdog) addRelish() {
	h.relish = true
}

// the benefit of this design pattern is that the attribute code is kept closed
// after the initial write, and new attribute code will not require a full
// rewrite of the object, just the addition of a new method.

// this has more recently been taken over by the factory pattern in my day to day
// code.
