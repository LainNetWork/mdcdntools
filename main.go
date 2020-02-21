package main

import (
	"flag"
	"fmt"
)
import (
	_ "github.com/LainNetWork/mdcdntools/common"
	_ "github.com/microcosm-cc/bluemonday"
	_ "github.com/russross/blackfriday/v2"
)

func main() {
	config:=new(ArgsConfig)
	flag.StringVar(&config.path,"p","","要处理的文件夹路径")
	flag.StringVar(&config.refer,"r","","从cdn拉取资源时需要的refer地址")
	fmt.Println("Hello,Lain！")
}
