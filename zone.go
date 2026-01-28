// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// ZoneService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneService] method instead.
type ZoneService struct {
	Options []option.RequestOption
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r *ZoneService) {
	r = &ZoneService{}
	r.Options = opts
	return
}

// Create Zone
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *ZoneNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edits a zone. Only one zone property can be changed at a time.
func (r *ZoneService) Edit(ctx context.Context, params ZoneEditParams, opts ...option.RequestOption) (res *ZoneEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Zone Details
func (r *ZoneService) Get(ctx context.Context, query ZoneGetParams, opts ...option.RequestOption) (res *ZoneGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneNewResponse struct {
	Errors   []ZoneNewResponseError   `json:"errors,required"`
	Messages []ZoneNewResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                  `json:"success,required"`
	Result  ZoneNewResponseResult `json:"result"`
	JSON    zoneNewResponseJSON   `json:"-"`
}

// zoneNewResponseJSON contains the JSON metadata for the struct [ZoneNewResponse]
type zoneNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneNewResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    zoneNewResponseErrorJSON `json:"-"`
}

// zoneNewResponseErrorJSON contains the JSON metadata for the struct
// [ZoneNewResponseError]
type zoneNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneNewResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    zoneNewResponseMessageJSON `json:"-"`
}

// zoneNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneNewResponseMessage]
type zoneNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneNewResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to.
	Account ZoneNewResponseResultAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active.
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone.
	Meta ZoneNewResponseResultMeta `json:"meta,required"`
	// When the zone was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name.
	Name string `json:"name,required"`
	// The name servers Cloudflare assigns to a zone.
	NameServers []string `json:"name_servers,required" format:"hostname"`
	// DNS host at the time of switching to Cloudflare.
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare.
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare.
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone.
	Owner ZoneNewResponseResultOwner `json:"owner,required"`
	// A Zones subscription information.
	//
	// Deprecated: Please use the `/zones/{zone_id}/subscription` API to update a
	// zone's plan. Changing this value will create/cancel associated subscriptions. To
	// view available plans for this zone, see
	// [Zone Plans](https://developers.cloudflare.com/api/resources/zones/subresources/plans/).
	Plan ZoneNewResponseResultPlan `json:"plan,required"`
	// Allows the customer to use a custom apex. _Tenants Only Configuration_.
	CnameSuffix string `json:"cname_suffix"`
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused bool `json:"paused"`
	// Legacy permissions based on legacy user membership information.
	//
	// Deprecated: This has been replaced by Account memberships.
	Permissions []string `json:"permissions"`
	// The zone status on Cloudflare.
	Status ZoneNewResponseResultStatus `json:"status"`
	// The root organizational unit that this zone belongs to (such as a tenant or
	// organization).
	Tenant ZoneNewResponseResultTenant `json:"tenant"`
	// The immediate parent organizational unit that this zone belongs to (such as
	// under a tenant or sub-organization).
	TenantUnit ZoneNewResponseResultTenantUnit `json:"tenant_unit"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type ZoneNewResponseResultType `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string `json:"vanity_name_servers" format:"hostname"`
	// Verification key for partial zone setup.
	VerificationKey string                    `json:"verification_key"`
	JSON            zoneNewResponseResultJSON `json:"-"`
}

// zoneNewResponseResultJSON contains the JSON metadata for the struct
// [ZoneNewResponseResult]
type zoneNewResponseResultJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	NameServers         apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	Plan                apijson.Field
	CnameSuffix         apijson.Field
	Paused              apijson.Field
	Permissions         apijson.Field
	Status              apijson.Field
	Tenant              apijson.Field
	TenantUnit          apijson.Field
	Type                apijson.Field
	VanityNameServers   apijson.Field
	VerificationKey     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// The account the zone belongs to.
type ZoneNewResponseResultAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account.
	Name string                           `json:"name"`
	JSON zoneNewResponseResultAccountJSON `json:"-"`
}

