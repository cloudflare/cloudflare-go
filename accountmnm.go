// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMnmService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountMnmService] method instead.
type AccountMnmService struct {
	Options []option.RequestOption
	Configs *AccountMnmConfigService
	Rules   *AccountMnmRuleService
}

// NewAccountMnmService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMnmService(opts ...option.RequestOption) (r *AccountMnmService) {
	r = &AccountMnmService{}
	r.Options = opts
	r.Configs = NewAccountMnmConfigService(opts...)
	r.Rules = NewAccountMnmRuleService(opts...)
	return
}
