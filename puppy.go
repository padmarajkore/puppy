package puppy

import (
	"github.com/padmarajkore/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() string {
	return ("I am version 1.1.0")
}

func From12() string {
	return ("I am version 1.2.0")
}

func From13() string {
	return ("I am version 1.3.0")
}

func Form14() string {
	return ("i am version 1.4.0")
}

func Form15() string {
	return ("i am version 1.5.0")
}
