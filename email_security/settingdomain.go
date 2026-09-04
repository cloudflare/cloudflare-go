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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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

// Protects a new email domain by adding it to Email Security. Accepts a flat
// configuration object covering all delivery modes. Returns the newly created
// domain configuration.
func (r *SettingDomainService) New(ctx context.Context, params SettingDomainNewParams, opts ...option.RequestOption) (res *SettingDomainNewResponse, err error) {
	var env SettingDomainNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces all mutable fields of a protected email domain in a single atomic
// operation. Unlike PATCH, all non-computed fields are required.
func (r *SettingDomainService) Update(ctx context.Context, domainID string, params SettingDomainUpdateParams, opts ...option.RequestOption) (res *SettingDomainUpdateResponse, err error) {
	var env SettingDomainUpdateResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
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

// Executes multiple domain operations in a single request. All four operation
// arrays (deletes, patches, puts, posts) are required and executed in order. Send
// empty arrays for unused operations.
func (r *SettingDomainService) Batch(ctx context.Context, params SettingDomainBatchParams, opts ...option.RequestOption) (res *SettingDomainBatchResponse, err error) {
	var env SettingDomainBatchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/batch", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes protection from multiple email domains. Deprecated; use the batch
// endpoint instead.
//
// Deprecated: This endpoint is deprecated. Use POST
// /accounts/{account_id}/email-security/settings/domains/batch instead.
func (r *SettingDomainService) BulkDelete(ctx context.Context, body SettingDomainBulkDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[SettingDomainBulkDeleteResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains", body.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Removes protection from multiple email domains. Deprecated; use the batch
// endpoint instead.
//
// Deprecated: This endpoint is deprecated. Use POST
// /accounts/{account_id}/email-security/settings/domains/batch instead.
func (r *SettingDomainService) BulkDeleteAutoPaging(ctx context.Context, body SettingDomainBulkDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SettingDomainBulkDeleteResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkDelete(ctx, body, opts...))
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

type SettingDomainNewResponse struct {
	// Domain identifier.
	ID                   string                                        `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainNewResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainNewResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                     `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainNewResponseDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                        `json:"domain"`
	DropDispositions     []SettingDomainNewResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainNewResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainNewResponseFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainNewResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                        `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                      `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                         `json:"last_modified" format:"date-time"`
	LookbackHops       int64                             `json:"lookback_hops"`
	ModifiedAt         time.Time                         `json:"modified_at" format:"date-time"`
	O365TenantID       string                            `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainNewResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                              `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                              `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainNewResponseSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainNewResponseStatus    `json:"status" api:"nullable"`
	Transport          string                            `json:"transport"`
	JSON               settingDomainNewResponseJSON      `json:"-"`
}

// settingDomainNewResponseJSON contains the JSON metadata for the struct
// [SettingDomainNewResponse]
type settingDomainNewResponseJSON struct {
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

func (r *SettingDomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainNewResponseAllowedDeliveryMode string

const (
	SettingDomainNewResponseAllowedDeliveryModeDirect    SettingDomainNewResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainNewResponseAllowedDeliveryModeBcc       SettingDomainNewResponseAllowedDeliveryMode = "BCC"
	SettingDomainNewResponseAllowedDeliveryModeJournal   SettingDomainNewResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainNewResponseAllowedDeliveryModeAPI       SettingDomainNewResponseAllowedDeliveryMode = "API"
	SettingDomainNewResponseAllowedDeliveryModeRetroScan SettingDomainNewResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainNewResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseAllowedDeliveryModeDirect, SettingDomainNewResponseAllowedDeliveryModeBcc, SettingDomainNewResponseAllowedDeliveryModeJournal, SettingDomainNewResponseAllowedDeliveryModeAPI, SettingDomainNewResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainNewResponseAuthorization struct {
	Authorized    bool                                      `json:"authorized" api:"required"`
	Timestamp     time.Time                                 `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                    `json:"status_message" api:"nullable"`
	JSON          settingDomainNewResponseAuthorizationJSON `json:"-"`
}

// settingDomainNewResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainNewResponseAuthorization]
type settingDomainNewResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainNewResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainNewResponseDMARCStatus string

const (
	SettingDomainNewResponseDMARCStatusNone    SettingDomainNewResponseDMARCStatus = "none"
	SettingDomainNewResponseDMARCStatusGood    SettingDomainNewResponseDMARCStatus = "good"
	SettingDomainNewResponseDMARCStatusInvalid SettingDomainNewResponseDMARCStatus = "invalid"
)

func (r SettingDomainNewResponseDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseDMARCStatusNone, SettingDomainNewResponseDMARCStatusGood, SettingDomainNewResponseDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainNewResponseDropDisposition string

const (
	SettingDomainNewResponseDropDispositionMalicious    SettingDomainNewResponseDropDisposition = "MALICIOUS"
	SettingDomainNewResponseDropDispositionMaliciousBec SettingDomainNewResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainNewResponseDropDispositionSuspicious   SettingDomainNewResponseDropDisposition = "SUSPICIOUS"
	SettingDomainNewResponseDropDispositionSpoof        SettingDomainNewResponseDropDisposition = "SPOOF"
	SettingDomainNewResponseDropDispositionSpam         SettingDomainNewResponseDropDisposition = "SPAM"
	SettingDomainNewResponseDropDispositionBulk         SettingDomainNewResponseDropDisposition = "BULK"
	SettingDomainNewResponseDropDispositionEncrypted    SettingDomainNewResponseDropDisposition = "ENCRYPTED"
	SettingDomainNewResponseDropDispositionExternal     SettingDomainNewResponseDropDisposition = "EXTERNAL"
	SettingDomainNewResponseDropDispositionUnknown      SettingDomainNewResponseDropDisposition = "UNKNOWN"
	SettingDomainNewResponseDropDispositionNone         SettingDomainNewResponseDropDisposition = "NONE"
)

func (r SettingDomainNewResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseDropDispositionMalicious, SettingDomainNewResponseDropDispositionMaliciousBec, SettingDomainNewResponseDropDispositionSuspicious, SettingDomainNewResponseDropDispositionSpoof, SettingDomainNewResponseDropDispositionSpam, SettingDomainNewResponseDropDispositionBulk, SettingDomainNewResponseDropDispositionEncrypted, SettingDomainNewResponseDropDispositionExternal, SettingDomainNewResponseDropDispositionUnknown, SettingDomainNewResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainNewResponseEmailsProcessed struct {
	Timestamp                    time.Time                                   `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                       `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                       `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainNewResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainNewResponseEmailsProcessedJSON contains the JSON metadata for the
// struct [SettingDomainNewResponseEmailsProcessed]
type settingDomainNewResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainNewResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainNewResponseFolder string

const (
	SettingDomainNewResponseFolderAllItems SettingDomainNewResponseFolder = "AllItems"
	SettingDomainNewResponseFolderInbox    SettingDomainNewResponseFolder = "Inbox"
)

func (r SettingDomainNewResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseFolderAllItems, SettingDomainNewResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainNewResponseInboxProvider string

const (
	SettingDomainNewResponseInboxProviderMicrosoft SettingDomainNewResponseInboxProvider = "Microsoft"
	SettingDomainNewResponseInboxProviderGoogle    SettingDomainNewResponseInboxProvider = "Google"
)

func (r SettingDomainNewResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseInboxProviderMicrosoft, SettingDomainNewResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainNewResponseRegion string

const (
	SettingDomainNewResponseRegionGlobal SettingDomainNewResponseRegion = "GLOBAL"
	SettingDomainNewResponseRegionAu     SettingDomainNewResponseRegion = "AU"
	SettingDomainNewResponseRegionDe     SettingDomainNewResponseRegion = "DE"
	SettingDomainNewResponseRegionIn     SettingDomainNewResponseRegion = "IN"
	SettingDomainNewResponseRegionUs     SettingDomainNewResponseRegion = "US"
)

func (r SettingDomainNewResponseRegion) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseRegionGlobal, SettingDomainNewResponseRegionAu, SettingDomainNewResponseRegionDe, SettingDomainNewResponseRegionIn, SettingDomainNewResponseRegionUs:
		return true
	}
	return false
}

type SettingDomainNewResponseSPFStatus string

const (
	SettingDomainNewResponseSPFStatusNone    SettingDomainNewResponseSPFStatus = "none"
	SettingDomainNewResponseSPFStatusGood    SettingDomainNewResponseSPFStatus = "good"
	SettingDomainNewResponseSPFStatusNeutral SettingDomainNewResponseSPFStatus = "neutral"
	SettingDomainNewResponseSPFStatusOpen    SettingDomainNewResponseSPFStatus = "open"
	SettingDomainNewResponseSPFStatusInvalid SettingDomainNewResponseSPFStatus = "invalid"
)

func (r SettingDomainNewResponseSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseSPFStatusNone, SettingDomainNewResponseSPFStatusGood, SettingDomainNewResponseSPFStatusNeutral, SettingDomainNewResponseSPFStatusOpen, SettingDomainNewResponseSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainNewResponseStatus string

const (
	SettingDomainNewResponseStatusPending SettingDomainNewResponseStatus = "PENDING"
	SettingDomainNewResponseStatusActive  SettingDomainNewResponseStatus = "ACTIVE"
	SettingDomainNewResponseStatusFailed  SettingDomainNewResponseStatus = "FAILED"
	SettingDomainNewResponseStatusTimeout SettingDomainNewResponseStatus = "TIMEOUT"
)

func (r SettingDomainNewResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseStatusPending, SettingDomainNewResponseStatusActive, SettingDomainNewResponseStatusFailed, SettingDomainNewResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainUpdateResponse struct {
	// Domain identifier.
	ID                   string                                           `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainUpdateResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainUpdateResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                        `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainUpdateResponseDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                           `json:"domain"`
	DropDispositions     []SettingDomainUpdateResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainUpdateResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainUpdateResponseFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainUpdateResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                           `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                         `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                            `json:"last_modified" format:"date-time"`
	LookbackHops       int64                                `json:"lookback_hops"`
	ModifiedAt         time.Time                            `json:"modified_at" format:"date-time"`
	O365TenantID       string                               `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainUpdateResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                                 `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                                 `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainUpdateResponseSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainUpdateResponseStatus    `json:"status" api:"nullable"`
	Transport          string                               `json:"transport"`
	JSON               settingDomainUpdateResponseJSON      `json:"-"`
}

// settingDomainUpdateResponseJSON contains the JSON metadata for the struct
// [SettingDomainUpdateResponse]
type settingDomainUpdateResponseJSON struct {
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

func (r *SettingDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainUpdateResponseAllowedDeliveryMode string

const (
	SettingDomainUpdateResponseAllowedDeliveryModeDirect    SettingDomainUpdateResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainUpdateResponseAllowedDeliveryModeBcc       SettingDomainUpdateResponseAllowedDeliveryMode = "BCC"
	SettingDomainUpdateResponseAllowedDeliveryModeJournal   SettingDomainUpdateResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainUpdateResponseAllowedDeliveryModeAPI       SettingDomainUpdateResponseAllowedDeliveryMode = "API"
	SettingDomainUpdateResponseAllowedDeliveryModeRetroScan SettingDomainUpdateResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainUpdateResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseAllowedDeliveryModeDirect, SettingDomainUpdateResponseAllowedDeliveryModeBcc, SettingDomainUpdateResponseAllowedDeliveryModeJournal, SettingDomainUpdateResponseAllowedDeliveryModeAPI, SettingDomainUpdateResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainUpdateResponseAuthorization struct {
	Authorized    bool                                         `json:"authorized" api:"required"`
	Timestamp     time.Time                                    `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                       `json:"status_message" api:"nullable"`
	JSON          settingDomainUpdateResponseAuthorizationJSON `json:"-"`
}

// settingDomainUpdateResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainUpdateResponseAuthorization]
type settingDomainUpdateResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainUpdateResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainUpdateResponseDMARCStatus string

const (
	SettingDomainUpdateResponseDMARCStatusNone    SettingDomainUpdateResponseDMARCStatus = "none"
	SettingDomainUpdateResponseDMARCStatusGood    SettingDomainUpdateResponseDMARCStatus = "good"
	SettingDomainUpdateResponseDMARCStatusInvalid SettingDomainUpdateResponseDMARCStatus = "invalid"
)

func (r SettingDomainUpdateResponseDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseDMARCStatusNone, SettingDomainUpdateResponseDMARCStatusGood, SettingDomainUpdateResponseDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainUpdateResponseDropDisposition string

const (
	SettingDomainUpdateResponseDropDispositionMalicious    SettingDomainUpdateResponseDropDisposition = "MALICIOUS"
	SettingDomainUpdateResponseDropDispositionMaliciousBec SettingDomainUpdateResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainUpdateResponseDropDispositionSuspicious   SettingDomainUpdateResponseDropDisposition = "SUSPICIOUS"
	SettingDomainUpdateResponseDropDispositionSpoof        SettingDomainUpdateResponseDropDisposition = "SPOOF"
	SettingDomainUpdateResponseDropDispositionSpam         SettingDomainUpdateResponseDropDisposition = "SPAM"
	SettingDomainUpdateResponseDropDispositionBulk         SettingDomainUpdateResponseDropDisposition = "BULK"
	SettingDomainUpdateResponseDropDispositionEncrypted    SettingDomainUpdateResponseDropDisposition = "ENCRYPTED"
	SettingDomainUpdateResponseDropDispositionExternal     SettingDomainUpdateResponseDropDisposition = "EXTERNAL"
	SettingDomainUpdateResponseDropDispositionUnknown      SettingDomainUpdateResponseDropDisposition = "UNKNOWN"
	SettingDomainUpdateResponseDropDispositionNone         SettingDomainUpdateResponseDropDisposition = "NONE"
)

func (r SettingDomainUpdateResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseDropDispositionMalicious, SettingDomainUpdateResponseDropDispositionMaliciousBec, SettingDomainUpdateResponseDropDispositionSuspicious, SettingDomainUpdateResponseDropDispositionSpoof, SettingDomainUpdateResponseDropDispositionSpam, SettingDomainUpdateResponseDropDispositionBulk, SettingDomainUpdateResponseDropDispositionEncrypted, SettingDomainUpdateResponseDropDispositionExternal, SettingDomainUpdateResponseDropDispositionUnknown, SettingDomainUpdateResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainUpdateResponseEmailsProcessed struct {
	Timestamp                    time.Time                                      `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                          `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                          `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainUpdateResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainUpdateResponseEmailsProcessedJSON contains the JSON metadata for
// the struct [SettingDomainUpdateResponseEmailsProcessed]
type settingDomainUpdateResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainUpdateResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainUpdateResponseFolder string

const (
	SettingDomainUpdateResponseFolderAllItems SettingDomainUpdateResponseFolder = "AllItems"
	SettingDomainUpdateResponseFolderInbox    SettingDomainUpdateResponseFolder = "Inbox"
)

func (r SettingDomainUpdateResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseFolderAllItems, SettingDomainUpdateResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainUpdateResponseInboxProvider string

const (
	SettingDomainUpdateResponseInboxProviderMicrosoft SettingDomainUpdateResponseInboxProvider = "Microsoft"
	SettingDomainUpdateResponseInboxProviderGoogle    SettingDomainUpdateResponseInboxProvider = "Google"
)

func (r SettingDomainUpdateResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseInboxProviderMicrosoft, SettingDomainUpdateResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainUpdateResponseRegion string

const (
	SettingDomainUpdateResponseRegionGlobal SettingDomainUpdateResponseRegion = "GLOBAL"
	SettingDomainUpdateResponseRegionAu     SettingDomainUpdateResponseRegion = "AU"
	SettingDomainUpdateResponseRegionDe     SettingDomainUpdateResponseRegion = "DE"
	SettingDomainUpdateResponseRegionIn     SettingDomainUpdateResponseRegion = "IN"
	SettingDomainUpdateResponseRegionUs     SettingDomainUpdateResponseRegion = "US"
)

func (r SettingDomainUpdateResponseRegion) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseRegionGlobal, SettingDomainUpdateResponseRegionAu, SettingDomainUpdateResponseRegionDe, SettingDomainUpdateResponseRegionIn, SettingDomainUpdateResponseRegionUs:
		return true
	}
	return false
}

type SettingDomainUpdateResponseSPFStatus string

const (
	SettingDomainUpdateResponseSPFStatusNone    SettingDomainUpdateResponseSPFStatus = "none"
	SettingDomainUpdateResponseSPFStatusGood    SettingDomainUpdateResponseSPFStatus = "good"
	SettingDomainUpdateResponseSPFStatusNeutral SettingDomainUpdateResponseSPFStatus = "neutral"
	SettingDomainUpdateResponseSPFStatusOpen    SettingDomainUpdateResponseSPFStatus = "open"
	SettingDomainUpdateResponseSPFStatusInvalid SettingDomainUpdateResponseSPFStatus = "invalid"
)

func (r SettingDomainUpdateResponseSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseSPFStatusNone, SettingDomainUpdateResponseSPFStatusGood, SettingDomainUpdateResponseSPFStatusNeutral, SettingDomainUpdateResponseSPFStatusOpen, SettingDomainUpdateResponseSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainUpdateResponseStatus string

const (
	SettingDomainUpdateResponseStatusPending SettingDomainUpdateResponseStatus = "PENDING"
	SettingDomainUpdateResponseStatusActive  SettingDomainUpdateResponseStatus = "ACTIVE"
	SettingDomainUpdateResponseStatusFailed  SettingDomainUpdateResponseStatus = "FAILED"
	SettingDomainUpdateResponseStatusTimeout SettingDomainUpdateResponseStatus = "TIMEOUT"
)

func (r SettingDomainUpdateResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseStatusPending, SettingDomainUpdateResponseStatusActive, SettingDomainUpdateResponseStatusFailed, SettingDomainUpdateResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainListResponse struct {
	// Domain identifier.
	ID                   string                                         `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainListResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainListResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                      `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainListResponseDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                         `json:"domain"`
	DropDispositions     []SettingDomainListResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainListResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainListResponseFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainListResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                         `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                       `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                          `json:"last_modified" format:"date-time"`
	LookbackHops       int64                              `json:"lookback_hops"`
	ModifiedAt         time.Time                          `json:"modified_at" format:"date-time"`
	O365TenantID       string                             `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainListResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                               `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                               `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainListResponseSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainListResponseStatus    `json:"status" api:"nullable"`
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
	SettingDomainListResponseStatusPending SettingDomainListResponseStatus = "PENDING"
	SettingDomainListResponseStatusActive  SettingDomainListResponseStatus = "ACTIVE"
	SettingDomainListResponseStatusFailed  SettingDomainListResponseStatus = "FAILED"
	SettingDomainListResponseStatusTimeout SettingDomainListResponseStatus = "TIMEOUT"
)

func (r SettingDomainListResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainListResponseStatusPending, SettingDomainListResponseStatusActive, SettingDomainListResponseStatusFailed, SettingDomainListResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainDeleteResponse struct {
	// Domain identifier.
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

type SettingDomainBatchResponse struct {
	Deletes []SettingDomainBatchResponseDelete `json:"deletes" api:"required"`
	Patches []SettingDomainBatchResponsePatch  `json:"patches" api:"required"`
	Posts   []SettingDomainBatchResponsePost   `json:"posts" api:"required"`
	Puts    []SettingDomainBatchResponsePut    `json:"puts" api:"required"`
	JSON    settingDomainBatchResponseJSON     `json:"-"`
}

// settingDomainBatchResponseJSON contains the JSON metadata for the struct
// [SettingDomainBatchResponse]
type settingDomainBatchResponseJSON struct {
	Deletes     apijson.Field
	Patches     apijson.Field
	Posts       apijson.Field
	Puts        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponseDelete struct {
	// Domain identifier.
	ID   string                               `json:"id" api:"required" format:"uuid"`
	JSON settingDomainBatchResponseDeleteJSON `json:"-"`
}

// settingDomainBatchResponseDeleteJSON contains the JSON metadata for the struct
// [SettingDomainBatchResponseDelete]
type settingDomainBatchResponseDeleteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainBatchResponseDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponseDeleteJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePatch struct {
	// Domain identifier.
	ID                   string                                                 `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainBatchResponsePatchesAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainBatchResponsePatchesAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                              `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainBatchResponsePatchesDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                                 `json:"domain"`
	DropDispositions     []SettingDomainBatchResponsePatchesDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainBatchResponsePatchesEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainBatchResponsePatchesFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainBatchResponsePatchesInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                                 `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                               `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                                  `json:"last_modified" format:"date-time"`
	LookbackHops       int64                                      `json:"lookback_hops"`
	ModifiedAt         time.Time                                  `json:"modified_at" format:"date-time"`
	O365TenantID       string                                     `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainBatchResponsePatchesRegion  `json:"regions"`
	RequireTLSInbound  bool                                       `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                                       `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainBatchResponsePatchesSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainBatchResponsePatchesStatus    `json:"status" api:"nullable"`
	Transport          string                                     `json:"transport"`
	JSON               settingDomainBatchResponsePatchJSON        `json:"-"`
}

// settingDomainBatchResponsePatchJSON contains the JSON metadata for the struct
// [SettingDomainBatchResponsePatch]
type settingDomainBatchResponsePatchJSON struct {
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

func (r *SettingDomainBatchResponsePatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePatchJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePatchesAllowedDeliveryMode string

const (
	SettingDomainBatchResponsePatchesAllowedDeliveryModeDirect    SettingDomainBatchResponsePatchesAllowedDeliveryMode = "DIRECT"
	SettingDomainBatchResponsePatchesAllowedDeliveryModeBcc       SettingDomainBatchResponsePatchesAllowedDeliveryMode = "BCC"
	SettingDomainBatchResponsePatchesAllowedDeliveryModeJournal   SettingDomainBatchResponsePatchesAllowedDeliveryMode = "JOURNAL"
	SettingDomainBatchResponsePatchesAllowedDeliveryModeAPI       SettingDomainBatchResponsePatchesAllowedDeliveryMode = "API"
	SettingDomainBatchResponsePatchesAllowedDeliveryModeRetroScan SettingDomainBatchResponsePatchesAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainBatchResponsePatchesAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesAllowedDeliveryModeDirect, SettingDomainBatchResponsePatchesAllowedDeliveryModeBcc, SettingDomainBatchResponsePatchesAllowedDeliveryModeJournal, SettingDomainBatchResponsePatchesAllowedDeliveryModeAPI, SettingDomainBatchResponsePatchesAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainBatchResponsePatchesAuthorization struct {
	Authorized    bool                                               `json:"authorized" api:"required"`
	Timestamp     time.Time                                          `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                             `json:"status_message" api:"nullable"`
	JSON          settingDomainBatchResponsePatchesAuthorizationJSON `json:"-"`
}

// settingDomainBatchResponsePatchesAuthorizationJSON contains the JSON metadata
// for the struct [SettingDomainBatchResponsePatchesAuthorization]
type settingDomainBatchResponsePatchesAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainBatchResponsePatchesAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePatchesAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePatchesDMARCStatus string

const (
	SettingDomainBatchResponsePatchesDMARCStatusNone    SettingDomainBatchResponsePatchesDMARCStatus = "none"
	SettingDomainBatchResponsePatchesDMARCStatusGood    SettingDomainBatchResponsePatchesDMARCStatus = "good"
	SettingDomainBatchResponsePatchesDMARCStatusInvalid SettingDomainBatchResponsePatchesDMARCStatus = "invalid"
)

func (r SettingDomainBatchResponsePatchesDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesDMARCStatusNone, SettingDomainBatchResponsePatchesDMARCStatusGood, SettingDomainBatchResponsePatchesDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainBatchResponsePatchesDropDisposition string

const (
	SettingDomainBatchResponsePatchesDropDispositionMalicious    SettingDomainBatchResponsePatchesDropDisposition = "MALICIOUS"
	SettingDomainBatchResponsePatchesDropDispositionMaliciousBec SettingDomainBatchResponsePatchesDropDisposition = "MALICIOUS-BEC"
	SettingDomainBatchResponsePatchesDropDispositionSuspicious   SettingDomainBatchResponsePatchesDropDisposition = "SUSPICIOUS"
	SettingDomainBatchResponsePatchesDropDispositionSpoof        SettingDomainBatchResponsePatchesDropDisposition = "SPOOF"
	SettingDomainBatchResponsePatchesDropDispositionSpam         SettingDomainBatchResponsePatchesDropDisposition = "SPAM"
	SettingDomainBatchResponsePatchesDropDispositionBulk         SettingDomainBatchResponsePatchesDropDisposition = "BULK"
	SettingDomainBatchResponsePatchesDropDispositionEncrypted    SettingDomainBatchResponsePatchesDropDisposition = "ENCRYPTED"
	SettingDomainBatchResponsePatchesDropDispositionExternal     SettingDomainBatchResponsePatchesDropDisposition = "EXTERNAL"
	SettingDomainBatchResponsePatchesDropDispositionUnknown      SettingDomainBatchResponsePatchesDropDisposition = "UNKNOWN"
	SettingDomainBatchResponsePatchesDropDispositionNone         SettingDomainBatchResponsePatchesDropDisposition = "NONE"
)

func (r SettingDomainBatchResponsePatchesDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesDropDispositionMalicious, SettingDomainBatchResponsePatchesDropDispositionMaliciousBec, SettingDomainBatchResponsePatchesDropDispositionSuspicious, SettingDomainBatchResponsePatchesDropDispositionSpoof, SettingDomainBatchResponsePatchesDropDispositionSpam, SettingDomainBatchResponsePatchesDropDispositionBulk, SettingDomainBatchResponsePatchesDropDispositionEncrypted, SettingDomainBatchResponsePatchesDropDispositionExternal, SettingDomainBatchResponsePatchesDropDispositionUnknown, SettingDomainBatchResponsePatchesDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainBatchResponsePatchesEmailsProcessed struct {
	Timestamp                    time.Time                                            `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                                `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                                `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainBatchResponsePatchesEmailsProcessedJSON `json:"-"`
}

// settingDomainBatchResponsePatchesEmailsProcessedJSON contains the JSON metadata
// for the struct [SettingDomainBatchResponsePatchesEmailsProcessed]
type settingDomainBatchResponsePatchesEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainBatchResponsePatchesEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePatchesEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePatchesFolder string

const (
	SettingDomainBatchResponsePatchesFolderAllItems SettingDomainBatchResponsePatchesFolder = "AllItems"
	SettingDomainBatchResponsePatchesFolderInbox    SettingDomainBatchResponsePatchesFolder = "Inbox"
)

func (r SettingDomainBatchResponsePatchesFolder) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesFolderAllItems, SettingDomainBatchResponsePatchesFolderInbox:
		return true
	}
	return false
}

type SettingDomainBatchResponsePatchesInboxProvider string

const (
	SettingDomainBatchResponsePatchesInboxProviderMicrosoft SettingDomainBatchResponsePatchesInboxProvider = "Microsoft"
	SettingDomainBatchResponsePatchesInboxProviderGoogle    SettingDomainBatchResponsePatchesInboxProvider = "Google"
)

func (r SettingDomainBatchResponsePatchesInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesInboxProviderMicrosoft, SettingDomainBatchResponsePatchesInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainBatchResponsePatchesRegion string

const (
	SettingDomainBatchResponsePatchesRegionGlobal SettingDomainBatchResponsePatchesRegion = "GLOBAL"
	SettingDomainBatchResponsePatchesRegionAu     SettingDomainBatchResponsePatchesRegion = "AU"
	SettingDomainBatchResponsePatchesRegionDe     SettingDomainBatchResponsePatchesRegion = "DE"
	SettingDomainBatchResponsePatchesRegionIn     SettingDomainBatchResponsePatchesRegion = "IN"
	SettingDomainBatchResponsePatchesRegionUs     SettingDomainBatchResponsePatchesRegion = "US"
)

func (r SettingDomainBatchResponsePatchesRegion) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesRegionGlobal, SettingDomainBatchResponsePatchesRegionAu, SettingDomainBatchResponsePatchesRegionDe, SettingDomainBatchResponsePatchesRegionIn, SettingDomainBatchResponsePatchesRegionUs:
		return true
	}
	return false
}

type SettingDomainBatchResponsePatchesSPFStatus string

const (
	SettingDomainBatchResponsePatchesSPFStatusNone    SettingDomainBatchResponsePatchesSPFStatus = "none"
	SettingDomainBatchResponsePatchesSPFStatusGood    SettingDomainBatchResponsePatchesSPFStatus = "good"
	SettingDomainBatchResponsePatchesSPFStatusNeutral SettingDomainBatchResponsePatchesSPFStatus = "neutral"
	SettingDomainBatchResponsePatchesSPFStatusOpen    SettingDomainBatchResponsePatchesSPFStatus = "open"
	SettingDomainBatchResponsePatchesSPFStatusInvalid SettingDomainBatchResponsePatchesSPFStatus = "invalid"
)

func (r SettingDomainBatchResponsePatchesSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesSPFStatusNone, SettingDomainBatchResponsePatchesSPFStatusGood, SettingDomainBatchResponsePatchesSPFStatusNeutral, SettingDomainBatchResponsePatchesSPFStatusOpen, SettingDomainBatchResponsePatchesSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainBatchResponsePatchesStatus string

const (
	SettingDomainBatchResponsePatchesStatusPending SettingDomainBatchResponsePatchesStatus = "PENDING"
	SettingDomainBatchResponsePatchesStatusActive  SettingDomainBatchResponsePatchesStatus = "ACTIVE"
	SettingDomainBatchResponsePatchesStatusFailed  SettingDomainBatchResponsePatchesStatus = "FAILED"
	SettingDomainBatchResponsePatchesStatusTimeout SettingDomainBatchResponsePatchesStatus = "TIMEOUT"
)

func (r SettingDomainBatchResponsePatchesStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePatchesStatusPending, SettingDomainBatchResponsePatchesStatusActive, SettingDomainBatchResponsePatchesStatusFailed, SettingDomainBatchResponsePatchesStatusTimeout:
		return true
	}
	return false
}

type SettingDomainBatchResponsePost struct {
	// Domain identifier.
	ID                   string                                               `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainBatchResponsePostsAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainBatchResponsePostsAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                            `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainBatchResponsePostsDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                               `json:"domain"`
	DropDispositions     []SettingDomainBatchResponsePostsDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainBatchResponsePostsEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainBatchResponsePostsFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainBatchResponsePostsInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                               `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                             `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                                `json:"last_modified" format:"date-time"`
	LookbackHops       int64                                    `json:"lookback_hops"`
	ModifiedAt         time.Time                                `json:"modified_at" format:"date-time"`
	O365TenantID       string                                   `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainBatchResponsePostsRegion  `json:"regions"`
	RequireTLSInbound  bool                                     `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                                     `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainBatchResponsePostsSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainBatchResponsePostsStatus    `json:"status" api:"nullable"`
	Transport          string                                   `json:"transport"`
	JSON               settingDomainBatchResponsePostJSON       `json:"-"`
}

// settingDomainBatchResponsePostJSON contains the JSON metadata for the struct
// [SettingDomainBatchResponsePost]
type settingDomainBatchResponsePostJSON struct {
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

func (r *SettingDomainBatchResponsePost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePostJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePostsAllowedDeliveryMode string

const (
	SettingDomainBatchResponsePostsAllowedDeliveryModeDirect    SettingDomainBatchResponsePostsAllowedDeliveryMode = "DIRECT"
	SettingDomainBatchResponsePostsAllowedDeliveryModeBcc       SettingDomainBatchResponsePostsAllowedDeliveryMode = "BCC"
	SettingDomainBatchResponsePostsAllowedDeliveryModeJournal   SettingDomainBatchResponsePostsAllowedDeliveryMode = "JOURNAL"
	SettingDomainBatchResponsePostsAllowedDeliveryModeAPI       SettingDomainBatchResponsePostsAllowedDeliveryMode = "API"
	SettingDomainBatchResponsePostsAllowedDeliveryModeRetroScan SettingDomainBatchResponsePostsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainBatchResponsePostsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsAllowedDeliveryModeDirect, SettingDomainBatchResponsePostsAllowedDeliveryModeBcc, SettingDomainBatchResponsePostsAllowedDeliveryModeJournal, SettingDomainBatchResponsePostsAllowedDeliveryModeAPI, SettingDomainBatchResponsePostsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainBatchResponsePostsAuthorization struct {
	Authorized    bool                                             `json:"authorized" api:"required"`
	Timestamp     time.Time                                        `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                           `json:"status_message" api:"nullable"`
	JSON          settingDomainBatchResponsePostsAuthorizationJSON `json:"-"`
}

// settingDomainBatchResponsePostsAuthorizationJSON contains the JSON metadata for
// the struct [SettingDomainBatchResponsePostsAuthorization]
type settingDomainBatchResponsePostsAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainBatchResponsePostsAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePostsAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePostsDMARCStatus string

const (
	SettingDomainBatchResponsePostsDMARCStatusNone    SettingDomainBatchResponsePostsDMARCStatus = "none"
	SettingDomainBatchResponsePostsDMARCStatusGood    SettingDomainBatchResponsePostsDMARCStatus = "good"
	SettingDomainBatchResponsePostsDMARCStatusInvalid SettingDomainBatchResponsePostsDMARCStatus = "invalid"
)

func (r SettingDomainBatchResponsePostsDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsDMARCStatusNone, SettingDomainBatchResponsePostsDMARCStatusGood, SettingDomainBatchResponsePostsDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainBatchResponsePostsDropDisposition string

const (
	SettingDomainBatchResponsePostsDropDispositionMalicious    SettingDomainBatchResponsePostsDropDisposition = "MALICIOUS"
	SettingDomainBatchResponsePostsDropDispositionMaliciousBec SettingDomainBatchResponsePostsDropDisposition = "MALICIOUS-BEC"
	SettingDomainBatchResponsePostsDropDispositionSuspicious   SettingDomainBatchResponsePostsDropDisposition = "SUSPICIOUS"
	SettingDomainBatchResponsePostsDropDispositionSpoof        SettingDomainBatchResponsePostsDropDisposition = "SPOOF"
	SettingDomainBatchResponsePostsDropDispositionSpam         SettingDomainBatchResponsePostsDropDisposition = "SPAM"
	SettingDomainBatchResponsePostsDropDispositionBulk         SettingDomainBatchResponsePostsDropDisposition = "BULK"
	SettingDomainBatchResponsePostsDropDispositionEncrypted    SettingDomainBatchResponsePostsDropDisposition = "ENCRYPTED"
	SettingDomainBatchResponsePostsDropDispositionExternal     SettingDomainBatchResponsePostsDropDisposition = "EXTERNAL"
	SettingDomainBatchResponsePostsDropDispositionUnknown      SettingDomainBatchResponsePostsDropDisposition = "UNKNOWN"
	SettingDomainBatchResponsePostsDropDispositionNone         SettingDomainBatchResponsePostsDropDisposition = "NONE"
)

func (r SettingDomainBatchResponsePostsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsDropDispositionMalicious, SettingDomainBatchResponsePostsDropDispositionMaliciousBec, SettingDomainBatchResponsePostsDropDispositionSuspicious, SettingDomainBatchResponsePostsDropDispositionSpoof, SettingDomainBatchResponsePostsDropDispositionSpam, SettingDomainBatchResponsePostsDropDispositionBulk, SettingDomainBatchResponsePostsDropDispositionEncrypted, SettingDomainBatchResponsePostsDropDispositionExternal, SettingDomainBatchResponsePostsDropDispositionUnknown, SettingDomainBatchResponsePostsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainBatchResponsePostsEmailsProcessed struct {
	Timestamp                    time.Time                                          `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                              `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                              `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainBatchResponsePostsEmailsProcessedJSON `json:"-"`
}

// settingDomainBatchResponsePostsEmailsProcessedJSON contains the JSON metadata
// for the struct [SettingDomainBatchResponsePostsEmailsProcessed]
type settingDomainBatchResponsePostsEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainBatchResponsePostsEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePostsEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePostsFolder string

const (
	SettingDomainBatchResponsePostsFolderAllItems SettingDomainBatchResponsePostsFolder = "AllItems"
	SettingDomainBatchResponsePostsFolderInbox    SettingDomainBatchResponsePostsFolder = "Inbox"
)

func (r SettingDomainBatchResponsePostsFolder) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsFolderAllItems, SettingDomainBatchResponsePostsFolderInbox:
		return true
	}
	return false
}

type SettingDomainBatchResponsePostsInboxProvider string

const (
	SettingDomainBatchResponsePostsInboxProviderMicrosoft SettingDomainBatchResponsePostsInboxProvider = "Microsoft"
	SettingDomainBatchResponsePostsInboxProviderGoogle    SettingDomainBatchResponsePostsInboxProvider = "Google"
)

func (r SettingDomainBatchResponsePostsInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsInboxProviderMicrosoft, SettingDomainBatchResponsePostsInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainBatchResponsePostsRegion string

const (
	SettingDomainBatchResponsePostsRegionGlobal SettingDomainBatchResponsePostsRegion = "GLOBAL"
	SettingDomainBatchResponsePostsRegionAu     SettingDomainBatchResponsePostsRegion = "AU"
	SettingDomainBatchResponsePostsRegionDe     SettingDomainBatchResponsePostsRegion = "DE"
	SettingDomainBatchResponsePostsRegionIn     SettingDomainBatchResponsePostsRegion = "IN"
	SettingDomainBatchResponsePostsRegionUs     SettingDomainBatchResponsePostsRegion = "US"
)

func (r SettingDomainBatchResponsePostsRegion) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsRegionGlobal, SettingDomainBatchResponsePostsRegionAu, SettingDomainBatchResponsePostsRegionDe, SettingDomainBatchResponsePostsRegionIn, SettingDomainBatchResponsePostsRegionUs:
		return true
	}
	return false
}

type SettingDomainBatchResponsePostsSPFStatus string

const (
	SettingDomainBatchResponsePostsSPFStatusNone    SettingDomainBatchResponsePostsSPFStatus = "none"
	SettingDomainBatchResponsePostsSPFStatusGood    SettingDomainBatchResponsePostsSPFStatus = "good"
	SettingDomainBatchResponsePostsSPFStatusNeutral SettingDomainBatchResponsePostsSPFStatus = "neutral"
	SettingDomainBatchResponsePostsSPFStatusOpen    SettingDomainBatchResponsePostsSPFStatus = "open"
	SettingDomainBatchResponsePostsSPFStatusInvalid SettingDomainBatchResponsePostsSPFStatus = "invalid"
)

func (r SettingDomainBatchResponsePostsSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsSPFStatusNone, SettingDomainBatchResponsePostsSPFStatusGood, SettingDomainBatchResponsePostsSPFStatusNeutral, SettingDomainBatchResponsePostsSPFStatusOpen, SettingDomainBatchResponsePostsSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainBatchResponsePostsStatus string

const (
	SettingDomainBatchResponsePostsStatusPending SettingDomainBatchResponsePostsStatus = "PENDING"
	SettingDomainBatchResponsePostsStatusActive  SettingDomainBatchResponsePostsStatus = "ACTIVE"
	SettingDomainBatchResponsePostsStatusFailed  SettingDomainBatchResponsePostsStatus = "FAILED"
	SettingDomainBatchResponsePostsStatusTimeout SettingDomainBatchResponsePostsStatus = "TIMEOUT"
)

func (r SettingDomainBatchResponsePostsStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePostsStatusPending, SettingDomainBatchResponsePostsStatusActive, SettingDomainBatchResponsePostsStatusFailed, SettingDomainBatchResponsePostsStatusTimeout:
		return true
	}
	return false
}

type SettingDomainBatchResponsePut struct {
	// Domain identifier.
	ID                   string                                              `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainBatchResponsePutsAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainBatchResponsePutsAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                           `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainBatchResponsePutsDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                              `json:"domain"`
	DropDispositions     []SettingDomainBatchResponsePutsDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainBatchResponsePutsEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainBatchResponsePutsFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainBatchResponsePutsInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                              `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                            `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                               `json:"last_modified" format:"date-time"`
	LookbackHops       int64                                   `json:"lookback_hops"`
	ModifiedAt         time.Time                               `json:"modified_at" format:"date-time"`
	O365TenantID       string                                  `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainBatchResponsePutsRegion  `json:"regions"`
	RequireTLSInbound  bool                                    `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                                    `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainBatchResponsePutsSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainBatchResponsePutsStatus    `json:"status" api:"nullable"`
	Transport          string                                  `json:"transport"`
	JSON               settingDomainBatchResponsePutJSON       `json:"-"`
}

// settingDomainBatchResponsePutJSON contains the JSON metadata for the struct
// [SettingDomainBatchResponsePut]
type settingDomainBatchResponsePutJSON struct {
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

func (r *SettingDomainBatchResponsePut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePutJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePutsAllowedDeliveryMode string

const (
	SettingDomainBatchResponsePutsAllowedDeliveryModeDirect    SettingDomainBatchResponsePutsAllowedDeliveryMode = "DIRECT"
	SettingDomainBatchResponsePutsAllowedDeliveryModeBcc       SettingDomainBatchResponsePutsAllowedDeliveryMode = "BCC"
	SettingDomainBatchResponsePutsAllowedDeliveryModeJournal   SettingDomainBatchResponsePutsAllowedDeliveryMode = "JOURNAL"
	SettingDomainBatchResponsePutsAllowedDeliveryModeAPI       SettingDomainBatchResponsePutsAllowedDeliveryMode = "API"
	SettingDomainBatchResponsePutsAllowedDeliveryModeRetroScan SettingDomainBatchResponsePutsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainBatchResponsePutsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsAllowedDeliveryModeDirect, SettingDomainBatchResponsePutsAllowedDeliveryModeBcc, SettingDomainBatchResponsePutsAllowedDeliveryModeJournal, SettingDomainBatchResponsePutsAllowedDeliveryModeAPI, SettingDomainBatchResponsePutsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainBatchResponsePutsAuthorization struct {
	Authorized    bool                                            `json:"authorized" api:"required"`
	Timestamp     time.Time                                       `json:"timestamp" api:"required" format:"date-time"`
	StatusMessage string                                          `json:"status_message" api:"nullable"`
	JSON          settingDomainBatchResponsePutsAuthorizationJSON `json:"-"`
}

// settingDomainBatchResponsePutsAuthorizationJSON contains the JSON metadata for
// the struct [SettingDomainBatchResponsePutsAuthorization]
type settingDomainBatchResponsePutsAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainBatchResponsePutsAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePutsAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePutsDMARCStatus string

const (
	SettingDomainBatchResponsePutsDMARCStatusNone    SettingDomainBatchResponsePutsDMARCStatus = "none"
	SettingDomainBatchResponsePutsDMARCStatusGood    SettingDomainBatchResponsePutsDMARCStatus = "good"
	SettingDomainBatchResponsePutsDMARCStatusInvalid SettingDomainBatchResponsePutsDMARCStatus = "invalid"
)

func (r SettingDomainBatchResponsePutsDMARCStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsDMARCStatusNone, SettingDomainBatchResponsePutsDMARCStatusGood, SettingDomainBatchResponsePutsDMARCStatusInvalid:
		return true
	}
	return false
}

type SettingDomainBatchResponsePutsDropDisposition string

const (
	SettingDomainBatchResponsePutsDropDispositionMalicious    SettingDomainBatchResponsePutsDropDisposition = "MALICIOUS"
	SettingDomainBatchResponsePutsDropDispositionMaliciousBec SettingDomainBatchResponsePutsDropDisposition = "MALICIOUS-BEC"
	SettingDomainBatchResponsePutsDropDispositionSuspicious   SettingDomainBatchResponsePutsDropDisposition = "SUSPICIOUS"
	SettingDomainBatchResponsePutsDropDispositionSpoof        SettingDomainBatchResponsePutsDropDisposition = "SPOOF"
	SettingDomainBatchResponsePutsDropDispositionSpam         SettingDomainBatchResponsePutsDropDisposition = "SPAM"
	SettingDomainBatchResponsePutsDropDispositionBulk         SettingDomainBatchResponsePutsDropDisposition = "BULK"
	SettingDomainBatchResponsePutsDropDispositionEncrypted    SettingDomainBatchResponsePutsDropDisposition = "ENCRYPTED"
	SettingDomainBatchResponsePutsDropDispositionExternal     SettingDomainBatchResponsePutsDropDisposition = "EXTERNAL"
	SettingDomainBatchResponsePutsDropDispositionUnknown      SettingDomainBatchResponsePutsDropDisposition = "UNKNOWN"
	SettingDomainBatchResponsePutsDropDispositionNone         SettingDomainBatchResponsePutsDropDisposition = "NONE"
)

func (r SettingDomainBatchResponsePutsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsDropDispositionMalicious, SettingDomainBatchResponsePutsDropDispositionMaliciousBec, SettingDomainBatchResponsePutsDropDispositionSuspicious, SettingDomainBatchResponsePutsDropDispositionSpoof, SettingDomainBatchResponsePutsDropDispositionSpam, SettingDomainBatchResponsePutsDropDispositionBulk, SettingDomainBatchResponsePutsDropDispositionEncrypted, SettingDomainBatchResponsePutsDropDispositionExternal, SettingDomainBatchResponsePutsDropDispositionUnknown, SettingDomainBatchResponsePutsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainBatchResponsePutsEmailsProcessed struct {
	Timestamp                    time.Time                                         `json:"timestamp" api:"required" format:"date-time"`
	TotalEmailsProcessed         int64                                             `json:"total_emails_processed" api:"required"`
	TotalEmailsProcessedPrevious int64                                             `json:"total_emails_processed_previous" api:"required"`
	JSON                         settingDomainBatchResponsePutsEmailsProcessedJSON `json:"-"`
}

// settingDomainBatchResponsePutsEmailsProcessedJSON contains the JSON metadata for
// the struct [SettingDomainBatchResponsePutsEmailsProcessed]
type settingDomainBatchResponsePutsEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainBatchResponsePutsEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponsePutsEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponsePutsFolder string

const (
	SettingDomainBatchResponsePutsFolderAllItems SettingDomainBatchResponsePutsFolder = "AllItems"
	SettingDomainBatchResponsePutsFolderInbox    SettingDomainBatchResponsePutsFolder = "Inbox"
)

func (r SettingDomainBatchResponsePutsFolder) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsFolderAllItems, SettingDomainBatchResponsePutsFolderInbox:
		return true
	}
	return false
}

type SettingDomainBatchResponsePutsInboxProvider string

const (
	SettingDomainBatchResponsePutsInboxProviderMicrosoft SettingDomainBatchResponsePutsInboxProvider = "Microsoft"
	SettingDomainBatchResponsePutsInboxProviderGoogle    SettingDomainBatchResponsePutsInboxProvider = "Google"
)

func (r SettingDomainBatchResponsePutsInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsInboxProviderMicrosoft, SettingDomainBatchResponsePutsInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainBatchResponsePutsRegion string

const (
	SettingDomainBatchResponsePutsRegionGlobal SettingDomainBatchResponsePutsRegion = "GLOBAL"
	SettingDomainBatchResponsePutsRegionAu     SettingDomainBatchResponsePutsRegion = "AU"
	SettingDomainBatchResponsePutsRegionDe     SettingDomainBatchResponsePutsRegion = "DE"
	SettingDomainBatchResponsePutsRegionIn     SettingDomainBatchResponsePutsRegion = "IN"
	SettingDomainBatchResponsePutsRegionUs     SettingDomainBatchResponsePutsRegion = "US"
)

func (r SettingDomainBatchResponsePutsRegion) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsRegionGlobal, SettingDomainBatchResponsePutsRegionAu, SettingDomainBatchResponsePutsRegionDe, SettingDomainBatchResponsePutsRegionIn, SettingDomainBatchResponsePutsRegionUs:
		return true
	}
	return false
}

type SettingDomainBatchResponsePutsSPFStatus string

const (
	SettingDomainBatchResponsePutsSPFStatusNone    SettingDomainBatchResponsePutsSPFStatus = "none"
	SettingDomainBatchResponsePutsSPFStatusGood    SettingDomainBatchResponsePutsSPFStatus = "good"
	SettingDomainBatchResponsePutsSPFStatusNeutral SettingDomainBatchResponsePutsSPFStatus = "neutral"
	SettingDomainBatchResponsePutsSPFStatusOpen    SettingDomainBatchResponsePutsSPFStatus = "open"
	SettingDomainBatchResponsePutsSPFStatusInvalid SettingDomainBatchResponsePutsSPFStatus = "invalid"
)

func (r SettingDomainBatchResponsePutsSPFStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsSPFStatusNone, SettingDomainBatchResponsePutsSPFStatusGood, SettingDomainBatchResponsePutsSPFStatusNeutral, SettingDomainBatchResponsePutsSPFStatusOpen, SettingDomainBatchResponsePutsSPFStatusInvalid:
		return true
	}
	return false
}

type SettingDomainBatchResponsePutsStatus string

const (
	SettingDomainBatchResponsePutsStatusPending SettingDomainBatchResponsePutsStatus = "PENDING"
	SettingDomainBatchResponsePutsStatusActive  SettingDomainBatchResponsePutsStatus = "ACTIVE"
	SettingDomainBatchResponsePutsStatusFailed  SettingDomainBatchResponsePutsStatus = "FAILED"
	SettingDomainBatchResponsePutsStatusTimeout SettingDomainBatchResponsePutsStatus = "TIMEOUT"
)

func (r SettingDomainBatchResponsePutsStatus) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponsePutsStatusPending, SettingDomainBatchResponsePutsStatusActive, SettingDomainBatchResponsePutsStatusFailed, SettingDomainBatchResponsePutsStatusTimeout:
		return true
	}
	return false
}

type SettingDomainBulkDeleteResponse struct {
	// Domain identifier.
	ID   string                              `json:"id" api:"required" format:"uuid"`
	JSON settingDomainBulkDeleteResponseJSON `json:"-"`
}

// settingDomainBulkDeleteResponseJSON contains the JSON metadata for the struct
// [SettingDomainBulkDeleteResponse]
type settingDomainBulkDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponse struct {
	// Domain identifier.
	ID                   string                                         `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainEditResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainEditResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                      `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainEditResponseDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                         `json:"domain"`
	DropDispositions     []SettingDomainEditResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainEditResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainEditResponseFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainEditResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                         `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                       `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                          `json:"last_modified" format:"date-time"`
	LookbackHops       int64                              `json:"lookback_hops"`
	ModifiedAt         time.Time                          `json:"modified_at" format:"date-time"`
	O365TenantID       string                             `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainEditResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                               `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                               `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainEditResponseSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainEditResponseStatus    `json:"status" api:"nullable"`
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
	SettingDomainEditResponseStatusPending SettingDomainEditResponseStatus = "PENDING"
	SettingDomainEditResponseStatusActive  SettingDomainEditResponseStatus = "ACTIVE"
	SettingDomainEditResponseStatusFailed  SettingDomainEditResponseStatus = "FAILED"
	SettingDomainEditResponseStatusTimeout SettingDomainEditResponseStatus = "TIMEOUT"
)

func (r SettingDomainEditResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseStatusPending, SettingDomainEditResponseStatusActive, SettingDomainEditResponseStatusFailed, SettingDomainEditResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainGetResponse struct {
	// Domain identifier.
	ID                   string                                        `json:"id" format:"uuid"`
	AllowedDeliveryModes []SettingDomainGetResponseAllowedDeliveryMode `json:"allowed_delivery_modes"`
	Authorization        SettingDomainGetResponseAuthorization         `json:"authorization"`
	CreatedAt            time.Time                                     `json:"created_at" format:"date-time"`
	DMARCStatus          SettingDomainGetResponseDMARCStatus           `json:"dmarc_status" api:"nullable"`
	Domain               string                                        `json:"domain"`
	DropDispositions     []SettingDomainGetResponseDropDisposition     `json:"drop_dispositions"`
	EmailsProcessed      SettingDomainGetResponseEmailsProcessed       `json:"emails_processed"`
	Folder               SettingDomainGetResponseFolder                `json:"folder" api:"nullable"`
	InboxProvider        SettingDomainGetResponseInboxProvider         `json:"inbox_provider" api:"nullable"`
	IntegrationID        string                                        `json:"integration_id" api:"nullable" format:"uuid"`
	IPRestrictions       []string                                      `json:"ip_restrictions"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified       time.Time                         `json:"last_modified" format:"date-time"`
	LookbackHops       int64                             `json:"lookback_hops"`
	ModifiedAt         time.Time                         `json:"modified_at" format:"date-time"`
	O365TenantID       string                            `json:"o365_tenant_id" api:"nullable"`
	Regions            []SettingDomainGetResponseRegion  `json:"regions"`
	RequireTLSInbound  bool                              `json:"require_tls_inbound" api:"nullable"`
	RequireTLSOutbound bool                              `json:"require_tls_outbound" api:"nullable"`
	SPFStatus          SettingDomainGetResponseSPFStatus `json:"spf_status" api:"nullable"`
	Status             SettingDomainGetResponseStatus    `json:"status" api:"nullable"`
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
	SettingDomainGetResponseStatusPending SettingDomainGetResponseStatus = "PENDING"
	SettingDomainGetResponseStatusActive  SettingDomainGetResponseStatus = "ACTIVE"
	SettingDomainGetResponseStatusFailed  SettingDomainGetResponseStatus = "FAILED"
	SettingDomainGetResponseStatusTimeout SettingDomainGetResponseStatus = "TIMEOUT"
)

func (r SettingDomainGetResponseStatus) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseStatusPending, SettingDomainGetResponseStatusActive, SettingDomainGetResponseStatusFailed, SettingDomainGetResponseStatusTimeout:
		return true
	}
	return false
}

type SettingDomainNewParams struct {
	// Identifier.
	AccountID            param.Field[string]                                      `path:"account_id" api:"required"`
	AllowedDeliveryModes param.Field[[]SettingDomainNewParamsAllowedDeliveryMode] `json:"allowed_delivery_modes" api:"required"`
	Domain               param.Field[string]                                      `json:"domain" api:"required"`
	DropDispositions     param.Field[[]SettingDomainNewParamsDropDisposition]     `json:"drop_dispositions" api:"required"`
	IPRestrictions       param.Field[[]string]                                    `json:"ip_restrictions" api:"required"`
	Regions              param.Field[[]SettingDomainNewParamsRegion]              `json:"regions" api:"required"`
	Folder               param.Field[SettingDomainNewParamsFolder]                `json:"folder"`
	IntegrationID        param.Field[string]                                      `json:"integration_id" format:"uuid"`
	LookbackHops         param.Field[int64]                                       `json:"lookback_hops"`
	RequireTLSInbound    param.Field[bool]                                        `json:"require_tls_inbound"`
	RequireTLSOutbound   param.Field[bool]                                        `json:"require_tls_outbound"`
	Transport            param.Field[string]                                      `json:"transport"`
}

func (r SettingDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainNewParamsAllowedDeliveryMode string

const (
	SettingDomainNewParamsAllowedDeliveryModeDirect    SettingDomainNewParamsAllowedDeliveryMode = "DIRECT"
	SettingDomainNewParamsAllowedDeliveryModeBcc       SettingDomainNewParamsAllowedDeliveryMode = "BCC"
	SettingDomainNewParamsAllowedDeliveryModeJournal   SettingDomainNewParamsAllowedDeliveryMode = "JOURNAL"
	SettingDomainNewParamsAllowedDeliveryModeAPI       SettingDomainNewParamsAllowedDeliveryMode = "API"
	SettingDomainNewParamsAllowedDeliveryModeRetroScan SettingDomainNewParamsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainNewParamsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainNewParamsAllowedDeliveryModeDirect, SettingDomainNewParamsAllowedDeliveryModeBcc, SettingDomainNewParamsAllowedDeliveryModeJournal, SettingDomainNewParamsAllowedDeliveryModeAPI, SettingDomainNewParamsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainNewParamsDropDisposition string

const (
	SettingDomainNewParamsDropDispositionMalicious    SettingDomainNewParamsDropDisposition = "MALICIOUS"
	SettingDomainNewParamsDropDispositionMaliciousBec SettingDomainNewParamsDropDisposition = "MALICIOUS-BEC"
	SettingDomainNewParamsDropDispositionSuspicious   SettingDomainNewParamsDropDisposition = "SUSPICIOUS"
	SettingDomainNewParamsDropDispositionSpoof        SettingDomainNewParamsDropDisposition = "SPOOF"
	SettingDomainNewParamsDropDispositionSpam         SettingDomainNewParamsDropDisposition = "SPAM"
	SettingDomainNewParamsDropDispositionBulk         SettingDomainNewParamsDropDisposition = "BULK"
	SettingDomainNewParamsDropDispositionEncrypted    SettingDomainNewParamsDropDisposition = "ENCRYPTED"
	SettingDomainNewParamsDropDispositionExternal     SettingDomainNewParamsDropDisposition = "EXTERNAL"
	SettingDomainNewParamsDropDispositionUnknown      SettingDomainNewParamsDropDisposition = "UNKNOWN"
	SettingDomainNewParamsDropDispositionNone         SettingDomainNewParamsDropDisposition = "NONE"
)

func (r SettingDomainNewParamsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainNewParamsDropDispositionMalicious, SettingDomainNewParamsDropDispositionMaliciousBec, SettingDomainNewParamsDropDispositionSuspicious, SettingDomainNewParamsDropDispositionSpoof, SettingDomainNewParamsDropDispositionSpam, SettingDomainNewParamsDropDispositionBulk, SettingDomainNewParamsDropDispositionEncrypted, SettingDomainNewParamsDropDispositionExternal, SettingDomainNewParamsDropDispositionUnknown, SettingDomainNewParamsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainNewParamsRegion string

const (
	SettingDomainNewParamsRegionGlobal SettingDomainNewParamsRegion = "GLOBAL"
	SettingDomainNewParamsRegionAu     SettingDomainNewParamsRegion = "AU"
	SettingDomainNewParamsRegionDe     SettingDomainNewParamsRegion = "DE"
	SettingDomainNewParamsRegionIn     SettingDomainNewParamsRegion = "IN"
	SettingDomainNewParamsRegionUs     SettingDomainNewParamsRegion = "US"
)

func (r SettingDomainNewParamsRegion) IsKnown() bool {
	switch r {
	case SettingDomainNewParamsRegionGlobal, SettingDomainNewParamsRegionAu, SettingDomainNewParamsRegionDe, SettingDomainNewParamsRegionIn, SettingDomainNewParamsRegionUs:
		return true
	}
	return false
}

type SettingDomainNewParamsFolder string

const (
	SettingDomainNewParamsFolderAllItems SettingDomainNewParamsFolder = "AllItems"
	SettingDomainNewParamsFolderInbox    SettingDomainNewParamsFolder = "Inbox"
)

func (r SettingDomainNewParamsFolder) IsKnown() bool {
	switch r {
	case SettingDomainNewParamsFolderAllItems, SettingDomainNewParamsFolderInbox:
		return true
	}
	return false
}

type SettingDomainNewResponseEnvelope struct {
	Errors   []SettingDomainNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingDomainNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingDomainNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingDomainNewResponse                `json:"result"`
	JSON    settingDomainNewResponseEnvelopeJSON    `json:"-"`
}

// settingDomainNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainNewResponseEnvelope]
type settingDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SettingDomainNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingDomainNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingDomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingDomainNewResponseEnvelopeErrors]
type settingDomainNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDomainNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    settingDomainNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingDomainNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingDomainNewResponseEnvelopeErrorsSource]
type settingDomainNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingDomainNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SettingDomainNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingDomainNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingDomainNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingDomainNewResponseEnvelopeMessages]
type settingDomainNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingDomainNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    settingDomainNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingDomainNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingDomainNewResponseEnvelopeMessagesSource]
type settingDomainNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingDomainNewResponseEnvelopeSuccess bool

const (
	SettingDomainNewResponseEnvelopeSuccessTrue SettingDomainNewResponseEnvelopeSuccess = true
)

func (r SettingDomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingDomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingDomainUpdateParams struct {
	// Identifier.
	AccountID            param.Field[string]                                         `path:"account_id" api:"required"`
	AllowedDeliveryModes param.Field[[]SettingDomainUpdateParamsAllowedDeliveryMode] `json:"allowed_delivery_modes" api:"required"`
	DropDispositions     param.Field[[]SettingDomainUpdateParamsDropDisposition]     `json:"drop_dispositions" api:"required"`
	IPRestrictions       param.Field[[]string]                                       `json:"ip_restrictions" api:"required"`
	Regions              param.Field[[]SettingDomainUpdateParamsRegion]              `json:"regions" api:"required"`
	Folder               param.Field[SettingDomainUpdateParamsFolder]                `json:"folder"`
	IntegrationID        param.Field[string]                                         `json:"integration_id" format:"uuid"`
	LookbackHops         param.Field[int64]                                          `json:"lookback_hops"`
	RequireTLSInbound    param.Field[bool]                                           `json:"require_tls_inbound"`
	RequireTLSOutbound   param.Field[bool]                                           `json:"require_tls_outbound"`
	Transport            param.Field[string]                                         `json:"transport"`
}

func (r SettingDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainUpdateParamsAllowedDeliveryMode string

const (
	SettingDomainUpdateParamsAllowedDeliveryModeDirect    SettingDomainUpdateParamsAllowedDeliveryMode = "DIRECT"
	SettingDomainUpdateParamsAllowedDeliveryModeBcc       SettingDomainUpdateParamsAllowedDeliveryMode = "BCC"
	SettingDomainUpdateParamsAllowedDeliveryModeJournal   SettingDomainUpdateParamsAllowedDeliveryMode = "JOURNAL"
	SettingDomainUpdateParamsAllowedDeliveryModeAPI       SettingDomainUpdateParamsAllowedDeliveryMode = "API"
	SettingDomainUpdateParamsAllowedDeliveryModeRetroScan SettingDomainUpdateParamsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainUpdateParamsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainUpdateParamsAllowedDeliveryModeDirect, SettingDomainUpdateParamsAllowedDeliveryModeBcc, SettingDomainUpdateParamsAllowedDeliveryModeJournal, SettingDomainUpdateParamsAllowedDeliveryModeAPI, SettingDomainUpdateParamsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainUpdateParamsDropDisposition string

const (
	SettingDomainUpdateParamsDropDispositionMalicious    SettingDomainUpdateParamsDropDisposition = "MALICIOUS"
	SettingDomainUpdateParamsDropDispositionMaliciousBec SettingDomainUpdateParamsDropDisposition = "MALICIOUS-BEC"
	SettingDomainUpdateParamsDropDispositionSuspicious   SettingDomainUpdateParamsDropDisposition = "SUSPICIOUS"
	SettingDomainUpdateParamsDropDispositionSpoof        SettingDomainUpdateParamsDropDisposition = "SPOOF"
	SettingDomainUpdateParamsDropDispositionSpam         SettingDomainUpdateParamsDropDisposition = "SPAM"
	SettingDomainUpdateParamsDropDispositionBulk         SettingDomainUpdateParamsDropDisposition = "BULK"
	SettingDomainUpdateParamsDropDispositionEncrypted    SettingDomainUpdateParamsDropDisposition = "ENCRYPTED"
	SettingDomainUpdateParamsDropDispositionExternal     SettingDomainUpdateParamsDropDisposition = "EXTERNAL"
	SettingDomainUpdateParamsDropDispositionUnknown      SettingDomainUpdateParamsDropDisposition = "UNKNOWN"
	SettingDomainUpdateParamsDropDispositionNone         SettingDomainUpdateParamsDropDisposition = "NONE"
)

func (r SettingDomainUpdateParamsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainUpdateParamsDropDispositionMalicious, SettingDomainUpdateParamsDropDispositionMaliciousBec, SettingDomainUpdateParamsDropDispositionSuspicious, SettingDomainUpdateParamsDropDispositionSpoof, SettingDomainUpdateParamsDropDispositionSpam, SettingDomainUpdateParamsDropDispositionBulk, SettingDomainUpdateParamsDropDispositionEncrypted, SettingDomainUpdateParamsDropDispositionExternal, SettingDomainUpdateParamsDropDispositionUnknown, SettingDomainUpdateParamsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainUpdateParamsRegion string

const (
	SettingDomainUpdateParamsRegionGlobal SettingDomainUpdateParamsRegion = "GLOBAL"
	SettingDomainUpdateParamsRegionAu     SettingDomainUpdateParamsRegion = "AU"
	SettingDomainUpdateParamsRegionDe     SettingDomainUpdateParamsRegion = "DE"
	SettingDomainUpdateParamsRegionIn     SettingDomainUpdateParamsRegion = "IN"
	SettingDomainUpdateParamsRegionUs     SettingDomainUpdateParamsRegion = "US"
)

func (r SettingDomainUpdateParamsRegion) IsKnown() bool {
	switch r {
	case SettingDomainUpdateParamsRegionGlobal, SettingDomainUpdateParamsRegionAu, SettingDomainUpdateParamsRegionDe, SettingDomainUpdateParamsRegionIn, SettingDomainUpdateParamsRegionUs:
		return true
	}
	return false
}

type SettingDomainUpdateParamsFolder string

const (
	SettingDomainUpdateParamsFolderAllItems SettingDomainUpdateParamsFolder = "AllItems"
	SettingDomainUpdateParamsFolderInbox    SettingDomainUpdateParamsFolder = "Inbox"
)

func (r SettingDomainUpdateParamsFolder) IsKnown() bool {
	switch r {
	case SettingDomainUpdateParamsFolderAllItems, SettingDomainUpdateParamsFolderInbox:
		return true
	}
	return false
}

type SettingDomainUpdateResponseEnvelope struct {
	Errors   []SettingDomainUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingDomainUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingDomainUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingDomainUpdateResponse                `json:"result"`
	JSON    settingDomainUpdateResponseEnvelopeJSON    `json:"-"`
}

// settingDomainUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingDomainUpdateResponseEnvelope]
type settingDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           SettingDomainUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingDomainUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingDomainUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingDomainUpdateResponseEnvelopeErrors]
type settingDomainUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDomainUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    settingDomainUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingDomainUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingDomainUpdateResponseEnvelopeErrorsSource]
type settingDomainUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingDomainUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingDomainUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingDomainUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingDomainUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingDomainUpdateResponseEnvelopeMessages]
type settingDomainUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingDomainUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingDomainUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingDomainUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingDomainUpdateResponseEnvelopeMessagesSource]
type settingDomainUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingDomainUpdateResponseEnvelopeSuccess bool

const (
	SettingDomainUpdateResponseEnvelopeSuccessTrue SettingDomainUpdateResponseEnvelopeSuccess = true
)

func (r SettingDomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingDomainUpdateResponseEnvelopeSuccessTrue:
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
	SettingDomainListParamsStatusPending SettingDomainListParamsStatus = "PENDING"
	SettingDomainListParamsStatusActive  SettingDomainListParamsStatus = "ACTIVE"
	SettingDomainListParamsStatusFailed  SettingDomainListParamsStatus = "FAILED"
	SettingDomainListParamsStatusTimeout SettingDomainListParamsStatus = "TIMEOUT"
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

type SettingDomainBatchParams struct {
	// Identifier.
	AccountID param.Field[string]                           `path:"account_id" api:"required"`
	Deletes   param.Field[[]SettingDomainBatchParamsDelete] `json:"deletes" api:"required"`
	Patches   param.Field[[]SettingDomainBatchParamsPatch]  `json:"patches" api:"required"`
	Posts     param.Field[[]SettingDomainBatchParamsPost]   `json:"posts" api:"required"`
	Puts      param.Field[[]SettingDomainBatchParamsPut]    `json:"puts" api:"required"`
}

func (r SettingDomainBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainBatchParamsDelete struct {
	// Domain identifier.
	ID param.Field[string] `json:"id" api:"required" format:"uuid"`
}

func (r SettingDomainBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainBatchParamsPatch struct {
	// Domain identifier.
	ID                   param.Field[string]                                               `json:"id" api:"required" format:"uuid"`
	AllowedDeliveryModes param.Field[[]SettingDomainBatchParamsPatchesAllowedDeliveryMode] `json:"allowed_delivery_modes"`
	DropDispositions     param.Field[[]SettingDomainBatchParamsPatchesDropDisposition]     `json:"drop_dispositions"`
	Folder               param.Field[SettingDomainBatchParamsPatchesFolder]                `json:"folder"`
	IntegrationID        param.Field[string]                                               `json:"integration_id" format:"uuid"`
	IPRestrictions       param.Field[[]string]                                             `json:"ip_restrictions"`
	LookbackHops         param.Field[int64]                                                `json:"lookback_hops"`
	Regions              param.Field[[]SettingDomainBatchParamsPatchesRegion]              `json:"regions"`
	RequireTLSInbound    param.Field[bool]                                                 `json:"require_tls_inbound"`
	RequireTLSOutbound   param.Field[bool]                                                 `json:"require_tls_outbound"`
	Transport            param.Field[string]                                               `json:"transport"`
}

func (r SettingDomainBatchParamsPatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainBatchParamsPatchesAllowedDeliveryMode string

const (
	SettingDomainBatchParamsPatchesAllowedDeliveryModeDirect    SettingDomainBatchParamsPatchesAllowedDeliveryMode = "DIRECT"
	SettingDomainBatchParamsPatchesAllowedDeliveryModeBcc       SettingDomainBatchParamsPatchesAllowedDeliveryMode = "BCC"
	SettingDomainBatchParamsPatchesAllowedDeliveryModeJournal   SettingDomainBatchParamsPatchesAllowedDeliveryMode = "JOURNAL"
	SettingDomainBatchParamsPatchesAllowedDeliveryModeAPI       SettingDomainBatchParamsPatchesAllowedDeliveryMode = "API"
	SettingDomainBatchParamsPatchesAllowedDeliveryModeRetroScan SettingDomainBatchParamsPatchesAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainBatchParamsPatchesAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPatchesAllowedDeliveryModeDirect, SettingDomainBatchParamsPatchesAllowedDeliveryModeBcc, SettingDomainBatchParamsPatchesAllowedDeliveryModeJournal, SettingDomainBatchParamsPatchesAllowedDeliveryModeAPI, SettingDomainBatchParamsPatchesAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainBatchParamsPatchesDropDisposition string

const (
	SettingDomainBatchParamsPatchesDropDispositionMalicious    SettingDomainBatchParamsPatchesDropDisposition = "MALICIOUS"
	SettingDomainBatchParamsPatchesDropDispositionMaliciousBec SettingDomainBatchParamsPatchesDropDisposition = "MALICIOUS-BEC"
	SettingDomainBatchParamsPatchesDropDispositionSuspicious   SettingDomainBatchParamsPatchesDropDisposition = "SUSPICIOUS"
	SettingDomainBatchParamsPatchesDropDispositionSpoof        SettingDomainBatchParamsPatchesDropDisposition = "SPOOF"
	SettingDomainBatchParamsPatchesDropDispositionSpam         SettingDomainBatchParamsPatchesDropDisposition = "SPAM"
	SettingDomainBatchParamsPatchesDropDispositionBulk         SettingDomainBatchParamsPatchesDropDisposition = "BULK"
	SettingDomainBatchParamsPatchesDropDispositionEncrypted    SettingDomainBatchParamsPatchesDropDisposition = "ENCRYPTED"
	SettingDomainBatchParamsPatchesDropDispositionExternal     SettingDomainBatchParamsPatchesDropDisposition = "EXTERNAL"
	SettingDomainBatchParamsPatchesDropDispositionUnknown      SettingDomainBatchParamsPatchesDropDisposition = "UNKNOWN"
	SettingDomainBatchParamsPatchesDropDispositionNone         SettingDomainBatchParamsPatchesDropDisposition = "NONE"
)

func (r SettingDomainBatchParamsPatchesDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPatchesDropDispositionMalicious, SettingDomainBatchParamsPatchesDropDispositionMaliciousBec, SettingDomainBatchParamsPatchesDropDispositionSuspicious, SettingDomainBatchParamsPatchesDropDispositionSpoof, SettingDomainBatchParamsPatchesDropDispositionSpam, SettingDomainBatchParamsPatchesDropDispositionBulk, SettingDomainBatchParamsPatchesDropDispositionEncrypted, SettingDomainBatchParamsPatchesDropDispositionExternal, SettingDomainBatchParamsPatchesDropDispositionUnknown, SettingDomainBatchParamsPatchesDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainBatchParamsPatchesFolder string

const (
	SettingDomainBatchParamsPatchesFolderAllItems SettingDomainBatchParamsPatchesFolder = "AllItems"
	SettingDomainBatchParamsPatchesFolderInbox    SettingDomainBatchParamsPatchesFolder = "Inbox"
)

func (r SettingDomainBatchParamsPatchesFolder) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPatchesFolderAllItems, SettingDomainBatchParamsPatchesFolderInbox:
		return true
	}
	return false
}

type SettingDomainBatchParamsPatchesRegion string

const (
	SettingDomainBatchParamsPatchesRegionGlobal SettingDomainBatchParamsPatchesRegion = "GLOBAL"
	SettingDomainBatchParamsPatchesRegionAu     SettingDomainBatchParamsPatchesRegion = "AU"
	SettingDomainBatchParamsPatchesRegionDe     SettingDomainBatchParamsPatchesRegion = "DE"
	SettingDomainBatchParamsPatchesRegionIn     SettingDomainBatchParamsPatchesRegion = "IN"
	SettingDomainBatchParamsPatchesRegionUs     SettingDomainBatchParamsPatchesRegion = "US"
)

func (r SettingDomainBatchParamsPatchesRegion) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPatchesRegionGlobal, SettingDomainBatchParamsPatchesRegionAu, SettingDomainBatchParamsPatchesRegionDe, SettingDomainBatchParamsPatchesRegionIn, SettingDomainBatchParamsPatchesRegionUs:
		return true
	}
	return false
}

type SettingDomainBatchParamsPost struct {
	AllowedDeliveryModes param.Field[[]SettingDomainBatchParamsPostsAllowedDeliveryMode] `json:"allowed_delivery_modes" api:"required"`
	Domain               param.Field[string]                                             `json:"domain" api:"required"`
	DropDispositions     param.Field[[]SettingDomainBatchParamsPostsDropDisposition]     `json:"drop_dispositions" api:"required"`
	IPRestrictions       param.Field[[]string]                                           `json:"ip_restrictions" api:"required"`
	Regions              param.Field[[]SettingDomainBatchParamsPostsRegion]              `json:"regions" api:"required"`
	Folder               param.Field[SettingDomainBatchParamsPostsFolder]                `json:"folder"`
	IntegrationID        param.Field[string]                                             `json:"integration_id" format:"uuid"`
	LookbackHops         param.Field[int64]                                              `json:"lookback_hops"`
	RequireTLSInbound    param.Field[bool]                                               `json:"require_tls_inbound"`
	RequireTLSOutbound   param.Field[bool]                                               `json:"require_tls_outbound"`
	Transport            param.Field[string]                                             `json:"transport"`
}

func (r SettingDomainBatchParamsPost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainBatchParamsPostsAllowedDeliveryMode string

const (
	SettingDomainBatchParamsPostsAllowedDeliveryModeDirect    SettingDomainBatchParamsPostsAllowedDeliveryMode = "DIRECT"
	SettingDomainBatchParamsPostsAllowedDeliveryModeBcc       SettingDomainBatchParamsPostsAllowedDeliveryMode = "BCC"
	SettingDomainBatchParamsPostsAllowedDeliveryModeJournal   SettingDomainBatchParamsPostsAllowedDeliveryMode = "JOURNAL"
	SettingDomainBatchParamsPostsAllowedDeliveryModeAPI       SettingDomainBatchParamsPostsAllowedDeliveryMode = "API"
	SettingDomainBatchParamsPostsAllowedDeliveryModeRetroScan SettingDomainBatchParamsPostsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainBatchParamsPostsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPostsAllowedDeliveryModeDirect, SettingDomainBatchParamsPostsAllowedDeliveryModeBcc, SettingDomainBatchParamsPostsAllowedDeliveryModeJournal, SettingDomainBatchParamsPostsAllowedDeliveryModeAPI, SettingDomainBatchParamsPostsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainBatchParamsPostsDropDisposition string

const (
	SettingDomainBatchParamsPostsDropDispositionMalicious    SettingDomainBatchParamsPostsDropDisposition = "MALICIOUS"
	SettingDomainBatchParamsPostsDropDispositionMaliciousBec SettingDomainBatchParamsPostsDropDisposition = "MALICIOUS-BEC"
	SettingDomainBatchParamsPostsDropDispositionSuspicious   SettingDomainBatchParamsPostsDropDisposition = "SUSPICIOUS"
	SettingDomainBatchParamsPostsDropDispositionSpoof        SettingDomainBatchParamsPostsDropDisposition = "SPOOF"
	SettingDomainBatchParamsPostsDropDispositionSpam         SettingDomainBatchParamsPostsDropDisposition = "SPAM"
	SettingDomainBatchParamsPostsDropDispositionBulk         SettingDomainBatchParamsPostsDropDisposition = "BULK"
	SettingDomainBatchParamsPostsDropDispositionEncrypted    SettingDomainBatchParamsPostsDropDisposition = "ENCRYPTED"
	SettingDomainBatchParamsPostsDropDispositionExternal     SettingDomainBatchParamsPostsDropDisposition = "EXTERNAL"
	SettingDomainBatchParamsPostsDropDispositionUnknown      SettingDomainBatchParamsPostsDropDisposition = "UNKNOWN"
	SettingDomainBatchParamsPostsDropDispositionNone         SettingDomainBatchParamsPostsDropDisposition = "NONE"
)

func (r SettingDomainBatchParamsPostsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPostsDropDispositionMalicious, SettingDomainBatchParamsPostsDropDispositionMaliciousBec, SettingDomainBatchParamsPostsDropDispositionSuspicious, SettingDomainBatchParamsPostsDropDispositionSpoof, SettingDomainBatchParamsPostsDropDispositionSpam, SettingDomainBatchParamsPostsDropDispositionBulk, SettingDomainBatchParamsPostsDropDispositionEncrypted, SettingDomainBatchParamsPostsDropDispositionExternal, SettingDomainBatchParamsPostsDropDispositionUnknown, SettingDomainBatchParamsPostsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainBatchParamsPostsRegion string

const (
	SettingDomainBatchParamsPostsRegionGlobal SettingDomainBatchParamsPostsRegion = "GLOBAL"
	SettingDomainBatchParamsPostsRegionAu     SettingDomainBatchParamsPostsRegion = "AU"
	SettingDomainBatchParamsPostsRegionDe     SettingDomainBatchParamsPostsRegion = "DE"
	SettingDomainBatchParamsPostsRegionIn     SettingDomainBatchParamsPostsRegion = "IN"
	SettingDomainBatchParamsPostsRegionUs     SettingDomainBatchParamsPostsRegion = "US"
)

func (r SettingDomainBatchParamsPostsRegion) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPostsRegionGlobal, SettingDomainBatchParamsPostsRegionAu, SettingDomainBatchParamsPostsRegionDe, SettingDomainBatchParamsPostsRegionIn, SettingDomainBatchParamsPostsRegionUs:
		return true
	}
	return false
}

type SettingDomainBatchParamsPostsFolder string

const (
	SettingDomainBatchParamsPostsFolderAllItems SettingDomainBatchParamsPostsFolder = "AllItems"
	SettingDomainBatchParamsPostsFolderInbox    SettingDomainBatchParamsPostsFolder = "Inbox"
)

func (r SettingDomainBatchParamsPostsFolder) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPostsFolderAllItems, SettingDomainBatchParamsPostsFolderInbox:
		return true
	}
	return false
}

// Request body for replacing an email domain. The `domain` field is intentionally
// absent — the domain name is immutable after creation.
type SettingDomainBatchParamsPut struct {
	// Domain identifier.
	ID                   param.Field[string]                                            `json:"id" api:"required" format:"uuid"`
	AllowedDeliveryModes param.Field[[]SettingDomainBatchParamsPutsAllowedDeliveryMode] `json:"allowed_delivery_modes" api:"required"`
	DropDispositions     param.Field[[]SettingDomainBatchParamsPutsDropDisposition]     `json:"drop_dispositions" api:"required"`
	IPRestrictions       param.Field[[]string]                                          `json:"ip_restrictions" api:"required"`
	Regions              param.Field[[]SettingDomainBatchParamsPutsRegion]              `json:"regions" api:"required"`
	Folder               param.Field[SettingDomainBatchParamsPutsFolder]                `json:"folder"`
	IntegrationID        param.Field[string]                                            `json:"integration_id" format:"uuid"`
	LookbackHops         param.Field[int64]                                             `json:"lookback_hops"`
	RequireTLSInbound    param.Field[bool]                                              `json:"require_tls_inbound"`
	RequireTLSOutbound   param.Field[bool]                                              `json:"require_tls_outbound"`
	Transport            param.Field[string]                                            `json:"transport"`
}

func (r SettingDomainBatchParamsPut) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainBatchParamsPutsAllowedDeliveryMode string

const (
	SettingDomainBatchParamsPutsAllowedDeliveryModeDirect    SettingDomainBatchParamsPutsAllowedDeliveryMode = "DIRECT"
	SettingDomainBatchParamsPutsAllowedDeliveryModeBcc       SettingDomainBatchParamsPutsAllowedDeliveryMode = "BCC"
	SettingDomainBatchParamsPutsAllowedDeliveryModeJournal   SettingDomainBatchParamsPutsAllowedDeliveryMode = "JOURNAL"
	SettingDomainBatchParamsPutsAllowedDeliveryModeAPI       SettingDomainBatchParamsPutsAllowedDeliveryMode = "API"
	SettingDomainBatchParamsPutsAllowedDeliveryModeRetroScan SettingDomainBatchParamsPutsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainBatchParamsPutsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPutsAllowedDeliveryModeDirect, SettingDomainBatchParamsPutsAllowedDeliveryModeBcc, SettingDomainBatchParamsPutsAllowedDeliveryModeJournal, SettingDomainBatchParamsPutsAllowedDeliveryModeAPI, SettingDomainBatchParamsPutsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainBatchParamsPutsDropDisposition string

const (
	SettingDomainBatchParamsPutsDropDispositionMalicious    SettingDomainBatchParamsPutsDropDisposition = "MALICIOUS"
	SettingDomainBatchParamsPutsDropDispositionMaliciousBec SettingDomainBatchParamsPutsDropDisposition = "MALICIOUS-BEC"
	SettingDomainBatchParamsPutsDropDispositionSuspicious   SettingDomainBatchParamsPutsDropDisposition = "SUSPICIOUS"
	SettingDomainBatchParamsPutsDropDispositionSpoof        SettingDomainBatchParamsPutsDropDisposition = "SPOOF"
	SettingDomainBatchParamsPutsDropDispositionSpam         SettingDomainBatchParamsPutsDropDisposition = "SPAM"
	SettingDomainBatchParamsPutsDropDispositionBulk         SettingDomainBatchParamsPutsDropDisposition = "BULK"
	SettingDomainBatchParamsPutsDropDispositionEncrypted    SettingDomainBatchParamsPutsDropDisposition = "ENCRYPTED"
	SettingDomainBatchParamsPutsDropDispositionExternal     SettingDomainBatchParamsPutsDropDisposition = "EXTERNAL"
	SettingDomainBatchParamsPutsDropDispositionUnknown      SettingDomainBatchParamsPutsDropDisposition = "UNKNOWN"
	SettingDomainBatchParamsPutsDropDispositionNone         SettingDomainBatchParamsPutsDropDisposition = "NONE"
)

func (r SettingDomainBatchParamsPutsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPutsDropDispositionMalicious, SettingDomainBatchParamsPutsDropDispositionMaliciousBec, SettingDomainBatchParamsPutsDropDispositionSuspicious, SettingDomainBatchParamsPutsDropDispositionSpoof, SettingDomainBatchParamsPutsDropDispositionSpam, SettingDomainBatchParamsPutsDropDispositionBulk, SettingDomainBatchParamsPutsDropDispositionEncrypted, SettingDomainBatchParamsPutsDropDispositionExternal, SettingDomainBatchParamsPutsDropDispositionUnknown, SettingDomainBatchParamsPutsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainBatchParamsPutsRegion string

const (
	SettingDomainBatchParamsPutsRegionGlobal SettingDomainBatchParamsPutsRegion = "GLOBAL"
	SettingDomainBatchParamsPutsRegionAu     SettingDomainBatchParamsPutsRegion = "AU"
	SettingDomainBatchParamsPutsRegionDe     SettingDomainBatchParamsPutsRegion = "DE"
	SettingDomainBatchParamsPutsRegionIn     SettingDomainBatchParamsPutsRegion = "IN"
	SettingDomainBatchParamsPutsRegionUs     SettingDomainBatchParamsPutsRegion = "US"
)

func (r SettingDomainBatchParamsPutsRegion) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPutsRegionGlobal, SettingDomainBatchParamsPutsRegionAu, SettingDomainBatchParamsPutsRegionDe, SettingDomainBatchParamsPutsRegionIn, SettingDomainBatchParamsPutsRegionUs:
		return true
	}
	return false
}

type SettingDomainBatchParamsPutsFolder string

const (
	SettingDomainBatchParamsPutsFolderAllItems SettingDomainBatchParamsPutsFolder = "AllItems"
	SettingDomainBatchParamsPutsFolderInbox    SettingDomainBatchParamsPutsFolder = "Inbox"
)

func (r SettingDomainBatchParamsPutsFolder) IsKnown() bool {
	switch r {
	case SettingDomainBatchParamsPutsFolderAllItems, SettingDomainBatchParamsPutsFolderInbox:
		return true
	}
	return false
}

type SettingDomainBatchResponseEnvelope struct {
	Errors   []SettingDomainBatchResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingDomainBatchResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingDomainBatchResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingDomainBatchResponse                `json:"result"`
	JSON    settingDomainBatchResponseEnvelopeJSON    `json:"-"`
}

// settingDomainBatchResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainBatchResponseEnvelope]
type settingDomainBatchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainBatchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SettingDomainBatchResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingDomainBatchResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingDomainBatchResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingDomainBatchResponseEnvelopeErrors]
type settingDomainBatchResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainBatchResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    settingDomainBatchResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingDomainBatchResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingDomainBatchResponseEnvelopeErrorsSource]
type settingDomainBatchResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainBatchResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           SettingDomainBatchResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingDomainBatchResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingDomainBatchResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingDomainBatchResponseEnvelopeMessages]
type settingDomainBatchResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingDomainBatchResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBatchResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    settingDomainBatchResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingDomainBatchResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingDomainBatchResponseEnvelopeMessagesSource]
type settingDomainBatchResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainBatchResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBatchResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingDomainBatchResponseEnvelopeSuccess bool

const (
	SettingDomainBatchResponseEnvelopeSuccessTrue SettingDomainBatchResponseEnvelopeSuccess = true
)

func (r SettingDomainBatchResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingDomainBatchResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingDomainBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingDomainEditParams struct {
	// Identifier.
	AccountID            param.Field[string]                                       `path:"account_id" api:"required"`
	AllowedDeliveryModes param.Field[[]SettingDomainEditParamsAllowedDeliveryMode] `json:"allowed_delivery_modes"`
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
