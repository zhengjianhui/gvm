package main

import (
	"fmt"
	"gvm/main/classpath"
	"gvm/main/rtda/heap"
	"gvm/main/cmd"
	"gvm/main/interpreter"
	"strings"
)


func main() {
	op := cmd.ParseCmd()

	if op.VersionFlag()  {
		fmt.Println("version 0.0.1")
	} else if op.HelpFlag() || op.Class() == "" {
		cmd.PrintUsage()
	} else {
		startJVM(op)
	}
}

func startJVM(op *cmd.Cmd) {
	cp := classpath.Parse(op.XjreOption(), op.CpOption())
	classLoader := heap.NewClassLoader(cp, op.VerboseClassFlag())

	className := strings.Replace(op.Class(), ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, op.VerboseInstFlag(), op.Args())
	} else {
		fmt.Printf("Main method not found in class %s\n", op.Class())
	}
}