package config

import (
	"os"

	"github.com/siddontang/go/log"
)

var Env string

func GetEnv() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Get system hostname error: ", err.Error())
	}
	if hostname == "yaozhijiedeMacBook-Air.local" {
		Env = "dev"
	} else {
		Env = "production"
	}
}
