package main

import (
	"bufio"
	"os"
	"strings"
)

func devConsole() {
	reader := bufio.NewReader(os.Stdin)
	for ;; {

		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.Replace(strings.Replace(input, "\r", "", -1), "\n", "", -1))

		if strings.HasPrefix(input, "debugmode") {
			tmp := strings.Replace(input, "debugmode", "", -1)
			if tmp == "" {
				printDebugModeState()
			} else if tmp == "=true" {
				debugmode = true
				printDebugModeState()
			} else if tmp == "=false" {
				debugmode = true
				printDebugModeState()
			} else {
				println("-> Invalid input. See in 'help'")
			}
		} else {
			printDevConsoleHelp()
		}
		println()
	}
}

func printDebugModeState() {
	if debugmode {
		println("-> Debug mode enabled")
	} else {
		println("-> Debug mode disabled")
	}
}

func printDevConsoleHelp() {
	println("----------------------------------")
	println("Developer console")
	println("-----------------")
	println("Commands:")
	println("- debugmode:")
	println("  debugmode          // Show the current state")
	println("  debugmode=true     // Enable debug mode")
	println("  debugmode=false    // Disable debug mode")
	println("----------------------------------")
}