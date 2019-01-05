package classfile

import "fmt"

/*
	ClassFile 结构体反映了 java 虚拟机规范定义的 class文件格式
 */
type ClassFile struct {
	// u4      magic
	//magic      uint32

	// u2      minor_version
	minorVersion uint16

	// u2      major_version
	majorVersion uint16

	// u2      constant_pool_count
	// cp_info constant_pool[constant_pool_count-1]
	constantPool ConstantPool

	// u2             access_flags
	accessFlags uint16

	// u2             this_class
	thisClass uint16
	// u2             super_class
	superClass uint16

	// u2             interfaces_count
	// u2             interfaces[interfaces_count]
	interfaces []uint16

	// u2             fields_count
	// field_info     fields[fields_count]
	fields []*MemberInfo

	// u2             methods_count
	// method_info    methods[methods_count]
	methods []*MemberInfo

	// u2             attributes_count
	// attribute_info attributes[attributes_count]
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)

	return
}


func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)

	// 类访问标志, 判断类是接口还是类(类的定义)
	self.accessFlags = reader.readUint16()
	// 类名
	self.thisClass = reader.readUint16()
	// 超类索引表 除了 Object 没有超类, 为0 余下都必须在常量池有索引
	self.superClass = reader.readUint16()
	// 接口索引表
	self.interfaces = reader.readUint16s()
	// 字段表
	self.fields = readMembers(reader, self.constantPool)
	// 方法表
	self.methods = readMembers(reader, self.constantPool)

	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	// java 魔数 0xCAFEBABE
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
	部分虚拟机只支持固定版本的 class 文件
	oracle 的虚拟机向后兼容
	1.2 后很少区分小版本
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 小版本号, 如 1.8.0_91 的 91
	self.minorVersion = reader.readUint16()
	// 大版本号 1.6 1.7 1.8
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	// jdk8 的 major 为 45
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}


func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}