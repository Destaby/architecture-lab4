package main

import (
	"fmt"
)

type printCommand struct {
	arg string
}

func (pc *printCommand) Execute(handler Handler) {
	fmt.Println(pc.arg)
}

type palindromCommand struct {
	arg string
}

func reverse(str string) (result string) {
	for _, v := range str {
			result = string(v) + result
	}
	return
}

func (pc *palindromCommand) Execute(handler Handler) {
	handler.Post(&printCommand{arg: pc.arg + reverse(pc.arg)})
}
