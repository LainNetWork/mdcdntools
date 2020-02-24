package main

import (
	"flag"
	_ "github.com/microcosm-cc/bluemonday"
	_ "github.com/russross/blackfriday/v2"
	"mdcdntools/common"
	"mdcdntools/processor"
)

func main() {
	flag.StringVar(&common.Config.Path, "p", "", "要处理的文件夹路径")
	flag.StringVar(&common.Config.Refer, "r", "", "从cdn拉取资源时需要的refer地址")
	flag.Parse()
	processor.Execute(*common.Config)
}
