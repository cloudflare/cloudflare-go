// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// MnmService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewMnmService] method instead.
type MnmService struct {
	Options []option.RequestOption
	Configs *MnmConfigService
	Rules   *MnmRuleService
}

// NewMnmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMnmService(opts ...option.RequestOption) (r *MnmService) {
	r = &MnmService{}
	r.Options = opts
	r.Configs = NewMnmConfigService(opts...)
	r.Rules = NewMnmRuleService(opts...)
	return
}
