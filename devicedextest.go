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

// Update a DEX test.
func (r *DeviceDEXTestService) Update(ctx context.Context, identifier interface{}, uuid string, body DeviceDEXTestUpdateParams, opts ...option.RequestOption) (res *DeviceDEXTestUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// Create a DEX test.
func (r *DeviceDEXTestService) DeviceDEXTestNewDeviceDEXTest(ctx context.Context, identifier interface{}, body DeviceDEXTestDeviceDEXTestNewDeviceDEXTestParams, opts ...option.RequestOption) (res *DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all DEX tests.
func (r *DeviceDEXTestService) DeviceDEXTestDetails(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DeviceDEXTestDeviceDEXTestDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDEXTestDeviceDEXTestDetailsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

type DeviceDEXTestUpdateResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestUpdateResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                          `json:"description"`
	JSON        deviceDEXTestUpdateResponseJSON `json:"-"`
}

// deviceDEXTestUpdateResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestUpdateResponse]
type deviceDEXTestUpdateResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestUpdateResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                              `json:"method"`
	JSON   deviceDEXTestUpdateResponseDataJSON `json:"-"`
}

// deviceDEXTestUpdateResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestUpdateResponseData]
type deviceDEXTestUpdateResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
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

type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                                 `json:"description"`
	JSON        deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseJSON contains the JSON
// metadata for the struct [DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponse]
type deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                                     `json:"method"`
	JSON   deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseDataJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseDataJSON contains the JSON
// metadata for the struct [DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseData]
type deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestDeviceDEXTestDetailsResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestDeviceDEXTestDetailsResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                        `json:"description"`
	JSON        deviceDEXTestDeviceDEXTestDetailsResponseJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestDetailsResponseJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeviceDEXTestDetailsResponse]
type deviceDEXTestDeviceDEXTestDetailsResponseJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestDeviceDEXTestDetailsResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                            `json:"method"`
	JSON   deviceDEXTestDeviceDEXTestDetailsResponseDataJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestDetailsResponseDataJSON contains the JSON metadata for
// the struct [DeviceDEXTestDeviceDEXTestDetailsResponseData]
type deviceDEXTestDeviceDEXTestDetailsResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
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

type DeviceDEXTestUpdateParams struct {
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
	Errors   []DeviceDEXTestUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceDEXTestUpdateResponse                   `json:"result,required,nullable"`
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

type DeviceDEXTestUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceDEXTestUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestUpdateResponseEnvelopeErrors]
type deviceDEXTestUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceDEXTestUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDEXTestUpdateResponseEnvelopeMessages]
type deviceDEXTestUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestUpdateResponseEnvelopeSuccess bool

const (
	DeviceDEXTestUpdateResponseEnvelopeSuccessTrue DeviceDEXTestUpdateResponseEnvelopeSuccess = true
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

type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestParams struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[DeviceDEXTestDeviceDEXTestNewDeviceDEXTestParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r DeviceDEXTestDeviceDEXTestNewDeviceDEXTestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r DeviceDEXTestDeviceDEXTestNewDeviceDEXTestParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelope struct {
	Errors   []DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelope]
type deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeErrors]
type deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeMessages]
type deviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeSuccess bool

const (
	DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeSuccessTrue DeviceDEXTestDeviceDEXTestNewDeviceDEXTestResponseEnvelopeSuccess = true
)

type DeviceDEXTestDeviceDEXTestDetailsResponseEnvelope struct {
	Errors   []DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   []DeviceDEXTestDeviceDEXTestDetailsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeJSON contains the JSON metadata
// for the struct [DeviceDEXTestDeviceDEXTestDetailsResponseEnvelope]
type deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeErrors]
type deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeMessages]
type deviceDEXTestDeviceDEXTestDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeSuccess bool

const (
	DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeSuccessTrue DeviceDEXTestDeviceDEXTestDetailsResponseEnvelopeSuccess = true
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
