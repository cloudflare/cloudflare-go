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

// AccountDeviceDexTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDeviceDexTestService]
// method instead.
type AccountDeviceDexTestService struct {
	Options []option.RequestOption
}

// NewAccountDeviceDexTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceDexTestService(opts ...option.RequestOption) (r *AccountDeviceDexTestService) {
	r = &AccountDeviceDexTestService{}
	r.Options = opts
	return
}

// Fetch a single DEX test.
func (r *AccountDeviceDexTestService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDeviceDexTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a DEX test.
func (r *AccountDeviceDexTestService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDeviceDexTestUpdateParams, opts ...option.RequestOption) (res *AccountDeviceDexTestUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *AccountDeviceDexTestService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDeviceDexTestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a DEX test.
func (r *AccountDeviceDexTestService) DeviceDexTestNewDeviceDexTest(ctx context.Context, identifier interface{}, body AccountDeviceDexTestDeviceDexTestNewDeviceDexTestParams, opts ...option.RequestOption) (res *AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch all DEX tests.
func (r *AccountDeviceDexTestService) DeviceDexTestDetails(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDeviceDexTestDeviceDexTestDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDeviceDexTestGetResponse struct {
	Errors   []AccountDeviceDexTestGetResponseError   `json:"errors"`
	Messages []AccountDeviceDexTestGetResponseMessage `json:"messages"`
	Result   AccountDeviceDexTestGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceDexTestGetResponseSuccess `json:"success"`
	JSON    accountDeviceDexTestGetResponseJSON    `json:"-"`
}

// accountDeviceDexTestGetResponseJSON contains the JSON metadata for the struct
// [AccountDeviceDexTestGetResponse]
type accountDeviceDexTestGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountDeviceDexTestGetResponseErrorJSON `json:"-"`
}

// accountDeviceDexTestGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestGetResponseError]
type accountDeviceDexTestGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountDeviceDexTestGetResponseMessageJSON `json:"-"`
}

// accountDeviceDexTestGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestGetResponseMessage]
type accountDeviceDexTestGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestGetResponseResult struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data AccountDeviceDexTestGetResponseResultData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                    `json:"description"`
	JSON        accountDeviceDexTestGetResponseResultJSON `json:"-"`
}

// accountDeviceDexTestGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestGetResponseResult]
type accountDeviceDexTestGetResponseResultJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type AccountDeviceDexTestGetResponseResultData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                        `json:"method"`
	JSON   accountDeviceDexTestGetResponseResultDataJSON `json:"-"`
}

// accountDeviceDexTestGetResponseResultDataJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestGetResponseResultData]
type accountDeviceDexTestGetResponseResultDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestGetResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceDexTestGetResponseSuccess bool

const (
	AccountDeviceDexTestGetResponseSuccessTrue AccountDeviceDexTestGetResponseSuccess = true
)

type AccountDeviceDexTestUpdateResponse struct {
	Errors   []AccountDeviceDexTestUpdateResponseError   `json:"errors"`
	Messages []AccountDeviceDexTestUpdateResponseMessage `json:"messages"`
	Result   AccountDeviceDexTestUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceDexTestUpdateResponseSuccess `json:"success"`
	JSON    accountDeviceDexTestUpdateResponseJSON    `json:"-"`
}

// accountDeviceDexTestUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDeviceDexTestUpdateResponse]
type accountDeviceDexTestUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDeviceDexTestUpdateResponseErrorJSON `json:"-"`
}

// accountDeviceDexTestUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestUpdateResponseError]
type accountDeviceDexTestUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountDeviceDexTestUpdateResponseMessageJSON `json:"-"`
}

// accountDeviceDexTestUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestUpdateResponseMessage]
type accountDeviceDexTestUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestUpdateResponseResult struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data AccountDeviceDexTestUpdateResponseResultData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                       `json:"description"`
	JSON        accountDeviceDexTestUpdateResponseResultJSON `json:"-"`
}

// accountDeviceDexTestUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestUpdateResponseResult]
type accountDeviceDexTestUpdateResponseResultJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type AccountDeviceDexTestUpdateResponseResultData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                           `json:"method"`
	JSON   accountDeviceDexTestUpdateResponseResultDataJSON `json:"-"`
}

// accountDeviceDexTestUpdateResponseResultDataJSON contains the JSON metadata for
// the struct [AccountDeviceDexTestUpdateResponseResultData]
type accountDeviceDexTestUpdateResponseResultDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestUpdateResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceDexTestUpdateResponseSuccess bool

const (
	AccountDeviceDexTestUpdateResponseSuccessTrue AccountDeviceDexTestUpdateResponseSuccess = true
)

