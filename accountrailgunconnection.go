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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *AccountRailgunConnectionService) Get(ctx context.Context, accountIdentifier string, railgunIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountRailgunConnectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections/%s", accountIdentifier, railgunIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable a connection.
func (r *AccountRailgunConnectionService) Update(ctx context.Context, accountIdentifier string, railgunIdentifier string, identifier string, body AccountRailgunConnectionUpdateParams, opts ...option.RequestOption) (res *AccountRailgunConnectionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections/%s", accountIdentifier, railgunIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disable and remove the connection to a zone.
func (r *AccountRailgunConnectionService) Delete(ctx context.Context, accountIdentifier string, railgunIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountRailgunConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections/%s", accountIdentifier, railgunIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Associates a zone to the Railgun.
func (r *AccountRailgunConnectionService) RailgunConnectionsNewConnection(ctx context.Context, accountIdentifier string, railgunIdentifier string, body AccountRailgunConnectionRailgunConnectionsNewConnectionParams, opts ...option.RequestOption) (res *AccountRailgunConnectionRailgunConnectionsNewConnectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections", accountIdentifier, railgunIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List connections associated with the Railgun.
func (r *AccountRailgunConnectionService) RailgunConnectionsListConnections(ctx context.Context, accountIdentifier string, railgunIdentifier string, query AccountRailgunConnectionRailgunConnectionsListConnectionsParams, opts ...option.RequestOption) (res *shared.Page[AccountRailgunConnectionRailgunConnectionsListConnectionsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s/connections", accountIdentifier, railgunIdentifier)
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

type AccountRailgunConnectionGetResponse struct {
	Errors   []AccountRailgunConnectionGetResponseError   `json:"errors"`
	Messages []AccountRailgunConnectionGetResponseMessage `json:"messages"`
	Result   interface{}                                  `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunConnectionGetResponseSuccess `json:"success"`
	JSON    accountRailgunConnectionGetResponseJSON    `json:"-"`
}

// accountRailgunConnectionGetResponseJSON contains the JSON metadata for the
// struct [AccountRailgunConnectionGetResponse]
type accountRailgunConnectionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionGetResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountRailgunConnectionGetResponseErrorJSON `json:"-"`
}

// accountRailgunConnectionGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountRailgunConnectionGetResponseError]
type accountRailgunConnectionGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionGetResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountRailgunConnectionGetResponseMessageJSON `json:"-"`
}

// accountRailgunConnectionGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountRailgunConnectionGetResponseMessage]
type accountRailgunConnectionGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunConnectionGetResponseSuccess bool

const (
	AccountRailgunConnectionGetResponseSuccessTrue AccountRailgunConnectionGetResponseSuccess = true
)

type AccountRailgunConnectionUpdateResponse struct {
	Errors   []AccountRailgunConnectionUpdateResponseError   `json:"errors"`
	Messages []AccountRailgunConnectionUpdateResponseMessage `json:"messages"`
	Result   interface{}                                     `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunConnectionUpdateResponseSuccess `json:"success"`
	JSON    accountRailgunConnectionUpdateResponseJSON    `json:"-"`
}

// accountRailgunConnectionUpdateResponseJSON contains the JSON metadata for the
// struct [AccountRailgunConnectionUpdateResponse]
type accountRailgunConnectionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountRailgunConnectionUpdateResponseErrorJSON `json:"-"`
}

// accountRailgunConnectionUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountRailgunConnectionUpdateResponseError]
type accountRailgunConnectionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountRailgunConnectionUpdateResponseMessageJSON `json:"-"`
}

// accountRailgunConnectionUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountRailgunConnectionUpdateResponseMessage]
type accountRailgunConnectionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunConnectionUpdateResponseSuccess bool

const (
	AccountRailgunConnectionUpdateResponseSuccessTrue AccountRailgunConnectionUpdateResponseSuccess = true
)

type AccountRailgunConnectionDeleteResponse struct {
	Errors   []AccountRailgunConnectionDeleteResponseError   `json:"errors"`
	Messages []AccountRailgunConnectionDeleteResponseMessage `json:"messages"`
	Result   AccountRailgunConnectionDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunConnectionDeleteResponseSuccess `json:"success"`
	JSON    accountRailgunConnectionDeleteResponseJSON    `json:"-"`
}

// accountRailgunConnectionDeleteResponseJSON contains the JSON metadata for the
// struct [AccountRailgunConnectionDeleteResponse]
type accountRailgunConnectionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionDeleteResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountRailgunConnectionDeleteResponseErrorJSON `json:"-"`
}

// accountRailgunConnectionDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountRailgunConnectionDeleteResponseError]
type accountRailgunConnectionDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionDeleteResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountRailgunConnectionDeleteResponseMessageJSON `json:"-"`
}

// accountRailgunConnectionDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountRailgunConnectionDeleteResponseMessage]
type accountRailgunConnectionDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionDeleteResponseResult struct {
	// Connection identifier tag.
	ID   string                                           `json:"id"`
	JSON accountRailgunConnectionDeleteResponseResultJSON `json:"-"`
}

// accountRailgunConnectionDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountRailgunConnectionDeleteResponseResult]
type accountRailgunConnectionDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunConnectionDeleteResponseSuccess bool

const (
	AccountRailgunConnectionDeleteResponseSuccessTrue AccountRailgunConnectionDeleteResponseSuccess = true
)

type AccountRailgunConnectionRailgunConnectionsNewConnectionResponse struct {
	Errors   []AccountRailgunConnectionRailgunConnectionsNewConnectionResponseError   `json:"errors"`
	Messages []AccountRailgunConnectionRailgunConnectionsNewConnectionResponseMessage `json:"messages"`
	Result   interface{}                                                              `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunConnectionRailgunConnectionsNewConnectionResponseSuccess `json:"success"`
	JSON    accountRailgunConnectionRailgunConnectionsNewConnectionResponseJSON    `json:"-"`
}

// accountRailgunConnectionRailgunConnectionsNewConnectionResponseJSON contains the
// JSON metadata for the struct
// [AccountRailgunConnectionRailgunConnectionsNewConnectionResponse]
type accountRailgunConnectionRailgunConnectionsNewConnectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionRailgunConnectionsNewConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionRailgunConnectionsNewConnectionResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accountRailgunConnectionRailgunConnectionsNewConnectionResponseErrorJSON `json:"-"`
}

// accountRailgunConnectionRailgunConnectionsNewConnectionResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountRailgunConnectionRailgunConnectionsNewConnectionResponseError]
type accountRailgunConnectionRailgunConnectionsNewConnectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionRailgunConnectionsNewConnectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionRailgunConnectionsNewConnectionResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountRailgunConnectionRailgunConnectionsNewConnectionResponseMessageJSON `json:"-"`
}

// accountRailgunConnectionRailgunConnectionsNewConnectionResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountRailgunConnectionRailgunConnectionsNewConnectionResponseMessage]
type accountRailgunConnectionRailgunConnectionsNewConnectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionRailgunConnectionsNewConnectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunConnectionRailgunConnectionsNewConnectionResponseSuccess bool

const (
	AccountRailgunConnectionRailgunConnectionsNewConnectionResponseSuccessTrue AccountRailgunConnectionRailgunConnectionsNewConnectionResponseSuccess = true
)

type AccountRailgunConnectionRailgunConnectionsListConnectionsResponse struct {
	// Connection identifier tag.
	ID string `json:"id,required"`
	// A value indicating whether the connection is enabled or not.
	Enabled bool                                                                  `json:"enabled,required"`
	Zone    AccountRailgunConnectionRailgunConnectionsListConnectionsResponseZone `json:"zone,required"`
	// When the connection was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// When the connection was last modified.
	ModifiedOn time.Time                                                             `json:"modified_on" format:"date-time"`
	JSON       accountRailgunConnectionRailgunConnectionsListConnectionsResponseJSON `json:"-"`
}

// accountRailgunConnectionRailgunConnectionsListConnectionsResponseJSON contains
// the JSON metadata for the struct
// [AccountRailgunConnectionRailgunConnectionsListConnectionsResponse]
type accountRailgunConnectionRailgunConnectionsListConnectionsResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Zone        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionRailgunConnectionsListConnectionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunConnectionRailgunConnectionsListConnectionsResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                                                                    `json:"name"`
	JSON accountRailgunConnectionRailgunConnectionsListConnectionsResponseZoneJSON `json:"-"`
}

// accountRailgunConnectionRailgunConnectionsListConnectionsResponseZoneJSON
// contains the JSON metadata for the struct
// [AccountRailgunConnectionRailgunConnectionsListConnectionsResponseZone]
type accountRailgunConnectionRailgunConnectionsListConnectionsResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunConnectionRailgunConnectionsListConnectionsResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
