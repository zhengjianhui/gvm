package base

import (
	"gvm/main/rtda"
	"gvm/main/rtda/heap"
)

// jvms 5.5
func InitClass(thread *rtda.Thread, class *heap.Class) {
	// 先将类的初始化状态设置为开始
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

/*
	获取初始化方法
 */
func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// 将初始化方法推入栈顶的一帧
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

/*
	先执行超类的初始化
 */
func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}