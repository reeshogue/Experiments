package main

import (
	"flag"
	"fmt"
)

/*In this sea of what's computable, I must create and find
A machine of logic, I do hope it's quite sublime.
The thought of an oasis begins to fill my mind.
Polygonal numbers, even turing machines, primes.
I know, I know. It's not ideal,
A simple sieve would be such an ordeal.*/

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
