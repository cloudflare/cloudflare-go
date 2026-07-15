// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DeviceDEXTestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceDEXTestService] method instead.
type DeviceDEXTestService struct {
	Options []option.RequestOption
}

// NewDeviceDEXTestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceDEXTestService(opts ...option.RequestOption) (r *DeviceDEXTestService) {
	r = &DeviceDEXTestService{}
	r.Options = opts
	return
}

// Create a DEX test.
func (r *DeviceDEXTestService) New(ctx context.Context, params DeviceDEXTestNewParams, opts ...option.RequestOption) (res *SchemaHTTP, err error) {
	var env DeviceDEXTestNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update a DEX test.
func (r *DeviceDEXTestService) Update(ctx context.Context, dexTestID string, params DeviceDEXTestUpdateParams, opts ...option.RequestOption) (res *SchemaHTTP, err error) {
	var env DeviceDEXTestUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests/%s", params.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetch all DEX tests.
func (r *DeviceDEXTestService) List(ctx context.Context, params DeviceDEXTestListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SchemaHTTP], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests", params.AccountID)
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

// Fetch all DEX tests.
func (r *DeviceDEXTestService) ListAutoPaging(ctx context.Context, params DeviceDEXTestListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SchemaHTTP] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *DeviceDEXTestService) Delete(ctx context.Context, dexTestID string, body DeviceDEXTestDeleteParams, opts ...option.RequestOption) (res *DeviceDEXTestDeleteResponse, err error) {
	var env DeviceDEXTestDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests/%s", body.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetch a single DEX test.
func (r *DeviceDEXTestService) Get(ctx context.Context, dexTestID string, query DeviceDEXTestGetParams, opts ...option.RequestOption) (res *SchemaHTTP, err error) {
	var env DeviceDEXTestGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests/%s", query.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type SchemaData struct {
	// The desired endpoint to test.
	Host string `json:"host" api:"required"`
	// The type of test.
	Kind SchemaDataKind `json:"kind" api:"required"`
	// The HTTP request method type.
	Method SchemaDataMethod `json:"method"`
	JSON   schemaDataJSON   `json:"-"`
}

// schemaDataJSON contains the JSON metadata for the struct [SchemaData]
type schemaDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaDataJSON) RawJSON() string {
	return r.raw
}

// The type of test.
type SchemaDataKind string

const (
	SchemaDataKindHTTP       SchemaDataKind = "http"
	SchemaDataKindTraceroute SchemaDataKind = "traceroute"
)

func (r SchemaDataKind) IsKnown() bool {
	switch r {
	case SchemaDataKindHTTP, SchemaDataKindTraceroute:
		return true
	}
	return false
}

// The HTTP request method type.
type SchemaDataMethod string

const (
	SchemaDataMethodGet SchemaDataMethod = "GET"
)

func (r SchemaDataMethod) IsKnown() bool {
	switch r {
	case SchemaDataMethodGet:
		return true
	}
	return false
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type SchemaDataParam struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host" api:"required"`
	// The type of test.
	Kind param.Field[SchemaDataKind] `json:"kind" api:"required"`
	// The HTTP request method type.
	Method param.Field[SchemaDataMethod] `json:"method"`
}

func (r SchemaDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SchemaHTTP struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data SchemaData `json:"data" api:"required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled" api:"required"`
	// How often the test will run.
	Interval string `json:"interval" api:"required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name" api:"required"`
	// Date the test was created, in RFC 3339 format.
	Created time.Time `json:"created" format:"date-time"`
	// Additional details about the test.
	Description string `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies []SchemaHTTPTargetPolicy `json:"target_policies"`
	Targeted       bool                     `json:"targeted"`
	// The unique identifier for the test.
	TestID string `json:"test_id"`
	// Date the test was last updated, in RFC 3339 format.
	Updated time.Time      `json:"updated" format:"date-time"`
	JSON    schemaHTTPJSON `json:"-"`
}

// schemaHTTPJSON contains the JSON metadata for the struct [SchemaHTTP]
type schemaHTTPJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Created        apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
	Updated        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SchemaHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaHTTPJSON) RawJSON() string {
	return r.raw
}

type SchemaHTTPTargetPolicy struct {
	// The id of the DEX rule.
	ID string `json:"id" api:"required"`
	// Whether the DEX rule is the account default.
	Default bool `json:"default"`
	// The name of the DEX rule.
	Name string                     `json:"name"`
	JSON schemaHTTPTargetPolicyJSON `json:"-"`
}

// schemaHTTPTargetPolicyJSON contains the JSON metadata for the struct
// [SchemaHTTPTargetPolicy]
type schemaHTTPTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaHTTPTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaHTTPTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type SchemaHTTPParam struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[SchemaDataParam] `json:"data" api:"required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval" api:"required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name" api:"required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies param.Field[[]SchemaHTTPTargetPolicyParam] `json:"target_policies"`
	Targeted       param.Field[bool]                          `json:"targeted"`
}

func (r SchemaHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SchemaHTTPTargetPolicyParam struct {
	// The id of the DEX rule.
	ID param.Field[string] `json:"id" api:"required"`
	// Whether the DEX rule is the account default.
	Default param.Field[bool] `json:"default"`
	// The name of the DEX rule.
	Name param.Field[string] `json:"name"`
}

func (r SchemaHTTPTargetPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestDeleteResponse struct {
	DEXTests []SchemaHTTP                    `json:"dex_tests"`
	JSON     deviceDEXTestDeleteResponseJSON `json:"-"`
}

// deviceDEXTestDeleteResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestDeleteResponse]
type deviceDEXTestDeleteResponseJSON struct {
	DEXTests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewParams struct {
	// Unique identifier linked to an account.
	AccountID  param.Field[string] `path:"account_id" api:"required"`
	SchemaHTTP SchemaHTTPParam     `json:"schema_http" api:"required"`
}

func (r DeviceDEXTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SchemaHTTP)
}

type DeviceDEXTestNewResponseEnvelope struct {
	Errors   []DeviceDEXTestNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDEXTestNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DeviceDEXTestNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SchemaHTTP                              `json:"result"`
	JSON    deviceDEXTestNewResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDEXTestNewResponseEnvelope]
type deviceDEXTestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DeviceDEXTestNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestNewResponseEnvelopeErrors]
type deviceDEXTestNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    deviceDEXTestNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DeviceDEXTestNewResponseEnvelopeErrorsSource]
type deviceDEXTestNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DeviceDEXTestNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDEXTestNewResponseEnvelopeMessages]
type deviceDEXTestNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    deviceDEXTestNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestNewResponseEnvelopeMessagesSource]
type deviceDEXTestNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestNewResponseEnvelopeSuccess bool

