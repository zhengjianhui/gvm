package base

import (
	"gvm/main/rtda"
	"gvm/main/rtda/heap"
)

/*
	在定位到需要调用的方法之后, java 虚拟机需要给这个方法创建一个新的帧并把它推入 java 虚拟机栈顶, 然后传递参数
 */
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	// 通过帧拿到线程
	thread := invokerFrame.Thread()
	// 新建一个帧
	newFrame := thread.NewFrame(method)
	// 将新建的帧推入栈顶
	thread.PushFrame(newFrame)

	/*
		参数的传递在局部变量表中占用多少位置(也就是本地变量表)
		这个数量并一定等于 java 代码中的参数个数
		因为 long 和 double 的参数要占用两个位置
		并且对于实例方法, java 编译器会在参数列表的前面添加一个 this
		假设知己参数占用 n 个位置
		一次把这 n 个变量冲调用者的操作数栈中弹出, 放进被调用的方法的局部标量表(本地变量表)
	 */
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	//// hack! 防止调用到 registerNatives 方法
	//if method.IsNative() {
	//	if method.Name() == "registerNatives" {
	//		thread.PopFrame()
	//	} else {
	//		panic(fmt.Sprintf("native method: %v.%v%v\n",
	//			method.Class().Name(), method.Name(), method.Descriptor()))
	//	}
	//}
}
