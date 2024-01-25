// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountLogpushService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountLogpushService] method
// instead.
type AccountLogpushService struct {
	Options    []option.RequestOption
	Datasets   *AccountLogpushDatasetService
	Jobs       *AccountLogpushJobService
	Ownerships *AccountLogpushOwnershipService
	Validates  *AccountLogpushValidateService
}

// NewAccountLogpushService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountLogpushService(opts ...option.RequestOption) (r *AccountLogpushService) {
	r = &AccountLogpushService{}
	r.Options = opts
	r.Datasets = NewAccountLogpushDatasetService(opts...)
	r.Jobs = NewAccountLogpushJobService(opts...)
	r.Ownerships = NewAccountLogpushOwnershipService(opts...)
	r.Validates = NewAccountLogpushValidateService(opts...)
	return
}
