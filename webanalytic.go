// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WebAnalyticService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWebAnalyticService] method
// instead.
type WebAnalyticService struct {
	Options  []option.RequestOption
	SiteInfo *WebAnalyticSiteInfoService
	Rules    *WebAnalyticRuleService
}

// NewWebAnalyticService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebAnalyticService(opts ...option.RequestOption) (r *WebAnalyticService) {
	r = &WebAnalyticService{}
	r.Options = opts
	r.SiteInfo = NewWebAnalyticSiteInfoService(opts...)
	r.Rules = NewWebAnalyticRuleService(opts...)
	return
}
