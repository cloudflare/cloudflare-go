// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
	Views   *SettingViewService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	r.Views = NewSettingViewService(opts...)
	return
}

// Update DNS settings for an account or zone
func (r *SettingService) Edit(ctx context.Context, params SettingEditParams, opts ...option.RequestOption) (res *SettingEditResponse, err error) {
	var env SettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/dns_settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show DNS settings for an account or zone
func (r *SettingService) Get(ctx context.Context, query SettingGetParams, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	var env SettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/dns_settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingEditResponse struct {
	ZoneDefaults SettingEditResponseZoneDefaults `json:"zone_defaults"`
	JSON         settingEditResponseJSON         `json:"-"`
}

// settingEditResponseJSON contains the JSON metadata for the struct
// [SettingEditResponse]
type settingEditResponseJSON struct {
	ZoneDefaults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingEditResponseZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers SettingEditResponseZoneDefaultsNameservers `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA SettingEditResponseZoneDefaultsSOA `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingEditResponseZoneDefaultsZoneMode `json:"zone_mode"`
	JSON     settingEditResponseZoneDefaultsJSON     `json:"-"`
}

// settingEditResponseZoneDefaultsJSON contains the JSON metadata for the struct
// [SettingEditResponseZoneDefaults]
type settingEditResponseZoneDefaultsJSON struct {
	FlattenAllCNAMEs   apijson.Field
	FoundationDNS      apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NSTTL              apijson.Field
	SecondaryOverrides apijson.Field
	SOA                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingEditResponseZoneDefaults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZoneDefaultsJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type SettingEditResponseZoneDefaultsNameservers struct {
	// Nameserver type
	Type SettingEditResponseZoneDefaultsNameserversType `json:"type,required"`
	JSON settingEditResponseZoneDefaultsNameserversJSON `json:"-"`
}

// settingEditResponseZoneDefaultsNameserversJSON contains the JSON metadata for
// the struct [SettingEditResponseZoneDefaultsNameservers]
type settingEditResponseZoneDefaultsNameserversJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZoneDefaultsNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZoneDefaultsNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type SettingEditResponseZoneDefaultsNameserversType string

const (
	SettingEditResponseZoneDefaultsNameserversTypeCloudflareStandard       SettingEditResponseZoneDefaultsNameserversType = "cloudflare.standard"
	SettingEditResponseZoneDefaultsNameserversTypeCloudflareStandardRandom SettingEditResponseZoneDefaultsNameserversType = "cloudflare.standard.random"
	SettingEditResponseZoneDefaultsNameserversTypeCustomAccount            SettingEditResponseZoneDefaultsNameserversType = "custom.account"
	SettingEditResponseZoneDefaultsNameserversTypeCustomTenant             SettingEditResponseZoneDefaultsNameserversType = "custom.tenant"
)

func (r SettingEditResponseZoneDefaultsNameserversType) IsKnown() bool {
	switch r {
	case SettingEditResponseZoneDefaultsNameserversTypeCloudflareStandard, SettingEditResponseZoneDefaultsNameserversTypeCloudflareStandardRandom, SettingEditResponseZoneDefaultsNameserversTypeCustomAccount, SettingEditResponseZoneDefaultsNameserversTypeCustomTenant:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingEditResponseZoneDefaultsSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL float64 `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	MNAME string `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME string `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL  float64                                `json:"ttl,required"`
	JSON settingEditResponseZoneDefaultsSOAJSON `json:"-"`
}

// settingEditResponseZoneDefaultsSOAJSON contains the JSON metadata for the struct
// [SettingEditResponseZoneDefaultsSOA]
type settingEditResponseZoneDefaultsSOAJSON struct {
	Expire      apijson.Field
	MinTTL      apijson.Field
	MNAME       apijson.Field
	Refresh     apijson.Field
	Retry       apijson.Field
	RNAME       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZoneDefaultsSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZoneDefaultsSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingEditResponseZoneDefaultsZoneMode string

const (
	SettingEditResponseZoneDefaultsZoneModeStandard SettingEditResponseZoneDefaultsZoneMode = "standard"
	SettingEditResponseZoneDefaultsZoneModeCDNOnly  SettingEditResponseZoneDefaultsZoneMode = "cdn_only"
	SettingEditResponseZoneDefaultsZoneModeDNSOnly  SettingEditResponseZoneDefaultsZoneMode = "dns_only"
)

func (r SettingEditResponseZoneDefaultsZoneMode) IsKnown() bool {
	switch r {
	case SettingEditResponseZoneDefaultsZoneModeStandard, SettingEditResponseZoneDefaultsZoneModeCDNOnly, SettingEditResponseZoneDefaultsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingGetResponse struct {
	ZoneDefaults SettingGetResponseZoneDefaults `json:"zone_defaults"`
	JSON         settingGetResponseJSON         `json:"-"`
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	ZoneDefaults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponseZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers SettingGetResponseZoneDefaultsNameservers `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA SettingGetResponseZoneDefaultsSOA `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingGetResponseZoneDefaultsZoneMode `json:"zone_mode"`
	JSON     settingGetResponseZoneDefaultsJSON     `json:"-"`
}

// settingGetResponseZoneDefaultsJSON contains the JSON metadata for the struct
// [SettingGetResponseZoneDefaults]
type settingGetResponseZoneDefaultsJSON struct {
	FlattenAllCNAMEs   apijson.Field
	FoundationDNS      apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NSTTL              apijson.Field
	SecondaryOverrides apijson.Field
	SOA                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingGetResponseZoneDefaults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZoneDefaultsJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type SettingGetResponseZoneDefaultsNameservers struct {
	// Nameserver type
	Type SettingGetResponseZoneDefaultsNameserversType `json:"type,required"`
	JSON settingGetResponseZoneDefaultsNameserversJSON `json:"-"`
}

// settingGetResponseZoneDefaultsNameserversJSON contains the JSON metadata for the
// struct [SettingGetResponseZoneDefaultsNameservers]
type settingGetResponseZoneDefaultsNameserversJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZoneDefaultsNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZoneDefaultsNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type SettingGetResponseZoneDefaultsNameserversType string

const (
	SettingGetResponseZoneDefaultsNameserversTypeCloudflareStandard       SettingGetResponseZoneDefaultsNameserversType = "cloudflare.standard"
	SettingGetResponseZoneDefaultsNameserversTypeCloudflareStandardRandom SettingGetResponseZoneDefaultsNameserversType = "cloudflare.standard.random"
	SettingGetResponseZoneDefaultsNameserversTypeCustomAccount            SettingGetResponseZoneDefaultsNameserversType = "custom.account"
	SettingGetResponseZoneDefaultsNameserversTypeCustomTenant             SettingGetResponseZoneDefaultsNameserversType = "custom.tenant"
)

func (r SettingGetResponseZoneDefaultsNameserversType) IsKnown() bool {
	switch r {
	case SettingGetResponseZoneDefaultsNameserversTypeCloudflareStandard, SettingGetResponseZoneDefaultsNameserversTypeCloudflareStandardRandom, SettingGetResponseZoneDefaultsNameserversTypeCustomAccount, SettingGetResponseZoneDefaultsNameserversTypeCustomTenant:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingGetResponseZoneDefaultsSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL float64 `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	MNAME string `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME string `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL  float64                               `json:"ttl,required"`
	JSON settingGetResponseZoneDefaultsSOAJSON `json:"-"`
}

// settingGetResponseZoneDefaultsSOAJSON contains the JSON metadata for the struct
// [SettingGetResponseZoneDefaultsSOA]
type settingGetResponseZoneDefaultsSOAJSON struct {
	Expire      apijson.Field
	MinTTL      apijson.Field
	MNAME       apijson.Field
	Refresh     apijson.Field
	Retry       apijson.Field
	RNAME       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZoneDefaultsSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZoneDefaultsSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingGetResponseZoneDefaultsZoneMode string

const (
	SettingGetResponseZoneDefaultsZoneModeStandard SettingGetResponseZoneDefaultsZoneMode = "standard"
	SettingGetResponseZoneDefaultsZoneModeCDNOnly  SettingGetResponseZoneDefaultsZoneMode = "cdn_only"
	SettingGetResponseZoneDefaultsZoneModeDNSOnly  SettingGetResponseZoneDefaultsZoneMode = "dns_only"
)

func (r SettingGetResponseZoneDefaultsZoneMode) IsKnown() bool {
	switch r {
	case SettingGetResponseZoneDefaultsZoneModeStandard, SettingGetResponseZoneDefaultsZoneModeCDNOnly, SettingGetResponseZoneDefaultsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingEditParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID       param.Field[string]                        `path:"zone_id"`
	ZoneDefaults param.Field[SettingEditParamsZoneDefaults] `json:"zone_defaults"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingEditParamsZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs param.Field[bool] `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS param.Field[bool] `json:"foundation_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider param.Field[bool] `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[SettingEditParamsZoneDefaultsNameservers] `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL param.Field[float64] `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA param.Field[SettingEditParamsZoneDefaultsSOA] `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode param.Field[SettingEditParamsZoneDefaultsZoneMode] `json:"zone_mode"`
}

func (r SettingEditParamsZoneDefaults) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings determining the nameservers through which the zone should be available.
type SettingEditParamsZoneDefaultsNameservers struct {
	// Nameserver type
	Type param.Field[SettingEditParamsZoneDefaultsNameserversType] `json:"type,required"`
}

func (r SettingEditParamsZoneDefaultsNameservers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nameserver type
type SettingEditParamsZoneDefaultsNameserversType string

const (
	SettingEditParamsZoneDefaultsNameserversTypeCloudflareStandard       SettingEditParamsZoneDefaultsNameserversType = "cloudflare.standard"
	SettingEditParamsZoneDefaultsNameserversTypeCloudflareStandardRandom SettingEditParamsZoneDefaultsNameserversType = "cloudflare.standard.random"
	SettingEditParamsZoneDefaultsNameserversTypeCustomAccount            SettingEditParamsZoneDefaultsNameserversType = "custom.account"
	SettingEditParamsZoneDefaultsNameserversTypeCustomTenant             SettingEditParamsZoneDefaultsNameserversType = "custom.tenant"
)

func (r SettingEditParamsZoneDefaultsNameserversType) IsKnown() bool {
	switch r {
	case SettingEditParamsZoneDefaultsNameserversTypeCloudflareStandard, SettingEditParamsZoneDefaultsNameserversTypeCloudflareStandardRandom, SettingEditParamsZoneDefaultsNameserversTypeCustomAccount, SettingEditParamsZoneDefaultsNameserversTypeCustomTenant:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingEditParamsZoneDefaultsSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire param.Field[float64] `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL param.Field[float64] `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	MNAME param.Field[string] `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh param.Field[float64] `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry param.Field[float64] `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME param.Field[string] `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL param.Field[float64] `json:"ttl,required"`
}

func (r SettingEditParamsZoneDefaultsSOA) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingEditParamsZoneDefaultsZoneMode string

const (
	SettingEditParamsZoneDefaultsZoneModeStandard SettingEditParamsZoneDefaultsZoneMode = "standard"
	SettingEditParamsZoneDefaultsZoneModeCDNOnly  SettingEditParamsZoneDefaultsZoneMode = "cdn_only"
	SettingEditParamsZoneDefaultsZoneModeDNSOnly  SettingEditParamsZoneDefaultsZoneMode = "dns_only"
)

func (r SettingEditParamsZoneDefaultsZoneMode) IsKnown() bool {
	switch r {
	case SettingEditParamsZoneDefaultsZoneModeStandard, SettingEditParamsZoneDefaultsZoneModeCDNOnly, SettingEditParamsZoneDefaultsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingEditResponse                `json:"result"`
	JSON    settingEditResponseEnvelopeJSON    `json:"-"`
}

// settingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingEditResponseEnvelope]
type settingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingEditResponseEnvelopeSuccess bool

const (
	SettingEditResponseEnvelopeSuccessTrue SettingEditResponseEnvelopeSuccess = true
)

func (r SettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type SettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingGetResponse                `json:"result"`
	JSON    settingGetResponseEnvelopeJSON    `json:"-"`
}

// settingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelope]
type settingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingGetResponseEnvelopeSuccess bool

const (
	SettingGetResponseEnvelopeSuccessTrue SettingGetResponseEnvelopeSuccess = true
)

func (r SettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
