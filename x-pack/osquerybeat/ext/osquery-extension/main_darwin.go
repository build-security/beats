// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build darwin
// +build darwin

package main

import (
	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/internal/tables"
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

func RegisterTables(server *osquery.ExtensionManagerServer) {
	server.RegisterPlugin(table.NewPlugin("host_users", tables.HostUsersColumns(), tables.GetHostUsersGenerateFunc()))
	server.RegisterPlugin(table.NewPlugin("host_groups", tables.HostGroupsColumns(), tables.GetHostGroupsGenerateFunc()))
	server.RegisterPlugin(table.NewPlugin("host_processes", tables.HostProcessesColumns(), tables.GetHostProcessesGenerateFunc()))
	server.RegisterPlugin(table.NewPlugin("k8s_pods", tables.KubePodsColumns(), tables.GetKubePodsGenerateFunc()))
	server.RegisterPlugin(table.NewPlugin("k8s_leases", tables.K8sLeaseColumns(), tables.GetK8sLeasesGenerateFunc()))
}
