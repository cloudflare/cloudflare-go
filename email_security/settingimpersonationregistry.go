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

// SettingImpersonationRegistryService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingImpersonationRegistryService] method instead.
type SettingImpersonationRegistryService struct {
	Options []option.RequestOption
}

// NewSettingImpersonationRegistryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingImpersonationRegistryService(opts ...option.RequestOption) (r *SettingImpersonationRegistryService) {
	r = &SettingImpersonationRegistryService{}
	r.Options = opts
	return
}

// Creates a new entry in the impersonation registry to protect against
// impersonation. Emails attempting to impersonate this identity will be flagged.
// Supports regex patterns for flexible email matching.
func (r *SettingImpersonationRegistryService) New(ctx context.Context, params SettingImpersonationRegistryNewParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryNewResponse, err error) {
	var env SettingImpersonationRegistryNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of protected identities in the impersonation registry.
// These entries define identities and email addresses to protect from
// impersonation attacks. Can be manually added or automatically synced from
// directory integrations.
func (r *SettingImpersonationRegistryService) List(ctx context.Context, params SettingImpersonationRegistryListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingImpersonationRegistryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry", params.AccountID)
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

// Returns a paginated list of protected identities in the impersonation registry.
// These entries define identities and email addresses to protect from
// impersonation attacks. Can be manually added or automatically synced from
// directory integrations.
func (r *SettingImpersonationRegistryService) ListAutoPaging(ctx context.Context, params SettingImpersonationRegistryListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingImpersonationRegistryListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// An impersonation registry entry
type SettingImpersonationRegistryNewResponse struct {
	// Impersonation registry entry identifier
	ID              string    `json:"id" format:"uuid"`
	Comments        string    `json:"comments" api:"nullable"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	DirectoryID     int64     `json:"directory_id" api:"nullable"`
	DirectoryNodeID int64     `json:"directory_node_id" api:"nullable"`
	Email           string    `json:"email"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string `json:"external_directory_node_id" api:"nullable"`
	IsEmailRegex            bool   `json:"is_email_regex"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                                         `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                                         `json:"modified_at" format:"date-time"`
	Name         string                                            `json:"name"`
	Provenance   SettingImpersonationRegistryNewResponseProvenance `json:"provenance"`
	JSON         settingImpersonationRegistryNewResponseJSON       `json:"-"`
}

// settingImpersonationRegistryNewResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryNewResponse]
type settingImpersonationRegistryNewResponseJSON struct {
	ID                      apijson.Field
	Comments                apijson.Field
	CreatedAt               apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	Email                   apijson.Field
	ExternalDirectoryNodeID apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	ModifiedAt              apijson.Field
	Name                    apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryNewResponseProvenance string

const (
	SettingImpersonationRegistryNewResponseProvenanceA1SInternal           SettingImpersonationRegistryNewResponseProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryNewResponseProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryNewResponseProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryNewResponseProvenanceSnoopyOffice365       SettingImpersonationRegistryNewResponseProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryNewResponseProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryNewResponseProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryNewResponseProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryNewResponseProvenanceA1SInternal, SettingImpersonationRegistryNewResponseProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryNewResponseProvenanceSnoopyOffice365, SettingImpersonationRegistryNewResponseProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}

// An impersonation registry entry
type SettingImpersonationRegistryListResponse struct {
	// Impersonation registry entry identifier
	ID              string    `json:"id" format:"uuid"`
	Comments        string    `json:"comments" api:"nullable"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	DirectoryID     int64     `json:"directory_id" api:"nullable"`
	DirectoryNodeID int64     `json:"directory_node_id" api:"nullable"`
	Email           string    `json:"email"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string `json:"external_directory_node_id" api:"nullable"`
	IsEmailRegex            bool   `json:"is_email_regex"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                                          `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                                          `json:"modified_at" format:"date-time"`
	Name         string                                             `json:"name"`
	Provenance   SettingImpersonationRegistryListResponseProvenance `json:"provenance"`
	JSON         settingImpersonationRegistryListResponseJSON       `json:"-"`
}

// settingImpersonationRegistryListResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryListResponse]
type settingImpersonationRegistryListResponseJSON struct {
	ID                      apijson.Field
	Comments                apijson.Field
	CreatedAt               apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	Email                   apijson.Field
	ExternalDirectoryNodeID apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	ModifiedAt              apijson.Field
	Name                    apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingImpersonationRegistryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryListResponseProvenance string

const (
	SettingImpersonationRegistryListResponseProvenanceA1SInternal           SettingImpersonationRegistryListResponseProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryListResponseProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryListResponseProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryListResponseProvenanceSnoopyOffice365       SettingImpersonationRegistryListResponseProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryListResponseProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryListResponseProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryListResponseProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryListResponseProvenanceA1SInternal, SettingImpersonationRegistryListResponseProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryListResponseProvenanceSnoopyOffice365, SettingImpersonationRegistryListResponseProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}

type SettingImpersonationRegistryNewParams struct {
	// Identifier.
	AccountID               param.Field[string]                                          `path:"account_id" api:"required"`
	Email                   param.Field[string]                                          `json:"email" api:"required"`
	IsEmailRegex            param.Field[bool]                                            `json:"is_email_regex" api:"required"`
	Name                    param.Field[string]                                          `json:"name" api:"required"`
	Comments                param.Field[string]                                          `json:"comments"`
	DirectoryID             param.Field[int64]                                           `json:"directory_id"`
	DirectoryNodeID         param.Field[int64]                                           `json:"directory_node_id"`
	ExternalDirectoryNodeID param.Field[string]                                          `json:"external_directory_node_id"`
	Provenance              param.Field[SettingImpersonationRegistryNewParamsProvenance] `json:"provenance"`
}

func (r SettingImpersonationRegistryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingImpersonationRegistryNewParamsProvenance string

const (
	SettingImpersonationRegistryNewParamsProvenanceA1SInternal           SettingImpersonationRegistryNewParamsProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryNewParamsProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryNewParamsProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryNewParamsProvenanceSnoopyOffice365       SettingImpersonationRegistryNewParamsProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryNewParamsProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryNewParamsProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryNewParamsProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryNewParamsProvenanceA1SInternal, SettingImpersonationRegistryNewParamsProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryNewParamsProvenanceSnoopyOffice365, SettingImpersonationRegistryNewParamsProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}

type SettingImpersonationRegistryNewResponseEnvelope struct {
	Errors   []SettingImpersonationRegistryNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingImpersonationRegistryNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingImpersonationRegistryNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// An impersonation registry entry
	Result SettingImpersonationRegistryNewResponse             `json:"result"`
	JSON   settingImpersonationRegistryNewResponseEnvelopeJSON `json:"-"`
}

// settingImpersonationRegistryNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingImpersonationRegistryNewResponseEnvelope]
type settingImpersonationRegistryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryNewResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           SettingImpersonationRegistryNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingImpersonationRegistryNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingImpersonationRegistryNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingImpersonationRegistryNewResponseEnvelopeErrors]
type settingImpersonationRegistryNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    settingImpersonationRegistryNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingImpersonationRegistryNewResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingImpersonationRegistryNewResponseEnvelopeErrorsSource]
type settingImpersonationRegistryNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryNewResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           SettingImpersonationRegistryNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingImpersonationRegistryNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingImpersonationRegistryNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingImpersonationRegistryNewResponseEnvelopeMessages]
type settingImpersonationRegistryNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    settingImpersonationRegistryNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingImpersonationRegistryNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SettingImpersonationRegistryNewResponseEnvelopeMessagesSource]
type settingImpersonationRegistryNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingImpersonationRegistryNewResponseEnvelopeSuccess bool

const (
	SettingImpersonationRegistryNewResponseEnvelopeSuccessTrue SettingImpersonationRegistryNewResponseEnvelopeSuccess = true
)

func (r SettingImpersonationRegistryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingImpersonationRegistryListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The sorting direction.
	Direction param.Field[SettingImpersonationRegistryListParamsDirection] `query:"direction"`
	// Field to sort by.
	Order param.Field[SettingImpersonationRegistryListParamsOrder] `query:"order"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage    param.Field[int64]                                            `query:"per_page"`
	Provenance param.Field[SettingImpersonationRegistryListParamsProvenance] `query:"provenance"`
	// Search term for filtering records. Behavior may change.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingImpersonationRegistryListParams]'s query parameters
// as `url.Values`.
func (r SettingImpersonationRegistryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingImpersonationRegistryListParamsDirection string

const (
	SettingImpersonationRegistryListParamsDirectionAsc  SettingImpersonationRegistryListParamsDirection = "asc"
	SettingImpersonationRegistryListParamsDirectionDesc SettingImpersonationRegistryListParamsDirection = "desc"
)

func (r SettingImpersonationRegistryListParamsDirection) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryListParamsDirectionAsc, SettingImpersonationRegistryListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by.
type SettingImpersonationRegistryListParamsOrder string

const (
	SettingImpersonationRegistryListParamsOrderName      SettingImpersonationRegistryListParamsOrder = "name"
	SettingImpersonationRegistryListParamsOrderEmail     SettingImpersonationRegistryListParamsOrder = "email"
	SettingImpersonationRegistryListParamsOrderCreatedAt SettingImpersonationRegistryListParamsOrder = "created_at"
)

func (r SettingImpersonationRegistryListParamsOrder) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryListParamsOrderName, SettingImpersonationRegistryListParamsOrderEmail, SettingImpersonationRegistryListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingImpersonationRegistryListParamsProvenance string

const (
	SettingImpersonationRegistryListParamsProvenanceA1SInternal           SettingImpersonationRegistryListParamsProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryListParamsProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryListParamsProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryListParamsProvenanceSnoopyOffice365       SettingImpersonationRegistryListParamsProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryListParamsProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryListParamsProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryListParamsProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryListParamsProvenanceA1SInternal, SettingImpersonationRegistryListParamsProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryListParamsProvenanceSnoopyOffice365, SettingImpersonationRegistryListParamsProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}
