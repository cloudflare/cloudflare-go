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
func (r *AccountDeviceDexTestService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DexSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a DEX test.
func (r *AccountDeviceDexTestService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDeviceDexTestUpdateParams, opts ...option.RequestOption) (res *DexSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *AccountDeviceDexTestService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DexResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a DEX test.
func (r *AccountDeviceDexTestService) DeviceDexTestNewDeviceDexTest(ctx context.Context, identifier interface{}, body AccountDeviceDexTestDeviceDexTestNewDeviceDexTestParams, opts ...option.RequestOption) (res *DexSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch all DEX tests.
func (r *AccountDeviceDexTestService) DeviceDexTestDetails(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *DexResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DexResponseCollection struct {
	Errors   []DexResponseCollectionError   `json:"errors"`
	Messages []DexResponseCollectionMessage `json:"messages"`
	Result   []DexResponseCollectionResult  `json:"result"`
	// Whether the API call was successful
	Success DexResponseCollectionSuccess `json:"success"`
	JSON    dexResponseCollectionJSON    `json:"-"`
}

// dexResponseCollectionJSON contains the JSON metadata for the struct
// [DexResponseCollection]
type dexResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexResponseCollectionError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    dexResponseCollectionErrorJSON `json:"-"`
}

// dexResponseCollectionErrorJSON contains the JSON metadata for the struct
// [DexResponseCollectionError]
type dexResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexResponseCollectionMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dexResponseCollectionMessageJSON `json:"-"`
}

// dexResponseCollectionMessageJSON contains the JSON metadata for the struct
// [DexResponseCollectionMessage]
type dexResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexResponseCollectionResult struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DexResponseCollectionResultData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                          `json:"description"`
	JSON        dexResponseCollectionResultJSON `json:"-"`
}

// dexResponseCollectionResultJSON contains the JSON metadata for the struct
// [DexResponseCollectionResult]
type dexResponseCollectionResultJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DexResponseCollectionResultData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                              `json:"method"`
	JSON   dexResponseCollectionResultDataJSON `json:"-"`
}

// dexResponseCollectionResultDataJSON contains the JSON metadata for the struct
// [DexResponseCollectionResultData]
type dexResponseCollectionResultDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexResponseCollectionResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexResponseCollectionSuccess bool

const (
	DexResponseCollectionSuccessTrue DexResponseCollectionSuccess = true
)

type DexSingleResponse struct {
	Errors   []DexSingleResponseError   `json:"errors"`
	Messages []DexSingleResponseMessage `json:"messages"`
	Result   DexSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexSingleResponseSuccess `json:"success"`
	JSON    dexSingleResponseJSON    `json:"-"`
}

// dexSingleResponseJSON contains the JSON metadata for the struct
// [DexSingleResponse]
type dexSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexSingleResponseError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    dexSingleResponseErrorJSON `json:"-"`
}

// dexSingleResponseErrorJSON contains the JSON metadata for the struct
// [DexSingleResponseError]
type dexSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexSingleResponseMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    dexSingleResponseMessageJSON `json:"-"`
}

// dexSingleResponseMessageJSON contains the JSON metadata for the struct
// [DexSingleResponseMessage]
type dexSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexSingleResponseResult struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DexSingleResponseResultData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string                      `json:"description"`
	JSON        dexSingleResponseResultJSON `json:"-"`
}

// dexSingleResponseResultJSON contains the JSON metadata for the struct
// [DexSingleResponseResult]
type dexSingleResponseResultJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DexSingleResponseResultData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                          `json:"method"`
	JSON   dexSingleResponseResultDataJSON `json:"-"`
}

// dexSingleResponseResultDataJSON contains the JSON metadata for the struct
// [DexSingleResponseResultData]
type dexSingleResponseResultDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexSingleResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexSingleResponseSuccess bool

const (
	DexSingleResponseSuccessTrue DexSingleResponseSuccess = true
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
