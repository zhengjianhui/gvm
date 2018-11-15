package classfile


/*
	字段或方法名称和描述符
	name_index + descriptor_index 可以确定一个方法或字段
	java 虚拟机规范定义了一种简单的语法来描述字段和方法

	1. 类型描述符
		1. 基本类型
			byte 		B
			short		S
			char 		C
			int			I
			long		J
			float		F
			dubbo		D

		2. 引用类型 L + 类的完全限定名
			Object 类型限定符为
				Ljava.lang.Object

		3. 数组类型的描述符是 [ + 数组元素类型描述符
			int[]   —> [I
			double[][] —> [[D
			String[] —> [Ljava.lang.String
	2. 方法描述符
		() 内为参数描述, () I 为返回值
		void run() 		    			—> ()V
		String toString()  				—> ()Ljava.lang.String
		int max(float x, floaty) 		—> (FF)I
		int binSearch(long[] a, long k) —> ([JJ)I


	CONSTANT_NameAndType_info {
    	u1 tag;
    	u2 name_index;
    	u2 descriptor_index;
	}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}