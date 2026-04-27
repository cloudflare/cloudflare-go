// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// SettingDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingDomainService] method instead.
type SettingDomainService struct {
	Options []option.RequestOption
}

// NewSettingDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingDomainService(opts ...option.RequestOption) (r *SettingDomainService) {
	r = &SettingDomainService{}
	r.Options = opts
	return
}

// Returns a paginated list of email domains protected by Email Security. Includes
// domain configuration, delivery modes, and authorization status. Supports
// filtering by delivery mode and integration ID.
func (r *SettingDomainService) List(ctx context.Context, params SettingDomainListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingDomainListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns a paginated list of email domains protected by Email Security. Includes
// domain configuration, delivery modes, and authorization status. Supports
// filtering by delivery mode and integration ID.
func (r *SettingDomainService) ListAutoPaging(ctx context.Context, params SettingDomainListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingDomainListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes email security protection from a domain. After deletion, emails for this
// domain will no longer be processed by Email Security. This action cannot be
// undone.
func (r *SettingDomainService) Delete(ctx context.Context, domainID string, body SettingDomainDeleteParams, opts ...option.RequestOption) (res *SettingDomainDeleteResponse, err error) {
	var env SettingDomainDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%s", body.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates configuration for a protected email domain. Only provided fields will be
// modified. Changes affect delivery mode, security settings, and regional
// processing.
func (r *SettingDomainService) Edit(ctx context.Context, domainID string, params SettingDomainEditParams, opts ...option.RequestOption) (res *SettingDomainEditResponse, err error) {
	var env SettingDomainEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%s", params.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves detailed information for a specific protected email domain including
// its delivery configuration, SPF/DMARC status, and authorization state.
func (r *SettingDomainService) Get(ctx context.Context, domainID string, query SettingDomainGetParams, opts ...option.RequestOption) (res *SettingDomainGetResponse, err error) {
	var env SettingDomainGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%s", query.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SettingDomainListResponse struct {
	// Domain identifier
	ID                   string                                         `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainListResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainListResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                      `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainListResponseDMARCStatus           `json:"dmarc_status"`
	Domain               string                                         `json:"domain"`
	DropDispositions     []SettingDomainListResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainListResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainListResponseFolder                `json:"folder"`
	InboxProvider        SettingDomainListResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                         `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                       `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified       time.Time                          `json:"last_modified" format:"date-time"`
	LookbackHops       int64                              `json:"lookback_hops"`
	ModifiedAt         time.Time                          `json:"modified_at" format:"date-time"`
	O365TenantID       string                             `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainListResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                               `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                               `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainListResponseSPFStatus `json:"spf_status"`
	Status             SettingDomainListResponseStatus    `json:"status"`
	Transport          string                             `json:"transport"`
	JSON               settingDomainListResponseJSON      `json:"-"`
}

// settingDomainListResponseJSON contains the JSON metadata for the struct
// [SettingDomainListResponse]
type settingDomainListResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	Authorization        apijson.Field
	CreatedAt            apijson.Field
	DMARCStatus          apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	ModifiedAt           apijson.Field
	O365TenantID         apijson.Field
	Regions              apijson.Field
	RequireTLSInbound    apijson.Field
	RequireTLSOutbound   apijson.Field
	SPFStatus            apijson.Field
	Status               apijson.Field
	Transport            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SettingDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainListResponseAllowedDeliveryMode string

const (
	SettingDomainListResponseAllowedDeliveryModeDirect    SettingDomainListResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainListResponseAllowedDeliveryModeBcc       SettingDomainListResponseAllowedDeliveryMode = "BCC"
	SettingDomainListResponseAllowedDeliveryModeJournal   SettingDomainListResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainListResponseAllowedDeliveryModeAPI       SettingDomainListResponseAllowedDeliveryMode = "API"
	SettingDomainListResponseAllowedDeliveryModeRetroScan SettingDomainListResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainListResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainListResponseAllowedDeliveryModeDirect, SettingDomainListResponseAllowedDeliveryModeBcc, SettingDomainListResponseAllowedDeliveryModeJournal, SettingDomainListResponseAllowedDeliveryModeAPI, SettingDomainListResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainListResponseAuthorization struct {
	Authorized    bool                                       `json:"authorized" api:"required"`
	Timestamp     time.Time                                  `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                     `json:"status_message" api:"nullable"`
	JSON          settingDomainListResponseAuthorizationJSON `json:"-"`
}

// settingDomainListResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainListResponseAuthorization]
type settingDomainListResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainListResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainListResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainListResponseDMARCStatus string

const (
	SettingDomainListResponseDMARCStatusNone    SettingDomainListResponseDMARCStatus = "none"
	SettingDomainListResponseDMARCStatusGood    SettingDomainListResponseDMARCStatus = "good"
	SettingDomainListResponseDMARCStatusInvalid SettingDomainListResponseDMARCStatus = "invalid"
)

func (r SettingDomainListResponseDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainListResponseDMARCStatusNone, SettingDomainListResponseDMARCStatusGood, SettingDomainListResponseDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainListResponseDropDisposition string

const (
	SettingDomainListResponseDropDispositionMalicious    SettingDomainListResponseDropDisposition = "MALICIOUS"
	SettingDomainListResponseDropDispositionMaliciousBec SettingDomainListResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainListResponseDropDispositionSuspicious   SettingDomainListResponseDropDisposition = "SUSPICIOUS"
	SettingDomainListResponseDropDispositionSpoof        SettingDomainListResponseDropDisposition = "SPOOF"
	SettingDomainListResponseDropDispositionSpam         SettingDomainListResponseDropDisposition = "SPAM"
	SettingDomainListResponseDropDispositionBulk         SettingDomainListResponseDropDisposition = "BULK"
	SettingDomainListResponseDropDispositionEncrypted    SettingDomainListResponseDropDisposition = "ENCRYPTED"
	SettingDomainListResponseDropDispositionExternal     SettingDomainListResponseDropDisposition = "EXTERNAL"
	SettingDomainListResponseDropDispositionUnknown      SettingDomainListResponseDropDisposition = "UNKNOWN"
	SettingDomainListResponseDropDispositionNone         SettingDomainListResponseDropDisposition = "NONE"
)

func (r SettingDomainListResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainListResponseDropDispositionMalicious, SettingDomainListResponseDropDispositionMaliciousBec, SettingDomainListResponseDropDispositionSuspicious, SettingDomainListResponseDropDispositionSpoof, SettingDomainListResponseDropDispositionSpam, SettingDomainListResponseDropDispositionBulk, SettingDomainListResponseDropDispositionEncrypted, SettingDomainListResponseDropDispositionExternal, SettingDomainListResponseDropDispositionUnknown, SettingDomainListResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainListResponseEmailsProcessed struct {
	Timestamp                    time.Time                                    `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                        `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                        `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainListResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainListResponseEmailsProcessedJSON contains the JSON metadata for the
// struct [SettingDomainListResponseEmailsProcessed]
type settingDomainListResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainListResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainListResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainListResponseFolder string

const (
	SettingDomainListResponseFolderAllItems SettingDomainListResponseFolder = "AllItems"
	SettingDomainListResponseFolderInbox    SettingDomainListResponseFolder = "Inbox"
)

func (r SettingDomainListResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainListResponseFolderAllItems, SettingDomainListResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainListResponseInboxProvider string

const (
	SettingDomainListResponseInboxProviderMicrosoft SettingDomainListResponseInboxProvider = "Microsoft"
	SettingDomainListResponseInboxProviderGoogle    SettingDomainListResponseInboxProvider = "Google"
)

func (r SettingDomainListResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainListResponseInboxProviderMicrosoft, SettingDomainListResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainListResponseRegion string

const (
	SettingDomainListResponseRegionGlobal SettingDomainListResponseRegion = "GLOBAL"
	SettingDomainListResponseRegionAu     SettingDomainListResponseRegion = "AU"
	SettingDomainListResponseRegionDe     SettingDomainListResponseRegion = "DE"
	SettingDomainListResponseRegionIn     SettingDomainListResponseRegion = "IN"
	SettingDomainListResponseRegionUs     SettingDomainListResponseRegion = "US"
)

func (r SettingDomainListResponseRegion) IsKnown() bool {
	switch r {
	case SettingDomainListResponseRegionGlobal, SettingDomainListResponseRegionAu, SettingDomainListResponseRegionDe, SettingDomainListResponseRegionIn, SettingDomainListResponseRegionUs:
		return true
	}
	return false
}

type SettingDomainListResponseSPFStatus string

const (
	SettingDomainListResponseSPFStatusNone    SettingDomainListResponseSPFStatus = "none"
	SettingDomainListResponseSPFStatusGood    SettingDomainListResponseSPFStatus = "good"
	SettingDomainListResponseSPFStatusNeutral SettingDomainListResponseSPFStatus = "neutral"
	SettingDomainListResponseSPFStatusOpen    SettingDomainListResponseSPFStatus = "open"
	SettingDomainListResponseSPFStatusInvalid SettingDomainListResponseSPFStatus = "invalid"
)

func (r SettingDomainListResponseSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainListResponseSPFStatusNone, SettingDomainListResponseSPFStatusGood, SettingDomainListResponseSPFStatusNeutral, SettingDomainListResponseSPFStatusOpen, SettingDomainListResponseSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainListResponseStatus string

const (
	SettingDomainListResponseStatusPending SettingDomainListResponseStatus = "pending"
	SettingDomainListResponseStatusActive  SettingDomainListResponseStatus = "active"
	SettingDomainListResponseStatusFailed  SettingDomainListResponseStatus = "failed"
	SettingDomainListResponseStatusTimeout SettingDomainListResponseStatus = "timeout"
)

func (r SettingDomainListResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainListResponseStatusPending, SettingDomainListResponseStatusActive, SettingDomainListResponseStatusFailed, SettingDomainListResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainDeleteResponse struct {
	// Domain identifier
	ID   string                          `json:"id" api:"required" format:"uuid"`
	JSON settingDomainDeleteResponseJSON `json:"-"`
}

// settingDomainDeleteResponseJSON contains the JSON metadata for the struct
// [SettingDomainDeleteResponse]
type settingDomainDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponse struct {
	// Domain identifier
	ID                   string                                         `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainEditResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainEditResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                      `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainEditResponseDMARCStatus           `json:"dmarc_status"`
	Domain               string                                         `json:"domain"`
	DropDispositions     []SettingDomainEditResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainEditResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainEditResponseFolder                `json:"folder"`
	InboxProvider        SettingDomainEditResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                         `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                       `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified       time.Time                          `json:"last_modified" format:"date-time"`
	LookbackHops       int64                              `json:"lookback_hops"`
	ModifiedAt         time.Time                          `json:"modified_at" format:"date-time"`
	O365TenantID       string                             `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainEditResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                               `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                               `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainEditResponseSPFStatus `json:"spf_status"`
	Status             SettingDomainEditResponseStatus    `json:"status"`
	Transport          string                             `json:"transport"`
	JSON               settingDomainEditResponseJSON      `json:"-"`
}

// settingDomainEditResponseJSON contains the JSON metadata for the struct
// [SettingDomainEditResponse]
type settingDomainEditResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	Authorization        apijson.Field
	CreatedAt            apijson.Field
	DMARCStatus          apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	ModifiedAt           apijson.Field
	O365TenantID         apijson.Field
	Regions              apijson.Field
	RequireTLSInbound    apijson.Field
	RequireTLSOutbound   apijson.Field
	SPFStatus            apijson.Field
	Status               apijson.Field
	Transport            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SettingDomainEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseAllowedDeliveryMode string

const (
	SettingDomainEditResponseAllowedDeliveryModeDirect    SettingDomainEditResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainEditResponseAllowedDeliveryModeBcc       SettingDomainEditResponseAllowedDeliveryMode = "BCC"
	SettingDomainEditResponseAllowedDeliveryModeJournal   SettingDomainEditResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainEditResponseAllowedDeliveryModeAPI       SettingDomainEditResponseAllowedDeliveryMode = "API"
	SettingDomainEditResponseAllowedDeliveryModeRetroScan SettingDomainEditResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainEditResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseAllowedDeliveryModeDirect, SettingDomainEditResponseAllowedDeliveryModeBcc, SettingDomainEditResponseAllowedDeliveryModeJournal, SettingDomainEditResponseAllowedDeliveryModeAPI, SettingDomainEditResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainEditResponseAuthorization struct {
	Authorized    bool                                       `json:"authorized" api:"required"`
	Timestamp     time.Time                                  `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                     `json:"status_message" api:"nullable"`
	JSON          settingDomainEditResponseAuthorizationJSON `json:"-"`
}

// settingDomainEditResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainEditResponseAuthorization]
type settingDomainEditResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainEditResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseDMARCStatus string

const (
	SettingDomainEditResponseDMARCStatusNone    SettingDomainEditResponseDMARCStatus = "none"
	SettingDomainEditResponseDMARCStatusGood    SettingDomainEditResponseDMARCStatus = "good"
	SettingDomainEditResponseDMARCStatusInvalid SettingDomainEditResponseDMARCStatus = "invalid"
)

func (r SettingDomainEditResponseDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseDMARCStatusNone, SettingDomainEditResponseDMARCStatusGood, SettingDomainEditResponseDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainEditResponseDropDisposition string

const (
	SettingDomainEditResponseDropDispositionMalicious    SettingDomainEditResponseDropDisposition = "MALICIOUS"
	SettingDomainEditResponseDropDispositionMaliciousBec SettingDomainEditResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainEditResponseDropDispositionSuspicious   SettingDomainEditResponseDropDisposition = "SUSPICIOUS"
	SettingDomainEditResponseDropDispositionSpoof        SettingDomainEditResponseDropDisposition = "SPOOF"
	SettingDomainEditResponseDropDispositionSpam         SettingDomainEditResponseDropDisposition = "SPAM"
	SettingDomainEditResponseDropDispositionBulk         SettingDomainEditResponseDropDisposition = "BULK"
	SettingDomainEditResponseDropDispositionEncrypted    SettingDomainEditResponseDropDisposition = "ENCRYPTED"
	SettingDomainEditResponseDropDispositionExternal     SettingDomainEditResponseDropDisposition = "EXTERNAL"
	SettingDomainEditResponseDropDispositionUnknown      SettingDomainEditResponseDropDisposition = "UNKNOWN"
	SettingDomainEditResponseDropDispositionNone         SettingDomainEditResponseDropDisposition = "NONE"
)

func (r SettingDomainEditResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseDropDispositionMalicious, SettingDomainEditResponseDropDispositionMaliciousBec, SettingDomainEditResponseDropDispositionSuspicious, SettingDomainEditResponseDropDispositionSpoof, SettingDomainEditResponseDropDispositionSpam, SettingDomainEditResponseDropDispositionBulk, SettingDomainEditResponseDropDispositionEncrypted, SettingDomainEditResponseDropDispositionExternal, SettingDomainEditResponseDropDispositionUnknown, SettingDomainEditResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainEditResponseEmailsProcessed struct {
	Timestamp                    time.Time                                    `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                        `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                        `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainEditResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainEditResponseEmailsProcessedJSON contains the JSON metadata for the
// struct [SettingDomainEditResponseEmailsProcessed]
type settingDomainEditResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainEditResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseFolder string

const (
	SettingDomainEditResponseFolderAllItems SettingDomainEditResponseFolder = "AllItems"
	SettingDomainEditResponseFolderInbox    SettingDomainEditResponseFolder = "Inbox"
)

func (r SettingDomainEditResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseFolderAllItems, SettingDomainEditResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainEditResponseInboxProvider string

const (
	SettingDomainEditResponseInboxProviderMicrosoft SettingDomainEditResponseInboxProvider = "Microsoft"
	SettingDomainEditResponseInboxProviderGoogle    SettingDomainEditResponseInboxProvider = "Google"
)

func (r SettingDomainEditResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseInboxProviderMicrosoft, SettingDomainEditResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainEditResponseRegion string

const (
	SettingDomainEditResponseRegionGlobal SettingDomainEditResponseRegion = "GLOBAL"
	SettingDomainEditResponseRegionAu     SettingDomainEditResponseRegion = "AU"
	SettingDomainEditResponseRegionDe     SettingDomainEditResponseRegion = "DE"
	SettingDomainEditResponseRegionIn     SettingDomainEditResponseRegion = "IN"
	SettingDomainEditResponseRegionUs     SettingDomainEditResponseRegion = "US"
)

func (r SettingDomainEditResponseRegion) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseRegionGlobal, SettingDomainEditResponseRegionAu, SettingDomainEditResponseRegionDe, SettingDomainEditResponseRegionIn, SettingDomainEditResponseRegionUs:
		return true
	}
	return false
}

type SettingDomainEditResponseSPFStatus string

const (
	SettingDomainEditResponseSPFStatusNone    SettingDomainEditResponseSPFStatus = "none"
	SettingDomainEditResponseSPFStatusGood    SettingDomainEditResponseSPFStatus = "good"
	SettingDomainEditResponseSPFStatusNeutral SettingDomainEditResponseSPFStatus = "neutral"
	SettingDomainEditResponseSPFStatusOpen    SettingDomainEditResponseSPFStatus = "open"
	SettingDomainEditResponseSPFStatusInvalid SettingDomainEditResponseSPFStatus = "invalid"
)

func (r SettingDomainEditResponseSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseSPFStatusNone, SettingDomainEditResponseSPFStatusGood, SettingDomainEditResponseSPFStatusNeutral, SettingDomainEditResponseSPFStatusOpen, SettingDomainEditResponseSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainEditResponseStatus string

const (
	SettingDomainEditResponseStatusPending SettingDomainEditResponseStatus = "pending"
	SettingDomainEditResponseStatusActive  SettingDomainEditResponseStatus = "active"
	SettingDomainEditResponseStatusFailed  SettingDomainEditResponseStatus = "failed"
	SettingDomainEditResponseStatusTimeout SettingDomainEditResponseStatus = "timeout"
)

func (r SettingDomainEditResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseStatusPending, SettingDomainEditResponseStatusActive, SettingDomainEditResponseStatusFailed, SettingDomainEditResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainGetResponse struct {
	// Domain identifier
	ID                   string                                        `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainGetResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainGetResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                     `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainGetResponseDMARCStatus           `json:"dmarc_status"`
	Domain               string                                        `json:"domain"`
	DropDispositions     []SettingDomainGetResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainGetResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainGetResponseFolder                `json:"folder"`
	InboxProvider        SettingDomainGetResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                        `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                      `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified       time.Time                         `json:"last_modified" format:"date-time"`
	LookbackHops       int64                             `json:"lookback_hops"`
	ModifiedAt         time.Time                         `json:"modified_at" format:"date-time"`
	O365TenantID       string                            `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainGetResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                              `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                              `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainGetResponseSPFStatus `json:"spf_status"`
	Status             SettingDomainGetResponseStatus    `json:"status"`
	Transport          string                            `json:"transport"`
	JSON               settingDomainGetResponseJSON      `json:"-"`
}

// settingDomainGetResponseJSON contains the JSON metadata for the struct
// [SettingDomainGetResponse]
type settingDomainGetResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	Authorization        apijson.Field
	CreatedAt            apijson.Field
	DMARCStatus          apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	ModifiedAt           apijson.Field
	O365TenantID         apijson.Field
	Regions              apijson.Field
	RequireTLSInbound    apijson.Field
	RequireTLSOutbound   apijson.Field
	SPFStatus            apijson.Field
	Status               apijson.Field
	Transport            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SettingDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseAllowedDeliveryMode string

const (
	SettingDomainGetResponseAllowedDeliveryModeDirect    SettingDomainGetResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainGetResponseAllowedDeliveryModeBcc       SettingDomainGetResponseAllowedDeliveryMode = "BCC"
	SettingDomainGetResponseAllowedDeliveryModeJournal   SettingDomainGetResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainGetResponseAllowedDeliveryModeAPI       SettingDomainGetResponseAllowedDeliveryMode = "API"
	SettingDomainGetResponseAllowedDeliveryModeRetroScan SettingDomainGetResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainGetResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseAllowedDeliveryModeDirect, SettingDomainGetResponseAllowedDeliveryModeBcc, SettingDomainGetResponseAllowedDeliveryModeJournal, SettingDomainGetResponseAllowedDeliveryModeAPI, SettingDomainGetResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainGetResponseAuthorization struct {
	Authorized    bool                                      `json:"authorized" api:"required"`
	Timestamp     time.Time                                 `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                    `json:"status_message" api:"nullable"`
	JSON          settingDomainGetResponseAuthorizationJSON `json:"-"`
}

// settingDomainGetResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainGetResponseAuthorization]
type settingDomainGetResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainGetResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseDMARCStatus string

const (
	SettingDomainGetResponseDMARCStatusNone    SettingDomainGetResponseDMARCStatus = "none"
	SettingDomainGetResponseDMARCStatusGood    SettingDomainGetResponseDMARCStatus = "good"
	SettingDomainGetResponseDMARCStatusInvalid SettingDomainGetResponseDMARCStatus = "invalid"
)

func (r SettingDomainGetResponseDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseDMARCStatusNone, SettingDomainGetResponseDMARCStatusGood, SettingDomainGetResponseDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainGetResponseDropDisposition string

const (
	SettingDomainGetResponseDropDispositionMalicious    SettingDomainGetResponseDropDisposition = "MALICIOUS"
	SettingDomainGetResponseDropDispositionMaliciousBec SettingDomainGetResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainGetResponseDropDispositionSuspicious   SettingDomainGetResponseDropDisposition = "SUSPICIOUS"
	SettingDomainGetResponseDropDispositionSpoof        SettingDomainGetResponseDropDisposition = "SPOOF"
	SettingDomainGetResponseDropDispositionSpam         SettingDomainGetResponseDropDisposition = "SPAM"
	SettingDomainGetResponseDropDispositionBulk         SettingDomainGetResponseDropDisposition = "BULK"
	SettingDomainGetResponseDropDispositionEncrypted    SettingDomainGetResponseDropDisposition = "ENCRYPTED"
	SettingDomainGetResponseDropDispositionExternal     SettingDomainGetResponseDropDisposition = "EXTERNAL"
	SettingDomainGetResponseDropDispositionUnknown      SettingDomainGetResponseDropDisposition = "UNKNOWN"
	SettingDomainGetResponseDropDispositionNone         SettingDomainGetResponseDropDisposition = "NONE"
)

func (r SettingDomainGetResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseDropDispositionMalicious, SettingDomainGetResponseDropDispositionMaliciousBec, SettingDomainGetResponseDropDispositionSuspicious, SettingDomainGetResponseDropDispositionSpoof, SettingDomainGetResponseDropDispositionSpam, SettingDomainGetResponseDropDispositionBulk, SettingDomainGetResponseDropDispositionEncrypted, SettingDomainGetResponseDropDispositionExternal, SettingDomainGetResponseDropDispositionUnknown, SettingDomainGetResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainGetResponseEmailsProcessed struct {
	Timestamp                    time.Time                                   `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                       `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                       `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainGetResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainGetResponseEmailsProcessedJSON contains the JSON metadata for the
// struct [SettingDomainGetResponseEmailsProcessed]
type settingDomainGetResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainGetResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseFolder string

const (
	SettingDomainGetResponseFolderAllItems SettingDomainGetResponseFolder = "AllItems"
	SettingDomainGetResponseFolderInbox    SettingDomainGetResponseFolder = "Inbox"
)

func (r SettingDomainGetResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseFolderAllItems, SettingDomainGetResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainGetResponseInboxProvider string

const (
	SettingDomainGetResponseInboxProviderMicrosoft SettingDomainGetResponseInboxProvider = "Microsoft"
	SettingDomainGetResponseInboxProviderGoogle    SettingDomainGetResponseInboxProvider = "Google"
)

func (r SettingDomainGetResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseInboxProviderMicrosoft, SettingDomainGetResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainGetResponseRegion string

const (
	SettingDomainGetResponseRegionGlobal SettingDomainGetResponseRegion = "GLOBAL"
	SettingDomainGetResponseRegionAu     SettingDomainGetResponseRegion = "AU"
	SettingDomainGetResponseRegionDe     SettingDomainGetResponseRegion = "DE"
	SettingDomainGetResponseRegionIn     SettingDomainGetResponseRegion = "IN"
	SettingDomainGetResponseRegionUs     SettingDomainGetResponseRegion = "US"
)

func (r SettingDomainGetResponseRegion) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseRegionGlobal, SettingDomainGetResponseRegionAu, SettingDomainGetResponseRegionDe, SettingDomainGetResponseRegionIn, SettingDomainGetResponseRegionUs:
		return true
	}
	return false
}

type SettingDomainGetResponseSPFStatus string

const (
	SettingDomainGetResponseSPFStatusNone    SettingDomainGetResponseSPFStatus = "none"
	SettingDomainGetResponseSPFStatusGood    SettingDomainGetResponseSPFStatus = "good"
	SettingDomainGetResponseSPFStatusNeutral SettingDomainGetResponseSPFStatus = "neutral"
	SettingDomainGetResponseSPFStatusOpen    SettingDomainGetResponseSPFStatus = "open"
	SettingDomainGetResponseSPFStatusInvalid SettingDomainGetResponseSPFStatus = "invalid"
)

func (r SettingDomainGetResponseSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseSPFStatusNone, SettingDomainGetResponseSPFStatusGood, SettingDomainGetResponseSPFStatusNeutral, SettingDomainGetResponseSPFStatusOpen, SettingDomainGetResponseSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainGetResponseStatus string

const (
	SettingDomainGetResponseStatusPending SettingDomainGetResponseStatus = "pending"
	SettingDomainGetResponseStatusActive  SettingDomainGetResponseStatus = "active"
	SettingDomainGetResponseStatusFailed  SettingDomainGetResponseStatus = "failed"
	SettingDomainGetResponseStatusTimeout SettingDomainGetResponseStatus = "timeout"
)

func (r SettingDomainGetResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseStatusPending, SettingDomainGetResponseStatusActive, SettingDomainGetResponseStatusFailed, SettingDomainGetResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Currently active delivery mode to filter by.
	ActiveDeliveryMode param.Field[SettingDomainListParamsActiveDeliveryMode] `query:"active_delivery_mode"`
	// Delivery mode to filter by.
	AllowedDeliveryMode param.Field[SettingDomainListParamsAllowedDeliveryMode] `query:"allowed_delivery_mode"`
	// The sorting direction.
	Direction param.Field[SettingDomainListParamsDirection] `query:"direction"`
	// Domain names to filter by.
	Domain param.Field[[]string] `query:"domain"`
	// Integration ID to filter by.
	IntegrationID param.Field[string] `query:"integration_id" format:"uuid"`
	// Field to sort by.
	Order param.Field[SettingDomainListParamsOrder] `query:"order"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
	// Search term for filtering records. Behavior may change.
	Search param.Field[string] `query:"search"`
	// Filters response to domains with the provided status.
	Status param.Field[SettingDomainListParamsStatus] `query:"status"`
}

// URLQuery serializes [SettingDomainListParams]'s query parameters as
// `url.Values`.
func (r SettingDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Currently active delivery mode to filter by.
type SettingDomainListParamsActiveDeliveryMode string

const (
	SettingDomainListParamsActiveDeliveryModeDirect    SettingDomainListParamsActiveDeliveryMode = "DIRECT"
	SettingDomainListParamsActiveDeliveryModeBcc       SettingDomainListParamsActiveDeliveryMode = "BCC"
	SettingDomainListParamsActiveDeliveryModeJournal   SettingDomainListParamsActiveDeliveryMode = "JOURNAL"
	SettingDomainListParamsActiveDeliveryModeAPI       SettingDomainListParamsActiveDeliveryMode = "API"
	SettingDomainListParamsActiveDeliveryModeRetroScan SettingDomainListParamsActiveDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainListParamsActiveDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainListParamsActiveDeliveryModeDirect, SettingDomainListParamsActiveDeliveryModeBcc, SettingDomainListParamsActiveDeliveryModeJournal, SettingDomainListParamsActiveDeliveryModeAPI, SettingDomainListParamsActiveDeliveryModeRetroScan:
		return true
	}
	return false
}

// Delivery mode to filter by.
type SettingDomainListParamsAllowedDeliveryMode string

const (
	SettingDomainListParamsAllowedDeliveryModeDirect    SettingDomainListParamsAllowedDeliveryMode = "DIRECT"
	SettingDomainListParamsAllowedDeliveryModeBcc       SettingDomainListParamsAllowedDeliveryMode = "BCC"
	SettingDomainListParamsAllowedDeliveryModeJournal   SettingDomainListParamsAllowedDeliveryMode = "JOURNAL"
	SettingDomainListParamsAllowedDeliveryModeAPI       SettingDomainListParamsAllowedDeliveryMode = "API"
	SettingDomainListParamsAllowedDeliveryModeRetroScan SettingDomainListParamsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainListParamsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainListParamsAllowedDeliveryModeDirect, SettingDomainListParamsAllowedDeliveryModeBcc, SettingDomainListParamsAllowedDeliveryModeJournal, SettingDomainListParamsAllowedDeliveryModeAPI, SettingDomainListParamsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

// The sorting direction.
type SettingDomainListParamsDirection string

const (
	SettingDomainListParamsDirectionAsc  SettingDomainListParamsDirection = "asc"
	SettingDomainListParamsDirectionDesc SettingDomainListParamsDirection = "desc"
)

func (r SettingDomainListParamsDirection) IsKnown() bool {
	switch r {
	case SettingDomainListParamsDirectionAsc, SettingDomainListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by.
type SettingDomainListParamsOrder string

const (
	SettingDomainListParamsOrderDomain    SettingDomainListParamsOrder = "domain"
	SettingDomainListParamsOrderCreatedAt SettingDomainListParamsOrder = "created_at"
)

func (r SettingDomainListParamsOrder) IsKnown() bool {
	switch r {
	case SettingDomainListParamsOrderDomain, SettingDomainListParamsOrderCreatedAt:
		return true
	}
	return false
}

// Filters response to domains with the provided status.
type SettingDomainListParamsStatus string

const (
	SettingDomainListParamsStatusPending SettingDomainListParamsStatus = "pending"
	SettingDomainListParamsStatusActive  SettingDomainListParamsStatus = "active"
	SettingDomainListParamsStatusFailed  SettingDomainListParamsStatus = "failed"
	SettingDomainListParamsStatusTimeout SettingDomainListParamsStatus = "timeout"
)

func (r SettingDomainListParamsStatus) IsKnown() bool {
	switch r {
	case SettingDomainListParamsStatusPending, SettingDomainListParamsStatusActive, SettingDomainListParamsStatusFailed, SettingDomainListParamsStatusTimeout:
		return true
	}
	return false
}

type SettingDomainDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingDomainDeleteResponseEnvelope struct {
	Errors   []SettingDomainDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingDomainDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingDomainDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingDomainDeleteResponse                `json:"result"`
	JSON    settingDomainDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingDomainDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingDomainDeleteResponseEnvelope]
type settingDomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainDeleteResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           SettingDomainDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingDomainDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingDomainDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingDomainDeleteResponseEnvelopeErrors]
type settingDomainDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDomainDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    settingDomainDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingDomainDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingDomainDeleteResponseEnvelopeErrorsSource]
type settingDomainDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingDomainDeleteResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingDomainDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingDomainDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingDomainDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingDomainDeleteResponseEnvelopeMessages]
type settingDomainDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingDomainDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingDomainDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingDomainDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingDomainDeleteResponseEnvelopeMessagesSource]
type settingDomainDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingDomainDeleteResponseEnvelopeSuccess bool

const (
	SettingDomainDeleteResponseEnvelopeSuccessTrue SettingDomainDeleteResponseEnvelopeSuccess = true
)

func (r SettingDomainDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingDomainDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingDomainEditParams struct {
	// Identifier.
	AccountID            param.Field[string]                                       `path:"account_id" api:"required"`
	AllowedDeliveryModes param.Field[[]SettingDomainEditParamsAllowedDeliveryMode] `json:"allowed_delivery_modes"`
	Domain               param.Field[string]                                       `json:"domain"`
	DropDispositions     param.Field[[]SettingDomainEditParamsDropDisposition]     `json:"drop_dispositions"`
	Folder               param.Field[SettingDomainEditParamsFolder]                `json:"folder"`
	IntegrationID        param.Field[string]                                       `json:"integration_id" format:"uuid"`
	IPRestrictions       param.Field[[]string]                                     `json:"ip_restrictions"`
	LookbackHops         param.Field[int64]                                        `json:"lookback_hops"`
	Regions              param.Field[[]SettingDomainEditParamsRegion]              `json:"regions"`
	RequireTLSInbound    param.Field[bool]                                         `json:"require_tls_inbound"`
	RequireTLSOutbound   param.Field[bool]                                         `json:"require_tls_outbound"`
	Transport            param.Field[string]                                       `json:"transport"`
}

func (r SettingDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainEditParamsAllowedDeliveryMode string

const (
	SettingDomainEditParamsAllowedDeliveryModeDirect    SettingDomainEditParamsAllowedDeliveryMode = "DIRECT"
	SettingDomainEditParamsAllowedDeliveryModeBcc       SettingDomainEditParamsAllowedDeliveryMode = "BCC"
	SettingDomainEditParamsAllowedDeliveryModeJournal   SettingDomainEditParamsAllowedDeliveryMode = "JOURNAL"
	SettingDomainEditParamsAllowedDeliveryModeAPI       SettingDomainEditParamsAllowedDeliveryMode = "API"
	SettingDomainEditParamsAllowedDeliveryModeRetroScan SettingDomainEditParamsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainEditParamsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainEditParamsAllowedDeliveryModeDirect, SettingDomainEditParamsAllowedDeliveryModeBcc, SettingDomainEditParamsAllowedDeliveryModeJournal, SettingDomainEditParamsAllowedDeliveryModeAPI, SettingDomainEditParamsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainEditParamsDropDisposition string

const (
	SettingDomainEditParamsDropDispositionMalicious    SettingDomainEditParamsDropDisposition = "MALICIOUS"
	SettingDomainEditParamsDropDispositionMaliciousBec SettingDomainEditParamsDropDisposition = "MALICIOUS-BEC"
	SettingDomainEditParamsDropDispositionSuspicious   SettingDomainEditParamsDropDisposition = "SUSPICIOUS"
	SettingDomainEditParamsDropDispositionSpoof        SettingDomainEditParamsDropDisposition = "SPOOF"
	SettingDomainEditParamsDropDispositionSpam         SettingDomainEditParamsDropDisposition = "SPAM"
	SettingDomainEditParamsDropDispositionBulk         SettingDomainEditParamsDropDisposition = "BULK"
	SettingDomainEditParamsDropDispositionEncrypted    SettingDomainEditParamsDropDisposition = "ENCRYPTED"
	SettingDomainEditParamsDropDispositionExternal     SettingDomainEditParamsDropDisposition = "EXTERNAL"
	SettingDomainEditParamsDropDispositionUnknown      SettingDomainEditParamsDropDisposition = "UNKNOWN"
	SettingDomainEditParamsDropDispositionNone         SettingDomainEditParamsDropDisposition = "NONE"
)

func (r SettingDomainEditParamsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainEditParamsDropDispositionMalicious, SettingDomainEditParamsDropDispositionMaliciousBec, SettingDomainEditParamsDropDispositionSuspicious, SettingDomainEditParamsDropDispositionSpoof, SettingDomainEditParamsDropDispositionSpam, SettingDomainEditParamsDropDispositionBulk, SettingDomainEditParamsDropDispositionEncrypted, SettingDomainEditParamsDropDispositionExternal, SettingDomainEditParamsDropDispositionUnknown, SettingDomainEditParamsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainEditParamsFolder string

const (
	SettingDomainEditParamsFolderAllItems SettingDomainEditParamsFolder = "AllItems"
	SettingDomainEditParamsFolderInbox    SettingDomainEditParamsFolder = "Inbox"
)

func (r SettingDomainEditParamsFolder) IsKnown() bool {
	switch r {
	case SettingDomainEditParamsFolderAllItems, SettingDomainEditParamsFolderInbox:
		return true
	}
	return false
}

type SettingDomainEditParamsRegion string

const (
	SettingDomainEditParamsRegionGlobal SettingDomainEditParamsRegion = "GLOBAL"
	SettingDomainEditParamsRegionAu     SettingDomainEditParamsRegion = "AU"
	SettingDomainEditParamsRegionDe     SettingDomainEditParamsRegion = "DE"
	SettingDomainEditParamsRegionIn     SettingDomainEditParamsRegion = "IN"
	SettingDomainEditParamsRegionUs     SettingDomainEditParamsRegion = "US"
)

func (r SettingDomainEditParamsRegion) IsKnown() bool {
	switch r {
	case SettingDomainEditParamsRegionGlobal, SettingDomainEditParamsRegionAu, SettingDomainEditParamsRegionDe, SettingDomainEditParamsRegionIn, SettingDomainEditParamsRegionUs:
		return true
	}
	return false
}

type SettingDomainEditResponseEnvelope struct {
	Errors   []SettingDomainEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingDomainEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingDomainEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingDomainEditResponse                `json:"result"`
	JSON    settingDomainEditResponseEnvelopeJSON    `json:"-"`
}

// settingDomainEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainEditResponseEnvelope]
type settingDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           SettingDomainEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingDomainEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingDomainEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingDomainEditResponseEnvelopeErrors]
type settingDomainEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    settingDomainEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingDomainEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingDomainEditResponseEnvelopeErrorsSource]
type settingDomainEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           SettingDomainEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingDomainEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingDomainEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingDomainEditResponseEnvelopeMessages]
type settingDomainEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    settingDomainEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingDomainEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingDomainEditResponseEnvelopeMessagesSource]
type settingDomainEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingDomainEditResponseEnvelopeSuccess bool

const (
	SettingDomainEditResponseEnvelopeSuccessTrue SettingDomainEditResponseEnvelopeSuccess = true
)

func (r SettingDomainEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingDomainGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingDomainGetResponseEnvelope struct {
	Errors   []SettingDomainGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingDomainGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingDomainGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingDomainGetResponse                `json:"result"`
	JSON    settingDomainGetResponseEnvelopeJSON    `json:"-"`
}

// settingDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainGetResponseEnvelope]
type settingDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SettingDomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingDomainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingDomainGetResponseEnvelopeErrors]
type settingDomainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    settingDomainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingDomainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingDomainGetResponseEnvelopeErrorsSource]
type settingDomainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SettingDomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingDomainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingDomainGetResponseEnvelopeMessages]
type settingDomainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    settingDomainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingDomainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingDomainGetResponseEnvelopeMessagesSource]
type settingDomainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingDomainGetResponseEnvelopeSuccess bool

const (
	SettingDomainGetResponseEnvelopeSuccessTrue SettingDomainGetResponseEnvelopeSuccess = true
)

func (r SettingDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