// zoneNewResponseResultAccountJSON contains the JSON metadata for the struct
// [ZoneNewResponseResultAccount]
type zoneNewResponseResultAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseResultAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseResultAccountJSON) RawJSON() string {
	return r.raw
}

// Metadata about the zone.
type ZoneNewResponseResultMeta struct {
	// The zone is only configured for CDN.
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have.
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS.
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS.
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have.
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing.
	PhishingDetected bool                          `json:"phishing_detected"`
	Step             int64                         `json:"step"`
	JSON             zoneNewResponseResultMetaJSON `json:"-"`
}

// zoneNewResponseResultMetaJSON contains the JSON metadata for the struct
// [ZoneNewResponseResultMeta]
type zoneNewResponseResultMetaJSON struct {
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

func (r *ZoneNewResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// The owner of the zone.
type ZoneNewResponseResultOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner.
	Name string `json:"name"`
	// The type of owner.
	Type string                         `json:"type"`
	JSON zoneNewResponseResultOwnerJSON `json:"-"`
}

// zoneNewResponseResultOwnerJSON contains the JSON metadata for the struct
// [ZoneNewResponseResultOwner]
type zoneNewResponseResultOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseResultOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseResultOwnerJSON) RawJSON() string {
	return r.raw
}

// A Zones subscription information.
//
// Deprecated: Please use the `/zones/{zone_id}/subscription` API to update a
// zone's plan. Changing this value will create/cancel associated subscriptions. To
// view available plans for this zone, see
// [Zone Plans](https://developers.cloudflare.com/api/resources/zones/subresources/plans/).
type ZoneNewResponseResultPlan struct {
	// Identifier
	ID string `json:"id"`
	// States if the subscription can be activated.
	CanSubscribe bool `json:"can_subscribe"`
	// The denomination of the customer.
	Currency string `json:"currency"`
	// If this Zone is managed by another company.
	ExternallyManaged bool `json:"externally_managed"`
	// How often the customer is billed.
	Frequency string `json:"frequency"`
	// States if the subscription active.
	IsSubscribed bool `json:"is_subscribed"`
	// If the legacy discount applies to this Zone.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy name of the plan.
	LegacyID string `json:"legacy_id"`
	// Name of the owner.
	Name string `json:"name"`
	// How much the customer is paying.
	Price float64                       `json:"price"`
	JSON  zoneNewResponseResultPlanJSON `json:"-"`
}

// zoneNewResponseResultPlanJSON contains the JSON metadata for the struct
// [ZoneNewResponseResultPlan]
type zoneNewResponseResultPlanJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneNewResponseResultPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseResultPlanJSON) RawJSON() string {
	return r.raw
}

// The zone status on Cloudflare.
type ZoneNewResponseResultStatus string

const (
	ZoneNewResponseResultStatusInitializing ZoneNewResponseResultStatus = "initializing"
	ZoneNewResponseResultStatusPending      ZoneNewResponseResultStatus = "pending"
	ZoneNewResponseResultStatusActive       ZoneNewResponseResultStatus = "active"
	ZoneNewResponseResultStatusMoved        ZoneNewResponseResultStatus = "moved"
)

func (r ZoneNewResponseResultStatus) IsKnown() bool {
	switch r {
	case ZoneNewResponseResultStatusInitializing, ZoneNewResponseResultStatusPending, ZoneNewResponseResultStatusActive, ZoneNewResponseResultStatusMoved:
		return true
	}
	return false
}

// The root organizational unit that this zone belongs to (such as a tenant or
// organization).
type ZoneNewResponseResultTenant struct {
	// Identifier
	ID string `json:"id"`
	// The name of the Tenant account.
	Name string                          `json:"name"`
	JSON zoneNewResponseResultTenantJSON `json:"-"`
}

// zoneNewResponseResultTenantJSON contains the JSON metadata for the struct
// [ZoneNewResponseResultTenant]
type zoneNewResponseResultTenantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseResultTenant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseResultTenantJSON) RawJSON() string {
	return r.raw
}

