package rtda


// stack frame 栈的节点
type Frame struct {
	// stack is implemented as linked list 下一个节点
	lower        *Frame
	// 本地变量表
	localVars    LocalVars
	// 操作数栈
	operandStack *OperandStack
	// todo
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}