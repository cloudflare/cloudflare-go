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

// AccountStorageKvNamespaceKeyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountStorageKvNamespaceKeyService] method instead.
type AccountStorageKvNamespaceKeyService struct {
	Options []option.RequestOption
}

// NewAccountStorageKvNamespaceKeyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceKeyService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceKeyService) {
	r = &AccountStorageKvNamespaceKeyService{}
	r.Options = opts
	return
}

// Lists a namespace's keys.
func (r *AccountStorageKvNamespaceKeyService) List(ctx context.Context, accountIdentifier string, namespaceIdentifier string, query AccountStorageKvNamespaceKeyListParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/keys", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountStorageKvNamespaceKeyListResponse struct {
	Errors     []AccountStorageKvNamespaceKeyListResponseError    `json:"errors"`
	Messages   []AccountStorageKvNamespaceKeyListResponseMessage  `json:"messages"`
	Result     []AccountStorageKvNamespaceKeyListResponseResult   `json:"result"`
	ResultInfo AccountStorageKvNamespaceKeyListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceKeyListResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceKeyListResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceKeyListResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceKeyListResponse]
type accountStorageKvNamespaceKeyListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceKeyListResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountStorageKvNamespaceKeyListResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceKeyListResponseErrorJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceKeyListResponseError]
type accountStorageKvNamespaceKeyListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceKeyListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceKeyListResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountStorageKvNamespaceKeyListResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceKeyListResponseMessageJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceKeyListResponseMessage]
type accountStorageKvNamespaceKeyListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceKeyListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A name for a value. A value stored under a given key may be retrieved via the
// same key.
type AccountStorageKvNamespaceKeyListResponseResult struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64 `json:"expiration"`
	// Arbitrary JSON that is associated with a key.
	Metadata interface{}                                        `json:"metadata"`
	JSON     accountStorageKvNamespaceKeyListResponseResultJSON `json:"-"`
}

// accountStorageKvNamespaceKeyListResponseResultJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceKeyListResponseResult]
type accountStorageKvNamespaceKeyListResponseResultJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceKeyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceKeyListResponseResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the cursors object
	// in the result_info structure.
	Cursor string                                                 `json:"cursor"`
	JSON   accountStorageKvNamespaceKeyListResponseResultInfoJSON `json:"-"`
}

// accountStorageKvNamespaceKeyListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceKeyListResponseResultInfo]
type accountStorageKvNamespaceKeyListResponseResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceKeyListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStorageKvNamespaceKeyListResponseSuccess bool

const (
	AccountStorageKvNamespaceKeyListResponseSuccessTrue AccountStorageKvNamespaceKeyListResponseSuccess = true
)

type AccountStorageKvNamespaceKeyListParams struct {
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the `cursors`
	// object in the `result_info` structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of keys to return. The cursor attribute may be used to iterate over
	// the next batch of keys if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
	// A string prefix used to filter down which keys will be returned. Exact matches
	// and any key names that begin with the prefix will be returned.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [AccountStorageKvNamespaceKeyListParams]'s query parameters
// as `url.Values`.
func (r AccountStorageKvNamespaceKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
