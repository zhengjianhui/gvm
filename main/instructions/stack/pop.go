package stack

import (
	"gvm/main/instructions/base"
	"gvm/main/rtda"
)

/*
	pop 指令只能用于弹出 int float 等占用一个操作数栈位置的变量
	pop2 则可以弹出 double 和 long
 */

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
