package main

import (
	"fmt"

	"github.com/anmolbabu/go-syscall-ex/pkg/utils"
)

func main() {

	uid := utils.GetUID()
	fmt.Println(uid)

	pid := utils.GetPID()
	fmt.Println(pid)
}
