package main

import (
	"fmt"
	"go-evm/evm"
)

func main() {
	evm := evm.NewEVM(nil)
	fmt.Println("EVM:", evm)
}
