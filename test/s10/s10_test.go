package s9

import (
	"fmt"
	"gvm/main/classpath"
	"gvm/main/interpreter"
	"gvm/main/rtda/heap"
	"strings"
	"testing"
)

func TestException(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("ParseIntTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, false, nil)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}





