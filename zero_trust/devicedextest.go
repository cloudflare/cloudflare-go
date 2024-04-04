// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DeviceDEXTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceDEXTestService] method
// instead.
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
func (r *DeviceDEXTestService) New(ctx context.Context, params DeviceDEXTestNewParams, opts ...option.RequestOption) (res *DEXTestSchemasHTTP, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/dex_tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a DEX test.
func (r *DeviceDEXTestService) Update(ctx context.Context, dexTestID string, params DeviceDEXTestUpdateParams, opts ...option.RequestOption) (res *DEXTestSchemasHTTP, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/dex_tests/%s", params.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all DEX tests.
func (r *DeviceDEXTestService) List(ctx context.Context, query DeviceDEXTestListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DEXTestSchemasHTTP], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices/dex_tests", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *DeviceDEXTestService) ListAutoPaging(ctx context.Context, query DeviceDEXTestListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DEXTestSchemasHTTP] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *DeviceDEXTestService) Delete(ctx context.Context, dexTestID string, body DeviceDEXTestDeleteParams, opts ...option.RequestOption) (res *[]DEXTestSchemasHTTP, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/dex_tests/%s", body.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single DEX test.
func (r *DeviceDEXTestService) Get(ctx context.Context, dexTestID string, query DeviceDEXTestGetParams, opts ...option.RequestOption) (res *DEXTestSchemasHTTP, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/dex_tests/%s", query.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXTestSchemasHTTP struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DEXTestSchemasHTTPData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                 `json:"description"`
	JSON        dexTestSchemasHTTPJSON `json:"-"`
}

// dexTestSchemasHTTPJSON contains the JSON metadata for the struct
// [DEXTestSchemasHTTP]
type dexTestSchemasHTTPJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestSchemasHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestSchemasHTTPJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DEXTestSchemasHTTPData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                     `json:"method"`
	JSON   dexTestSchemasHTTPDataJSON `json:"-"`
}

// dexTestSchemasHTTPDataJSON contains the JSON metadata for the struct
// [DEXTestSchemasHTTPData]
type dexTestSchemasHTTPDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestSchemasHTTPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestSchemasHTTPDataJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[DeviceDEXTestNewParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r DeviceDEXTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestNewParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r DeviceDEXTestNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DEXTestSchemasHTTP                                        `json:"result,required,nullable"`
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
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[DeviceDEXTestUpdateParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r DeviceDEXTestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestUpdateParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r DeviceDEXTestUpdateParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DEXTestSchemasHTTP                                        `json:"result,required,nullable"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []DEXTestSchemasHTTP                                      `json:"result,required,nullable"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DEXTestSchemasHTTP                                        `json:"result,required,nullable"`
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
