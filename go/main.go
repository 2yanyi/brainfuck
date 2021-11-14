package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

var compileError = errors.New("compile error")

const (
	_OPR = '>' // Pointer right  将指针向右移动
	_OPL = '<' // Pointer left   将指针向左移动
	_ADD = '+' // Add unit       增加指针处的内存单元
	_SUB = '-' // Sub unit       减少指针处的内存单元
	_OUT = '.' // Output         输出指针所在单元格所代表的字符
	_OIN = ',' // Input          输入一个字符并将其存储在指针所在的单元格中
	_OJF = '[' // Jump forward   如果指针处的单元格为零，则跳过匹配项
	_OJB = ']' // Jump back      如果指针处的单元格非零，则跳回匹配项
)

type opcode struct {
	code rune
	jump int
}

func compile(input []rune) ([]opcode, error) {
	program := make([]opcode, 0, 100)
	stack := make([]int, 0, 100)
	jump := 0
	pc := 0
	for _, char := range input {
		switch char {
		case _OPR:
			program = append(program, opcode{code: _OPR})
		case _OPL:
			program = append(program, opcode{code: _OPL})
		case _ADD:
			program = append(program, opcode{code: _ADD})
		case _SUB:
			program = append(program, opcode{code: _SUB})
		case _OUT:
			program = append(program, opcode{code: _OUT})
		case _OIN:
			program = append(program, opcode{code: _OIN})
		case _OJF:
			program = append(program, opcode{code: _OJF})
			stack = append(stack, pc)
		case _OJB:
			if len(stack) == 0 {
				return nil, compileError
			}
			jump = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			program = append(program, opcode{code: _OJB, jump: jump})
			program[jump].jump = pc
		default:
			pc--
		}
		pc++
	}
	if len(stack) != 0 {
		return nil, compileError
	}
	return program, nil
}

func execute(program []opcode) {
	reader := bufio.NewReader(os.Stdin)
	data := make([]int, math.MaxUint16)
	dataPtr := 0
	for i := 0; i < len(program); i++ {
		switch program[i].code {
		case _OPR:
			dataPtr++
		case _OPL:
			dataPtr--
		case _ADD:
			data[dataPtr]++
		case _SUB:
			data[dataPtr]--
		case _OUT:
			fmt.Printf("%c", data[dataPtr])
		case _OIN:
			value, err := reader.ReadByte()
			if err != nil {
				panic(compileError)
			}
			data[dataPtr] = int(value)
		case _OJF:
			if data[dataPtr] == 0 {
				i = program[i].jump
			}
		case _OJB:
			if data[dataPtr] > 0 {
				i = program[i].jump
			}
		default:
			panic(compileError)
		}
	}
}

func run(code string) {
	program, err := compile([]rune(code))
	if err != nil {
		panic(err)
	}
	execute(program)
}

func main() {
	run(cat("testdata/hello_world.bf"))
	run(cat("testdata/mandelbrot.bf"))
}

func cat(fp string) string {
	data, _ := ioutil.ReadFile(fp)
	return string(data)
}
