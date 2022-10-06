package configs

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type server struct {
	Host      string
	Port      string
	FeWrokDir string
}

var Server = server{}

func init() {
	cfg, err := ini.Load("configs.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Server = server{
		Host:      cfg.Section("server").Key("host").String(),
		Port:      cfg.Section("server").Key("port").String(),
		FeWrokDir: cfg.Section("server").Key("feWorkDir").String(),
	}
}
