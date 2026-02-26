// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acm

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

// CustomTrustStoreService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomTrustStoreService] method instead.
type CustomTrustStoreService struct {
	Options []option.RequestOption
}

// NewCustomTrustStoreService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomTrustStoreService(opts ...option.RequestOption) (r *CustomTrustStoreService) {
	r = &CustomTrustStoreService{}
	r.Options = opts
	return
}

// Add Custom Origin Trust Store for a Zone.
func (r *CustomTrustStoreService) New(ctx context.Context, params CustomTrustStoreNewParams, opts ...option.RequestOption) (res *CustomTrustStore, err error) {
	var env CustomTrustStoreNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/custom_trust_store", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Custom Origin Trust Store for a Zone.
func (r *CustomTrustStoreService) List(ctx context.Context, params CustomTrustStoreListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CustomTrustStore], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/custom_trust_store", params.ZoneID)
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

// Get Custom Origin Trust Store for a Zone.
func (r *CustomTrustStoreService) ListAutoPaging(ctx context.Context, params CustomTrustStoreListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CustomTrustStore] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes a CA certificate from the custom origin trust store. Origins using
// certificates signed by this CA will no longer be trusted.
func (r *CustomTrustStoreService) Delete(ctx context.Context, customOriginTrustStoreID string, body CustomTrustStoreDeleteParams, opts ...option.RequestOption) (res *CustomTrustStoreDeleteResponse, err error) {
	var env CustomTrustStoreDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customOriginTrustStoreID == "" {
		err = errors.New("missing required custom_origin_trust_store_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/custom_trust_store/%s", body.ZoneID, customOriginTrustStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves details about a specific certificate in the custom origin trust store,
// including expiration and subject information.
func (r *CustomTrustStoreService) Get(ctx context.Context, customOriginTrustStoreID string, query CustomTrustStoreGetParams, opts ...option.RequestOption) (res *CustomTrustStore, err error) {
	var env CustomTrustStoreGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customOriginTrustStoreID == "" {
		err = errors.New("missing required custom_origin_trust_store_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/custom_trust_store/%s", query.ZoneID, customOriginTrustStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomTrustStore struct {
	// Identifier.
	ID string `json:"id,required"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate string `json:"certificate,required"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on,required" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer,required"`
	// The type of hash used for the certificate.
	Signature string `json:"signature,required"`
	// Status of the zone's custom SSL.
	Status CustomTrustStoreStatus `json:"status,required"`
	// When the certificate was last modified.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// When the certificate was uploaded to Cloudflare.
	UploadedOn time.Time            `json:"uploaded_on,required" format:"date-time"`
	JSON       customTrustStoreJSON `json:"-"`
}

// customTrustStoreJSON contains the JSON metadata for the struct
// [CustomTrustStore]
type customTrustStoreJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	ExpiresOn   apijson.Field
	Issuer      apijson.Field
	Signature   apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	UploadedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreJSON) RawJSON() string {
	return r.raw
}

// Status of the zone's custom SSL.
type CustomTrustStoreStatus string

const (
	CustomTrustStoreStatusInitializing      CustomTrustStoreStatus = "initializing"
	CustomTrustStoreStatusPendingDeployment CustomTrustStoreStatus = "pending_deployment"
	CustomTrustStoreStatusActive            CustomTrustStoreStatus = "active"
	CustomTrustStoreStatusPendingDeletion   CustomTrustStoreStatus = "pending_deletion"
	CustomTrustStoreStatusDeleted           CustomTrustStoreStatus = "deleted"
	CustomTrustStoreStatusExpired           CustomTrustStoreStatus = "expired"
)

func (r CustomTrustStoreStatus) IsKnown() bool {
	switch r {
	case CustomTrustStoreStatusInitializing, CustomTrustStoreStatusPendingDeployment, CustomTrustStoreStatusActive, CustomTrustStoreStatusPendingDeletion, CustomTrustStoreStatusDeleted, CustomTrustStoreStatusExpired:
		return true
	}
	return false
}

type CustomTrustStoreDeleteResponse struct {
	// Identifier.
	ID   string                             `json:"id"`
	JSON customTrustStoreDeleteResponseJSON `json:"-"`
}

// customTrustStoreDeleteResponseJSON contains the JSON metadata for the struct
// [CustomTrustStoreDeleteResponse]
type customTrustStoreDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
}

func (r CustomTrustStoreNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomTrustStoreNewResponseEnvelope struct {
	Errors   []CustomTrustStoreNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomTrustStoreNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomTrustStoreNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomTrustStore                           `json:"result"`
	JSON    customTrustStoreNewResponseEnvelopeJSON    `json:"-"`
}

// customTrustStoreNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomTrustStoreNewResponseEnvelope]
type customTrustStoreNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           CustomTrustStoreNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             customTrustStoreNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// customTrustStoreNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomTrustStoreNewResponseEnvelopeErrors]
type customTrustStoreNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomTrustStoreNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    customTrustStoreNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customTrustStoreNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CustomTrustStoreNewResponseEnvelopeErrorsSource]
type customTrustStoreNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           CustomTrustStoreNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             customTrustStoreNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// customTrustStoreNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomTrustStoreNewResponseEnvelopeMessages]
type customTrustStoreNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomTrustStoreNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    customTrustStoreNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customTrustStoreNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CustomTrustStoreNewResponseEnvelopeMessagesSource]
type customTrustStoreNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomTrustStoreNewResponseEnvelopeSuccess bool

const (
	CustomTrustStoreNewResponseEnvelopeSuccessTrue CustomTrustStoreNewResponseEnvelopeSuccess = true
)

func (r CustomTrustStoreNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomTrustStoreNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomTrustStoreListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Limit to the number of records returned.
	Limit param.Field[int64] `query:"limit"`
	// Offset the results
	Offset param.Field[int64] `query:"offset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of records per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [CustomTrustStoreListParams]'s query parameters as
// `url.Values`.
func (r CustomTrustStoreListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CustomTrustStoreDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomTrustStoreDeleteResponseEnvelope struct {
	Errors   []CustomTrustStoreDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomTrustStoreDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomTrustStoreDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomTrustStoreDeleteResponse                `json:"result"`
	JSON    customTrustStoreDeleteResponseEnvelopeJSON    `json:"-"`
}

// customTrustStoreDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomTrustStoreDeleteResponseEnvelope]
type customTrustStoreDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreDeleteResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           CustomTrustStoreDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             customTrustStoreDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// customTrustStoreDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomTrustStoreDeleteResponseEnvelopeErrors]
type customTrustStoreDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomTrustStoreDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    customTrustStoreDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customTrustStoreDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CustomTrustStoreDeleteResponseEnvelopeErrorsSource]
type customTrustStoreDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreDeleteResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           CustomTrustStoreDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             customTrustStoreDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// customTrustStoreDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CustomTrustStoreDeleteResponseEnvelopeMessages]
type customTrustStoreDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomTrustStoreDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    customTrustStoreDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customTrustStoreDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CustomTrustStoreDeleteResponseEnvelopeMessagesSource]
type customTrustStoreDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomTrustStoreDeleteResponseEnvelopeSuccess bool

const (
	CustomTrustStoreDeleteResponseEnvelopeSuccessTrue CustomTrustStoreDeleteResponseEnvelopeSuccess = true
)

func (r CustomTrustStoreDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomTrustStoreDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomTrustStoreGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomTrustStoreGetResponseEnvelope struct {
	Errors   []CustomTrustStoreGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomTrustStoreGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomTrustStoreGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomTrustStore                           `json:"result"`
	JSON    customTrustStoreGetResponseEnvelopeJSON    `json:"-"`
}

// customTrustStoreGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomTrustStoreGetResponseEnvelope]
type customTrustStoreGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           CustomTrustStoreGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             customTrustStoreGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// customTrustStoreGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomTrustStoreGetResponseEnvelopeErrors]
type customTrustStoreGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomTrustStoreGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    customTrustStoreGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customTrustStoreGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CustomTrustStoreGetResponseEnvelopeErrorsSource]
type customTrustStoreGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           CustomTrustStoreGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             customTrustStoreGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// customTrustStoreGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomTrustStoreGetResponseEnvelopeMessages]
type customTrustStoreGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomTrustStoreGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomTrustStoreGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    customTrustStoreGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customTrustStoreGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CustomTrustStoreGetResponseEnvelopeMessagesSource]
type customTrustStoreGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomTrustStoreGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customTrustStoreGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomTrustStoreGetResponseEnvelopeSuccess bool

const (
	CustomTrustStoreGetResponseEnvelopeSuccessTrue CustomTrustStoreGetResponseEnvelopeSuccess = true
)

func (r CustomTrustStoreGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomTrustStoreGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
