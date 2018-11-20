package heap

// symbolic reference
type SymRef struct {
	// 存放符号引用所在的运行时常量池指针
	cp        *ConstantPool
	// 存放类完全限定名
	className string
	// 缓存解析后的类结构体指针
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

/*
	如果类 D 通过符号引用 N 引用类 C 需要解析 N
	先用 D 的类加载器加载 C 然后检查 D 是否有权限访问 C 如果没有 抛出 java.lang.IllegalAccessError

 */
// jvms8 5.4.3.1
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}