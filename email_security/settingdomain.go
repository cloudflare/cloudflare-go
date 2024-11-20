// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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

// Lists, searches, and sorts an account’s email domains.
func (r *SettingDomainService) List(ctx context.Context, params SettingDomainListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingDomainListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
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

// Lists, searches, and sorts an account’s email domains.
func (r *SettingDomainService) ListAutoPaging(ctx context.Context, params SettingDomainListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingDomainListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Unprotect an email domain
func (r *SettingDomainService) Delete(ctx context.Context, domainID int64, body SettingDomainDeleteParams, opts ...option.RequestOption) (res *SettingDomainDeleteResponse, err error) {
	var env SettingDomainDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", body.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an email domain
func (r *SettingDomainService) Edit(ctx context.Context, domainID int64, params SettingDomainEditParams, opts ...option.RequestOption) (res *SettingDomainEditResponse, err error) {
	var env SettingDomainEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", params.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an email domain
func (r *SettingDomainService) Get(ctx context.Context, domainID int64, query SettingDomainGetParams, opts ...option.RequestOption) (res *SettingDomainGetResponse, err error) {
	var env SettingDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", query.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingDomainListResponse struct {
	// The unique identifier for the domain.
	ID                   int64                                          `json:"id,required"`
	AllowedDeliveryModes []SettingDomainListResponseAllowedDeliveryMode `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                      `json:"created_at,required" format:"date-time"`
	Domain               string                                         `json:"domain,required"`
	LastModified         time.Time                                      `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                          `json:"lookback_hops,required"`
	Folder               SettingDomainListResponseFolder                `json:"folder,nullable"`
	InboxProvider        SettingDomainListResponseInboxProvider         `json:"inbox_provider,nullable"`
	IntegrationID        string                                         `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                         `json:"o365_tenant_id,nullable"`
	JSON                 settingDomainListResponseJSON                  `json:"-"`
}

// settingDomainListResponseJSON contains the JSON metadata for the struct
// [SettingDomainListResponse]
type settingDomainListResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
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

type SettingDomainDeleteResponse struct {
	// The unique identifier for the domain.
	ID   int64                           `json:"id,required"`
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
	// The unique identifier for the domain.
	ID                   int64                                          `json:"id,required"`
	AllowedDeliveryModes []SettingDomainEditResponseAllowedDeliveryMode `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                      `json:"created_at,required" format:"date-time"`
	Domain               string                                         `json:"domain,required"`
	LastModified         time.Time                                      `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                          `json:"lookback_hops,required"`
	Folder               SettingDomainEditResponseFolder                `json:"folder,nullable"`
	InboxProvider        SettingDomainEditResponseInboxProvider         `json:"inbox_provider,nullable"`
	IntegrationID        string                                         `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                         `json:"o365_tenant_id,nullable"`
	JSON                 settingDomainEditResponseJSON                  `json:"-"`
}

// settingDomainEditResponseJSON contains the JSON metadata for the struct
// [SettingDomainEditResponse]
type settingDomainEditResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
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

type SettingDomainGetResponse struct {
	// The unique identifier for the domain.
	ID                   int64                                         `json:"id,required"`
	AllowedDeliveryModes []SettingDomainGetResponseAllowedDeliveryMode `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                     `json:"created_at,required" format:"date-time"`
	Domain               string                                        `json:"domain,required"`
	LastModified         time.Time                                     `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                         `json:"lookback_hops,required"`
	Folder               SettingDomainGetResponseFolder                `json:"folder,nullable"`
	InboxProvider        SettingDomainGetResponseInboxProvider         `json:"inbox_provider,nullable"`
	IntegrationID        string                                        `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                        `json:"o365_tenant_id,nullable"`
	JSON                 settingDomainGetResponseJSON                  `json:"-"`
}

// settingDomainGetResponseJSON contains the JSON metadata for the struct
// [SettingDomainGetResponse]
type settingDomainGetResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
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

type SettingDomainListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Filters response to domains with the provided delivery mode.
	AllowedDeliveryMode param.Field[SettingDomainListParamsAllowedDeliveryMode] `query:"allowed_delivery_mode"`
	// The sorting direction.
	Direction param.Field[SettingDomainListParamsDirection] `query:"direction"`
	// Filters results by the provided domains, allowing for multiple occurrences.
	Domain param.Field[[]string] `query:"domain"`
	// The field to sort by.
	Order param.Field[SettingDomainListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingDomainListParams]'s query parameters as
// `url.Values`.
func (r SettingDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters response to domains with the provided delivery mode.
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

// The field to sort by.
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

type SettingDomainDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingDomainDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                   `json:"errors,required"`
	Messages []shared.ResponseInfo                   `json:"messages,required"`
	Result   SettingDomainDeleteResponse             `json:"result,required"`
	Success  bool                                    `json:"success,required"`
	JSON     settingDomainDeleteResponseEnvelopeJSON `json:"-"`
}

// settingDomainDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingDomainDeleteResponseEnvelope]
type settingDomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditParams struct {
	// Account Identifier
	AccountID     param.Field[string]                        `path:"account_id,required"`
	Domain        param.Field[string]                        `json:"domain"`
	Folder        param.Field[SettingDomainEditParamsFolder] `json:"folder"`
	IntegrationID param.Field[string]                        `json:"integration_id" format:"uuid"`
	LookbackHops  param.Field[int64]                         `json:"lookback_hops"`
}

func (r SettingDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type SettingDomainEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []shared.ResponseInfo                 `json:"messages,required"`
	Result   SettingDomainEditResponse             `json:"result,required"`
	Success  bool                                  `json:"success,required"`
	JSON     settingDomainEditResponseEnvelopeJSON `json:"-"`
}

// settingDomainEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainEditResponseEnvelope]
type settingDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                `json:"errors,required"`
	Messages []shared.ResponseInfo                `json:"messages,required"`
	Result   SettingDomainGetResponse             `json:"result,required"`
	Success  bool                                 `json:"success,required"`
	JSON     settingDomainGetResponseEnvelopeJSON `json:"-"`
}

// settingDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainGetResponseEnvelope]
type settingDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