// The immediate parent organizational unit that this zone belongs to (such as
// under a tenant or sub-organization).
type ZoneNewResponseResultTenantUnit struct {
	// Identifier
	ID   string                              `json:"id"`
	JSON zoneNewResponseResultTenantUnitJSON `json:"-"`
}

// zoneNewResponseResultTenantUnitJSON contains the JSON metadata for the struct
// [ZoneNewResponseResultTenantUnit]
type zoneNewResponseResultTenantUnitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseResultTenantUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseResultTenantUnitJSON) RawJSON() string {
	return r.raw
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup.
type ZoneNewResponseResultType string

const (
	ZoneNewResponseResultTypeFull      ZoneNewResponseResultType = "full"
	ZoneNewResponseResultTypePartial   ZoneNewResponseResultType = "partial"
	ZoneNewResponseResultTypeSecondary ZoneNewResponseResultType = "secondary"
	ZoneNewResponseResultTypeInternal  ZoneNewResponseResultType = "internal"
)

func (r ZoneNewResponseResultType) IsKnown() bool {
	switch r {
	case ZoneNewResponseResultTypeFull, ZoneNewResponseResultTypePartial, ZoneNewResponseResultTypeSecondary, ZoneNewResponseResultTypeInternal:
		return true
	}
	return false
}

type ZoneEditResponse struct {
	Errors   []ZoneEditResponseError   `json:"errors,required"`
	Messages []ZoneEditResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                   `json:"success,required"`
	Result  ZoneEditResponseResult `json:"result"`
	JSON    zoneEditResponseJSON   `json:"-"`
}

// zoneEditResponseJSON contains the JSON metadata for the struct
// [ZoneEditResponse]
type zoneEditResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneEditResponseError struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    zoneEditResponseErrorJSON `json:"-"`
}

// zoneEditResponseErrorJSON contains the JSON metadata for the struct
// [ZoneEditResponseError]
type zoneEditResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneEditResponseMessage struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    zoneEditResponseMessageJSON `json:"-"`
}

// zoneEditResponseMessageJSON contains the JSON metadata for the struct
// [ZoneEditResponseMessage]
type zoneEditResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneEditResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to.
	Account ZoneEditResponseResultAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active.
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone.
	Meta ZoneEditResponseResultMeta `json:"meta,required"`
	// When the zone was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name.
	Name string `json:"name,required"`
	// The name servers Cloudflare assigns to a zone.
	NameServers []string `json:"name_servers,required" format:"hostname"`
	// DNS host at the time of switching to Cloudflare.
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare.
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare.
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone.
	Owner ZoneEditResponseResultOwner `json:"owner,required"`
	// A Zones subscription information.
	//
	// Deprecated: Please use the `/zones/{zone_id}/subscription` API to update a
	// zone's plan. Changing this value will create/cancel associated subscriptions. To
	// view available plans for this zone, see
	// [Zone Plans](https://developers.cloudflare.com/api/resources/zones/subresources/plans/).
	Plan ZoneEditResponseResultPlan `json:"plan,required"`
	// Allows the customer to use a custom apex. _Tenants Only Configuration_.
	CnameSuffix string `json:"cname_suffix"`
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused bool `json:"paused"`
	// Legacy permissions based on legacy user membership information.
	//
	// Deprecated: This has been replaced by Account memberships.
	Permissions []string `json:"permissions"`
	// The zone status on Cloudflare.
	Status ZoneEditResponseResultStatus `json:"status"`
	// The root organizational unit that this zone belongs to (such as a tenant or
	// organization).
	Tenant ZoneEditResponseResultTenant `json:"tenant"`
	// The immediate parent organizational unit that this zone belongs to (such as
	// under a tenant or sub-organization).
	TenantUnit ZoneEditResponseResultTenantUnit `json:"tenant_unit"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type ZoneEditResponseResultType `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string `json:"vanity_name_servers" format:"hostname"`
	// Verification key for partial zone setup.
	VerificationKey string                     `json:"verification_key"`
	JSON            zoneEditResponseResultJSON `json:"-"`
}

