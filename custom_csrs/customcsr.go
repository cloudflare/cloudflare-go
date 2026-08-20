// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_csrs

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

// CustomCsrService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomCsrService] method instead.
type CustomCsrService struct {
	Options []option.RequestOption
}

// NewCustomCsrService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomCsrService(opts ...option.RequestOption) (r *CustomCsrService) {
	r = &CustomCsrService{}
	r.Options = opts
	return
}

// Generate a new custom Certificate Signing Request (CSR) for an account or zone.
// Cloudflare generates and securely stores the private key associated with the
// CSR.
func (r *CustomCsrService) New(ctx context.Context, params CustomCsrNewParams, opts ...option.RequestOption) (res *CustomCsrNewResponse, err error) {
	var env CustomCsrNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("%s/%s/custom_csrs", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all custom Certificate Signing Requests (CSRs) for an account or zone.
func (r *CustomCsrService) List(ctx context.Context, params CustomCsrListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CustomCsrListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("%s/%s/custom_csrs", accountOrZone, accountOrZoneID)
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

// List all custom Certificate Signing Requests (CSRs) for an account or zone.
func (r *CustomCsrService) ListAutoPaging(ctx context.Context, params CustomCsrListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CustomCsrListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a custom Certificate Signing Request (CSR) and its associated private
// key.
func (r *CustomCsrService) Delete(ctx context.Context, customCsrID string, body CustomCsrDeleteParams, opts ...option.RequestOption) (res *CustomCsrDeleteResponse, err error) {
	var env CustomCsrDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if customCsrID == "" {
		err = errors.New("missing required custom_csr_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/custom_csrs/%s", accountOrZone, accountOrZoneID, customCsrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve details for a specific custom Certificate Signing Request (CSR).
func (r *CustomCsrService) Get(ctx context.Context, customCsrID string, query CustomCsrGetParams, opts ...option.RequestOption) (res *CustomCsrGetResponse, err error) {
	var env CustomCsrGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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
	if customCsrID == "" {
		err = errors.New("missing required custom_csr_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/custom_csrs/%s", accountOrZone, accountOrZoneID, customCsrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A custom Certificate Signing Request (CSR).
type CustomCsrNewResponse struct {
	// Custom CSR identifier tag.
	ID string `json:"id" api:"required"`
	// When the CSR was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The key algorithm used to generate the CSR.
	KeyType CustomCsrNewResponseKeyType `json:"key_type" api:"required"`
	// Account identifier associated with this CSR.
	AccountTag string `json:"account_tag"`
	// The common name (domain) for the CSR.
	CommonName string `json:"common_name"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country"`
	// The PEM-encoded Certificate Signing Request.
	Csr string `json:"csr"`
	// Optional description for the CSR.
	Description string `json:"description"`
	// City or locality name.
	Locality string `json:"locality"`
	// Human-readable name for the CSR.
	Name string `json:"name"`
	// Organization name.
	Organization string `json:"organization"`
	// Organizational unit name.
	OrganizationalUnit string `json:"organizational_unit"`
	// Subject Alternative Names included in the CSR.
	Sans []string `json:"sans"`
	// State or province name.
	State string                   `json:"state"`
	JSON  customCsrNewResponseJSON `json:"-"`
}

// customCsrNewResponseJSON contains the JSON metadata for the struct
// [CustomCsrNewResponse]
type customCsrNewResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	KeyType            apijson.Field
	AccountTag         apijson.Field
	CommonName         apijson.Field
	Country            apijson.Field
	Csr                apijson.Field
	Description        apijson.Field
	Locality           apijson.Field
	Name               apijson.Field
	Organization       apijson.Field
	OrganizationalUnit apijson.Field
	Sans               apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomCsrNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrNewResponseJSON) RawJSON() string {
	return r.raw
}

// The key algorithm used to generate the CSR.
type CustomCsrNewResponseKeyType string

const (
	CustomCsrNewResponseKeyTypeRsa2048 CustomCsrNewResponseKeyType = "rsa2048"
	CustomCsrNewResponseKeyTypeP256v1  CustomCsrNewResponseKeyType = "p256v1"
)

func (r CustomCsrNewResponseKeyType) IsKnown() bool {
	switch r {
	case CustomCsrNewResponseKeyTypeRsa2048, CustomCsrNewResponseKeyTypeP256v1:
		return true
	}
	return false
}

// A custom Certificate Signing Request (CSR).
type CustomCsrListResponse struct {
	// Custom CSR identifier tag.
	ID string `json:"id" api:"required"`
	// When the CSR was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The key algorithm used to generate the CSR.
	KeyType CustomCsrListResponseKeyType `json:"key_type" api:"required"`
	// Account identifier associated with this CSR.
	AccountTag string `json:"account_tag"`
	// The common name (domain) for the CSR.
	CommonName string `json:"common_name"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country"`
	// The PEM-encoded Certificate Signing Request.
	Csr string `json:"csr"`
	// Optional description for the CSR.
	Description string `json:"description"`
	// City or locality name.
	Locality string `json:"locality"`
	// Human-readable name for the CSR.
	Name string `json:"name"`
	// Organization name.
	Organization string `json:"organization"`
	// Organizational unit name.
	OrganizationalUnit string `json:"organizational_unit"`
	// Subject Alternative Names included in the CSR.
	Sans []string `json:"sans"`
	// State or province name.
	State string                    `json:"state"`
	JSON  customCsrListResponseJSON `json:"-"`
}

// customCsrListResponseJSON contains the JSON metadata for the struct
// [CustomCsrListResponse]
type customCsrListResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	KeyType            apijson.Field
	AccountTag         apijson.Field
	CommonName         apijson.Field
	Country            apijson.Field
	Csr                apijson.Field
	Description        apijson.Field
	Locality           apijson.Field
	Name               apijson.Field
	Organization       apijson.Field
	OrganizationalUnit apijson.Field
	Sans               apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomCsrListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrListResponseJSON) RawJSON() string {
	return r.raw
}

// The key algorithm used to generate the CSR.
type CustomCsrListResponseKeyType string

const (
	CustomCsrListResponseKeyTypeRsa2048 CustomCsrListResponseKeyType = "rsa2048"
	CustomCsrListResponseKeyTypeP256v1  CustomCsrListResponseKeyType = "p256v1"
)

func (r CustomCsrListResponseKeyType) IsKnown() bool {
	switch r {
	case CustomCsrListResponseKeyTypeRsa2048, CustomCsrListResponseKeyTypeP256v1:
		return true
	}
	return false
}

type CustomCsrDeleteResponse struct {
	// Custom CSR identifier tag.
	ID   string                      `json:"id"`
	JSON customCsrDeleteResponseJSON `json:"-"`
}

// customCsrDeleteResponseJSON contains the JSON metadata for the struct
// [CustomCsrDeleteResponse]
type customCsrDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A custom Certificate Signing Request (CSR).
type CustomCsrGetResponse struct {
	// Custom CSR identifier tag.
	ID string `json:"id" api:"required"`
	// When the CSR was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The key algorithm used to generate the CSR.
	KeyType CustomCsrGetResponseKeyType `json:"key_type" api:"required"`
	// Account identifier associated with this CSR.
	AccountTag string `json:"account_tag"`
	// The common name (domain) for the CSR.
	CommonName string `json:"common_name"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country"`
	// The PEM-encoded Certificate Signing Request.
	Csr string `json:"csr"`
	// Optional description for the CSR.
	Description string `json:"description"`
	// City or locality name.
	Locality string `json:"locality"`
	// Human-readable name for the CSR.
	Name string `json:"name"`
	// Organization name.
	Organization string `json:"organization"`
	// Organizational unit name.
	OrganizationalUnit string `json:"organizational_unit"`
	// Subject Alternative Names included in the CSR.
	Sans []string `json:"sans"`
	// State or province name.
	State string                   `json:"state"`
	JSON  customCsrGetResponseJSON `json:"-"`
}

// customCsrGetResponseJSON contains the JSON metadata for the struct
// [CustomCsrGetResponse]
type customCsrGetResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	KeyType            apijson.Field
	AccountTag         apijson.Field
	CommonName         apijson.Field
	Country            apijson.Field
	Csr                apijson.Field
	Description        apijson.Field
	Locality           apijson.Field
	Name               apijson.Field
	Organization       apijson.Field
	OrganizationalUnit apijson.Field
	Sans               apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomCsrGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrGetResponseJSON) RawJSON() string {
	return r.raw
}

// The key algorithm used to generate the CSR.
type CustomCsrGetResponseKeyType string

const (
	CustomCsrGetResponseKeyTypeRsa2048 CustomCsrGetResponseKeyType = "rsa2048"
	CustomCsrGetResponseKeyTypeP256v1  CustomCsrGetResponseKeyType = "p256v1"
)

func (r CustomCsrGetResponseKeyType) IsKnown() bool {
	switch r {
	case CustomCsrGetResponseKeyTypeRsa2048, CustomCsrGetResponseKeyTypeP256v1:
		return true
	}
	return false
}

type CustomCsrNewParams struct {
	// The common name (domain) for the CSR. Must be at most 64 characters.
	CommonName param.Field[string] `json:"common_name" api:"required"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	Country param.Field[string] `json:"country" api:"required"`
	// City or locality name.
	Locality param.Field[string] `json:"locality" api:"required"`
	// Organization name.
	Organization param.Field[string] `json:"organization" api:"required"`
	// Subject Alternative Names for the CSR. At least one SAN is required.
	Sans param.Field[[]string] `json:"sans" api:"required"`
	// State or province name.
	State param.Field[string] `json:"state" api:"required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Optional description for the CSR.
	Description param.Field[string] `json:"description"`
	// Key algorithm to use for the CSR. Defaults to rsa2048 if not specified.
	KeyType param.Field[CustomCsrNewParamsKeyType] `json:"key_type"`
	// Human-readable name for the CSR.
	Name param.Field[string] `json:"name"`
	// Organizational unit name.
	OrganizationalUnit param.Field[string] `json:"organizational_unit"`
}

func (r CustomCsrNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Key algorithm to use for the CSR. Defaults to rsa2048 if not specified.
type CustomCsrNewParamsKeyType string

const (
	CustomCsrNewParamsKeyTypeRsa2048 CustomCsrNewParamsKeyType = "rsa2048"
	CustomCsrNewParamsKeyTypeP256v1  CustomCsrNewParamsKeyType = "p256v1"
)

func (r CustomCsrNewParamsKeyType) IsKnown() bool {
	switch r {
	case CustomCsrNewParamsKeyTypeRsa2048, CustomCsrNewParamsKeyTypeP256v1:
		return true
	}
	return false
}

type CustomCsrNewResponseEnvelope struct {
	Errors   []CustomCsrNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CustomCsrNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success CustomCsrNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A custom Certificate Signing Request (CSR).
	Result CustomCsrNewResponse             `json:"result"`
	JSON   customCsrNewResponseEnvelopeJSON `json:"-"`
}

// customCsrNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomCsrNewResponseEnvelope]
type customCsrNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomCsrNewResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           CustomCsrNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             customCsrNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// customCsrNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CustomCsrNewResponseEnvelopeErrors]
type customCsrNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomCsrNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomCsrNewResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    customCsrNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customCsrNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [CustomCsrNewResponseEnvelopeErrorsSource]
type customCsrNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomCsrNewResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           CustomCsrNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             customCsrNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// customCsrNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomCsrNewResponseEnvelopeMessages]
type customCsrNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomCsrNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomCsrNewResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    customCsrNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customCsrNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CustomCsrNewResponseEnvelopeMessagesSource]
type customCsrNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomCsrNewResponseEnvelopeSuccess bool

const (
	CustomCsrNewResponseEnvelopeSuccessTrue CustomCsrNewResponseEnvelopeSuccess = true
)

func (r CustomCsrNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomCsrNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomCsrListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of custom CSRs per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [CustomCsrListParams]'s query parameters as `url.Values`.
func (r CustomCsrListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CustomCsrDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type CustomCsrDeleteResponseEnvelope struct {
	Errors   []CustomCsrDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CustomCsrDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success CustomCsrDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  CustomCsrDeleteResponse                `json:"result"`
	JSON    customCsrDeleteResponseEnvelopeJSON    `json:"-"`
}

// customCsrDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomCsrDeleteResponseEnvelope]
type customCsrDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomCsrDeleteResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           CustomCsrDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             customCsrDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// customCsrDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomCsrDeleteResponseEnvelopeErrors]
type customCsrDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomCsrDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomCsrDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    customCsrDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customCsrDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CustomCsrDeleteResponseEnvelopeErrorsSource]
type customCsrDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomCsrDeleteResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           CustomCsrDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             customCsrDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// customCsrDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomCsrDeleteResponseEnvelopeMessages]
type customCsrDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomCsrDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomCsrDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    customCsrDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customCsrDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CustomCsrDeleteResponseEnvelopeMessagesSource]
type customCsrDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomCsrDeleteResponseEnvelopeSuccess bool

const (
	CustomCsrDeleteResponseEnvelopeSuccessTrue CustomCsrDeleteResponseEnvelopeSuccess = true
)

func (r CustomCsrDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomCsrDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomCsrGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type CustomCsrGetResponseEnvelope struct {
	Errors   []CustomCsrGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CustomCsrGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success CustomCsrGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A custom Certificate Signing Request (CSR).
	Result CustomCsrGetResponse             `json:"result"`
	JSON   customCsrGetResponseEnvelopeJSON `json:"-"`
}

// customCsrGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomCsrGetResponseEnvelope]
type customCsrGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomCsrGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           CustomCsrGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             customCsrGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// customCsrGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CustomCsrGetResponseEnvelopeErrors]
type customCsrGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomCsrGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomCsrGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    customCsrGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customCsrGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [CustomCsrGetResponseEnvelopeErrorsSource]
type customCsrGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomCsrGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           CustomCsrGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             customCsrGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// customCsrGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomCsrGetResponseEnvelopeMessages]
type customCsrGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomCsrGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomCsrGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    customCsrGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customCsrGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CustomCsrGetResponseEnvelopeMessagesSource]
type customCsrGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCsrGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCsrGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomCsrGetResponseEnvelopeSuccess bool

const (
	CustomCsrGetResponseEnvelopeSuccessTrue CustomCsrGetResponseEnvelopeSuccess = true
)

func (r CustomCsrGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomCsrGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
