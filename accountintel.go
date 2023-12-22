// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountIntelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelService] method
// instead.
type AccountIntelService struct {
	Options            []option.RequestOption
	ASNs               *AccountIntelASNService
	DNS                *AccountIntelDNSService
	Domains            *AccountIntelDomainService
	DomainHistories    *AccountIntelDomainHistoryService
	IPs                *AccountIntelIPService
	IPLists            *AccountIntelIPListService
	Miscategorizations *AccountIntelMiscategorizationService
	Whois              *AccountIntelWhoisService
}

// NewAccountIntelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountIntelService(opts ...option.RequestOption) (r *AccountIntelService) {
	r = &AccountIntelService{}
	r.Options = opts
	r.ASNs = NewAccountIntelASNService(opts...)
	r.DNS = NewAccountIntelDNSService(opts...)
	r.Domains = NewAccountIntelDomainService(opts...)
	r.DomainHistories = NewAccountIntelDomainHistoryService(opts...)
	r.IPs = NewAccountIntelIPService(opts...)
	r.IPLists = NewAccountIntelIPListService(opts...)
	r.Miscategorizations = NewAccountIntelMiscategorizationService(opts...)
	r.Whois = NewAccountIntelWhoisService(opts...)
	return
}
