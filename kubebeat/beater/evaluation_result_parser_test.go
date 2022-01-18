package beater

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEvaluationResultParserParseResult(t *testing.T) {

	var result map[string]interface{}
	json.Unmarshal([]byte(jsonExample), &result)
	runId, _ := uuid.NewV4()
	timestamp := time.Now()
	//Creating a new evaluation parser
	parser, _ := NewEvaluationResultParser()

	parsedResult, err := parser.ParseResult(nil, result, runId, timestamp)
	if err != nil {
		assert.Fail(t, "error during parsing of the json", err)
	}

	for _, event := range parsedResult {

		assert.Equal(t, timestamp, event.Timestamp, `event timestamp is not correct`)
		assert.Equal(t, runId, event.Fields["run_id"], "event run_id is not correct")
		assert.NotEmpty(t, event.Fields["result"], "event result is missing")
		assert.NotEmpty(t, runId, event.Fields["rule"], "event rule is missing")
		assert.NotEmpty(t, runId, event.Fields["resource"], "event resource is missing")
	}
}

var jsonExample = `{
"findings":
[
{
	"result": {
	"evaluation": "failed",
	"evidence": {
		"filemode": "700"
	}
},
"rule": {
"benchmark": "CIS Kubernetes",
"description": "The scheduler.conf file is the kubeconfig file for the Scheduler. You should restrict its file permissions to maintain the integrity of the file. The file should be writable by only the administrators on the system.",
"impact": "None",
"name": "Ensure that the scheduler.conf file permissions are set to 644 or more restrictive",
"remediation": "chmod 644 /etc/kubernetes/scheduler.conf",
"tags": [
"CIS",
"CIS v1.6.0",
"Kubernetes",
"CIS 1.1.15",
"Master Node Configuration"
]
}
},
{
"result": {
"evaluation": "passed",
"evidence": {
"gid": "root",
"uid": "root"
}
},
"rule": {
"benchmark": "CIS Kubernetes",
"description": "The scheduler.conf file is the kubeconfig file for the Scheduler. You should set its file ownership to maintain the integrity of the file. The file should be owned by root:root.",
"impact": "None",
"name": "Ensure that the scheduler.conf file ownership is set to root:root",
"remediation": "chown root:root /etc/kubernetes/scheduler.conf",
"tags": [
"CIS",
"CIS v1.6.0",
"Kubernetes",
"CIS 1.1.16",
"Master Node Configuration"
]
}
}
],
"resource": {
"filename": "scheduler.conf",
"gid": "root",
"mode": "700",
"path": "/hostfs/etc/kubernetes/scheduler.conf",
"type": "file-system",
"uid": "root"
}
}
`
