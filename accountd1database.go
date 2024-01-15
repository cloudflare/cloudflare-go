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

// AccountD1DatabaseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountD1DatabaseService] method
// instead.
type AccountD1DatabaseService struct {
	Options []option.RequestOption
}

// NewAccountD1DatabaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountD1DatabaseService(opts ...option.RequestOption) (r *AccountD1DatabaseService) {
	r = &AccountD1DatabaseService{}
	r.Options = opts
	return
}

// Returns the created D1 database.
func (r *AccountD1DatabaseService) New(ctx context.Context, accountIdentifier string, body AccountD1DatabaseNewParams, opts ...option.RequestOption) (res *AccountD1DatabaseNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/d1/database", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the specified D1 database.
func (r *AccountD1DatabaseService) Get(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *AccountD1DatabaseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of D1 databases.
func (r *AccountD1DatabaseService) List(ctx context.Context, accountIdentifier string, query AccountD1DatabaseListParams, opts ...option.RequestOption) (res *AccountD1DatabaseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/d1/database", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the specified D1 database.
func (r *AccountD1DatabaseService) Delete(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *AccountD1DatabaseDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns the query result.
func (r *AccountD1DatabaseService) Query(ctx context.Context, accountIdentifier string, databaseIdentifier string, body AccountD1DatabaseQueryParams, opts ...option.RequestOption) (res *AccountD1DatabaseQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/d1/database/%s/query", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountD1DatabaseNewResponse struct {
	Errors   []AccountD1DatabaseNewResponseError   `json:"errors"`
	Messages []AccountD1DatabaseNewResponseMessage `json:"messages"`
	Result   AccountD1DatabaseNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountD1DatabaseNewResponseSuccess `json:"success"`
	JSON    accountD1DatabaseNewResponseJSON    `json:"-"`
}

// accountD1DatabaseNewResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseNewResponse]
type accountD1DatabaseNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseNewResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountD1DatabaseNewResponseErrorJSON `json:"-"`
}

// accountD1DatabaseNewResponseErrorJSON contains the JSON metadata for the struct
// [AccountD1DatabaseNewResponseError]
type accountD1DatabaseNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseNewResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountD1DatabaseNewResponseMessageJSON `json:"-"`
}

// accountD1DatabaseNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountD1DatabaseNewResponseMessage]
type accountD1DatabaseNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseNewResponseResult struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{}                            `json:"created_at"`
	Name      string                                 `json:"name"`
	Uuid      string                                 `json:"uuid"`
	Version   string                                 `json:"version"`
	JSON      accountD1DatabaseNewResponseResultJSON `json:"-"`
}

// accountD1DatabaseNewResponseResultJSON contains the JSON metadata for the struct
// [AccountD1DatabaseNewResponseResult]
type accountD1DatabaseNewResponseResultJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountD1DatabaseNewResponseSuccess bool

const (
	AccountD1DatabaseNewResponseSuccessTrue AccountD1DatabaseNewResponseSuccess = true
)

type AccountD1DatabaseGetResponse struct {
	Errors   []AccountD1DatabaseGetResponseError   `json:"errors"`
	Messages []AccountD1DatabaseGetResponseMessage `json:"messages"`
	Result   AccountD1DatabaseGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountD1DatabaseGetResponseSuccess `json:"success"`
	JSON    accountD1DatabaseGetResponseJSON    `json:"-"`
}

// accountD1DatabaseGetResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseGetResponse]
type accountD1DatabaseGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountD1DatabaseGetResponseErrorJSON `json:"-"`
}

// accountD1DatabaseGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountD1DatabaseGetResponseError]
type accountD1DatabaseGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountD1DatabaseGetResponseMessageJSON `json:"-"`
}

// accountD1DatabaseGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountD1DatabaseGetResponseMessage]
type accountD1DatabaseGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseGetResponseResult struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{} `json:"created_at"`
	// The D1 database's size, in bytes.
	FileSize  float64                                `json:"file_size"`
	Name      string                                 `json:"name"`
	NumTables float64                                `json:"num_tables"`
	Uuid      string                                 `json:"uuid"`
	Version   string                                 `json:"version"`
	JSON      accountD1DatabaseGetResponseResultJSON `json:"-"`
}

// accountD1DatabaseGetResponseResultJSON contains the JSON metadata for the struct
// [AccountD1DatabaseGetResponseResult]
type accountD1DatabaseGetResponseResultJSON struct {
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Name        apijson.Field
	NumTables   apijson.Field
	Uuid        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountD1DatabaseGetResponseSuccess bool

const (
	AccountD1DatabaseGetResponseSuccessTrue AccountD1DatabaseGetResponseSuccess = true
)

type AccountD1DatabaseListResponse struct {
	Errors   []AccountD1DatabaseListResponseError   `json:"errors"`
	Messages []AccountD1DatabaseListResponseMessage `json:"messages"`
	Result   []AccountD1DatabaseListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountD1DatabaseListResponseSuccess `json:"success"`
	JSON    accountD1DatabaseListResponseJSON    `json:"-"`
}

// accountD1DatabaseListResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseListResponse]
type accountD1DatabaseListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountD1DatabaseListResponseErrorJSON `json:"-"`
}

