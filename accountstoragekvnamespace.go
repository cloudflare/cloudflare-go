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

// AccountStorageKvNamespaceService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountStorageKvNamespaceService] method instead.
type AccountStorageKvNamespaceService struct {
	Options  []option.RequestOption
	Bulks    *AccountStorageKvNamespaceBulkService
	Keys     *AccountStorageKvNamespaceKeyService
	Metadata *AccountStorageKvNamespaceMetadataService
	Values   *AccountStorageKvNamespaceValueService
}

// NewAccountStorageKvNamespaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceService) {
	r = &AccountStorageKvNamespaceService{}
	r.Options = opts
	r.Bulks = NewAccountStorageKvNamespaceBulkService(opts...)
	r.Keys = NewAccountStorageKvNamespaceKeyService(opts...)
	r.Metadata = NewAccountStorageKvNamespaceMetadataService(opts...)
	r.Values = NewAccountStorageKvNamespaceValueService(opts...)
	return
}

// Modifies a namespace's title.
func (r *AccountStorageKvNamespaceService) Update(ctx context.Context, accountIdentifier string, namespaceIdentifier string, body AccountStorageKvNamespaceUpdateParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the namespaces owned by an account.
func (r *AccountStorageKvNamespaceService) List(ctx context.Context, accountIdentifier string, query AccountStorageKvNamespaceListParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the namespace corresponding to the given ID.
func (r *AccountStorageKvNamespaceService) Delete(ctx context.Context, accountIdentifier string, namespaceIdentifier string, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a namespace under the given title. A 400 is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *AccountStorageKvNamespaceService) WorkersKvNamespaceNewANamespace(ctx context.Context, accountIdentifier string, body AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountStorageKvNamespaceListResponse struct {
	Errors     []AccountStorageKvNamespaceListResponseError    `json:"errors"`
	Messages   []AccountStorageKvNamespaceListResponseMessage  `json:"messages"`
	Result     []AccountStorageKvNamespaceListResponseResult   `json:"result"`
	ResultInfo AccountStorageKvNamespaceListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceListResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceListResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceListResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceListResponse]
type accountStorageKvNamespaceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountStorageKvNamespaceListResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceListResponseErrorJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceListResponseError]
type accountStorageKvNamespaceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountStorageKvNamespaceListResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceListResponseMessageJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceListResponseMessage]
type accountStorageKvNamespaceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceListResponseResult struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                                            `json:"supports_url_encoding"`
	JSON                accountStorageKvNamespaceListResponseResultJSON `json:"-"`
}

// accountStorageKvNamespaceListResponseResultJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceListResponseResult]
type accountStorageKvNamespaceListResponseResultJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       accountStorageKvNamespaceListResponseResultInfoJSON `json:"-"`
}

// accountStorageKvNamespaceListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceListResponseResultInfo]
type accountStorageKvNamespaceListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStorageKvNamespaceListResponseSuccess bool

const (
	AccountStorageKvNamespaceListResponseSuccessTrue AccountStorageKvNamespaceListResponseSuccess = true
)

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse struct {
	Errors   []AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON contains
// the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                                                                       `json:"supports_url_encoding"`
	JSON                accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResultJSON `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResultJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResultJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccess bool

const (
	AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccessTrue AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccess = true
)

type AccountStorageKvNamespaceUpdateParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r AccountStorageKvNamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStorageKvNamespaceListParams struct {
	// Direction to order namespaces.
	Direction param.Field[AccountStorageKvNamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[AccountStorageKvNamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountStorageKvNamespaceListParams]'s query parameters as
// `url.Values`.
func (r AccountStorageKvNamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order namespaces.
type AccountStorageKvNamespaceListParamsDirection string

const (
	AccountStorageKvNamespaceListParamsDirectionAsc  AccountStorageKvNamespaceListParamsDirection = "asc"
	AccountStorageKvNamespaceListParamsDirectionDesc AccountStorageKvNamespaceListParamsDirection = "desc"
)

// Field to order results by.
type AccountStorageKvNamespaceListParamsOrder string

const (
	AccountStorageKvNamespaceListParamsOrderID    AccountStorageKvNamespaceListParamsOrder = "id"
	AccountStorageKvNamespaceListParamsOrderTitle AccountStorageKvNamespaceListParamsOrder = "title"
)

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
