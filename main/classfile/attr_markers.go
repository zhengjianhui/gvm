package classfile


/*
	Deprecated 用于支出类, 接口, 字段或方法已经不建议使用

	@Deprecated
	public void oldMethod() {}

	Deprecated_attribute {
    	u2 attribute_name_index;
    	u4 attribute_length;
	}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
	Synthetic 用于标记源文件中不存在, 由编译器生成的类成员
	Synthetic_attribute {
    	u2 attribute_name_index;
    	u4 attribute_length;
	}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

/*
	只起到标记做用所以是空的实现
 */
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}