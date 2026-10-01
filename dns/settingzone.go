// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// SettingZoneService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingZoneService] method instead.
type SettingZoneService struct {
	Options []option.RequestOption
}

// NewSettingZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingZoneService(opts ...option.RequestOption) (r *SettingZoneService) {
	r = &SettingZoneService{}
	r.Options = opts
	return
}

// Update DNS settings for a zone
func (r *SettingZoneService) Edit(ctx context.Context, params SettingZoneEditParams, opts ...option.RequestOption) (res *SettingZoneEditResponse, err error) {
	var env SettingZoneEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/dns_settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Show DNS settings for a zone
func (r *SettingZoneService) Get(ctx context.Context, query SettingZoneGetParams, opts ...option.RequestOption) (res *SettingZoneGetResponse, err error) {
	var env SettingZoneGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/dns_settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SettingZoneEditResponse struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames" api:"required"`
	// Deprecated. Use nameservers.type to configure Advanced Nameservers.
	//
	// Deprecated: foundation_dns is deprecated. Use nameservers.type:
	// cloudflare.advanced to turn on Advanced Nameservers and cloudflare.standard to
	// turn it off. This field will be removed in a future API version.
	FoundationDNS bool `json:"foundation_dns" api:"required"`
	// Settings for this internal zone.
	InternalDNS SettingZoneEditResponseInternalDNS `json:"internal_dns" api:"required"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider" api:"required"`
	// Controls the nameservers through which the zone is available.
	Nameservers SettingZoneEditResponseNameservers `json:"nameservers" api:"required"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl" api:"required"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides" api:"required"`
	// Components of the zone's SOA record.
	SOA SettingZoneEditResponseSOA `json:"soa" api:"required"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingZoneEditResponseZoneMode `json:"zone_mode" api:"required"`
	JSON     settingZoneEditResponseJSON     `json:"-"`
}

// settingZoneEditResponseJSON contains the JSON metadata for the struct
// [SettingZoneEditResponse]
type settingZoneEditResponseJSON struct {
	FlattenAllCNAMEs   apijson.Field
	FoundationDNS      apijson.Field
	InternalDNS        apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NSTTL              apijson.Field
	SecondaryOverrides apijson.Field
	SOA                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingZoneEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type SettingZoneEditResponseInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                                 `json:"reference_zone_id"`
	JSON            settingZoneEditResponseInternalDNSJSON `json:"-"`
}

// settingZoneEditResponseInternalDNSJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseInternalDNS]
type settingZoneEditResponseInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingZoneEditResponseInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseInternalDNSJSON) RawJSON() string {
	return r.raw
}

// Controls the nameservers through which the zone is available.
type SettingZoneEditResponseNameservers struct {
	// Nameserver type.
	Type SettingZoneEditResponseNameserversType `json:"type" api:"required"`
	// Identifier of the account-owned Custom Nameserver Set to use for this zone.
	NameserverSetID string `json:"nameserver_set_id"`
	// Configured nameserver set number to use for this zone.
	NSSet int64                                  `json:"ns_set"`
	JSON  settingZoneEditResponseNameserversJSON `json:"-"`
	union SettingZoneEditResponseNameserversUnion
}

// settingZoneEditResponseNameserversJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseNameservers]
type settingZoneEditResponseNameserversJSON struct {
	Type            apijson.Field
	NameserverSetID apijson.Field
	NSSet           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r settingZoneEditResponseNameserversJSON) RawJSON() string {
	return r.raw
}

func (r *SettingZoneEditResponseNameservers) UnmarshalJSON(data []byte) (err error) {
	*r = SettingZoneEditResponseNameservers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SettingZoneEditResponseNameserversUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflare],
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExisting],
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSet].
func (r SettingZoneEditResponseNameservers) AsUnion() SettingZoneEditResponseNameserversUnion {
	return r.union
}

// Controls the nameservers through which the zone is available.
//
// Union satisfied by
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflare],
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExisting] or
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSet].
type SettingZoneEditResponseNameserversUnion interface {
	implementsSettingZoneEditResponseNameservers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingZoneEditResponseNameserversUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflare{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExisting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSet{}),
		},
	)
}

type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflare struct {
	// Nameserver type.
	Type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareType `json:"type" api:"required"`
	JSON settingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareJSON `json:"-"`
}

// settingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareJSON
// contains the JSON metadata for the struct
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflare]
type settingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareJSON) RawJSON() string {
	return r.raw
}

func (r SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflare) implementsSettingZoneEditResponseNameservers() {
}

// Nameserver type.
type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareType string

const (
	SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareStandard SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareType = "cloudflare.standard"
	SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareAdvanced SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareType = "cloudflare.advanced"
)

func (r SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareType) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareStandard, SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareAdvanced:
		return true
	}
	return false
}

type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExisting struct {
	// Nameserver type.
	Type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingType `json:"type" api:"required"`
	// Configured nameserver set number to use for this zone.
	NSSet int64                                                                          `json:"ns_set"`
	JSON  settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON `json:"-"`
}

// settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON
// contains the JSON metadata for the struct
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExisting]
type settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON struct {
	Type        apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExisting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON) RawJSON() string {
	return r.raw
}

func (r SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExisting) implementsSettingZoneEditResponseNameservers() {
}

// Nameserver type.
type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingType string

const (
	SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomAccount SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.account"
	SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomTenant  SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.tenant"
	SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomZone    SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.zone"
)

func (r SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingType) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomAccount, SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomTenant, SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomZone:
		return true
	}
	return false
}

type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSet struct {
	// Identifier of the account-owned Custom Nameserver Set to use for this zone.
	NameserverSetID string `json:"nameserver_set_id" api:"required"`
	// Nameserver type.
	Type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetType `json:"type" api:"required"`
	JSON settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetJSON `json:"-"`
}

// settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetJSON
// contains the JSON metadata for the struct
// [SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSet]
type settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetJSON struct {
	NameserverSetID apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetJSON) RawJSON() string {
	return r.raw
}

func (r SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSet) implementsSettingZoneEditResponseNameservers() {
}

// Nameserver type.
type SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetType string

const (
	SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetTypeCustom SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetType = "custom"
)

func (r SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetType) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseNameserversDNSSettingsZoneNameserversCustomSetTypeCustom:
		return true
	}
	return false
}

// Nameserver type.
type SettingZoneEditResponseNameserversType string

const (
	SettingZoneEditResponseNameserversTypeCloudflareStandard SettingZoneEditResponseNameserversType = "cloudflare.standard"
	SettingZoneEditResponseNameserversTypeCloudflareAdvanced SettingZoneEditResponseNameserversType = "cloudflare.advanced"
	SettingZoneEditResponseNameserversTypeCustomAccount      SettingZoneEditResponseNameserversType = "custom.account"
	SettingZoneEditResponseNameserversTypeCustomTenant       SettingZoneEditResponseNameserversType = "custom.tenant"
	SettingZoneEditResponseNameserversTypeCustomZone         SettingZoneEditResponseNameserversType = "custom.zone"
	SettingZoneEditResponseNameserversTypeCustom             SettingZoneEditResponseNameserversType = "custom"
)

func (r SettingZoneEditResponseNameserversType) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseNameserversTypeCloudflareStandard, SettingZoneEditResponseNameserversTypeCloudflareAdvanced, SettingZoneEditResponseNameserversTypeCustomAccount, SettingZoneEditResponseNameserversTypeCustomTenant, SettingZoneEditResponseNameserversTypeCustomZone, SettingZoneEditResponseNameserversTypeCustom:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingZoneEditResponseSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL float64 `json:"min_ttl"`
	// The primary nameserver, which may be used for outbound zone transfers. If null,
	// a Cloudflare-assigned value will be used.
	MNAME string `json:"mname" api:"nullable"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME string `json:"rname"`
	// The time to live (TTL) of the SOA record itself.
	TTL  float64                        `json:"ttl"`
	JSON settingZoneEditResponseSOAJSON `json:"-"`
}

// settingZoneEditResponseSOAJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseSOA]
type settingZoneEditResponseSOAJSON struct {
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

func (r *SettingZoneEditResponseSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingZoneEditResponseZoneMode string

const (
	SettingZoneEditResponseZoneModeStandard SettingZoneEditResponseZoneMode = "standard"
	SettingZoneEditResponseZoneModeCDNOnly  SettingZoneEditResponseZoneMode = "cdn_only"
	SettingZoneEditResponseZoneModeDNSOnly  SettingZoneEditResponseZoneMode = "dns_only"
)

func (r SettingZoneEditResponseZoneMode) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseZoneModeStandard, SettingZoneEditResponseZoneModeCDNOnly, SettingZoneEditResponseZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingZoneGetResponse struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames" api:"required"`
	// Deprecated. Use nameservers.type to configure Advanced Nameservers.
	//
	// Deprecated: foundation_dns is deprecated. Use nameservers.type:
	// cloudflare.advanced to turn on Advanced Nameservers and cloudflare.standard to
	// turn it off. This field will be removed in a future API version.
	FoundationDNS bool `json:"foundation_dns" api:"required"`
	// Settings for this internal zone.
	InternalDNS SettingZoneGetResponseInternalDNS `json:"internal_dns" api:"required"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider" api:"required"`
	// Controls the nameservers through which the zone is available.
	Nameservers SettingZoneGetResponseNameservers `json:"nameservers" api:"required"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl" api:"required"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides" api:"required"`
	// Components of the zone's SOA record.
	SOA SettingZoneGetResponseSOA `json:"soa" api:"required"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingZoneGetResponseZoneMode `json:"zone_mode" api:"required"`
	JSON     settingZoneGetResponseJSON     `json:"-"`
}

// settingZoneGetResponseJSON contains the JSON metadata for the struct
// [SettingZoneGetResponse]
type settingZoneGetResponseJSON struct {
	FlattenAllCNAMEs   apijson.Field
	FoundationDNS      apijson.Field
	InternalDNS        apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NSTTL              apijson.Field
	SecondaryOverrides apijson.Field
	SOA                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingZoneGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type SettingZoneGetResponseInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                                `json:"reference_zone_id"`
	JSON            settingZoneGetResponseInternalDNSJSON `json:"-"`
}

// settingZoneGetResponseInternalDNSJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseInternalDNS]
type settingZoneGetResponseInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingZoneGetResponseInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseInternalDNSJSON) RawJSON() string {
	return r.raw
}

// Controls the nameservers through which the zone is available.
type SettingZoneGetResponseNameservers struct {
	// Nameserver type.
	Type SettingZoneGetResponseNameserversType `json:"type" api:"required"`
	// Identifier of the account-owned Custom Nameserver Set to use for this zone.
	NameserverSetID string `json:"nameserver_set_id"`
	// Configured nameserver set number to use for this zone.
	NSSet int64                                 `json:"ns_set"`
	JSON  settingZoneGetResponseNameserversJSON `json:"-"`
	union SettingZoneGetResponseNameserversUnion
}

// settingZoneGetResponseNameserversJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseNameservers]
type settingZoneGetResponseNameserversJSON struct {
	Type            apijson.Field
	NameserverSetID apijson.Field
	NSSet           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r settingZoneGetResponseNameserversJSON) RawJSON() string {
	return r.raw
}

func (r *SettingZoneGetResponseNameservers) UnmarshalJSON(data []byte) (err error) {
	*r = SettingZoneGetResponseNameservers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SettingZoneGetResponseNameserversUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflare],
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExisting],
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSet].
func (r SettingZoneGetResponseNameservers) AsUnion() SettingZoneGetResponseNameserversUnion {
	return r.union
}

// Controls the nameservers through which the zone is available.
//
// Union satisfied by
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflare],
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExisting] or
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSet].
type SettingZoneGetResponseNameserversUnion interface {
	implementsSettingZoneGetResponseNameservers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingZoneGetResponseNameserversUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflare{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExisting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSet{}),
		},
	)
}

type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflare struct {
	// Nameserver type.
	Type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareType `json:"type" api:"required"`
	JSON settingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareJSON `json:"-"`
}

// settingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareJSON
// contains the JSON metadata for the struct
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflare]
type settingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareJSON) RawJSON() string {
	return r.raw
}

func (r SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflare) implementsSettingZoneGetResponseNameservers() {
}

// Nameserver type.
type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareType string

const (
	SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareStandard SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareType = "cloudflare.standard"
	SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareAdvanced SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareType = "cloudflare.advanced"
)

func (r SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareType) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareStandard, SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareAdvanced:
		return true
	}
	return false
}

type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExisting struct {
	// Nameserver type.
	Type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingType `json:"type" api:"required"`
	// Configured nameserver set number to use for this zone.
	NSSet int64                                                                         `json:"ns_set"`
	JSON  settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON `json:"-"`
}

// settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON
// contains the JSON metadata for the struct
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExisting]
type settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON struct {
	Type        apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExisting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingJSON) RawJSON() string {
	return r.raw
}

func (r SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExisting) implementsSettingZoneGetResponseNameservers() {
}

// Nameserver type.
type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingType string

const (
	SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomAccount SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.account"
	SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomTenant  SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.tenant"
	SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomZone    SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.zone"
)

func (r SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingType) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomAccount, SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomTenant, SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomZone:
		return true
	}
	return false
}

type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSet struct {
	// Identifier of the account-owned Custom Nameserver Set to use for this zone.
	NameserverSetID string `json:"nameserver_set_id" api:"required"`
	// Nameserver type.
	Type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetType `json:"type" api:"required"`
	JSON settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetJSON `json:"-"`
}

// settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetJSON
// contains the JSON metadata for the struct
// [SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSet]
type settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetJSON struct {
	NameserverSetID apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetJSON) RawJSON() string {
	return r.raw
}

func (r SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSet) implementsSettingZoneGetResponseNameservers() {
}

// Nameserver type.
type SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetType string

const (
	SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetTypeCustom SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetType = "custom"
)

func (r SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetType) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseNameserversDNSSettingsZoneNameserversCustomSetTypeCustom:
		return true
	}
	return false
}

// Nameserver type.
type SettingZoneGetResponseNameserversType string

const (
	SettingZoneGetResponseNameserversTypeCloudflareStandard SettingZoneGetResponseNameserversType = "cloudflare.standard"
	SettingZoneGetResponseNameserversTypeCloudflareAdvanced SettingZoneGetResponseNameserversType = "cloudflare.advanced"
	SettingZoneGetResponseNameserversTypeCustomAccount      SettingZoneGetResponseNameserversType = "custom.account"
	SettingZoneGetResponseNameserversTypeCustomTenant       SettingZoneGetResponseNameserversType = "custom.tenant"
	SettingZoneGetResponseNameserversTypeCustomZone         SettingZoneGetResponseNameserversType = "custom.zone"
	SettingZoneGetResponseNameserversTypeCustom             SettingZoneGetResponseNameserversType = "custom"
)

func (r SettingZoneGetResponseNameserversType) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseNameserversTypeCloudflareStandard, SettingZoneGetResponseNameserversTypeCloudflareAdvanced, SettingZoneGetResponseNameserversTypeCustomAccount, SettingZoneGetResponseNameserversTypeCustomTenant, SettingZoneGetResponseNameserversTypeCustomZone, SettingZoneGetResponseNameserversTypeCustom:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingZoneGetResponseSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL float64 `json:"min_ttl"`
	// The primary nameserver, which may be used for outbound zone transfers. If null,
	// a Cloudflare-assigned value will be used.
	MNAME string `json:"mname" api:"nullable"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME string `json:"rname"`
	// The time to live (TTL) of the SOA record itself.
	TTL  float64                       `json:"ttl"`
	JSON settingZoneGetResponseSOAJSON `json:"-"`
}

// settingZoneGetResponseSOAJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseSOA]
type settingZoneGetResponseSOAJSON struct {
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

func (r *SettingZoneGetResponseSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingZoneGetResponseZoneMode string

const (
	SettingZoneGetResponseZoneModeStandard SettingZoneGetResponseZoneMode = "standard"
	SettingZoneGetResponseZoneModeCDNOnly  SettingZoneGetResponseZoneMode = "cdn_only"
	SettingZoneGetResponseZoneModeDNSOnly  SettingZoneGetResponseZoneMode = "dns_only"
)

func (r SettingZoneGetResponseZoneMode) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseZoneModeStandard, SettingZoneGetResponseZoneModeCDNOnly, SettingZoneGetResponseZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingZoneEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs param.Field[bool] `json:"flatten_all_cnames"`
	// Deprecated. Use nameservers.type to configure Advanced Nameservers.
	FoundationDNS param.Field[bool] `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS param.Field[SettingZoneEditParamsInternalDNS] `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider param.Field[bool] `json:"multi_provider"`
	// Controls the nameservers through which the zone is available.
	Nameservers param.Field[SettingZoneEditParamsNameserversUnion] `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL param.Field[float64] `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA param.Field[SettingZoneEditParamsSOA] `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode param.Field[SettingZoneEditParamsZoneMode] `json:"zone_mode"`
}

func (r SettingZoneEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for this internal zone.
type SettingZoneEditParamsInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID param.Field[string] `json:"reference_zone_id"`
}

func (r SettingZoneEditParamsInternalDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls the nameservers through which the zone is available.
type SettingZoneEditParamsNameservers struct {
	// Nameserver type.
	Type param.Field[SettingZoneEditParamsNameserversType] `json:"type" api:"required"`
	// Identifier of the account-owned Custom Nameserver Set to use for this zone.
	NameserverSetID param.Field[string] `json:"nameserver_set_id"`
	// Configured nameserver set number to use for this zone.
	NSSet param.Field[int64] `json:"ns_set"`
}

func (r SettingZoneEditParamsNameservers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingZoneEditParamsNameservers) implementsSettingZoneEditParamsNameserversUnion() {}

// Controls the nameservers through which the zone is available.
//
// Satisfied by
// [dns.SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflare],
// [dns.SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExisting],
// [dns.SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSet],
// [SettingZoneEditParamsNameservers].
type SettingZoneEditParamsNameserversUnion interface {
	implementsSettingZoneEditParamsNameserversUnion()
}

type SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflare struct {
	// Nameserver type.
	Type param.Field[SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareType] `json:"type" api:"required"`
}

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflare) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflare) implementsSettingZoneEditParamsNameserversUnion() {
}

// Nameserver type.
type SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareType string

const (
	SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareStandard SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareType = "cloudflare.standard"
	SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareAdvanced SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareType = "cloudflare.advanced"
)

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareType) IsKnown() bool {
	switch r {
	case SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareStandard, SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCloudflareTypeCloudflareAdvanced:
		return true
	}
	return false
}

type SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExisting struct {
	// Nameserver type.
	Type param.Field[SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingType] `json:"type" api:"required"`
	// Configured nameserver set number to use for this zone.
	NSSet param.Field[int64] `json:"ns_set"`
}

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExisting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExisting) implementsSettingZoneEditParamsNameserversUnion() {
}

// Nameserver type.
type SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingType string

const (
	SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomAccount SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.account"
	SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomTenant  SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.tenant"
	SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomZone    SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingType = "custom.zone"
)

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingType) IsKnown() bool {
	switch r {
	case SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomAccount, SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomTenant, SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomExistingTypeCustomZone:
		return true
	}
	return false
}

type SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSet struct {
	// Identifier of the account-owned Custom Nameserver Set to use for this zone.
	NameserverSetID param.Field[string] `json:"nameserver_set_id" api:"required"`
	// Nameserver type.
	Type param.Field[SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSetType] `json:"type" api:"required"`
}

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSet) implementsSettingZoneEditParamsNameserversUnion() {
}

// Nameserver type.
type SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSetType string

const (
	SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSetTypeCustom SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSetType = "custom"
)

func (r SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSetType) IsKnown() bool {
	switch r {
	case SettingZoneEditParamsNameserversDNSSettingsZoneNameserversCustomSetTypeCustom:
		return true
	}
	return false
}

// Nameserver type.
type SettingZoneEditParamsNameserversType string

const (
	SettingZoneEditParamsNameserversTypeCloudflareStandard SettingZoneEditParamsNameserversType = "cloudflare.standard"
	SettingZoneEditParamsNameserversTypeCloudflareAdvanced SettingZoneEditParamsNameserversType = "cloudflare.advanced"
	SettingZoneEditParamsNameserversTypeCustomAccount      SettingZoneEditParamsNameserversType = "custom.account"
	SettingZoneEditParamsNameserversTypeCustomTenant       SettingZoneEditParamsNameserversType = "custom.tenant"
	SettingZoneEditParamsNameserversTypeCustomZone         SettingZoneEditParamsNameserversType = "custom.zone"
	SettingZoneEditParamsNameserversTypeCustom             SettingZoneEditParamsNameserversType = "custom"
)

func (r SettingZoneEditParamsNameserversType) IsKnown() bool {
	switch r {
	case SettingZoneEditParamsNameserversTypeCloudflareStandard, SettingZoneEditParamsNameserversTypeCloudflareAdvanced, SettingZoneEditParamsNameserversTypeCustomAccount, SettingZoneEditParamsNameserversTypeCustomTenant, SettingZoneEditParamsNameserversTypeCustomZone, SettingZoneEditParamsNameserversTypeCustom:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingZoneEditParamsSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire param.Field[float64] `json:"expire"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL param.Field[float64] `json:"min_ttl"`
	// The primary nameserver, which may be used for outbound zone transfers. If null,
	// a Cloudflare-assigned value will be used.
	MNAME param.Field[string] `json:"mname"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh param.Field[float64] `json:"refresh"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry param.Field[float64] `json:"retry"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME param.Field[string] `json:"rname"`
	// The time to live (TTL) of the SOA record itself.
	TTL param.Field[float64] `json:"ttl"`
}

func (r SettingZoneEditParamsSOA) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingZoneEditParamsZoneMode string

const (
	SettingZoneEditParamsZoneModeStandard SettingZoneEditParamsZoneMode = "standard"
	SettingZoneEditParamsZoneModeCDNOnly  SettingZoneEditParamsZoneMode = "cdn_only"
	SettingZoneEditParamsZoneModeDNSOnly  SettingZoneEditParamsZoneMode = "dns_only"
)

func (r SettingZoneEditParamsZoneMode) IsKnown() bool {
	switch r {
	case SettingZoneEditParamsZoneModeStandard, SettingZoneEditParamsZoneModeCDNOnly, SettingZoneEditParamsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingZoneEditResponseEnvelope struct {
	Errors   []SettingZoneEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingZoneEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingZoneEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingZoneEditResponse                `json:"result"`
	JSON    settingZoneEditResponseEnvelopeJSON    `json:"-"`
}

// settingZoneEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseEnvelope]
type settingZoneEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           SettingZoneEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingZoneEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingZoneEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZoneEditResponseEnvelopeErrors]
type settingZoneEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    settingZoneEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingZoneEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingZoneEditResponseEnvelopeErrorsSource]
type settingZoneEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           SettingZoneEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingZoneEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingZoneEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingZoneEditResponseEnvelopeMessages]
type settingZoneEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    settingZoneEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingZoneEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SettingZoneEditResponseEnvelopeMessagesSource]
type settingZoneEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingZoneEditResponseEnvelopeSuccess bool

const (
	SettingZoneEditResponseEnvelopeSuccessTrue SettingZoneEditResponseEnvelopeSuccess = true
)

func (r SettingZoneEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingZoneGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SettingZoneGetResponseEnvelope struct {
	Errors   []SettingZoneGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingZoneGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingZoneGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingZoneGetResponse                `json:"result"`
	JSON    settingZoneGetResponseEnvelopeJSON    `json:"-"`
}

// settingZoneGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseEnvelope]
type settingZoneGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SettingZoneGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingZoneGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingZoneGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZoneGetResponseEnvelopeErrors]
type settingZoneGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    settingZoneGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingZoneGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingZoneGetResponseEnvelopeErrorsSource]
type settingZoneGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SettingZoneGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingZoneGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingZoneGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingZoneGetResponseEnvelopeMessages]
type settingZoneGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    settingZoneGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingZoneGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SettingZoneGetResponseEnvelopeMessagesSource]
type settingZoneGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingZoneGetResponseEnvelopeSuccess bool

const (
	SettingZoneGetResponseEnvelopeSuccessTrue SettingZoneGetResponseEnvelopeSuccess = true
)

func (r SettingZoneGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
