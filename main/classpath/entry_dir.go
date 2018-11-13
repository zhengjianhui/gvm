package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}


func newDirEntry(path string) *DirEntry {
	// 转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		// 如果发生异常调用 panic 结束函数
		panic(err)
	}

	return &DirEntry{absDir}
}

/*
	绑定 readClass 方法
 */
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

/*
	绑定 toString 方法
 */
func (self *DirEntry) String() string {
	return self.absDir
}