// zoneEditResponseResultJSON contains the JSON metadata for the struct
// [ZoneEditResponseResult]
type zoneEditResponseResultJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	NameServers         apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	Plan                apijson.Field
	CnameSuffix         apijson.Field
	Paused              apijson.Field
	Permissions         apijson.Field
	Status              apijson.Field
	Tenant              apijson.Field
	TenantUnit          apijson.Field
	Type                apijson.Field
	VanityNameServers   apijson.Field
	VerificationKey     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneEditResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseResultJSON) RawJSON() string {
	return r.raw
}

// The account the zone belongs to.
type ZoneEditResponseResultAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account.
	Name string                            `json:"name"`
	JSON zoneEditResponseResultAccountJSON `json:"-"`
}

// zoneEditResponseResultAccountJSON contains the JSON metadata for the struct
// [ZoneEditResponseResultAccount]
type zoneEditResponseResultAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseResultAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseResultAccountJSON) RawJSON() string {
	return r.raw
}

// Metadata about the zone.
type ZoneEditResponseResultMeta struct {
	// The zone is only configured for CDN.
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have.
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS.
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS.
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have.
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing.
	PhishingDetected bool                           `json:"phishing_detected"`
	Step             int64                          `json:"step"`
	JSON             zoneEditResponseResultMetaJSON `json:"-"`
}

// zoneEditResponseResultMetaJSON contains the JSON metadata for the struct
// [ZoneEditResponseResultMeta]
type zoneEditResponseResultMetaJSON struct {
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

func (r *ZoneEditResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// The owner of the zone.
type ZoneEditResponseResultOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner.
	Name string `json:"name"`
	// The type of owner.
	Type string                          `json:"type"`
	JSON zoneEditResponseResultOwnerJSON `json:"-"`
}

// zoneEditResponseResultOwnerJSON contains the JSON metadata for the struct
// [ZoneEditResponseResultOwner]
type zoneEditResponseResultOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseResultOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseResultOwnerJSON) RawJSON() string {
	return r.raw
}

// A Zones subscription information.
//
// Deprecated: Please use the `/zones/{zone_id}/subscription` API to update a
// zone's plan. Changing this value will create/cancel associated subscriptions. To
// view available plans for this zone, see
// [Zone Plans](https://developers.cloudflare.com/api/resources/zones/subresources/plans/).
type ZoneEditResponseResultPlan struct {
	// Identifier
	ID string `json:"id"`
	// States if the subscription can be activated.
	CanSubscribe bool `json:"can_subscribe"`
	// The denomination of the customer.
	Currency string `json:"currency"`
	// If this Zone is managed by another company.
	ExternallyManaged bool `json:"externally_managed"`
	// How often the customer is billed.
	Frequency string `json:"frequency"`
	// States if the subscription active.
	IsSubscribed bool `json:"is_subscribed"`
	// If the legacy discount applies to this Zone.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy name of the plan.
	LegacyID string `json:"legacy_id"`
	// Name of the owner.
	Name string `json:"name"`
	// How much the customer is paying.
	Price float64                        `json:"price"`
	JSON  zoneEditResponseResultPlanJSON `json:"-"`
}

// zoneEditResponseResultPlanJSON contains the JSON metadata for the struct
// [ZoneEditResponseResultPlan]
type zoneEditResponseResultPlanJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneEditResponseResultPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseResultPlanJSON) RawJSON() string {
	return r.raw
}

// The zone status on Cloudflare.
type ZoneEditResponseResultStatus string

