package cmd

import (
	"github.com/elastic/beats/v7/kubebeat/beater"

	cmd "github.com/elastic/beats/v7/libbeat/cmd"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"
	"github.com/elastic/beats/v7/libbeat/publisher/processing"

	// Register the fleet management
	_ "github.com/elastic/beats/v7/x-pack/libbeat/include"
)

// Name of this beat
var Name = "kubebeat"

var RootCmd = Kubebeat()

func Kubebeat() *cmd.BeatsRootCmd {
	settings := instance.Settings{
		Name:            Name,
		Processing:      processing.MakeDefaultBeatSupport(true),
		ElasticLicensed: true,
	}
	command := cmd.GenRootCmdWithSettings(beater.New, settings)

	return command
}
