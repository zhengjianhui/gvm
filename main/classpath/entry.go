/*
	java 虚拟机规范并没有规定虚拟机应该从哪里寻找类, 因此不同的虚拟机实现可以采用不同的方法.
	oracle 的 java 虚拟机实现根据类路径(class path) 来搜索类. 按照搜索先后顺序可以分为以下 3 个部分
		- 启动类路径(bootstarp classpath) —> 对应 jre/lib (大部分在 rt.jar) java 标准库
		- 扩展类路径(extension classpath) —> 对应 jre/lib/ext 目录 使用 java 扩展机制的类位于该路径
		- 用户类路径(user classpath) 第三方类库则位于用户类路径, 可以通过 -Xbootclasspath 选线修改启动路径, 正常用不到
 */
package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/*
		寻找和加载 class 文件
	 */
	readClass(className string) ([]byte, Entry, error)

	/*
		toString
	 */
	String() string

}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	return newDirEntry(path)
}
