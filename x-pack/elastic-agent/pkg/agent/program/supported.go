// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/elastic-agent/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/kubebeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
