// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IntelService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewIntelService] method instead.
type IntelService struct {
	Options            []option.RequestOption
	ASN                *IntelASNService
	DNS                *IntelDNSService
	Domains            *IntelDomainService
	DomainHistories    *IntelDomainHistoryService
	IPs                *IntelIPService
	IPLists            *IntelIPListService
	Miscategorizations *IntelMiscategorizationService
	Whois              *IntelWhoisService
	IndicatorFeeds     *IntelIndicatorFeedService
	Sinkholes          *IntelSinkholeService
}

// NewIntelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIntelService(opts ...option.RequestOption) (r *IntelService) {
	r = &IntelService{}
	r.Options = opts
	r.ASN = NewIntelASNService(opts...)
	r.DNS = NewIntelDNSService(opts...)
	r.Domains = NewIntelDomainService(opts...)
	r.DomainHistories = NewIntelDomainHistoryService(opts...)
	r.IPs = NewIntelIPService(opts...)
	r.IPLists = NewIntelIPListService(opts...)
	r.Miscategorizations = NewIntelMiscategorizationService(opts...)
	r.Whois = NewIntelWhoisService(opts...)
	r.IndicatorFeeds = NewIntelIndicatorFeedService(opts...)
	r.Sinkholes = NewIntelSinkholeService(opts...)
	return
}
