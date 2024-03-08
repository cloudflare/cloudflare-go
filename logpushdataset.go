// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// LogpushDatasetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogpushDatasetService] method
// instead.
type LogpushDatasetService struct {
	Options []option.RequestOption
	Fields  *LogpushDatasetFieldService
	Jobs    *LogpushDatasetJobService
}

// NewLogpushDatasetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogpushDatasetService(opts ...option.RequestOption) (r *LogpushDatasetService) {
	r = &LogpushDatasetService{}
	r.Options = opts
	r.Fields = NewLogpushDatasetFieldService(opts...)
	r.Jobs = NewLogpushDatasetJobService(opts...)
	return
}
