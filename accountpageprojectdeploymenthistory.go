// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPageProjectDeploymentHistoryService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPageProjectDeploymentHistoryService] method instead.
type AccountPageProjectDeploymentHistoryService struct {
	Options []option.RequestOption
	Logs    *AccountPageProjectDeploymentHistoryLogService
}

// NewAccountPageProjectDeploymentHistoryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountPageProjectDeploymentHistoryService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentHistoryService) {
	r = &AccountPageProjectDeploymentHistoryService{}
	r.Options = opts
	r.Logs = NewAccountPageProjectDeploymentHistoryLogService(opts...)
	return
}
