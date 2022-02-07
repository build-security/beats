package main

import (
	"fmt"
	"os"

	"github.com/elastic/beats/v7/kubebeat/beater"
	"github.com/elastic/beats/v7/kubebeat/cmd"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"

	_ "github.com/elastic/beats/v7/kubebeat/include"
)

func main1() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	err := instance.Run(instance.Settings{Name: cmd.Name}, beater.New)
	if err != nil {
		fmt.Printf("Main Beat Error %v", err)
	}
}
