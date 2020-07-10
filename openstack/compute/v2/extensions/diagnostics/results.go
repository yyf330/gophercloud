package diagnostics

import (
	"github.com/yyf330/gophercloud"
)

type serverDiagnosticsResult struct {
	gophercloud.Result
}

// Extract interprets any diagnostic response as a map
func (r serverDiagnosticsResult) Extract() (map[string]interface{}, error) {
	var s map[string]interface{}
	err := r.ExtractInto(&s)
	return s, err
}
