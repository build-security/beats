package main

import (
	"fmt"
	"os"

<<<<<<< HEAD:kubebeat/main.go
	"github.com/elastic/beats/v7/kubebeat/beater"
	"github.com/elastic/beats/v7/kubebeat/cmd"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"
=======
	"github.com/elastic/beats/v7/cloudbeat/cmd"
>>>>>>> master:cloudbeat/main.go

	_ "github.com/elastic/beats/v7/cloudbeat/include"
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
