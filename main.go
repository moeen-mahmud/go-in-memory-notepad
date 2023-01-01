package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func askForCommand() {
	fmt.Print("Enter a command and data:")
}

func verifyCommands(scannedCommand string) bool {
	commands := []string{"create", "exit", "list", "clear", "update", "delete"}
	for _, command := range commands {
		if scannedCommand == command {
			return true
		}
		return false
	}
	return false
}

func hasWhiteSpace(str string, command string) bool {
	trimmedStr := strings.TrimSpace(str)

	words := strings.Split(trimmedStr, " ")

	if len(words) == 1 {
		if words[0] == command {
			return true
		}
	}
	return false
}

func verifyPositionBoundaries(position string, notes []string) bool {
	parsedPosition, _ := strconv.Atoi(position)
	if parsedPosition-1 < len(notes) {
		return true
	}
	return false
}

func main() {
	// 	write your code here
	var noteCapacity int

	fmt.Print("Enter the maximum number of notes:")
	fmt.Scan(&noteCapacity)
	// askForCommand()
	scanner := bufio.NewScanner(os.Stdin)

	// make note slice
	notes := make([]string, 0, noteCapacity-1)

	askForCommand()
	for scanner.Scan() {
		scannedText := scanner.Text()
		trimmedStrings := strings.TrimSpace(scannedText)
		splittedString := strings.Split(trimmedStrings, " ")
		scannedCommand := splittedString[0]

		if scannedCommand == "create" {
			if len(notes) <= noteCapacity-1 && hasWhiteSpace(scannedText, "create") == false {
				trimmedNotes := strings.TrimLeft(scannedText, "create ")
				notes = append(notes, trimmedNotes)
				fmt.Print("[OK] The note was successfully created\n")
			} else if hasWhiteSpace(scannedText, "create") == true {
				fmt.Print("[Error] Missing note argument\n")
			} else {
				fmt.Print("[Error] Notepad is full\n")
			}
			askForCommand()
		} else if scannedCommand == "update" {
			if hasWhiteSpace(scannedText, "update") == true {
				fmt.Println("[Error] Missing position argument")
			} else if len(notes) < 1 {
				fmt.Println("[Error] There is nothing to update")
			} else if _, err := strconv.Atoi(splittedString[1]); err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", splittedString[1])
				// if the position command has no space after it
			} else if len(notes) > 0 && hasWhiteSpace(scannedText, "update") == false && verifyPositionBoundaries(splittedString[1], notes) == true && len(splittedString) > 2 {
				position, _ := strconv.Atoi(splittedString[1])
				trimmedNotes := strings.TrimLeft(scannedText, "update "+splittedString[1]+" ")
				notes[position-1] = trimmedNotes
				fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
			} else if verifyPositionBoundaries(splittedString[1], notes) == false {
				position, _ := strconv.Atoi(splittedString[1])
				fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, cap(notes)+1)
			} else if len(splittedString) == 2 && hasWhiteSpace(scannedText, "update") == false {
				fmt.Printf("[Error] Missing note argument\n")
			}
			askForCommand()
		} else if scannedCommand == "delete" {
			if hasWhiteSpace(scannedText, "delete") == true {
				fmt.Println("[Error] Missing position argument")
			} else if len(notes) < 1 {
				fmt.Println("[Error] There is nothing to delete")
			} else if _, err := strconv.Atoi(splittedString[1]); err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", splittedString[1])
			} else if len(notes) > 0 && hasWhiteSpace(scannedText, "delete") == false && verifyPositionBoundaries(splittedString[1], notes) == true && len(splittedString) == 2 {
				position, _ := strconv.Atoi(splittedString[1])
				notes = append(notes[:position-1], notes[position:]...)
				fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
			} else if verifyPositionBoundaries(splittedString[1], notes) == false {
				position, _ := strconv.Atoi(splittedString[1])
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, len(notes))
			}
			askForCommand()

		} else if scannedCommand == "clear" {
			notes = nil
			fmt.Print("[OK] All notes were successfully deleted\n")
			askForCommand()
		} else if scannedCommand == "list" {
			if len(notes) <= 0 {
				fmt.Print("[Info] Notepad is empty\n")
			} else {
				for index, note := range notes {
					fmt.Printf("[Info] %d: %s\n", index+1, note)
				}
			}
			askForCommand()
		} else if scannedCommand == "exit" {
			fmt.Print("[Info] Bye!\n")
			break
		} else if verifyCommands(scannedCommand) == false && len(scannedCommand) > 0 {
			fmt.Print("[Error] Unknown command\n")
			askForCommand()
		}
	}
}
