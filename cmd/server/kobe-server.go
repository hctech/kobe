package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/KubeOperator/kobe/pkg/config"
	"github.com/KubeOperator/kobe/pkg/constant"
	"log"
)

func main() {
	config.Init()
	constant.Init()
	if err := prepareStart(); err != nil {
		log.Fatal(err)
	}
	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")
	address := fmt.Sprintf("%s:%d", host, port)
	lis, err := newTcpListener(address)
	if err != nil {
		log.Fatal(err)
	}
	server := newServer()
	log.Printf("kobe server listen at: %s", address)
	if err := server.Serve(*lis); err != nil {
		log.Fatal(err)
	}
}
