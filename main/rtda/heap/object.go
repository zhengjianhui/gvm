package heap

/*
	数组的相关指令
	public class ArrayDemo {

		public static void main(String[] args) {
			// newarray
			int[] arr = new int[10];

			// new anewarray
			String[] objArr = new String[10];

			// multianewarray
			int[][] muArr = new int[10][10];

			// arraylenght
			int x = arr.lenght;

			// iastore
			arr[0]  = 100;

			// iaload
			int y = arr[0];

			// aastore
			objArr[0] = “aaaa”;

			// aaload
			String s = objArr[0];
		}

	}
 */
type Object struct {
	class  *Class
	data  interface{} // Slots for Object, []int32 for int[] ...
}

// 用于分配内存
func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
