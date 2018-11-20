package references

import (
	"gvm/main/instructions/base"
	"gvm/main/rtda"
	"gvm/main/rtda/heap"
)

/*
	new 指令的操作数是一个 uint16 索引, 来自字节码.
	通过这个索引可以从当前类的运行时常量池中找到一个类符号引用, 解析这个类符号引用, 拿到类数据, 然后创建兑现, 并把对象引用推入栈顶
 */
// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// todo: init class

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}