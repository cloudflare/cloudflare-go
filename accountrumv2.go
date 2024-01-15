// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRumV2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRumV2Service] method
// instead.
type AccountRumV2Service struct {
	Options []option.RequestOption
	Rule    *AccountRumV2RuleService
}

// NewAccountRumV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRumV2Service(opts ...option.RequestOption) (r *AccountRumV2Service) {
	r = &AccountRumV2Service{}
	r.Options = opts
	r.Rule = NewAccountRumV2RuleService(opts...)
	return
}
