package adapter

import "encoding/json"

/*
the adapter design pattern is used when dealing with code that is already
written, and too integrated to refactor, or using a third-party codebase,
or perhaps just code that you don't want to refactor.

it is like a socket wrench, adding a socket between a tool and surface to
provide functionality that could not be achieved without.
*/

type LegacyComponent struct {
	x, y    int
	text    string
	subType string
}

type NewComponent struct {
	x, y     int64
	text     []string
	subType  string
	newthing struct {
		value float32
		text  []string
	}
}

func (c *NewComponent) getNewThing() map[string]interface{} {
	bytes, err := json.Marshal(c.newthing)
	if err != nil {
		panic(err)
	}

	var output map[string]interface{}
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		panic(err)
	}

	return output
}

/*
now the problem when looking at something like a generic component function of
getNewThing() would be if it were a LegacyComponent there would not be a
compatible interface for it, because the data nor the function exists.  So
writing our own bridge interface will allow us to call a single function in our
codebase, and not need to check what type of component we are interacting with
*/

// we craeate the interface
type Component interface {
	getNewThing() map[string]interface{}
}

// and the functional adapter for the LegacyComponent
func (lc *LegacyComponent) getNewThing() map[string]interface{} {
	return make(map[string]interface{})
}

/*
this is similar to a bridge pattern, except that it is written after the
initial code
*/
