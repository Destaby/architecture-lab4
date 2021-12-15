package main

import (
	"bufio"
	"strings"
	"os"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type Loop struct {

}

func (l *Loop) Post(cmd Command) {
	cmd.Execute(l);
}

func (l *Loop) Start() {

}

func (l *Loop) AwaitFinish() {

}

func parse(line string) Command {
	parts := strings.Fields(line)
	var cmd Command;
	if len(parts) == 2 && parts[0] == "print" {
		cmd = &printCommand{arg: parts[1]}
	} else if len(parts) == 2 && parts[0] == "palindrom" {
		cmd = &palindromCommand{arg: parts[1]}
	} else {
	  var errMsg [2]string
		if len(parts) != 2 {
			errMsg[0] = "\nReason: Comand should have one argument"
		}
		if parts[0] != "print" || parts[0] != "palindrom" {
			errMsg[1] = "\nReason: Unknown command"
		}
		cmd = &printCommand{arg: "SYNTAX ERROR in line: " + line + errMsg[0] + errMsg[1]}
	}
	return cmd
}

func main() {
	loop := new(Loop) 
	loop.Start() 
	
	if input, err := os.Open("test.txt"); err == nil {  
		defer input.Close() 
		scanner := bufio.NewScanner(input) 
		for scanner.Scan() { 
			commandLine := scanner.Text() 
			cmd := parse(commandLine) // parse the line to get a Command  
			loop.Post(cmd) 
		} 
	} 
	loop.AwaitFinish() 
}
