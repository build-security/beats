// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package tables

import (
	"github.com/osquery/osquery-go/plugin/table"

	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/internal/k8s"
)

func KubePodsColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("uid"),
		table.TextColumn("name"),
		table.TextColumn("namespace"),
		table.TextColumn("ip"),
		table.TextColumn("service_account"),
		table.TextColumn("node_name"),
		table.TextColumn("phase"),
		table.TextColumn("security_context"),
	}
}

func GetKubePodsGenerateFunc() table.GenerateFunc {
	return k8s.Pods
}
