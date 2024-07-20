package main

import (
	"github.com/KubeOperator/kobe/cmd/client/root"
	"github.com/KubeOperator/kobe/pkg/config"
	"os"
)

func main() {
	config.Init()
	if err := root.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
