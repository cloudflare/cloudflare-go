// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAccessServiceTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAccessServiceTokenService]
// method instead.
type ZoneAccessServiceTokenService struct {
	Options []option.RequestOption
}

// NewZoneAccessServiceTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAccessServiceTokenService(opts ...option.RequestOption) (r *ZoneAccessServiceTokenService) {
	r = &ZoneAccessServiceTokenService{}
	r.Options = opts
	return
}

// Updates a configured service token.
func (r *ZoneAccessServiceTokenService) Update(ctx context.Context, identifier string, uuid string, body ZoneAccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *ServiceTokensSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a service token.
func (r *ZoneAccessServiceTokenService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ServiceTokensSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to create a new
// service token.
func (r *ZoneAccessServiceTokenService) ZoneLevelAccessServiceTokensNewAServiceToken(ctx context.Context, identifier string, body ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenParams, opts ...option.RequestOption) (res *CreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all service tokens.
func (r *ZoneAccessServiceTokenService) ZoneLevelAccessServiceTokensListServiceTokens(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/service_tokens", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse struct {
	Errors     []ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError    `json:"errors"`
	Messages   []ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage  `json:"messages"`
	Result     []ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult   `json:"result"`
	ResultInfo ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccess `json:"success"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseJSON    `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseErrorJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessageJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The name of the service token.
	Name      string                                                                                `json:"name"`
	UpdatedAt time.Time                                                                             `json:"updated_at" format:"date-time"`
	JSON      zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                   `json:"total_count"`
	JSON       zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfoJSON `json:"-"`
}

// zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo]
type zoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccess bool

const (
	ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccessTrue ZoneAccessServiceTokenZoneLevelAccessServiceTokensListServiceTokensResponseSuccess = true
)

type ZoneAccessServiceTokenUpdateParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r ZoneAccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessServiceTokenZoneLevelAccessServiceTokensNewAServiceTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
