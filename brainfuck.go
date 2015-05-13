package main

import (
"bufio"
"fmt"
"os"
"io/ioutil"
)

func main() {
	var buf [30000]byte
	var ptr, instr int

	filename := os.Args[1]
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
				fmt.Print(string(buf[ptr]))
			case ',':
				buf[ptr], err = in.ReadByte()
				check(err)
		}

		instr++
		if instr > len(program)-1 {
			break;
		}
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
