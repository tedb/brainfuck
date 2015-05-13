package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"io"
)

func main() {
	var buf [30000]byte
	var ptr, instr int

	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		fmt.Println("no filename specified")
		os.Exit(1)
	}
	program, err := ioutil.ReadFile(filename)
	check(err)

	in := bufio.NewReader(os.Stdin)

	for {
		switch program[instr] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			buf[ptr]++
		case '-':
			buf[ptr]--
		case '.':
			//fmt.Print(ptr, buf[ptr], "|", string(buf[ptr]), "|")
			fmt.Print(string(buf[ptr]))
		case ',':
			buf[ptr], err = in.ReadByte()
			//fmt.Print("read #", buf[ptr], "#")
			if err == io.EOF {
				buf[ptr] = 0
				//break
			}
			check(err)
		case '[':
			if buf[ptr] == 0 {
				instr = next_close_loop_index(program, instr) + 1
			}
		case ']':
			if buf[ptr] != 0 {
				instr = prev_open_loop_index(program, instr)
			}
		default:
		}

		instr++
		if instr > len(program)-1 {
			break
		}
	}
	fmt.Print("\n")
}

func next_close_loop_index(program []byte, ptr int) int {
	nested := 0
	for pos, char := range program[ptr+1:] {
		if char == '[' {
			nested++
		} else if char == ']' && nested > 0 {
			nested--
		} else if char == ']' {
			return ptr + 1 + pos
		}
	}
	panic("Could not find matching ]")
}

func prev_open_loop_index(program []byte, ptr int) int {
	nested := 0
	program = program[:ptr]
	for pos := len(program)-1; pos >= 0; pos-- {
		char := program[pos]
		if char == ']' {
			nested++
		} else if char == '[' && nested > 0 {
			nested--
		} else if char == '[' {
			return pos
		}
	}
	panic("Could not find matching [")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
