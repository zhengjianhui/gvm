package base

import "gvm/main/rtda"

/*
	指令接口
 */
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 方法执行指令逻辑
	Execute(frame *rtda.Frame)
}

/*
	表示没有操作数的指令
 */
type NoOperandsInstruction struct {
	// empty
}


func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

/*
	表示跳转指令
	Offset 存放跳转偏移量
 */
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/*
	存储和加载类指令需要根据索引存取局部变量表, 索引由单字节操作数给出
	把这类指令抽象成 Index8Instruction 用 Index 字段表示局部变量索引
 */
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/*
	有些指令需要访问运行时常量池, 常量池索引由两字节操作数给出
	把这类指令抽象成 Index16Instruction 用 Index 字段表示常量池索引
 */
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}