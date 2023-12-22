// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRailgunConnectionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountRailgunConnectionService] method instead.
type AccountRailgunConnectionService struct {
	Options []option.RequestOption
}

// NewAccountRailgunConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountRailgunConnectionService(opts ...option.RequestOption) (r *AccountRailgunConnectionService) {
	r = &AccountRailgunConnectionService{}
	r.Options = opts
	return
}

// Get a connection by ID.
func (r *AccountRailgunConnectionService) Get(ctx context.Context, accountIdentifier string, railgunIdentifier string, identifier string, opts ...option.RequestOption) (res *ConnectionSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections/%s", accountIdentifier, railgunIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable a connection.
func (r *AccountRailgunConnectionService) Update(ctx context.Context, accountIdentifier string, railgunIdentifier string, identifier string, body AccountRailgunConnectionUpdateParams, opts ...option.RequestOption) (res *ConnectionSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections/%s", accountIdentifier, railgunIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disable and remove the connection to a zone.
func (r *AccountRailgunConnectionService) Delete(ctx context.Context, accountIdentifier string, railgunIdentifier string, identifier string, opts ...option.RequestOption) (res *ConnectionSingleIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections/%s", accountIdentifier, railgunIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Associates a zone to the Railgun.
func (r *AccountRailgunConnectionService) RailgunConnectionsNewConnection(ctx context.Context, accountIdentifier string, railgunIdentifier string, body AccountRailgunConnectionRailgunConnectionsNewConnectionParams, opts ...option.RequestOption) (res *ConnectionSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections", accountIdentifier, railgunIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List connections associated with the Railgun.
func (r *AccountRailgunConnectionService) RailgunConnectionsListConnections(ctx context.Context, accountIdentifier string, railgunIdentifier string, query AccountRailgunConnectionRailgunConnectionsListConnectionsParams, opts ...option.RequestOption) (res *ConnectionCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections", accountIdentifier, railgunIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ConnectionCollectionResponse struct {
	Errors     []ConnectionCollectionResponseError    `json:"errors"`
	Messages   []ConnectionCollectionResponseMessage  `json:"messages"`
	Result     []ConnectionCollectionResponseResult   `json:"result"`
	ResultInfo ConnectionCollectionResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ConnectionCollectionResponseSuccess `json:"success"`
	JSON    connectionCollectionResponseJSON    `json:"-"`
}

// connectionCollectionResponseJSON contains the JSON metadata for the struct
// [ConnectionCollectionResponse]
type connectionCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionCollectionResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    connectionCollectionResponseErrorJSON `json:"-"`
}

// connectionCollectionResponseErrorJSON contains the JSON metadata for the struct
// [ConnectionCollectionResponseError]
type connectionCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionCollectionResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    connectionCollectionResponseMessageJSON `json:"-"`
}

// connectionCollectionResponseMessageJSON contains the JSON metadata for the
// struct [ConnectionCollectionResponseMessage]
type connectionCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionCollectionResponseResult struct {
	// Connection identifier tag.
	ID string `json:"id,required"`
	// A value indicating whether the connection is enabled or not.
	Enabled bool                                   `json:"enabled,required"`
	Zone    ConnectionCollectionResponseResultZone `json:"zone,required"`
	// When the connection was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// When the connection was last modified.
	ModifiedOn time.Time                              `json:"modified_on" format:"date-time"`
	JSON       connectionCollectionResponseResultJSON `json:"-"`
}

// connectionCollectionResponseResultJSON contains the JSON metadata for the struct
// [ConnectionCollectionResponseResult]
type connectionCollectionResponseResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Zone        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionCollectionResponseResultZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                                     `json:"name"`
	JSON connectionCollectionResponseResultZoneJSON `json:"-"`
}

// connectionCollectionResponseResultZoneJSON contains the JSON metadata for the
// struct [ConnectionCollectionResponseResultZone]
type connectionCollectionResponseResultZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionCollectionResponseResultZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionCollectionResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       connectionCollectionResponseResultInfoJSON `json:"-"`
}

