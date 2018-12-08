package references

import "gvm/main/instructions/base"
import "gvm/main/rtda"
import "gvm/main/rtda/heap"

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()

	// 如果类没有初始化则先初始化类
	// 由于指令已经执行过了, pc 会指向下一条指令
	// 所以这边重新设置帧的 pc (用线程的 pc 重置帧的pc)
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}
