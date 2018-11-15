package rtda

/*
 	局部变量表的数据结构
	局部变量表是按索引访问的, 所以很自然可以把他想象为一个数组,
	这个数组的每个元素至少可以容纳一个 int 或 一个引用, 两个连续的元素可以容纳一个 long 或 double
 */
type Slot struct {
	// 保存 32 位整数
	num int32
	// 保存一个引用
	ref *Object
}