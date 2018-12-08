package references

import (
	"gvm/main/instructions/base"
	"gvm/main/rtda"
	"gvm/main/rtda/heap"
)

/*
	new 指令的操作数是一个 uint16 索引, 来自字节码.
	通过这个索引可以从当前类的运行时常量池中找到一个类符号引用, 解析这个类符号引用, 拿到类数据, 然后创建兑现, 并把对象引用推入栈顶

	new 指令创建类实例, 但是类还没有被初始化
	执行 putstatic getstatic 指令存取类的静态变量, 但声明的该字段的类还没有被执行初始化
	执行 invokestatic 调用类的静态方法, 但声明改方法的类还没有被初始化
	当初始化一个类的时候,  超类没有被初始化要先初始化超类
	执行某些反射操作需要判断类是否初始化了
	@see putstatic
	@see getstatic
	@see invokestatic
 */
// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	// 如果类没有初始化则先初始化类
	// 由于指令已经执行过了, pc 会指向下一条指令
	// 所以这边重新设置帧的 pc (用线程的 pc 重置帧的pc)
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}