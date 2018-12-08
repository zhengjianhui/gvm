package rtda

// jvm stack jvm 的栈
type Stack struct {
	// 最大栈深
	maxSize uint
	// 当前栈
	size uint
	// stack is implemented as linked list 栈的实现
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

/*
	push 栈
 */
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

/*
	弹出栈顶
 */
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

/*
	返回栈顶, 不弹出
 */
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}
