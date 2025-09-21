package main

import (
	//Global Imports
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	c4 "github.com/accal/GoLangProjects/Connect4"
)

// ----------------------------------------------------------------
// -------- Global Variables for the main application loop --------
// ----------------------------------------------------------------
// Simple Structure to act like a tuple for the Dictionary for the main selection of Programs
// From the list of Modules withinq the project
type Program struct {
	Name          string
	MainExecution func()
}

// Mapping for the main dictionary for the Programs that are available to run
// For the main loop of the function
var programs = map[int]Program{
	0: {Name: "Quit", MainExecution: func() { fmt.Println("-------- Ending Simulation -------") }},
	1: {Name: "Connect4", MainExecution: func() { c4.PlayConnect4() }},
}

// ----------------------------------------------------------------
// --------------------- End of Global Variables ------------------
// ----------------------------------------------------------------

/*
Main application function that you want to run from the root of the project
inputs: none
outputs: none
This function will display a menu to the user to select which program that they want to run
from the list of programs available in the project. It will then call the main function of that program to start the execution of that program.
The user can select a number from the list provided from the programs dictionary to select the program they want to run.
If the user selects a number outside of that range, it will prompt them to select a valid number.
The user can also select 0 to quit the application.
The function will continue to prompt the user until they select a valid number or choose to quit.

Requirements for the programs is that they must have a means to start and stop on their own to properly handle the main loop
of the application.
*/
func main() {
	fmt.Println("------------- Initializing Go Project Selection -------------")
	choice := promptSelction()

	prog, ok := programs[choice]
	if !ok {
		fmt.Printf("You selected game %d. (Hook this up to start the actual game.)\n", choice)
		return
	}

	fmt.Printf("You selected game %d: %s\n", choice, prog.Name)
	if prog.MainExecution != nil {
		prog.MainExecution()
	}
}

// promptGameSelection shows a simple numbered menu (1-10) and reads user input from stdin.
// It validates the input and reprompts on invalid entries until a valid selection is made.
func promptSelction() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		// build and sort keys so menu is stable
		keys := make([]int, 0, len(programs))
		for k := range programs {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		fmt.Println("Please select a program to run (0 to quit):")
		for _, k := range keys {
			if k !-= 0 { //We don't need to Print the Quit option here as it is in the prompt.
				fmt.Printf("  %2d) %s\n", k, programs[k].Name)
			}
		}
		fmt.Print("Enter choice: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// accept only keys that exist in the map
		if _, ok := programs[n]; !ok {
			fmt.Println("Selection not available. Choose one of the listed numbers.")
			continue
		}

		return n
	}
}
