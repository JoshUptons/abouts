package factory

import (
	"fmt"
	"runtime"
)

/*
The factory pattern eliminates the issue where you have multiple objects of
similar types, which have a lot of the same underlying code, but perhaps a
small set of differences that are quite important.
To elimate writing a bunch of creation methods for specific attributes, we write
one factory that takes in an argument and passes out the object that coorisponds
to the given type.
*/

/*
A good example is cross-platform objects.
If we have an iOS and an Android application, but we don't want to constantly
be writing `(os.platform == 'darwin') ? doIosThing : doAndroidThing`, we can
write a factory for each that helps build some boilerplate without too much
monotonous if statements.
*/

// we might have something like this

type IOSButton struct {
	id           string
	clicked      bool
	text         string
	iosOnlyThing interface{}
}

func (b *IOSButton) toString() string {
	return fmt.Sprintf("<button id='%s' clicked='%t'>%s</button>", b.id, b.clicked, b.text)
}

type AndroidButton struct {
	id               string
	clicked          bool
	text             string
	androidOnlyThing interface{}
}

func (b *AndroidButton) toString() string {
	return fmt.Sprintf("<button id='%s' clicked='%t'>%s</button>", b.id, b.clicked, b.text)
}

// instead of writing a function like this
func getLayout() string {
	layoutText := "<div>this is some layout</div>"
	if runtime.GOOS == "darwin" {
		button := IOSButton{
			"ios-button-1",
			false,
			"click me",
			"ios related thing",
		}
		layoutText += button.toString()
	} else {
		button := AndroidButton{
			"android-button-1",
			false,
			"click me",
			"android thing",
		}
		layoutText += button.toString()
	}

	layoutText += "<div>this is some more layout</div>"
	if runtime.GOOS == "darwin" {
		button := IOSButton{
			"ios-button-2",
			false,
			"click me",
			"ios related thing",
		}
		layoutText += button.toString()
	} else {
		button := AndroidButton{
			"android-button-2",
			false,
			"click me",
			"android thing",
		}
		layoutText += button.toString()
	}
	layoutText += "<div>last layout thing</div>"
	return layoutText
}

// we can instead write something cleaner with an abstraction level that helps
// write easier to maintain code

type Button interface {
	toString() string
}

type ButtonFactory struct {
}

func (bf *ButtonFactory) createButton(id string, clicked bool,
	text string, relatedItem interface{}) Button {
	if runtime.GOOS == "darwin" {
		b := IOSButton{
			fmt.Sprintf("ios-%s", id),
			clicked,
			text,
			relatedItem,
		}
		return Button(&b)
	} else {
		b := AndroidButton{
			fmt.Sprintf("android-%s", id),
			clicked,
			text,
			relatedItem,
		}
		return Button(&b)
	}
}

func getLayoutWithFactory() string {
	var bf ButtonFactory

	layoutText := "<div>some layout</div>"
	layoutText += bf.createButton("button-1", false, "click me", "my thing").toString()
	layoutText += "<div>some more layout stuff</div>"
	layoutText += bf.createButton("button-2", false, "click me", "my other thing").toString()
	layoutText += "<div>some final layout stuff</div>"

	return layoutText
}

/*
this way we only need to maintain the check within one factory code base and
call simple interface code in our actual layout.

this can be further abstracted to individual factories of say:
type IosButtonFactory and type AndroidButtonFactory respectively that will
abstract some of the code into those factories that will handle os specific things
while the main ButtonFactory just creates an instance of the specific Factory
per runtime.
*/
