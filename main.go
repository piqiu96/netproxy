package main

import (
	"fmt"
	"github.com/piqiu96/netproxy/util"
	"github.com/piqiu96/netproxy/bootstrap"
)

func main() {
	a := util.GotStr()
	fmt.Println(a)
	fmt.Println("vim-go")

	bootstrap.Hello()
}
