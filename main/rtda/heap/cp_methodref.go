package heap

import "gvm/main/classfile"

/*
	方法符号引用
 */
type MethodRef struct {
	MemberRef
	// 方法符号引用
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

/*
	如果解析过方法的符号引用直接返回
	没有则解析
 */
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	// 如果类 D 想通过方法的符号引用访问类C的某个方法, 先要解析符号引用得到类C
	// 根据方法名和描述符查找方法

	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	// 先从 class 的继承体系中去找
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		// 承上, 找不到的情况下从 interfaces 中查找
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}