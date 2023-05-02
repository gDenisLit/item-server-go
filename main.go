package main

import (
	"github.com/gDenisLit/item-server-go/cmd"
	"github.com/gDenisLit/item-server-go/cmd/services"
)

func main() {
	services.LoadEnv()
	cmd.InitServer()
}
