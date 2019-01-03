package s8

import (
	"fmt"
	"gvm/main/classpath"
	"gvm/main/interpreter"
	"gvm/main/rtda/heap"
	"strings"
	"testing"
)

func TestArr(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("BubbleSortTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, false, nil)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}

func TestStr(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("StrTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, false, nil)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}
