package main

import (
	"flag"
	"fmt"
)

func InterpretProgram(program string) int {
	pointer := 0
	memory := make([]int, 2048)
	for i, char := range program {
		if char == '<' {
			i++
			pointer -= int(program[i]) - 48
		} else if char == '>' {
			i++
			pointer += int(program[i]) - 48
		} else if char == '+' {
			i++
			memory[pointer] += int(program[i]) - 48
		} else if char == '*' {
			i++
			memory[pointer] *= int(program[i]) - 48
		} else if char == '-' {
			memory[pointer]--
		} else if char == 'p' {
			fmt.Printf("%d\n", memory[pointer])
		} else if char == 'i' {
			var scanned int
			fmt.Scan(&scanned)
			memory[pointer] = scanned
		}
	}
	return 0
}
func main() {
	program := flag.String("program", "", "Program to give to the interpreter.")
	flag.Parse()
	programNoPointer := *program
	InterpretProgram(programNoPointer)
}
