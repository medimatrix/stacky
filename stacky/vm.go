package stacky

import (
	"fmt"
	"log"
)

const maxStackSize = 128

type stackVal int
type stack []stackVal

// A virtual machine interprets instructions and pushes values to the stack (and
// pops them later) when instructed.
type VM struct {
	stack stack
}

func assert(test bool, message string) {
	if test == false {
		log.Fatalln(message)
	}
}

func push(s *stack, data stackVal) {
	assert(len(*s) < maxStackSize, "Stack overflow")
	*s = append(*s, data)
}

func pop(s *stack) stackVal {
	assert(len(*s) > 0, "Stack is not big enough")
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

// Interpret runs instructions (bytecode).
func (vm *VM) Interpret(instructions instructions) {
	skip := false

	for i, instruction := range instructions {
		if skip {
			skip = false
			continue
		}

		switch instruction {
		case instPrint:
			strLen := int(instructions[i+1])
			assert(len(vm.stack) >= strLen, "Stack is not big enough")
			chars := vm.stack[len(vm.stack)-strLen : len(vm.stack)]

			// Note: string([]byte) works, but go does not want to convert the
			// stack type.
			for _, char := range chars {
				pop(&vm.stack)
				fmt.Print(string(char))
			}
			// Append a newline
			fmt.Println("")
		case instAdd:
			n1 := pop(&vm.stack)
			n2 := pop(&vm.stack)
			push(&vm.stack, n1+n2)
		case instMult:
			n1 := pop(&vm.stack)
			n2 := pop(&vm.stack)
			push(&vm.stack, n1*n2)
		case instLiteral:
			val := stackVal(instructions[i+1])
			push(&vm.stack, val)
			skip = true
		case instDBGSTK:
			fmt.Println(vm.stack)
		}
	}
}

// NewVM returns a new virtual machine that has its own stack.
func NewVM() *VM {
	stack := make(stack, 0)
	return &VM{stack}
}