type AccountDeviceDexTestDeleteResponse struct {
	Errors   []AccountDeviceDexTestDeleteResponseError   `json:"errors"`
	Messages []AccountDeviceDexTestDeleteResponseMessage `json:"messages"`
	Result   []AccountDeviceDexTestDeleteResponseResult  `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceDexTestDeleteResponseSuccess `json:"success"`
	JSON    accountDeviceDexTestDeleteResponseJSON    `json:"-"`
}

// accountDeviceDexTestDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDeviceDexTestDeleteResponse]
type accountDeviceDexTestDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDeviceDexTestDeleteResponseErrorJSON `json:"-"`
}

// accountDeviceDexTestDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestDeleteResponseError]
type accountDeviceDexTestDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountDeviceDexTestDeleteResponseMessageJSON `json:"-"`
}

// accountDeviceDexTestDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestDeleteResponseMessage]
type accountDeviceDexTestDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeleteResponseResult struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data AccountDeviceDexTestDeleteResponseResultData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                       `json:"description"`
	JSON        accountDeviceDexTestDeleteResponseResultJSON `json:"-"`
}

// accountDeviceDexTestDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestDeleteResponseResult]
type accountDeviceDexTestDeleteResponseResultJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type AccountDeviceDexTestDeleteResponseResultData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                           `json:"method"`
	JSON   accountDeviceDexTestDeleteResponseResultDataJSON `json:"-"`
}

// accountDeviceDexTestDeleteResponseResultDataJSON contains the JSON metadata for
// the struct [AccountDeviceDexTestDeleteResponseResultData]
type accountDeviceDexTestDeleteResponseResultDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeleteResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceDexTestDeleteResponseSuccess bool

const (
	AccountDeviceDexTestDeleteResponseSuccessTrue AccountDeviceDexTestDeleteResponseSuccess = true
)

type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponse struct {
	Errors   []AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseError   `json:"errors"`
	Messages []AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseMessage `json:"messages"`
	Result   AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseSuccess `json:"success"`
	JSON    accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseJSON    `json:"-"`
}

// accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseJSON contains the JSON
// metadata for the struct
// [AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponse]
type accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseErrorJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseError]
type accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseMessageJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseMessage]
type accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResult struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                                              `json:"description"`
	JSON        accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultJSON contains the
// JSON metadata for the struct
// [AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResult]
type accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                                                  `json:"method"`
	JSON   accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultDataJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultDataJSON contains
// the JSON metadata for the struct
// [AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultData]
type accountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseSuccess bool

const (
	AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseSuccessTrue AccountDeviceDexTestDeviceDexTestNewDeviceDexTestResponseSuccess = true
)

type AccountDeviceDexTestDeviceDexTestDetailsResponse struct {
	Errors   []AccountDeviceDexTestDeviceDexTestDetailsResponseError   `json:"errors"`
	Messages []AccountDeviceDexTestDeviceDexTestDetailsResponseMessage `json:"messages"`
	Result   []AccountDeviceDexTestDeviceDexTestDetailsResponseResult  `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceDexTestDeviceDexTestDetailsResponseSuccess `json:"success"`
	JSON    accountDeviceDexTestDeviceDexTestDetailsResponseJSON    `json:"-"`
}

// accountDeviceDexTestDeviceDexTestDetailsResponseJSON contains the JSON metadata
// for the struct [AccountDeviceDexTestDeviceDexTestDetailsResponse]
type accountDeviceDexTestDeviceDexTestDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeviceDexTestDetailsResponseError struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountDeviceDexTestDeviceDexTestDetailsResponseErrorJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestDetailsResponseErrorJSON contains the JSON
// metadata for the struct [AccountDeviceDexTestDeviceDexTestDetailsResponseError]
type accountDeviceDexTestDeviceDexTestDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeviceDexTestDetailsResponseMessage struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountDeviceDexTestDeviceDexTestDetailsResponseMessageJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestDetailsResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountDeviceDexTestDeviceDexTestDetailsResponseMessage]
type accountDeviceDexTestDeviceDexTestDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDexTestDeviceDexTestDetailsResponseResult struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data AccountDeviceDexTestDeviceDexTestDetailsResponseResultData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                                                     `json:"description"`
	JSON        accountDeviceDexTestDeviceDexTestDetailsResponseResultJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestDetailsResponseResultJSON contains the JSON
// metadata for the struct [AccountDeviceDexTestDeviceDexTestDetailsResponseResult]
type accountDeviceDexTestDeviceDexTestDetailsResponseResultJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type AccountDeviceDexTestDeviceDexTestDetailsResponseResultData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                                         `json:"method"`
	JSON   accountDeviceDexTestDeviceDexTestDetailsResponseResultDataJSON `json:"-"`
}

// accountDeviceDexTestDeviceDexTestDetailsResponseResultDataJSON contains the JSON
// metadata for the struct
// [AccountDeviceDexTestDeviceDexTestDetailsResponseResultData]
type accountDeviceDexTestDeviceDexTestDetailsResponseResultDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeviceDexTestDetailsResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceDexTestDeviceDexTestDetailsResponseSuccess bool

const (
	AccountDeviceDexTestDeviceDexTestDetailsResponseSuccessTrue AccountDeviceDexTestDeviceDexTestDetailsResponseSuccess = true
)

type AccountDeviceDexTestUpdateParams struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[AccountDeviceDexTestUpdateParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r AccountDeviceDexTestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type AccountDeviceDexTestUpdateParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r AccountDeviceDexTestUpdateParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestParams struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[AccountDeviceDexTestDeviceDexTestNewDeviceDexTestParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
}

func (r AccountDeviceDexTestDeviceDexTestNewDeviceDexTestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type AccountDeviceDexTestDeviceDexTestNewDeviceDexTestParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r AccountDeviceDexTestDeviceDexTestNewDeviceDexTestParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
