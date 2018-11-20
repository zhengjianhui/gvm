package stack

import (
	"gvm/main/instructions/base"
	"gvm/main/rtda"
)

/*
	交换栈顶两个元素
 */
// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}