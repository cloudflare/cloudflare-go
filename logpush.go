// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// LogpushService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewLogpushService] method instead.
type LogpushService struct {
	Options   []option.RequestOption
	Datasets  *LogpushDatasetService
	Edge      *LogpushEdgeService
	Jobs      *LogpushJobService
	Ownership *LogpushOwnershipService
	Validate  *LogpushValidateService
}

// NewLogpushService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogpushService(opts ...option.RequestOption) (r *LogpushService) {
	r = &LogpushService{}
	r.Options = opts
	r.Datasets = NewLogpushDatasetService(opts...)
	r.Edge = NewLogpushEdgeService(opts...)
	r.Jobs = NewLogpushJobService(opts...)
	r.Ownership = NewLogpushOwnershipService(opts...)
	r.Validate = NewLogpushValidateService(opts...)
	return
}
