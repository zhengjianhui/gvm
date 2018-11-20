package main

import (
	"fmt"
	"gvm/main/instructions"
	"gvm/main/instructions/base"
	"gvm/main/rtda"
	"gvm/main/rtda/heap"
)

/*
	main -cp "/Users/zhengjianhui/Desktop" GaussTest -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_91.jdk/Contents/Home/jre"
 */
func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		// 获取标识位转换为对应的指令
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		// 从字节码中提取操作数
		inst.FetchOperands(reader)
		// 设置 字节码读取位置
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		// 方法执行指令逻辑
		inst.Execute(frame)
	}
}
