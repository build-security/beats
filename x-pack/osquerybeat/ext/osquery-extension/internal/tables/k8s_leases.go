// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package tables

import (
	"github.com/osquery/osquery-go/plugin/table"

	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/internal/k8s"
)

func K8sLeaseColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("uid"),
		table.TextColumn("name"),
		table.TextColumn("namespace"),
		table.TextColumn("holderIdentity"),
		table.TextColumn("leaseDurationSeconds"),
		table.TextColumn("acquireTime"),
		table.TextColumn("renewTime"),
		table.TextColumn("leaseTransitions"),
	}
}

func GetK8sLeasesGenerateFunc() table.GenerateFunc {
	return k8s.Leases
}
