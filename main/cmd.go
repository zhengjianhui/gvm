package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	main -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_91.jdk/Contents/Home/jre"
 */

type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	cpOption         string
	XjreOption       string
	class            string
	args             []string
}

/*
	1. 设置 flag.Usage
	2. 调用 flag 包的 var() 设置解析的各种选项
	3. 调用 Parse() 解析
	4. 失败后通过 Usage 绑定的 printUsage 输出定义的用法到控制台给出警告
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}

	// go 语言内置 flag 包, 包可以处理命令行 args 变量
	// 1. 设置 flag.Usage
	flag.Usage = printUsage

	// 2. 调用 flag 包的 var() 设置解析的各种选项
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	// 3. 调用 Parse() 解析
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
