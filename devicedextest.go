// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *DeviceDEXTestService) New(ctx context.Context, identifier interface{}, body DeviceDEXTestNewParams, opts ...option.RequestOption) (res *DeviceDEXTestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all DEX tests.
func (r *DeviceDEXTestService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DeviceDEXTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *DeviceDEXTestService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DeviceDEXTestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single DEX test.
func (r *DeviceDEXTestService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DeviceDEXTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a DEX test.
func (r *DeviceDEXTestService) Replace(ctx context.Context, identifier interface{}, uuid string, body DeviceDEXTestReplaceParams, opts ...option.RequestOption) (res *DeviceDEXTestReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceDEXTestNewResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestNewResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                       `json:"description"`
	JSON        deviceDEXTestNewResponseJSON `json:"-"`
}

// deviceDEXTestNewResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestNewResponse]
type deviceDEXTestNewResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestNewResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                           `json:"method"`
	JSON   deviceDEXTestNewResponseDataJSON `json:"-"`
}

// deviceDEXTestNewResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestNewResponseData]
type deviceDEXTestNewResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestListResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestListResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                        `json:"description"`
	JSON        deviceDEXTestListResponseJSON `json:"-"`
}

// deviceDEXTestListResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestListResponse]
type deviceDEXTestListResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestListResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                            `json:"method"`
	JSON   deviceDEXTestListResponseDataJSON `json:"-"`
}

// deviceDEXTestListResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestListResponseData]
type deviceDEXTestListResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestDeleteResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestDeleteResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                          `json:"description"`
	JSON        deviceDEXTestDeleteResponseJSON `json:"-"`
}

// deviceDEXTestDeleteResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestDeleteResponse]
type deviceDEXTestDeleteResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestDeleteResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                              `json:"method"`
	JSON   deviceDEXTestDeleteResponseDataJSON `json:"-"`
}

// deviceDEXTestDeleteResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestDeleteResponseData]
type deviceDEXTestDeleteResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestGetResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestGetResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                       `json:"description"`
	JSON        deviceDEXTestGetResponseJSON `json:"-"`
}

// deviceDEXTestGetResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestGetResponse]
type deviceDEXTestGetResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestGetResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                           `json:"method"`
	JSON   deviceDEXTestGetResponseDataJSON `json:"-"`
}

// deviceDEXTestGetResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestGetResponseData]
type deviceDEXTestGetResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestReplaceResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestReplaceResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                           `json:"description"`
	JSON        deviceDEXTestReplaceResponseJSON `json:"-"`
}

// deviceDEXTestReplaceResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestReplaceResponse]
type deviceDEXTestReplaceResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestReplaceResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                               `json:"method"`
	JSON   deviceDEXTestReplaceResponseDataJSON `json:"-"`
}

// deviceDEXTestReplaceResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestReplaceResponseData]
type deviceDEXTestReplaceResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestReplaceResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestNewParams struct {
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
	Errors   []DeviceDEXTestNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceDEXTestNewResponse                   `json:"result,required,nullable"`
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

type DeviceDEXTestNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    deviceDEXTestNewResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestNewResponseEnvelopeErrors]
type deviceDEXTestNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    deviceDEXTestNewResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDEXTestNewResponseEnvelopeMessages]
type deviceDEXTestNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestNewResponseEnvelopeSuccess bool

const (
	DeviceDEXTestNewResponseEnvelopeSuccessTrue DeviceDEXTestNewResponseEnvelopeSuccess = true
)

type DeviceDEXTestListResponseEnvelope struct {
	Errors   []DeviceDEXTestListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DeviceDEXTestListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceDEXTestListResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestListResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDEXTestListResponseEnvelope]
type deviceDEXTestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    deviceDEXTestListResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestListResponseEnvelopeErrors]
type deviceDEXTestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceDEXTestListResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDEXTestListResponseEnvelopeMessages]
type deviceDEXTestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestListResponseEnvelopeSuccess bool

const (
	DeviceDEXTestListResponseEnvelopeSuccessTrue DeviceDEXTestListResponseEnvelopeSuccess = true
)

type DeviceDEXTestDeleteResponseEnvelope struct {
	Errors   []DeviceDEXTestDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []DeviceDEXTestDeleteResponse                 `json:"result,required,nullable"`
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

type DeviceDEXTestDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceDEXTestDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeleteResponseEnvelopeErrors]
type deviceDEXTestDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceDEXTestDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDEXTestDeleteResponseEnvelopeMessages]
type deviceDEXTestDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestDeleteResponseEnvelopeSuccess bool

const (
	DeviceDEXTestDeleteResponseEnvelopeSuccessTrue DeviceDEXTestDeleteResponseEnvelopeSuccess = true
)

type DeviceDEXTestGetResponseEnvelope struct {
	Errors   []DeviceDEXTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceDEXTestGetResponse                   `json:"result,required,nullable"`
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

type DeviceDEXTestGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    deviceDEXTestGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestGetResponseEnvelopeErrors]
type deviceDEXTestGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    deviceDEXTestGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDEXTestGetResponseEnvelopeMessages]
type deviceDEXTestGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestGetResponseEnvelopeSuccess bool

const (
	DeviceDEXTestGetResponseEnvelopeSuccessTrue DeviceDEXTestGetResponseEnvelopeSuccess = true
)

type DeviceDEXTestReplaceParams struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[DeviceDEXTestReplaceParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r DeviceDEXTestReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestReplaceParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r DeviceDEXTestReplaceParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestReplaceResponseEnvelope struct {
	Errors   []DeviceDEXTestReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceDEXTestReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceDEXTestReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestReplaceResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDEXTestReplaceResponseEnvelope]
type deviceDEXTestReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestReplaceResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    deviceDEXTestReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeviceDEXTestReplaceResponseEnvelopeErrors]
type deviceDEXTestReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestReplaceResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    deviceDEXTestReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDEXTestReplaceResponseEnvelopeMessages]
type deviceDEXTestReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestReplaceResponseEnvelopeSuccess bool

const (
	DeviceDEXTestReplaceResponseEnvelopeSuccessTrue DeviceDEXTestReplaceResponseEnvelopeSuccess = true
)
