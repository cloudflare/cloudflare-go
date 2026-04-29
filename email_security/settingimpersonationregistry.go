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

// Removes an entry from the impersonation registry. After deletion, this identity
// will no longer be protected from impersonation.
func (r *SettingImpersonationRegistryService) Delete(ctx context.Context, impersonationRegistryID string, body SettingImpersonationRegistryDeleteParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryDeleteResponse, err error) {
	var env SettingImpersonationRegistryDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if impersonationRegistryID == "" {
		err = errors.New("missing required impersonation_registry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%s", body.AccountID, impersonationRegistryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing impersonation registry entry. Only provided fields will be
// modified. Directory-synced entries can't be updated.
func (r *SettingImpersonationRegistryService) Edit(ctx context.Context, impersonationRegistryID string, params SettingImpersonationRegistryEditParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryEditResponse, err error) {
	var env SettingImpersonationRegistryEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if impersonationRegistryID == "" {
		err = errors.New("missing required impersonation_registry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%s", params.AccountID, impersonationRegistryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific impersonation registry entry including the
// protected identity, email pattern, and synchronization source if
// directory-synced.
func (r *SettingImpersonationRegistryService) Get(ctx context.Context, impersonationRegistryID string, query SettingImpersonationRegistryGetParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryGetResponse, err error) {
	var env SettingImpersonationRegistryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if impersonationRegistryID == "" {
		err = errors.New("missing required impersonation_registry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%s", query.AccountID, impersonationRegistryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
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

type SettingImpersonationRegistryDeleteResponse struct {
	// Impersonation registry entry identifier
	ID   string                                         `json:"id" api:"required" format:"uuid"`
	JSON settingImpersonationRegistryDeleteResponseJSON `json:"-"`
}

// settingImpersonationRegistryDeleteResponseJSON contains the JSON metadata for
// the struct [SettingImpersonationRegistryDeleteResponse]
type settingImpersonationRegistryDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// An impersonation registry entry
type SettingImpersonationRegistryEditResponse struct {
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
	Provenance   SettingImpersonationRegistryEditResponseProvenance `json:"provenance"`
	JSON         settingImpersonationRegistryEditResponseJSON       `json:"-"`
}

// settingImpersonationRegistryEditResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryEditResponse]
type settingImpersonationRegistryEditResponseJSON struct {
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

func (r *SettingImpersonationRegistryEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryEditResponseProvenance string

const (
	SettingImpersonationRegistryEditResponseProvenanceA1SInternal           SettingImpersonationRegistryEditResponseProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryEditResponseProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryEditResponseProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryEditResponseProvenanceSnoopyOffice365       SettingImpersonationRegistryEditResponseProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryEditResponseProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryEditResponseProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryEditResponseProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryEditResponseProvenanceA1SInternal, SettingImpersonationRegistryEditResponseProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryEditResponseProvenanceSnoopyOffice365, SettingImpersonationRegistryEditResponseProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}

// An impersonation registry entry
type SettingImpersonationRegistryGetResponse struct {
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
	Provenance   SettingImpersonationRegistryGetResponseProvenance `json:"provenance"`
	JSON         settingImpersonationRegistryGetResponseJSON       `json:"-"`
}

// settingImpersonationRegistryGetResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryGetResponse]
type settingImpersonationRegistryGetResponseJSON struct {
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

func (r *SettingImpersonationRegistryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryGetResponseProvenance string

const (
	SettingImpersonationRegistryGetResponseProvenanceA1SInternal           SettingImpersonationRegistryGetResponseProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryGetResponseProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryGetResponseProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryGetResponseProvenanceSnoopyOffice365       SettingImpersonationRegistryGetResponseProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryGetResponseProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryGetResponseProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryGetResponseProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryGetResponseProvenanceA1SInternal, SettingImpersonationRegistryGetResponseProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryGetResponseProvenanceSnoopyOffice365, SettingImpersonationRegistryGetResponseProvenanceSnoopyGoogleDirectory:
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

type SettingImpersonationRegistryDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingImpersonationRegistryDeleteResponseEnvelope struct {
	Errors   []SettingImpersonationRegistryDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingImpersonationRegistryDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingImpersonationRegistryDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingImpersonationRegistryDeleteResponse                `json:"result"`
	JSON    settingImpersonationRegistryDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingImpersonationRegistryDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingImpersonationRegistryDeleteResponseEnvelope]
type settingImpersonationRegistryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryDeleteResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           SettingImpersonationRegistryDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingImpersonationRegistryDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingImpersonationRegistryDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingImpersonationRegistryDeleteResponseEnvelopeErrors]
type settingImpersonationRegistryDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    settingImpersonationRegistryDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingImpersonationRegistryDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingImpersonationRegistryDeleteResponseEnvelopeErrorsSource]
type settingImpersonationRegistryDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryDeleteResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code" api:"required"`
	Message          string                                                           `json:"message" api:"required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           SettingImpersonationRegistryDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingImpersonationRegistryDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingImpersonationRegistryDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingImpersonationRegistryDeleteResponseEnvelopeMessages]
type settingImpersonationRegistryDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    settingImpersonationRegistryDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingImpersonationRegistryDeleteResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [SettingImpersonationRegistryDeleteResponseEnvelopeMessagesSource]
type settingImpersonationRegistryDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingImpersonationRegistryDeleteResponseEnvelopeSuccess bool

const (
	SettingImpersonationRegistryDeleteResponseEnvelopeSuccessTrue SettingImpersonationRegistryDeleteResponseEnvelopeSuccess = true
)

func (r SettingImpersonationRegistryDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingImpersonationRegistryEditParams struct {
	// Identifier.
	AccountID               param.Field[string]                                           `path:"account_id" api:"required"`
	Comments                param.Field[string]                                           `json:"comments"`
	DirectoryID             param.Field[int64]                                            `json:"directory_id"`
	DirectoryNodeID         param.Field[int64]                                            `json:"directory_node_id"`
	Email                   param.Field[string]                                           `json:"email"`
	ExternalDirectoryNodeID param.Field[string]                                           `json:"external_directory_node_id"`
	IsEmailRegex            param.Field[bool]                                             `json:"is_email_regex"`
	Name                    param.Field[string]                                           `json:"name"`
	Provenance              param.Field[SettingImpersonationRegistryEditParamsProvenance] `json:"provenance"`
}

func (r SettingImpersonationRegistryEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingImpersonationRegistryEditParamsProvenance string

const (
	SettingImpersonationRegistryEditParamsProvenanceA1SInternal           SettingImpersonationRegistryEditParamsProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryEditParamsProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryEditParamsProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryEditParamsProvenanceSnoopyOffice365       SettingImpersonationRegistryEditParamsProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryEditParamsProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryEditParamsProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryEditParamsProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryEditParamsProvenanceA1SInternal, SettingImpersonationRegistryEditParamsProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryEditParamsProvenanceSnoopyOffice365, SettingImpersonationRegistryEditParamsProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}

type SettingImpersonationRegistryEditResponseEnvelope struct {
	Errors   []SettingImpersonationRegistryEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingImpersonationRegistryEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingImpersonationRegistryEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// An impersonation registry entry
	Result SettingImpersonationRegistryEditResponse             `json:"result"`
	JSON   settingImpersonationRegistryEditResponseEnvelopeJSON `json:"-"`
}

// settingImpersonationRegistryEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingImpersonationRegistryEditResponseEnvelope]
type settingImpersonationRegistryEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryEditResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           SettingImpersonationRegistryEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingImpersonationRegistryEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingImpersonationRegistryEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingImpersonationRegistryEditResponseEnvelopeErrors]
type settingImpersonationRegistryEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    settingImpersonationRegistryEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingImpersonationRegistryEditResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingImpersonationRegistryEditResponseEnvelopeErrorsSource]
type settingImpersonationRegistryEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryEditResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           SettingImpersonationRegistryEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingImpersonationRegistryEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingImpersonationRegistryEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingImpersonationRegistryEditResponseEnvelopeMessages]
type settingImpersonationRegistryEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    settingImpersonationRegistryEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingImpersonationRegistryEditResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SettingImpersonationRegistryEditResponseEnvelopeMessagesSource]
type settingImpersonationRegistryEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingImpersonationRegistryEditResponseEnvelopeSuccess bool

const (
	SettingImpersonationRegistryEditResponseEnvelopeSuccessTrue SettingImpersonationRegistryEditResponseEnvelopeSuccess = true
)

func (r SettingImpersonationRegistryEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingImpersonationRegistryGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingImpersonationRegistryGetResponseEnvelope struct {
	Errors   []SettingImpersonationRegistryGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingImpersonationRegistryGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingImpersonationRegistryGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// An impersonation registry entry
	Result SettingImpersonationRegistryGetResponse             `json:"result"`
	JSON   settingImpersonationRegistryGetResponseEnvelopeJSON `json:"-"`
}

// settingImpersonationRegistryGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingImpersonationRegistryGetResponseEnvelope]
type settingImpersonationRegistryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryGetResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           SettingImpersonationRegistryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingImpersonationRegistryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingImpersonationRegistryGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingImpersonationRegistryGetResponseEnvelopeErrors]
type settingImpersonationRegistryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    settingImpersonationRegistryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingImpersonationRegistryGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingImpersonationRegistryGetResponseEnvelopeErrorsSource]
type settingImpersonationRegistryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryGetResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           SettingImpersonationRegistryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingImpersonationRegistryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingImpersonationRegistryGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingImpersonationRegistryGetResponseEnvelopeMessages]
type settingImpersonationRegistryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingImpersonationRegistryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    settingImpersonationRegistryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingImpersonationRegistryGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SettingImpersonationRegistryGetResponseEnvelopeMessagesSource]
type settingImpersonationRegistryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingImpersonationRegistryGetResponseEnvelopeSuccess bool

const (
	SettingImpersonationRegistryGetResponseEnvelopeSuccessTrue SettingImpersonationRegistryGetResponseEnvelopeSuccess = true
)

func (r SettingImpersonationRegistryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
