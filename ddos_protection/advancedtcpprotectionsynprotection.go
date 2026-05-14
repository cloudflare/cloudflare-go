// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ddos_protection

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AdvancedTCPProtectionSynProtectionService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionSynProtectionService] method instead.
type AdvancedTCPProtectionSynProtectionService struct {
	Options []option.RequestOption
	Filters *AdvancedTCPProtectionSynProtectionFilterService
	Rules   *AdvancedTCPProtectionSynProtectionRuleService
}

// NewAdvancedTCPProtectionSynProtectionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionSynProtectionService(opts ...option.RequestOption) (r *AdvancedTCPProtectionSynProtectionService) {
	r = &AdvancedTCPProtectionSynProtectionService{}
	r.Options = opts
	r.Filters = NewAdvancedTCPProtectionSynProtectionFilterService(opts...)
	r.Rules = NewAdvancedTCPProtectionSynProtectionRuleService(opts...)
	return
}
