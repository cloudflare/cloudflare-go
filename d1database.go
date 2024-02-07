// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// D1DatabaseService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewD1DatabaseService] method instead.
type D1DatabaseService struct {
	Options []option.RequestOption
}

// NewD1DatabaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewD1DatabaseService(opts ...option.RequestOption) (r *D1DatabaseService) {
	r = &D1DatabaseService{}
	r.Options = opts
	return
}

// Returns the created D1 database.
func (r *D1DatabaseService) New(ctx context.Context, accountID string, body D1DatabaseNewParams, opts ...option.RequestOption) (res *D1DatabaseNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/d1/database", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of D1 databases.
func (r *D1DatabaseService) List(ctx context.Context, accountID string, query D1DatabaseListParams, opts ...option.RequestOption) (res *D1DatabaseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/d1/database", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type D1DatabaseNewResponse struct {
	Errors   []D1DatabaseNewResponseError   `json:"errors"`
	Messages []D1DatabaseNewResponseMessage `json:"messages"`
	Result   D1DatabaseNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success D1DatabaseNewResponseSuccess `json:"success"`
	JSON    d1DatabaseNewResponseJSON    `json:"-"`
}

// d1DatabaseNewResponseJSON contains the JSON metadata for the struct
// [D1DatabaseNewResponse]
type d1DatabaseNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseNewResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    d1DatabaseNewResponseErrorJSON `json:"-"`
}

// d1DatabaseNewResponseErrorJSON contains the JSON metadata for the struct
// [D1DatabaseNewResponseError]
type d1DatabaseNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseNewResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    d1DatabaseNewResponseMessageJSON `json:"-"`
}

// d1DatabaseNewResponseMessageJSON contains the JSON metadata for the struct
// [D1DatabaseNewResponseMessage]
type d1DatabaseNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseNewResponseResult struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{}                     `json:"created_at"`
	Name      string                          `json:"name"`
	Uuid      string                          `json:"uuid"`
	Version   string                          `json:"version"`
	JSON      d1DatabaseNewResponseResultJSON `json:"-"`
}

// d1DatabaseNewResponseResultJSON contains the JSON metadata for the struct
// [D1DatabaseNewResponseResult]
type d1DatabaseNewResponseResultJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type D1DatabaseNewResponseSuccess bool

const (
	D1DatabaseNewResponseSuccessTrue D1DatabaseNewResponseSuccess = true
)

type D1DatabaseListResponse struct {
	Errors   []D1DatabaseListResponseError   `json:"errors"`
	Messages []D1DatabaseListResponseMessage `json:"messages"`
	Result   []D1DatabaseListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success D1DatabaseListResponseSuccess `json:"success"`
	JSON    d1DatabaseListResponseJSON    `json:"-"`
}

// d1DatabaseListResponseJSON contains the JSON metadata for the struct
// [D1DatabaseListResponse]
type d1DatabaseListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseListResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    d1DatabaseListResponseErrorJSON `json:"-"`
}

// d1DatabaseListResponseErrorJSON contains the JSON metadata for the struct
// [D1DatabaseListResponseError]
type d1DatabaseListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseListResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    d1DatabaseListResponseMessageJSON `json:"-"`
}

// d1DatabaseListResponseMessageJSON contains the JSON metadata for the struct
// [D1DatabaseListResponseMessage]
type d1DatabaseListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseListResponseResult struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{}                      `json:"created_at"`
	Name      string                           `json:"name"`
	Uuid      string                           `json:"uuid"`
	Version   string                           `json:"version"`
	JSON      d1DatabaseListResponseResultJSON `json:"-"`
}

// d1DatabaseListResponseResultJSON contains the JSON metadata for the struct
// [D1DatabaseListResponseResult]
type d1DatabaseListResponseResultJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type D1DatabaseListResponseSuccess bool

const (
	D1DatabaseListResponseSuccessTrue D1DatabaseListResponseSuccess = true
)

type D1DatabaseNewParams struct {
	Name param.Field[string] `json:"name,required"`
}

func (r D1DatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type D1DatabaseListParams struct {
	// a database name to search for.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [D1DatabaseListParams]'s query parameters as `url.Values`.
func (r D1DatabaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