// accountD1DatabaseListResponseErrorJSON contains the JSON metadata for the struct
// [AccountD1DatabaseListResponseError]
type accountD1DatabaseListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountD1DatabaseListResponseMessageJSON `json:"-"`
}

// accountD1DatabaseListResponseMessageJSON contains the JSON metadata for the
// struct [AccountD1DatabaseListResponseMessage]
type accountD1DatabaseListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseListResponseResult struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{}                             `json:"created_at"`
	Name      string                                  `json:"name"`
	Uuid      string                                  `json:"uuid"`
	Version   string                                  `json:"version"`
	JSON      accountD1DatabaseListResponseResultJSON `json:"-"`
}

// accountD1DatabaseListResponseResultJSON contains the JSON metadata for the
// struct [AccountD1DatabaseListResponseResult]
type accountD1DatabaseListResponseResultJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountD1DatabaseListResponseSuccess bool

const (
	AccountD1DatabaseListResponseSuccessTrue AccountD1DatabaseListResponseSuccess = true
)

type AccountD1DatabaseDeleteResponse struct {
	Errors   []AccountD1DatabaseDeleteResponseError   `json:"errors"`
	Messages []AccountD1DatabaseDeleteResponseMessage `json:"messages"`
	Result   interface{}                              `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountD1DatabaseDeleteResponseSuccess `json:"success"`
	JSON    accountD1DatabaseDeleteResponseJSON    `json:"-"`
}

// accountD1DatabaseDeleteResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseDeleteResponse]
type accountD1DatabaseDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseDeleteResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountD1DatabaseDeleteResponseErrorJSON `json:"-"`
}

// accountD1DatabaseDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountD1DatabaseDeleteResponseError]
type accountD1DatabaseDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseDeleteResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountD1DatabaseDeleteResponseMessageJSON `json:"-"`
}

// accountD1DatabaseDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountD1DatabaseDeleteResponseMessage]
type accountD1DatabaseDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountD1DatabaseDeleteResponseSuccess bool

const (
	AccountD1DatabaseDeleteResponseSuccessTrue AccountD1DatabaseDeleteResponseSuccess = true
)

type AccountD1DatabaseQueryResponse struct {
	Errors   []AccountD1DatabaseQueryResponseError   `json:"errors"`
	Messages []AccountD1DatabaseQueryResponseMessage `json:"messages"`
	Result   []AccountD1DatabaseQueryResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountD1DatabaseQueryResponseSuccess `json:"success"`
	JSON    accountD1DatabaseQueryResponseJSON    `json:"-"`
}

// accountD1DatabaseQueryResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseQueryResponse]
type accountD1DatabaseQueryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseQueryResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountD1DatabaseQueryResponseErrorJSON `json:"-"`
}

// accountD1DatabaseQueryResponseErrorJSON contains the JSON metadata for the
// struct [AccountD1DatabaseQueryResponseError]
type accountD1DatabaseQueryResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseQueryResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseQueryResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountD1DatabaseQueryResponseMessageJSON `json:"-"`
}

// accountD1DatabaseQueryResponseMessageJSON contains the JSON metadata for the
// struct [AccountD1DatabaseQueryResponseMessage]
type accountD1DatabaseQueryResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseQueryResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseQueryResponseResult struct {
	Meta    AccountD1DatabaseQueryResponseResultMeta `json:"meta"`
	Results []interface{}                            `json:"results"`
	Success bool                                     `json:"success"`
	JSON    accountD1DatabaseQueryResponseResultJSON `json:"-"`
}

// accountD1DatabaseQueryResponseResultJSON contains the JSON metadata for the
// struct [AccountD1DatabaseQueryResponseResult]
type accountD1DatabaseQueryResponseResultJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountD1DatabaseQueryResponseResultMeta struct {
	ChangedDB   bool                                         `json:"changed_db"`
	Changes     float64                                      `json:"changes"`
	Duration    float64                                      `json:"duration"`
	LastRowID   float64                                      `json:"last_row_id"`
	RowsRead    float64                                      `json:"rows_read"`
	RowsWritten float64                                      `json:"rows_written"`
	SizeAfter   float64                                      `json:"size_after"`
	JSON        accountD1DatabaseQueryResponseResultMetaJSON `json:"-"`
}

// accountD1DatabaseQueryResponseResultMetaJSON contains the JSON metadata for the
// struct [AccountD1DatabaseQueryResponseResultMeta]
type accountD1DatabaseQueryResponseResultMetaJSON struct {
	ChangedDB   apijson.Field
	Changes     apijson.Field
	Duration    apijson.Field
	LastRowID   apijson.Field
	RowsRead    apijson.Field
	RowsWritten apijson.Field
	SizeAfter   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseQueryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountD1DatabaseQueryResponseSuccess bool

const (
	AccountD1DatabaseQueryResponseSuccessTrue AccountD1DatabaseQueryResponseSuccess = true
)

type AccountD1DatabaseNewParams struct {
	Name param.Field[string] `json:"name,required"`
}

func (r AccountD1DatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountD1DatabaseListParams struct {
	// a database name to search for.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountD1DatabaseListParams]'s query parameters as
// `url.Values`.
func (r AccountD1DatabaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountD1DatabaseQueryParams struct {
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r AccountD1DatabaseQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
