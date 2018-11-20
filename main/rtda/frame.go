package rtda

// stack frame 栈的节点
type Frame struct {
	// stack is implemented as linked list 下一个节点
	lower *Frame
	// 本地变量表
	localVars LocalVars
	// 操作数栈
	operandStack *OperandStack
	thread       *Thread
	// the next instruction after the call
	nextPC int

	// todo
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