const (
	DeviceDEXTestNewResponseEnvelopeSuccessTrue DeviceDEXTestNewResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceDEXTestUpdateParams struct {
	// Unique identifier linked to an account.
	AccountID  param.Field[string] `path:"account_id" api:"required"`
	SchemaHTTP SchemaHTTPParam     `json:"schema_http" api:"required"`
}

func (r DeviceDEXTestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SchemaHTTP)
}

type DeviceDEXTestUpdateResponseEnvelope struct {
	Errors   []DeviceDEXTestUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDEXTestUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DeviceDEXTestUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SchemaHTTP                                 `json:"result"`
	JSON    deviceDEXTestUpdateResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDEXTestUpdateResponseEnvelope]
type deviceDEXTestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DeviceDEXTestUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestUpdateResponseEnvelopeErrors]
type deviceDEXTestUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestUpdateResponseEnvelopeErrorsSource]
type deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DeviceDEXTestUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDEXTestUpdateResponseEnvelopeMessages]
type deviceDEXTestUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestUpdateResponseEnvelopeMessagesSource]
type deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestUpdateResponseEnvelopeSuccess bool

const (
	DeviceDEXTestUpdateResponseEnvelopeSuccessTrue DeviceDEXTestUpdateResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceDEXTestListParams struct {
	// Unique identifier linked to an account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by test type.
	Kind param.Field[DeviceDEXTestListParamsKind] `query:"kind"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by test name.
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [DeviceDEXTestListParams]'s query parameters as
// `url.Values`.
func (r DeviceDEXTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by test type.
type DeviceDEXTestListParamsKind string

const (
	DeviceDEXTestListParamsKindHTTP       DeviceDEXTestListParamsKind = "http"
	DeviceDEXTestListParamsKindTraceroute DeviceDEXTestListParamsKind = "traceroute"
)

func (r DeviceDEXTestListParamsKind) IsKnown() bool {
	switch r {
	case DeviceDEXTestListParamsKindHTTP, DeviceDEXTestListParamsKindTraceroute:
		return true
	}
	return false
}

type DeviceDEXTestDeleteParams struct {
	// Unique identifier linked to an account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DeviceDEXTestDeleteResponseEnvelope struct {
	Errors   []DeviceDEXTestDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDEXTestDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DeviceDEXTestDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DeviceDEXTestDeleteResponse                `json:"result"`
	JSON    deviceDEXTestDeleteResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeleteResponseEnvelope]
type deviceDEXTestDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DeviceDEXTestDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeleteResponseEnvelopeErrors]
type deviceDEXTestDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestDeleteResponseEnvelopeErrorsSource]
type deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DeviceDEXTestDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDEXTestDeleteResponseEnvelopeMessages]
type deviceDEXTestDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestDeleteResponseEnvelopeMessagesSource]
type deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestDeleteResponseEnvelopeSuccess bool

const (
	DeviceDEXTestDeleteResponseEnvelopeSuccessTrue DeviceDEXTestDeleteResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceDEXTestGetParams struct {
	// Unique identifier linked to an account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DeviceDEXTestGetResponseEnvelope struct {
	Errors   []DeviceDEXTestGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDEXTestGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DeviceDEXTestGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SchemaHTTP                              `json:"result"`
	JSON    deviceDEXTestGetResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDEXTestGetResponseEnvelope]
type deviceDEXTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DeviceDEXTestGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestGetResponseEnvelopeErrors]
type deviceDEXTestGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    deviceDEXTestGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DeviceDEXTestGetResponseEnvelopeErrorsSource]
type deviceDEXTestGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DeviceDEXTestGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDEXTestGetResponseEnvelopeMessages]
type deviceDEXTestGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    deviceDEXTestGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestGetResponseEnvelopeMessagesSource]
type deviceDEXTestGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestGetResponseEnvelopeSuccess bool

const (
	DeviceDEXTestGetResponseEnvelopeSuccessTrue DeviceDEXTestGetResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
