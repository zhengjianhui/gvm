package arr

import (
	"fmt"
	"gvm/main/classpath"
	"gvm/main/instructions"
	"gvm/main/instructions/base"
	"gvm/main/rtda"
	"gvm/main/rtda/heap"
	"strings"
	"testing"
)

func TestStack(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("BubbleSortTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, false)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}



func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		// 获取标识位转换为对应的指令
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		// 从字节码中提取操作数
		inst.FetchOperands(reader)
		// 设置 字节码读取位置
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// 方法执行指令逻辑
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}