package puppy

import (
	"fmt"

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

func From11() {
	fmt.Println("I am version 1.1.0")
}

func From12() {
	fmt.Println("I am version 1.2.0")
}

func From13() {
	fmt.Println("I am version 1.3.0")
}

func Form14() {
	fmt.Println("i am version 1.4.0")
}
