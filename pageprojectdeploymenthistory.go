// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// PageProjectDeploymentHistoryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewPageProjectDeploymentHistoryService] method instead.
type PageProjectDeploymentHistoryService struct {
	Options []option.RequestOption
	Logs    *PageProjectDeploymentHistoryLogService
}

// NewPageProjectDeploymentHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPageProjectDeploymentHistoryService(opts ...option.RequestOption) (r *PageProjectDeploymentHistoryService) {
	r = &PageProjectDeploymentHistoryService{}
	r.Options = opts
	r.Logs = NewPageProjectDeploymentHistoryLogService(opts...)
	return
}
