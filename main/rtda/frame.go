package rtda

import "gvm/main/rtda/heap"

/*
	stack frame 栈的节点
 	每一个方法从调用开始到执行完成的过程，就对应着一个栈帧在虚拟机栈里面从入栈到出栈的过程。
	也就是栈帧
*/
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

	method *heap.Method

	// todo
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
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
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}
