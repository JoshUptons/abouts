package composite

/*
Composte design pattern works best on structures that are tree-like
What this pattern is used to defeat, is the issue where specific code on
a sub-product will require finding specific modules and altering the code.
On a small data set, this is not all that difficult, but at scale, trying to
maintain all products modules would get out of hand.
Insert the Composite pattern. In golang we utilize a composite with a common
interface.  When a composite interface is called, it passes the function calls
down the tree.

Also this pattern allows easily adding new sub-products without breaking code up
the tree.
*/

type Box struct {
	Items []Product
	Boxes []Box
}

type Product struct {
	Price float64
	Name  string
}

func (p Product) CalcPrice() float64 {
	return p.Price
}

func (b Box) CalcPrice() float64 {
	var total float64
	for _, prod := range b.Items {
		total += prod.CalcPrice()
	}
	for _, box := range b.Boxes {
		total += box.CalcPrice()
	}
	return total
}

func (b *Box) AddItem(item Product) {
	b.Items = append(b.Items, item)
}

func (b *Box) AddBox(box Box) {
	b.Boxes = append(b.Boxes, box)
}

type CompositeBox interface {
	CalcPrice() float64
}

func CalcPrice(c []CompositeBox) float64 {
	var total float64
	for _, c := range c {
		total += c.CalcPrice()
	}
	return total
}

/*
We can now create a box, that can contain a mix of boxes and products,
all of which implement the CalcPrice method.  which means when we call
the base CalcPrice function which takes a slice of CompositeBoxes, the sub
functions will call their child element CalcPrice functions, returning their
price to the top level of the tree
*/
