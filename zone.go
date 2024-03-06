// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewZoneService] method instead.
type ZoneService struct {
	Options           []option.RequestOption
	ActivationCheck   *ZoneActivationCheckService
	Settings          *ZoneSettingService
	CustomNameservers *ZoneCustomNameserverService
	Holds             *ZoneHoldService
	Workers           *ZoneWorkerService
	Subscriptions     *ZoneSubscriptionService
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r *ZoneService) {
	r = &ZoneService{}
	r.Options = opts
	r.ActivationCheck = NewZoneActivationCheckService(opts...)
	r.Settings = NewZoneSettingService(opts...)
	r.CustomNameservers = NewZoneCustomNameserverService(opts...)
	r.Holds = NewZoneHoldService(opts...)
	r.Workers = NewZoneWorkerService(opts...)
	r.Subscriptions = NewZoneSubscriptionService(opts...)
	return
}

// Create Zone
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *ZoneNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneNewResponseEnvelope
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists, searches, sorts, and filters your zones.
func (r *ZoneService) List(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[ZoneListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "zones"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists, searches, sorts, and filters your zones.
func (r *ZoneService) ListAutoPaging(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[ZoneListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an existing zone.
func (r *ZoneService) Delete(ctx context.Context, body ZoneDeleteParams, opts ...option.RequestOption) (res *ZoneDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edits a zone. Only one zone property can be changed at a time.
func (r *ZoneService) Edit(ctx context.Context, params ZoneEditParams, opts ...option.RequestOption) (res *ZoneEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneEditResponseEnvelope
	path := fmt.Sprintf("zones/%s", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zone Details
func (r *ZoneService) Get(ctx context.Context, query ZoneGetParams, opts ...option.RequestOption) (res *ZoneGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneGetResponseEnvelope
	path := fmt.Sprintf("zones/%s", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to
	Account ZoneNewResponseAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone
	Meta ZoneNewResponseMeta `json:"meta,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone
	Owner ZoneNewResponseOwner `json:"owner,required"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string            `json:"vanity_name_servers" format:"hostname"`
	JSON              zoneNewResponseJSON `json:"-"`
}

// zoneNewResponseJSON contains the JSON metadata for the struct [ZoneNewResponse]
type zoneNewResponseJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	VanityNameServers   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The account the zone belongs to
type ZoneNewResponseAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account
	Name string                     `json:"name"`
	JSON zoneNewResponseAccountJSON `json:"-"`
}

// zoneNewResponseAccountJSON contains the JSON metadata for the struct
// [ZoneNewResponseAccount]
type zoneNewResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the zone
type ZoneNewResponseMeta struct {
	// The zone is only configured for CDN
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing
	PhishingDetected bool                    `json:"phishing_detected"`
	Step             int64                   `json:"step"`
	JSON             zoneNewResponseMetaJSON `json:"-"`
}

// zoneNewResponseMetaJSON contains the JSON metadata for the struct
// [ZoneNewResponseMeta]
type zoneNewResponseMetaJSON struct {
	CdnOnly                apijson.Field
	CustomCertificateQuota apijson.Field
	DNSOnly                apijson.Field
	FoundationDNS          apijson.Field
	PageRuleQuota          apijson.Field
	PhishingDetected       apijson.Field
	Step                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneNewResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The owner of the zone
type ZoneNewResponseOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner
	Name string `json:"name"`
	// The type of owner
	Type string                   `json:"type"`
	JSON zoneNewResponseOwnerJSON `json:"-"`
}

// zoneNewResponseOwnerJSON contains the JSON metadata for the struct
// [ZoneNewResponseOwner]
type zoneNewResponseOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to
	Account ZoneListResponseAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone
	Meta ZoneListResponseMeta `json:"meta,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone
	Owner ZoneListResponseOwner `json:"owner,required"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string             `json:"vanity_name_servers" format:"hostname"`
	JSON              zoneListResponseJSON `json:"-"`
}

// zoneListResponseJSON contains the JSON metadata for the struct
// [ZoneListResponse]
type zoneListResponseJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	VanityNameServers   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The account the zone belongs to
type ZoneListResponseAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account
	Name string                      `json:"name"`
	JSON zoneListResponseAccountJSON `json:"-"`
}

// zoneListResponseAccountJSON contains the JSON metadata for the struct
// [ZoneListResponseAccount]
type zoneListResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the zone
type ZoneListResponseMeta struct {
	// The zone is only configured for CDN
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing
	PhishingDetected bool                     `json:"phishing_detected"`
	Step             int64                    `json:"step"`
	JSON             zoneListResponseMetaJSON `json:"-"`
}

// zoneListResponseMetaJSON contains the JSON metadata for the struct
// [ZoneListResponseMeta]
type zoneListResponseMetaJSON struct {
	CdnOnly                apijson.Field
	CustomCertificateQuota apijson.Field
	DNSOnly                apijson.Field
	FoundationDNS          apijson.Field
	PageRuleQuota          apijson.Field
	PhishingDetected       apijson.Field
	Step                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The owner of the zone
type ZoneListResponseOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner
	Name string `json:"name"`
	// The type of owner
	Type string                    `json:"type"`
	JSON zoneListResponseOwnerJSON `json:"-"`
}

// zoneListResponseOwnerJSON contains the JSON metadata for the struct
// [ZoneListResponseOwner]
type zoneListResponseOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponseOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDeleteResponse struct {
	// Identifier
	ID   string                 `json:"id,required"`
	JSON zoneDeleteResponseJSON `json:"-"`
}

// zoneDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneDeleteResponse]
type zoneDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to
	Account ZoneEditResponseAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone
	Meta ZoneEditResponseMeta `json:"meta,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone
	Owner ZoneEditResponseOwner `json:"owner,required"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string             `json:"vanity_name_servers" format:"hostname"`
	JSON              zoneEditResponseJSON `json:"-"`
}

// zoneEditResponseJSON contains the JSON metadata for the struct
// [ZoneEditResponse]
type zoneEditResponseJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	VanityNameServers   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The account the zone belongs to
type ZoneEditResponseAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account
	Name string                      `json:"name"`
	JSON zoneEditResponseAccountJSON `json:"-"`
}

// zoneEditResponseAccountJSON contains the JSON metadata for the struct
// [ZoneEditResponseAccount]
type zoneEditResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the zone
type ZoneEditResponseMeta struct {
	// The zone is only configured for CDN
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing
	PhishingDetected bool                     `json:"phishing_detected"`
	Step             int64                    `json:"step"`
	JSON             zoneEditResponseMetaJSON `json:"-"`
}

// zoneEditResponseMetaJSON contains the JSON metadata for the struct
// [ZoneEditResponseMeta]
type zoneEditResponseMetaJSON struct {
	CdnOnly                apijson.Field
	CustomCertificateQuota apijson.Field
	DNSOnly                apijson.Field
	FoundationDNS          apijson.Field
	PageRuleQuota          apijson.Field
	PhishingDetected       apijson.Field
	Step                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneEditResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The owner of the zone
type ZoneEditResponseOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner
	Name string `json:"name"`
	// The type of owner
	Type string                    `json:"type"`
	JSON zoneEditResponseOwnerJSON `json:"-"`
}

// zoneEditResponseOwnerJSON contains the JSON metadata for the struct
// [ZoneEditResponseOwner]
type zoneEditResponseOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to
	Account ZoneGetResponseAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone
	Meta ZoneGetResponseMeta `json:"meta,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone
	Owner ZoneGetResponseOwner `json:"owner,required"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string            `json:"vanity_name_servers" format:"hostname"`
	JSON              zoneGetResponseJSON `json:"-"`
}

// zoneGetResponseJSON contains the JSON metadata for the struct [ZoneGetResponse]
type zoneGetResponseJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	VanityNameServers   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The account the zone belongs to
type ZoneGetResponseAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account
	Name string                     `json:"name"`
	JSON zoneGetResponseAccountJSON `json:"-"`
}

// zoneGetResponseAccountJSON contains the JSON metadata for the struct
// [ZoneGetResponseAccount]
type zoneGetResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the zone
type ZoneGetResponseMeta struct {
	// The zone is only configured for CDN
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing
	PhishingDetected bool                    `json:"phishing_detected"`
	Step             int64                   `json:"step"`
	JSON             zoneGetResponseMetaJSON `json:"-"`
}

// zoneGetResponseMetaJSON contains the JSON metadata for the struct
// [ZoneGetResponseMeta]
type zoneGetResponseMetaJSON struct {
	CdnOnly                apijson.Field
	CustomCertificateQuota apijson.Field
	DNSOnly                apijson.Field
	FoundationDNS          apijson.Field
	PageRuleQuota          apijson.Field
	PhishingDetected       apijson.Field
	Step                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The owner of the zone
type ZoneGetResponseOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner
	Name string `json:"name"`
	// The type of owner
	Type string                   `json:"type"`
	JSON zoneGetResponseOwnerJSON `json:"-"`
}

// zoneGetResponseOwnerJSON contains the JSON metadata for the struct
// [ZoneGetResponseOwner]
type zoneGetResponseOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewParams struct {
	Account param.Field[ZoneNewParamsAccount] `json:"account,required"`
	// The domain name
	Name param.Field[string] `json:"name,required"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type param.Field[ZoneNewParamsType] `json:"type"`
}

func (r ZoneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneNewParamsAccount struct {
	// Identifier
	ID param.Field[string] `json:"id"`
}

func (r ZoneNewParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup.
type ZoneNewParamsType string

const (
	ZoneNewParamsTypeFull      ZoneNewParamsType = "full"
	ZoneNewParamsTypePartial   ZoneNewParamsType = "partial"
	ZoneNewParamsTypeSecondary ZoneNewParamsType = "secondary"
)

type ZoneNewResponseEnvelope struct {
	Errors   []ZoneNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                        `json:"success,required"`
	Result  ZoneNewResponse             `json:"result"`
	JSON    zoneNewResponseEnvelopeJSON `json:"-"`
}

// zoneNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneNewResponseEnvelope]
type zoneNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZoneNewResponseEnvelopeErrors]
type zoneNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ZoneNewResponseEnvelopeMessages]
type zoneNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListParams struct {
	Account param.Field[ZoneListParamsAccount] `query:"account"`
	// Direction to order zones.
	Direction param.Field[ZoneListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[ZoneListParamsMatch] `query:"match"`
	// A domain name. Optional filter operators can be provided to extend refine the
	// search:
	//
	// - `equal` (default)
	// - `not_equal`
	// - `starts_with`
	// - `ends_with`
	// - `contains`
	// - `starts_with_case_sensitive`
	// - `ends_with_case_sensitive`
	// - `contains_case_sensitive`
	Name param.Field[string] `query:"name"`
	// Field to order zones by.
	Order param.Field[ZoneListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A zone status
	Status param.Field[ZoneListParamsStatus] `query:"status"`
}

// URLQuery serializes [ZoneListParams]'s query parameters as `url.Values`.
func (r ZoneListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneListParamsAccount struct {
	// An account ID
	ID param.Field[string] `query:"id"`
	// An account Name. Optional filter operators can be provided to extend refine the
	// search:
	//
	// - `equal` (default)
	// - `not_equal`
	// - `starts_with`
	// - `ends_with`
	// - `contains`
	// - `starts_with_case_sensitive`
	// - `ends_with_case_sensitive`
	// - `contains_case_sensitive`
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [ZoneListParamsAccount]'s query parameters as `url.Values`.
func (r ZoneListParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order zones.
type ZoneListParamsDirection string

const (
	ZoneListParamsDirectionAsc  ZoneListParamsDirection = "asc"
	ZoneListParamsDirectionDesc ZoneListParamsDirection = "desc"
)

// Whether to match all search requirements or at least one (any).
type ZoneListParamsMatch string

const (
	ZoneListParamsMatchAny ZoneListParamsMatch = "any"
	ZoneListParamsMatchAll ZoneListParamsMatch = "all"
)

// Field to order zones by.
type ZoneListParamsOrder string

const (
	ZoneListParamsOrderName        ZoneListParamsOrder = "name"
	ZoneListParamsOrderStatus      ZoneListParamsOrder = "status"
	ZoneListParamsOrderAccountID   ZoneListParamsOrder = "account.id"
	ZoneListParamsOrderAccountName ZoneListParamsOrder = "account.name"
)

// A zone status
type ZoneListParamsStatus string

const (
	ZoneListParamsStatusInitializing ZoneListParamsStatus = "initializing"
	ZoneListParamsStatusPending      ZoneListParamsStatus = "pending"
	ZoneListParamsStatusActive       ZoneListParamsStatus = "active"
	ZoneListParamsStatusMoved        ZoneListParamsStatus = "moved"
)

type ZoneDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneDeleteResponseEnvelope struct {
	Errors   []ZoneDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                           `json:"success,required"`
	Result  ZoneDeleteResponse             `json:"result,nullable"`
	JSON    zoneDeleteResponseEnvelopeJSON `json:"-"`
}

// zoneDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseEnvelope]
type zoneDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseEnvelopeErrors]
type zoneDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseEnvelopeMessages]
type zoneDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// (Deprecated) Please use the `/zones/{zone_id}/subscription` API to update a
	// zone's plan. Changing this value will create/cancel associated subscriptions. To
	// view available plans for this zone, see Zone Plans.
	Plan param.Field[ZoneEditParamsPlan] `json:"plan"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup. This parameter is only
	// available to Enterprise customers or if it has been explicitly enabled on a
	// zone.
	Type param.Field[ZoneEditParamsType] `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers param.Field[[]string] `json:"vanity_name_servers" format:"hostname"`
}

func (r ZoneEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// (Deprecated) Please use the `/zones/{zone_id}/subscription` API to update a
// zone's plan. Changing this value will create/cancel associated subscriptions. To
// view available plans for this zone, see Zone Plans.
type ZoneEditParamsPlan struct {
	// Identifier
	ID param.Field[string] `json:"id"`
}

func (r ZoneEditParamsPlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup. This parameter is only
// available to Enterprise customers or if it has been explicitly enabled on a
// zone.
type ZoneEditParamsType string

const (
	ZoneEditParamsTypeFull      ZoneEditParamsType = "full"
	ZoneEditParamsTypePartial   ZoneEditParamsType = "partial"
	ZoneEditParamsTypeSecondary ZoneEditParamsType = "secondary"
)

type ZoneEditResponseEnvelope struct {
	Errors   []ZoneEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                         `json:"success,required"`
	Result  ZoneEditResponse             `json:"result"`
	JSON    zoneEditResponseEnvelopeJSON `json:"-"`
}

// zoneEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneEditResponseEnvelope]
type zoneEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEditResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZoneEditResponseEnvelopeErrors]
type zoneEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEditResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ZoneEditResponseEnvelopeMessages]
type zoneEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneGetResponseEnvelope struct {
	Errors   []ZoneGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                        `json:"success,required"`
	Result  ZoneGetResponse             `json:"result"`
	JSON    zoneGetResponseEnvelopeJSON `json:"-"`
}

// zoneGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneGetResponseEnvelope]
type zoneGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZoneGetResponseEnvelopeErrors]
type zoneGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ZoneGetResponseEnvelopeMessages]
type zoneGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
