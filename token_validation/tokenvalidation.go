// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// TokenValidationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenValidationService] method instead.
type TokenValidationService struct {
	Options       []option.RequestOption
	Configuration *ConfigurationService
	Rules         *RuleService
}

// NewTokenValidationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenValidationService(opts ...option.RequestOption) (r *TokenValidationService) {
	r = &TokenValidationService{}
	r.Options = opts
	r.Configuration = NewConfigurationService(opts...)
	r.Rules = NewRuleService(opts...)
	return
}
