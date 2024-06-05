// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DNSSettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSSettingService] method instead.
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
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	Soa DNSSettingSoa `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode DNSSettingZoneMode `json:"zone_mode"`
	JSON     dnsSettingJSON     `json:"-"`
}

// dnsSettingJSON contains the JSON metadata for the struct [DNSSetting]
type dnsSettingJSON struct {
	FoundationDNS      apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NSTTL              apijson.Field
	SecondaryOverrides apijson.Field
	Soa                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingJSON) RawJSON() string {
	return r.raw
}

// Components of the zone's SOA record.
type DNSSettingSoa struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL float64 `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	Mname string `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	Rname string `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL  float64           `json:"ttl,required"`
	JSON dnsSettingSoaJSON `json:"-"`
}

// dnsSettingSoaJSON contains the JSON metadata for the struct [DNSSettingSoa]
type dnsSettingSoaJSON struct {
	Expire      apijson.Field
	MinTTL      apijson.Field
	Mname       apijson.Field
	Refresh     apijson.Field
	Retry       apijson.Field
	Rname       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingSoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingSoaJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type DNSSettingZoneMode string

const (
	DNSSettingZoneModeStandard DNSSettingZoneMode = "standard"
	DNSSettingZoneModeCDNOnly  DNSSettingZoneMode = "cdn_only"
	DNSSettingZoneModeDNSOnly  DNSSettingZoneMode = "dns_only"
)

func (r DNSSettingZoneMode) IsKnown() bool {
	switch r {
	case DNSSettingZoneModeStandard, DNSSettingZoneModeCDNOnly, DNSSettingZoneModeDNSOnly:
		return true
	}
	return false
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
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL param.Field[float64] `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	Soa param.Field[DNSSettingSoaParam] `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode param.Field[DNSSettingZoneMode] `json:"zone_mode"`
}

func (r DNSSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Components of the zone's SOA record.
type DNSSettingSoaParam struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire param.Field[float64] `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL param.Field[float64] `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	Mname param.Field[string] `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh param.Field[float64] `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry param.Field[float64] `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	Rname param.Field[string] `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL param.Field[float64] `json:"ttl,required"`
}

func (r DNSSettingSoaParam) MarshalJSON() (data []byte, err error) {
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSSetting                            `json:"result"`
	JSON    dnsSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingEditResponseEnvelope]
type dnsSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSSetting                           `json:"result"`
	JSON    dnsSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingGetResponseEnvelope]
type dnsSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