const (
	ZoneEditResponseResultStatusInitializing ZoneEditResponseResultStatus = "initializing"
	ZoneEditResponseResultStatusPending      ZoneEditResponseResultStatus = "pending"
	ZoneEditResponseResultStatusActive       ZoneEditResponseResultStatus = "active"
	ZoneEditResponseResultStatusMoved        ZoneEditResponseResultStatus = "moved"
)

func (r ZoneEditResponseResultStatus) IsKnown() bool {
	switch r {
	case ZoneEditResponseResultStatusInitializing, ZoneEditResponseResultStatusPending, ZoneEditResponseResultStatusActive, ZoneEditResponseResultStatusMoved:
		return true
	}
	return false
}

// The root organizational unit that this zone belongs to (such as a tenant or
// organization).
type ZoneEditResponseResultTenant struct {
	// Identifier
	ID string `json:"id"`
	// The name of the Tenant account.
	Name string                           `json:"name"`
	JSON zoneEditResponseResultTenantJSON `json:"-"`
}

// zoneEditResponseResultTenantJSON contains the JSON metadata for the struct
// [ZoneEditResponseResultTenant]
type zoneEditResponseResultTenantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseResultTenant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseResultTenantJSON) RawJSON() string {
	return r.raw
}

// The immediate parent organizational unit that this zone belongs to (such as
// under a tenant or sub-organization).
type ZoneEditResponseResultTenantUnit struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON zoneEditResponseResultTenantUnitJSON `json:"-"`
}

// zoneEditResponseResultTenantUnitJSON contains the JSON metadata for the struct
// [ZoneEditResponseResultTenantUnit]
type zoneEditResponseResultTenantUnitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseResultTenantUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseResultTenantUnitJSON) RawJSON() string {
	return r.raw
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup.
type ZoneEditResponseResultType string

const (
	ZoneEditResponseResultTypeFull      ZoneEditResponseResultType = "full"
	ZoneEditResponseResultTypePartial   ZoneEditResponseResultType = "partial"
	ZoneEditResponseResultTypeSecondary ZoneEditResponseResultType = "secondary"
	ZoneEditResponseResultTypeInternal  ZoneEditResponseResultType = "internal"
)

func (r ZoneEditResponseResultType) IsKnown() bool {
	switch r {
	case ZoneEditResponseResultTypeFull, ZoneEditResponseResultTypePartial, ZoneEditResponseResultTypeSecondary, ZoneEditResponseResultTypeInternal:
		return true
	}
	return false
}

type ZoneGetResponse struct {
	Errors   []ZoneGetResponseError   `json:"errors,required"`
	Messages []ZoneGetResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                  `json:"success,required"`
	Result  ZoneGetResponseResult `json:"result"`
	JSON    zoneGetResponseJSON   `json:"-"`
}

// zoneGetResponseJSON contains the JSON metadata for the struct [ZoneGetResponse]
type zoneGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneGetResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    zoneGetResponseErrorJSON `json:"-"`
}

// zoneGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneGetResponseError]
type zoneGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneGetResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    zoneGetResponseMessageJSON `json:"-"`
}

// zoneGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneGetResponseMessage]
type zoneGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneGetResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to.
	Account ZoneGetResponseResultAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active.
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone.
	Meta ZoneGetResponseResultMeta `json:"meta,required"`
	// When the zone was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name.
	Name string `json:"name,required"`
	// The name servers Cloudflare assigns to a zone.
	NameServers []string `json:"name_servers,required" format:"hostname"`
	// DNS host at the time of switching to Cloudflare.
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare.
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare.
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone.
	Owner ZoneGetResponseResultOwner `json:"owner,required"`
	// A Zones subscription information.
	//
	// Deprecated: Please use the `/zones/{zone_id}/subscription` API to update a
	// zone's plan. Changing this value will create/cancel associated subscriptions. To
	// view available plans for this zone, see
	// [Zone Plans](https://developers.cloudflare.com/api/resources/zones/subresources/plans/).
	Plan ZoneGetResponseResultPlan `json:"plan,required"`
	// Allows the customer to use a custom apex. _Tenants Only Configuration_.
	CnameSuffix string `json:"cname_suffix"`
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused bool `json:"paused"`
	// Legacy permissions based on legacy user membership information.
	//
	// Deprecated: This has been replaced by Account memberships.
	Permissions []string `json:"permissions"`
	// The zone status on Cloudflare.
	Status ZoneGetResponseResultStatus `json:"status"`
	// The root organizational unit that this zone belongs to (such as a tenant or
	// organization).
	Tenant ZoneGetResponseResultTenant `json:"tenant"`
	// The immediate parent organizational unit that this zone belongs to (such as
	// under a tenant or sub-organization).
	TenantUnit ZoneGetResponseResultTenantUnit `json:"tenant_unit"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type ZoneGetResponseResultType `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string `json:"vanity_name_servers" format:"hostname"`
	// Verification key for partial zone setup.
	VerificationKey string                    `json:"verification_key"`
	JSON            zoneGetResponseResultJSON `json:"-"`
}

// zoneGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneGetResponseResult]
type zoneGetResponseResultJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	NameServers         apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	Plan                apijson.Field
	CnameSuffix         apijson.Field
	Paused              apijson.Field
	Permissions         apijson.Field
	Status              apijson.Field
	Tenant              apijson.Field
	TenantUnit          apijson.Field
	Type                apijson.Field
	VanityNameServers   apijson.Field
	VerificationKey     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// The account the zone belongs to.
type ZoneGetResponseResultAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account.
	Name string                           `json:"name"`
	JSON zoneGetResponseResultAccountJSON `json:"-"`
}

// zoneGetResponseResultAccountJSON contains the JSON metadata for the struct
// [ZoneGetResponseResultAccount]
type zoneGetResponseResultAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseResultAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseResultAccountJSON) RawJSON() string {
	return r.raw
}

// Metadata about the zone.
type ZoneGetResponseResultMeta struct {
	// The zone is only configured for CDN.
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have.
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS.
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS.
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have.
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing.
	PhishingDetected bool                          `json:"phishing_detected"`
	Step             int64                         `json:"step"`
	JSON             zoneGetResponseResultMetaJSON `json:"-"`
}

// zoneGetResponseResultMetaJSON contains the JSON metadata for the struct
// [ZoneGetResponseResultMeta]
type zoneGetResponseResultMetaJSON struct {
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

func (r *ZoneGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// The owner of the zone.
type ZoneGetResponseResultOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner.
	Name string `json:"name"`
	// The type of owner.
	Type string                         `json:"type"`
	JSON zoneGetResponseResultOwnerJSON `json:"-"`
}

// zoneGetResponseResultOwnerJSON contains the JSON metadata for the struct
// [ZoneGetResponseResultOwner]
type zoneGetResponseResultOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseResultOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseResultOwnerJSON) RawJSON() string {
	return r.raw
}

// A Zones subscription information.
//
// Deprecated: Please use the `/zones/{zone_id}/subscription` API to update a
// zone's plan. Changing this value will create/cancel associated subscriptions. To
// view available plans for this zone, see
// [Zone Plans](https://developers.cloudflare.com/api/resources/zones/subresources/plans/).
type ZoneGetResponseResultPlan struct {
	// Identifier
	ID string `json:"id"`
	// States if the subscription can be activated.
	CanSubscribe bool `json:"can_subscribe"`
	// The denomination of the customer.
	Currency string `json:"currency"`
	// If this Zone is managed by another company.
	ExternallyManaged bool `json:"externally_managed"`
	// How often the customer is billed.
	Frequency string `json:"frequency"`
	// States if the subscription active.
	IsSubscribed bool `json:"is_subscribed"`
	// If the legacy discount applies to this Zone.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy name of the plan.
	LegacyID string `json:"legacy_id"`
	// Name of the owner.
	Name string `json:"name"`
	// How much the customer is paying.
	Price float64                       `json:"price"`
	JSON  zoneGetResponseResultPlanJSON `json:"-"`
}

