package main

import "./Context"
import "os"
import "bufio"
import "strings"

func main() {
	// set the default state
	context := new(IRPCgame.Context)
	context.State = new(IRPCgame.RestingState)
	
	println("Welcome to the State Game!")
	println("You are standing here looking hesitant!")
	
	// a reader to read the keys from the bash (or command line)
	stdinReader := bufio.NewReader(os.Stdin)
	var key string
	var error os.Error

	// end the loop (and the game) if e pressed
	for key != "e" {
		println("what do you like to do now?")
		print(" Move(m)\n Attack(a)\n Stop(s)\n Run(r)\n Panic(p)\n CalmDown(c)\n Exit(e) ==> ")

		for {
			// read rhe keys from the bash (or command line)
	 	  	key, error = stdinReader.ReadString('\n')
			
			if len(key) > 2 {
				println("please enter a valid key compination")
				continue
			}
			
			if key != "\n" {
				break
			}
			
			// only for the compiler
			if error != nil {
            break
        	}
	 	}
	 	
	 	// remove the new line
	 	key = RemoveNewLine(key)
	 	context.Request(key)
	}
}

func RemoveNewLine(currentEntry string) string {
	var splittedString []string 
	splittedString = strings.Split(currentEntry, "\n",0)
	return splittedString[0]
}
