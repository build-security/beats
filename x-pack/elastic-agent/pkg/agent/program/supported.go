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
	// spec/cloudbeat.yml
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