// zoneGetResponseResultPlanJSON contains the JSON metadata for the struct
// [ZoneGetResponseResultPlan]
type zoneGetResponseResultPlanJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneGetResponseResultPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseResultPlanJSON) RawJSON() string {
	return r.raw
}

// The zone status on Cloudflare.
type ZoneGetResponseResultStatus string

const (
	ZoneGetResponseResultStatusInitializing ZoneGetResponseResultStatus = "initializing"
	ZoneGetResponseResultStatusPending      ZoneGetResponseResultStatus = "pending"
	ZoneGetResponseResultStatusActive       ZoneGetResponseResultStatus = "active"
	ZoneGetResponseResultStatusMoved        ZoneGetResponseResultStatus = "moved"
)

func (r ZoneGetResponseResultStatus) IsKnown() bool {
	switch r {
	case ZoneGetResponseResultStatusInitializing, ZoneGetResponseResultStatusPending, ZoneGetResponseResultStatusActive, ZoneGetResponseResultStatusMoved:
		return true
	}
	return false
}

// The root organizational unit that this zone belongs to (such as a tenant or
// organization).
type ZoneGetResponseResultTenant struct {
	// Identifier
	ID string `json:"id"`
	// The name of the Tenant account.
	Name string                          `json:"name"`
	JSON zoneGetResponseResultTenantJSON `json:"-"`
}

// zoneGetResponseResultTenantJSON contains the JSON metadata for the struct
// [ZoneGetResponseResultTenant]
type zoneGetResponseResultTenantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseResultTenant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseResultTenantJSON) RawJSON() string {
	return r.raw
}

// The immediate parent organizational unit that this zone belongs to (such as
// under a tenant or sub-organization).
type ZoneGetResponseResultTenantUnit struct {
	// Identifier
	ID   string                              `json:"id"`
	JSON zoneGetResponseResultTenantUnitJSON `json:"-"`
}

// zoneGetResponseResultTenantUnitJSON contains the JSON metadata for the struct
// [ZoneGetResponseResultTenantUnit]
type zoneGetResponseResultTenantUnitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseResultTenantUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseResultTenantUnitJSON) RawJSON() string {
	return r.raw
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup.
type ZoneGetResponseResultType string

const (
	ZoneGetResponseResultTypeFull      ZoneGetResponseResultType = "full"
	ZoneGetResponseResultTypePartial   ZoneGetResponseResultType = "partial"
	ZoneGetResponseResultTypeSecondary ZoneGetResponseResultType = "secondary"
	ZoneGetResponseResultTypeInternal  ZoneGetResponseResultType = "internal"
)

func (r ZoneGetResponseResultType) IsKnown() bool {
	switch r {
	case ZoneGetResponseResultTypeFull, ZoneGetResponseResultTypePartial, ZoneGetResponseResultTypeSecondary, ZoneGetResponseResultTypeInternal:
		return true
	}
	return false
}

type ZoneNewParams struct {
	Account param.Field[ZoneNewParamsAccount] `json:"account,required"`
	// The domain name.
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
	ZoneNewParamsTypeInternal  ZoneNewParamsType = "internal"
)

func (r ZoneNewParamsType) IsKnown() bool {
	switch r {
	case ZoneNewParamsTypeFull, ZoneNewParamsTypePartial, ZoneNewParamsTypeSecondary, ZoneNewParamsTypeInternal:
		return true
	}
	return false
}

type ZoneEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused param.Field[bool] `json:"paused"`
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
	ZoneEditParamsTypeInternal  ZoneEditParamsType = "internal"
)

func (r ZoneEditParamsType) IsKnown() bool {
	switch r {
	case ZoneEditParamsTypeFull, ZoneEditParamsTypePartial, ZoneEditParamsTypeSecondary, ZoneEditParamsTypeInternal:
		return true
	}
	return false
}

type ZoneGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
