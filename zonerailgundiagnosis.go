// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRailgunDiagnosisService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRailgunDiagnosisService]
// method instead.
type ZoneRailgunDiagnosisService struct {
	Options []option.RequestOption
}

// NewZoneRailgunDiagnosisService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneRailgunDiagnosisService(opts ...option.RequestOption) (r *ZoneRailgunDiagnosisService) {
	r = &ZoneRailgunDiagnosisService{}
	r.Options = opts
	return
}

// Tests the Railgun connection to the zone.
func (r *ZoneRailgunDiagnosisService) RailgunConnectionsForAZoneTestRailgunConnection(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *TestConnectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/railguns/%s/diagnose", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TestConnectionResponse struct {
	Errors   []TestConnectionResponseError   `json:"errors"`
	Messages []TestConnectionResponseMessage `json:"messages"`
	Result   TestConnectionResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TestConnectionResponseSuccess `json:"success"`
	JSON    testConnectionResponseJSON    `json:"-"`
}

// testConnectionResponseJSON contains the JSON metadata for the struct
// [TestConnectionResponse]
type testConnectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TestConnectionResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    testConnectionResponseErrorJSON `json:"-"`
}

// testConnectionResponseErrorJSON contains the JSON metadata for the struct
// [TestConnectionResponseError]
type testConnectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestConnectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TestConnectionResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    testConnectionResponseMessageJSON `json:"-"`
}

// testConnectionResponseMessageJSON contains the JSON metadata for the struct
// [TestConnectionResponseMessage]
type testConnectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestConnectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TestConnectionResponseResult struct {
	// Hash version of body.
	BodyHash string `json:"body_hash"`
	// Size of the body in bytes.
	BodySize string `json:"body_size"`
	// Lists any `cf-cache-status` present.
	CfCacheStatus string `json:"cf-cache-status"`
	// Lists any `cf-ray` present.
	CfRay string `json:"cf-ray"`
	// Lists any `cf-wan-error` present.
	CfWanError string `json:"cf-wan-error"`
	// Whether Cloudflare is enabled on the host.
	Cloudflare string `json:"cloudflare"`
	// Connection closed or open.
	ConnectionClose bool `json:"connection_close"`
	// Amount of seconds that the test lasted.
	ElapsedTime string `json:"elapsed_time"`
	// The hostname queried.
	HostName string `json:"host_name"`
	// The HTTP status response code.
	HTTPStatus float64 `json:"http_status"`
	// HTTP Method used to test the connection.
	Method TestConnectionResponseResultMethod `json:"method"`
	// What headers are missing.
	MissingHeaders string `json:"missing_headers"`
	// Protocol used to test the connection.
	Protocol string `json:"protocol"`
	// Indicates if Railgun is enabled on the queried hostname.
	Railgun string `json:"railgun"`
	// HTTP Status code.
	ResponseStatus string `json:"response_status"`
	// Url of the domain you can compare the connection to.
	URL  string                           `json:"url"`
	JSON testConnectionResponseResultJSON `json:"-"`
}

// testConnectionResponseResultJSON contains the JSON metadata for the struct
// [TestConnectionResponseResult]
type testConnectionResponseResultJSON struct {
	BodyHash        apijson.Field
	BodySize        apijson.Field
	CfCacheStatus   apijson.Field
	CfRay           apijson.Field
	CfWanError      apijson.Field
	Cloudflare      apijson.Field
	ConnectionClose apijson.Field
	ElapsedTime     apijson.Field
	HostName        apijson.Field
	HTTPStatus      apijson.Field
	Method          apijson.Field
	MissingHeaders  apijson.Field
	Protocol        apijson.Field
	Railgun         apijson.Field
	ResponseStatus  apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TestConnectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP Method used to test the connection.
type TestConnectionResponseResultMethod string

const (
	TestConnectionResponseResultMethodGet  TestConnectionResponseResultMethod = "GET"
	TestConnectionResponseResultMethodPost TestConnectionResponseResultMethod = "POST"
)

// Whether the API call was successful
type TestConnectionResponseSuccess bool

const (
	TestConnectionResponseSuccessTrue TestConnectionResponseSuccess = true
)
