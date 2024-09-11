// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options               []option.RequestOption
	AllowPatterns         *SettingAllowPatternService
	BlockSenders          *SettingBlockSenderService
	Domains               *SettingDomainService
	ImpersonationRegistry *SettingImpersonationRegistryService
	TrustedDomains        *SettingTrustedDomainService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	r.AllowPatterns = NewSettingAllowPatternService(opts...)
	r.BlockSenders = NewSettingBlockSenderService(opts...)
	r.Domains = NewSettingDomainService(opts...)
	r.ImpersonationRegistry = NewSettingImpersonationRegistryService(opts...)
	r.TrustedDomains = NewSettingTrustedDomainService(opts...)
	return
}