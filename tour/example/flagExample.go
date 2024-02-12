package example

import (
	"flag"
	"log"
)

func flagExample() {
	//==============flag 基本使用和长短选项==================
	// $ go run main.go -name=eddycjy -n= 煎鱼
	// name: 煎鱼
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
	//=================子命令的实现=========================
	// go run main.go go -name=eddycjy
	// name: eddycjy
	flag.Parse() //将命令行解析为定义的标志，便于我们后续的参数使用。

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		//相当于就是创建一个新的命令集了去支持子命令,并且遇到错误类型会进行处理
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
