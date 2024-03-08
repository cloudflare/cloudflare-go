// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustDeviceDEXTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDeviceDEXTestService]
// method instead.
type ZeroTrustDeviceDEXTestService struct {
	Options []option.RequestOption
}

// NewZeroTrustDeviceDEXTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDeviceDEXTestService(opts ...option.RequestOption) (r *ZeroTrustDeviceDEXTestService) {
	r = &ZeroTrustDeviceDEXTestService{}
	r.Options = opts
	return
}

// Create a DEX test.
func (r *ZeroTrustDeviceDEXTestService) New(ctx context.Context, params ZeroTrustDeviceDEXTestNewParams, opts ...option.RequestOption) (res *ZeroTrustDeviceDEXTestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceDEXTestNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a DEX test.
func (r *ZeroTrustDeviceDEXTestService) Update(ctx context.Context, dexTestID string, params ZeroTrustDeviceDEXTestUpdateParams, opts ...option.RequestOption) (res *ZeroTrustDeviceDEXTestUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceDEXTestUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", params.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all DEX tests.
func (r *ZeroTrustDeviceDEXTestService) List(ctx context.Context, query ZeroTrustDeviceDEXTestListParams, opts ...option.RequestOption) (res *[]ZeroTrustDeviceDEXTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceDEXTestListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *ZeroTrustDeviceDEXTestService) Delete(ctx context.Context, dexTestID string, body ZeroTrustDeviceDEXTestDeleteParams, opts ...option.RequestOption) (res *[]ZeroTrustDeviceDEXTestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceDEXTestDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", body.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single DEX test.
func (r *ZeroTrustDeviceDEXTestService) Get(ctx context.Context, dexTestID string, query ZeroTrustDeviceDEXTestGetParams, opts ...option.RequestOption) (res *ZeroTrustDeviceDEXTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceDEXTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", query.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDeviceDEXTestNewResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data ZeroTrustDeviceDEXTestNewResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                `json:"description"`
	JSON        zeroTrustDeviceDEXTestNewResponseJSON `json:"-"`
}

// zeroTrustDeviceDEXTestNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceDEXTestNewResponse]
type zeroTrustDeviceDEXTestNewResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestNewResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type ZeroTrustDeviceDEXTestNewResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                    `json:"method"`
	JSON   zeroTrustDeviceDEXTestNewResponseDataJSON `json:"-"`
}

// zeroTrustDeviceDEXTestNewResponseDataJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestNewResponseData]
type zeroTrustDeviceDEXTestNewResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestUpdateResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data ZeroTrustDeviceDEXTestUpdateResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                   `json:"description"`
	JSON        zeroTrustDeviceDEXTestUpdateResponseJSON `json:"-"`
}

// zeroTrustDeviceDEXTestUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestUpdateResponse]
type zeroTrustDeviceDEXTestUpdateResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type ZeroTrustDeviceDEXTestUpdateResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                       `json:"method"`
	JSON   zeroTrustDeviceDEXTestUpdateResponseDataJSON `json:"-"`
}

// zeroTrustDeviceDEXTestUpdateResponseDataJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestUpdateResponseData]
type zeroTrustDeviceDEXTestUpdateResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestUpdateResponseDataJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestListResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data ZeroTrustDeviceDEXTestListResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                 `json:"description"`
	JSON        zeroTrustDeviceDEXTestListResponseJSON `json:"-"`
}

// zeroTrustDeviceDEXTestListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceDEXTestListResponse]
type zeroTrustDeviceDEXTestListResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestListResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type ZeroTrustDeviceDEXTestListResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                     `json:"method"`
	JSON   zeroTrustDeviceDEXTestListResponseDataJSON `json:"-"`
}

// zeroTrustDeviceDEXTestListResponseDataJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestListResponseData]
type zeroTrustDeviceDEXTestListResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestListResponseDataJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestDeleteResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data ZeroTrustDeviceDEXTestDeleteResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                   `json:"description"`
	JSON        zeroTrustDeviceDEXTestDeleteResponseJSON `json:"-"`
}

// zeroTrustDeviceDEXTestDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestDeleteResponse]
type zeroTrustDeviceDEXTestDeleteResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type ZeroTrustDeviceDEXTestDeleteResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                       `json:"method"`
	JSON   zeroTrustDeviceDEXTestDeleteResponseDataJSON `json:"-"`
}

// zeroTrustDeviceDEXTestDeleteResponseDataJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestDeleteResponseData]
type zeroTrustDeviceDEXTestDeleteResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestDeleteResponseDataJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestGetResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data ZeroTrustDeviceDEXTestGetResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                `json:"description"`
	JSON        zeroTrustDeviceDEXTestGetResponseJSON `json:"-"`
}

// zeroTrustDeviceDEXTestGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceDEXTestGetResponse]
type zeroTrustDeviceDEXTestGetResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestGetResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type ZeroTrustDeviceDEXTestGetResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                    `json:"method"`
	JSON   zeroTrustDeviceDEXTestGetResponseDataJSON `json:"-"`
}

// zeroTrustDeviceDEXTestGetResponseDataJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestGetResponseData]
type zeroTrustDeviceDEXTestGetResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[ZeroTrustDeviceDEXTestNewParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r ZeroTrustDeviceDEXTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type ZeroTrustDeviceDEXTestNewParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r ZeroTrustDeviceDEXTestNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDeviceDEXTestNewResponseEnvelope struct {
	Errors   []ZeroTrustDeviceDEXTestNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceDEXTestNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceDEXTestNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceDEXTestNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceDEXTestNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceDEXTestNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestNewResponseEnvelope]
type zeroTrustDeviceDEXTestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceDEXTestNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceDEXTestNewResponseEnvelopeErrors]
type zeroTrustDeviceDEXTestNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceDEXTestNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceDEXTestNewResponseEnvelopeMessages]
type zeroTrustDeviceDEXTestNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceDEXTestNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceDEXTestNewResponseEnvelopeSuccessTrue ZeroTrustDeviceDEXTestNewResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceDEXTestUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[ZeroTrustDeviceDEXTestUpdateParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r ZeroTrustDeviceDEXTestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type ZeroTrustDeviceDEXTestUpdateParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r ZeroTrustDeviceDEXTestUpdateParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDeviceDEXTestUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDeviceDEXTestUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceDEXTestUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceDEXTestUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceDEXTestUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceDEXTestUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceDEXTestUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceDEXTestUpdateResponseEnvelope]
type zeroTrustDeviceDEXTestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceDEXTestUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceDEXTestUpdateResponseEnvelopeErrors]
type zeroTrustDeviceDEXTestUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceDEXTestUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceDEXTestUpdateResponseEnvelopeMessages]
type zeroTrustDeviceDEXTestUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceDEXTestUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceDEXTestUpdateResponseEnvelopeSuccessTrue ZeroTrustDeviceDEXTestUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceDEXTestListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceDEXTestListResponseEnvelope struct {
	Errors   []ZeroTrustDeviceDEXTestListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceDEXTestListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDeviceDEXTestListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceDEXTestListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceDEXTestListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceDEXTestListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceDEXTestListResponseEnvelope]
type zeroTrustDeviceDEXTestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestListResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceDEXTestListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceDEXTestListResponseEnvelopeErrors]
type zeroTrustDeviceDEXTestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestListResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceDEXTestListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceDEXTestListResponseEnvelopeMessages]
type zeroTrustDeviceDEXTestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceDEXTestListResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceDEXTestListResponseEnvelopeSuccessTrue ZeroTrustDeviceDEXTestListResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceDEXTestDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceDEXTestDeleteResponseEnvelope struct {
	Errors   []ZeroTrustDeviceDEXTestDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceDEXTestDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDeviceDEXTestDeleteResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceDEXTestDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceDEXTestDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceDEXTestDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceDEXTestDeleteResponseEnvelope]
type zeroTrustDeviceDEXTestDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceDEXTestDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceDEXTestDeleteResponseEnvelopeErrors]
type zeroTrustDeviceDEXTestDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceDEXTestDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceDEXTestDeleteResponseEnvelopeMessages]
type zeroTrustDeviceDEXTestDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceDEXTestDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceDEXTestDeleteResponseEnvelopeSuccessTrue ZeroTrustDeviceDEXTestDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceDEXTestGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceDEXTestGetResponseEnvelope struct {
	Errors   []ZeroTrustDeviceDEXTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceDEXTestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceDEXTestGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceDEXTestGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceDEXTestGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceDEXTestGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceDEXTestGetResponseEnvelope]
type zeroTrustDeviceDEXTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceDEXTestGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceDEXTestGetResponseEnvelopeErrors]
type zeroTrustDeviceDEXTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceDEXTestGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDeviceDEXTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceDEXTestGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceDEXTestGetResponseEnvelopeMessages]
type zeroTrustDeviceDEXTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceDEXTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceDEXTestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceDEXTestGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceDEXTestGetResponseEnvelopeSuccessTrue ZeroTrustDeviceDEXTestGetResponseEnvelopeSuccess = true
)