// connectionCollectionResponseResultInfoJSON contains the JSON metadata for the
// struct [ConnectionCollectionResponseResultInfo]
type connectionCollectionResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionCollectionResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ConnectionCollectionResponseSuccess bool

const (
	ConnectionCollectionResponseSuccessTrue ConnectionCollectionResponseSuccess = true
)

type ConnectionSingleIDResponse struct {
	Errors   []ConnectionSingleIDResponseError   `json:"errors"`
	Messages []ConnectionSingleIDResponseMessage `json:"messages"`
	Result   ConnectionSingleIDResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ConnectionSingleIDResponseSuccess `json:"success"`
	JSON    connectionSingleIDResponseJSON    `json:"-"`
}

// connectionSingleIDResponseJSON contains the JSON metadata for the struct
// [ConnectionSingleIDResponse]
type connectionSingleIDResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionSingleIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionSingleIDResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    connectionSingleIDResponseErrorJSON `json:"-"`
}

// connectionSingleIDResponseErrorJSON contains the JSON metadata for the struct
// [ConnectionSingleIDResponseError]
type connectionSingleIDResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionSingleIDResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionSingleIDResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    connectionSingleIDResponseMessageJSON `json:"-"`
}

// connectionSingleIDResponseMessageJSON contains the JSON metadata for the struct
// [ConnectionSingleIDResponseMessage]
type connectionSingleIDResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionSingleIDResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionSingleIDResponseResult struct {
	// Connection identifier tag.
	ID   string                               `json:"id"`
	JSON connectionSingleIDResponseResultJSON `json:"-"`
}

// connectionSingleIDResponseResultJSON contains the JSON metadata for the struct
// [ConnectionSingleIDResponseResult]
type connectionSingleIDResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionSingleIDResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ConnectionSingleIDResponseSuccess bool

const (
	ConnectionSingleIDResponseSuccessTrue ConnectionSingleIDResponseSuccess = true
)

type ConnectionSingleResponse struct {
	Errors   []ConnectionSingleResponseError   `json:"errors"`
	Messages []ConnectionSingleResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success ConnectionSingleResponseSuccess `json:"success"`
	JSON    connectionSingleResponseJSON    `json:"-"`
}

// connectionSingleResponseJSON contains the JSON metadata for the struct
// [ConnectionSingleResponse]
type connectionSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionSingleResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    connectionSingleResponseErrorJSON `json:"-"`
}

// connectionSingleResponseErrorJSON contains the JSON metadata for the struct
// [ConnectionSingleResponseError]
type connectionSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionSingleResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    connectionSingleResponseMessageJSON `json:"-"`
}

// connectionSingleResponseMessageJSON contains the JSON metadata for the struct
// [ConnectionSingleResponseMessage]
type connectionSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ConnectionSingleResponseSuccess bool

const (
	ConnectionSingleResponseSuccessTrue ConnectionSingleResponseSuccess = true
)

type AccountRailgunConnectionUpdateParams struct {
	// A value indicating whether the connection is enabled or not.
	Enabled param.Field[bool]                                     `json:"enabled,required"`
	Zone    param.Field[AccountRailgunConnectionUpdateParamsZone] `json:"zone,required"`
}

func (r AccountRailgunConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRailgunConnectionUpdateParamsZone struct {
}

func (r AccountRailgunConnectionUpdateParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRailgunConnectionRailgunConnectionsNewConnectionParams struct {
	Zone param.Field[AccountRailgunConnectionRailgunConnectionsNewConnectionParamsZone] `json:"zone,required"`
	// A value indicating whether the connection is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountRailgunConnectionRailgunConnectionsNewConnectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRailgunConnectionRailgunConnectionsNewConnectionParamsZone struct {
}

func (r AccountRailgunConnectionRailgunConnectionsNewConnectionParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRailgunConnectionRailgunConnectionsListConnectionsParams struct {
	// A value indicating whether the connection is enabled or not.
	Enabled param.Field[bool] `query:"enabled"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [AccountRailgunConnectionRailgunConnectionsListConnectionsParams]'s query
// parameters as `url.Values`.
func (r AccountRailgunConnectionRailgunConnectionsListConnectionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
