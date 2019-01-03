package s9

import (
	"fmt"
	"gvm/main/classpath"
	"gvm/main/interpreter"
	"gvm/main/rtda/heap"
	"strings"
	"testing"
)

func TestGetClass(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("GetClassTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, false, nil)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}


func TestHashCode(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("ObjectTest", ".", "/", -1)
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

	className := strings.Replace("StringTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, false, nil)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}

func TestClone(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("CloneTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, false, nil)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}


func TestBox(t *testing.T) {

	cp := classpath.Parse("", "/Users/zhengjianhui/go/src/gvm/java")
	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace("BoxTest", ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, false, nil)
	} else {
		fmt.Printf("Main method not found in class %s\n", "GaussTest")
	}

}
