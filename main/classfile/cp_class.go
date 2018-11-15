package classfile

/*
	类和超类以及接口表中的接口索引指向 Constant_Class_Info
	CONSTANT_Class_info {
    	u1 tag;
    	u2 name_index;
	}
*/
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
