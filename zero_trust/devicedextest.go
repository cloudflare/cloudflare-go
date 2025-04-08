// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
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
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/dex_tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a DEX test.
func (r *DeviceDEXTestService) Update(ctx context.Context, dexTestID string, params DeviceDEXTestUpdateParams, opts ...option.RequestOption) (res *SchemaHTTP, err error) {
	var env DeviceDEXTestUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/dex_tests/%s", params.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all DEX tests.
func (r *DeviceDEXTestService) List(ctx context.Context, query DeviceDEXTestListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SchemaHTTP], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/dex_tests", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *DeviceDEXTestService) ListAutoPaging(ctx context.Context, query DeviceDEXTestListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SchemaHTTP] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *DeviceDEXTestService) Delete(ctx context.Context, dexTestID string, body DeviceDEXTestDeleteParams, opts ...option.RequestOption) (res *DeviceDEXTestDeleteResponse, err error) {
	var env DeviceDEXTestDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/dex_tests/%s", body.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single DEX test.
func (r *DeviceDEXTestService) Get(ctx context.Context, dexTestID string, query DeviceDEXTestGetParams, opts ...option.RequestOption) (res *SchemaHTTP, err error) {
	var env DeviceDEXTestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/dex_tests/%s", query.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type SchemaData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string         `json:"method"`
	JSON   schemaDataJSON `json:"-"`
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

// The configuration object which contains the details for the WARP client to
// conduct the test.
type SchemaDataParam struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r SchemaDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SchemaHTTP struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data SchemaData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string `json:"description"`
	// Device settings profiles targeted by this test
	TargetPolicies []SchemaHTTPTargetPolicy `json:"target_policies"`
	Targeted       bool                     `json:"targeted"`
	// The unique identifier for the test.
	TestID string         `json:"test_id"`
	JSON   schemaHTTPJSON `json:"-"`
}

// schemaHTTPJSON contains the JSON metadata for the struct [SchemaHTTP]
type schemaHTTPJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
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
	// The id of the device settings profile
	ID string `json:"id"`
	// Whether the profile is the account default
	Default bool `json:"default"`
	// The name of the device settings profile
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
	Data param.Field[SchemaDataParam] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
	// Device settings profiles targeted by this test
	TargetPolicies param.Field[[]SchemaHTTPTargetPolicyParam] `json:"target_policies"`
	Targeted       param.Field[bool]                          `json:"targeted"`
}

func (r SchemaHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SchemaHTTPTargetPolicyParam struct {
	// The id of the device settings profile
	ID param.Field[string] `json:"id"`
	// Whether the profile is the account default
	Default param.Field[bool] `json:"default"`
	// The name of the device settings profile
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
	AccountID  param.Field[string] `path:"account_id,required"`
	SchemaHTTP SchemaHTTPParam     `json:"schema_http,required"`
}

func (r DeviceDEXTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SchemaHTTP)
}

type DeviceDEXTestNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SchemaHTTP            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceDEXTestNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestNewResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDEXTestNewResponseEnvelope]
type deviceDEXTestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeJSON) RawJSON() string {
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
	AccountID  param.Field[string] `path:"account_id,required"`
	SchemaHTTP SchemaHTTPParam     `json:"schema_http,required"`
}

func (r DeviceDEXTestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SchemaHTTP)
}

type DeviceDEXTestUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SchemaHTTP            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceDEXTestUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestUpdateResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDEXTestUpdateResponseEnvelope]
type deviceDEXTestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeJSON) RawJSON() string {
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDEXTestDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDEXTestDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   DeviceDEXTestDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success DeviceDEXTestDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestDeleteResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeleteResponseEnvelope]
type deviceDEXTestDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeJSON) RawJSON() string {
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDEXTestGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SchemaHTTP            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceDEXTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestGetResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDEXTestGetResponseEnvelope]
type deviceDEXTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeJSON) RawJSON() string {
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
