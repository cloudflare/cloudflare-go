// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountLogpushDatasetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountLogpushDatasetService]
// method instead.
type AccountLogpushDatasetService struct {
	Options []option.RequestOption
	Fields  *AccountLogpushDatasetFieldService
	Jobs    *AccountLogpushDatasetJobService
}

// NewAccountLogpushDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushDatasetService(opts ...option.RequestOption) (r *AccountLogpushDatasetService) {
	r = &AccountLogpushDatasetService{}
	r.Options = opts
	r.Fields = NewAccountLogpushDatasetFieldService(opts...)
	r.Jobs = NewAccountLogpushDatasetJobService(opts...)
	return
}
