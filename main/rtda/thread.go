package rtda

import "gvm/main/rtda/heap"

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
	// 当前执行的指令地址
	pc int
	// 栈指针
	stack *Stack
}


func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/*
	返回当前帧
 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}


func (self *Thread) ClearStack() {
	self.stack.clear()
}

func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}