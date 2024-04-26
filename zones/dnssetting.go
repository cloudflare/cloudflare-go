// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DNSSettingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDNSSettingService] method instead.
type DNSSettingService struct {
	Options []option.RequestOption
}

// NewDNSSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSSettingService(opts ...option.RequestOption) (r *DNSSettingService) {
	r = &DNSSettingService{}
	r.Options = opts
	return
}

// Update DNS settings for a zone
func (r *DNSSettingService) Edit(ctx context.Context, params DNSSettingEditParams, opts ...option.RequestOption) (res *DNSSetting, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show DNS settings for a zone
func (r *DNSSettingService) Get(ctx context.Context, query DNSSettingGetParams, opts ...option.RequestOption) (res *DNSSetting, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSSetting struct {
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers Nameserver `json:"nameservers"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool           `json:"secondary_overrides"`
	JSON               dnsSettingJSON `json:"-"`
}

// dnsSettingJSON contains the JSON metadata for the struct [DNSSetting]
type dnsSettingJSON struct {
	FoundationDNS      apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	SecondaryOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingJSON) RawJSON() string {
	return r.raw
}

type DNSSettingParam struct {
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS param.Field[bool] `json:"foundation_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider param.Field[bool] `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[NameserverParam] `json:"nameservers"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
}

func (r DNSSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings determining the nameservers through which the zone should be available.
type Nameserver struct {
	// Nameserver type
	Type NameserverType `json:"type,required"`
	JSON nameserverJSON `json:"-"`
}

// nameserverJSON contains the JSON metadata for the struct [Nameserver]
type nameserverJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Nameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nameserverJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type NameserverType string

const (
	NameserverTypeCloudflareStandard NameserverType = "cloudflare.standard"
)

func (r NameserverType) IsKnown() bool {
	switch r {
	case NameserverTypeCloudflareStandard:
		return true
	}
	return false
}

// Settings determining the nameservers through which the zone should be available.
type NameserverParam struct {
	// Nameserver type
	Type param.Field[NameserverType] `json:"type,required"`
}

func (r NameserverParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSSettingEditParams struct {
	// Identifier
	ZoneID     param.Field[string] `path:"zone_id,required"`
	DNSSetting DNSSettingParam     `json:"dns_setting,required"`
}

func (r DNSSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSSetting)
}

type DNSSettingEditResponseEnvelope struct {
	Errors             []shared.ResponseInfo `json:"errors,required"`
	FoundationDNS      interface{}           `json:"foundation_dns,required"`
	Messages           []shared.ResponseInfo `json:"messages,required"`
	MultiProvider      interface{}           `json:"multi_provider,required"`
	Nameservers        interface{}           `json:"nameservers,required"`
	SecondaryOverrides interface{}           `json:"secondary_overrides,required"`
	// Whether the API call was successful
	Success DNSSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSSetting                            `json:"result"`
	JSON    dnsSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingEditResponseEnvelope]
type dnsSettingEditResponseEnvelopeJSON struct {
	Errors             apijson.Field
	FoundationDNS      apijson.Field
	Messages           apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	SecondaryOverrides apijson.Field
	Success            apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSSettingEditResponseEnvelopeSuccess bool

const (
	DNSSettingEditResponseEnvelopeSuccessTrue DNSSettingEditResponseEnvelopeSuccess = true
)

func (r DNSSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSSettingGetResponseEnvelope struct {
	Errors             []shared.ResponseInfo `json:"errors,required"`
	FoundationDNS      interface{}           `json:"foundation_dns,required"`
	Messages           []shared.ResponseInfo `json:"messages,required"`
	MultiProvider      interface{}           `json:"multi_provider,required"`
	Nameservers        interface{}           `json:"nameservers,required"`
	SecondaryOverrides interface{}           `json:"secondary_overrides,required"`
	// Whether the API call was successful
	Success DNSSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSSetting                           `json:"result"`
	JSON    dnsSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingGetResponseEnvelope]
type dnsSettingGetResponseEnvelopeJSON struct {
	Errors             apijson.Field
	FoundationDNS      apijson.Field
	Messages           apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	SecondaryOverrides apijson.Field
	Success            apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSSettingGetResponseEnvelopeSuccess bool

const (
	DNSSettingGetResponseEnvelopeSuccessTrue DNSSettingGetResponseEnvelopeSuccess = true
)

func (r DNSSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
