// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDlpPatternService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDlpPatternService] method
// instead.
type AccountDlpPatternService struct {
	Options   []option.RequestOption
	Validates *AccountDlpPatternValidateService
}

// NewAccountDlpPatternService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpPatternService(opts ...option.RequestOption) (r *AccountDlpPatternService) {
	r = &AccountDlpPatternService{}
	r.Options = opts
	r.Validates = NewAccountDlpPatternValidateService(opts...)
	return
}
