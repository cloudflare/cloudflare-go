// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccountStorageKvNamespaceBulkService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountStorageKvNamespaceBulkService] method instead.
type AccountStorageKvNamespaceBulkService struct {
	Options []option.RequestOption
}

// NewAccountStorageKvNamespaceBulkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceBulkService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceBulkService) {
	r = &AccountStorageKvNamespaceBulkService{}
	r.Options = opts
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *AccountStorageKvNamespaceBulkService) Delete(ctx context.Context, accountIdentifier string, namespaceIdentifier string, body AccountStorageKvNamespaceBulkDeleteParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *AccountStorageKvNamespaceBulkService) WorkersKvNamespaceWriteMultipleKeyValuePairs(ctx context.Context, accountIdentifier string, namespaceIdentifier string, body AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountStorageKvNamespaceBulkDeleteResponse struct {
	Errors   []AccountStorageKvNamespaceBulkDeleteResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceBulkDeleteResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceBulkDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceBulkDeleteResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceBulkDeleteResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceBulkDeleteResponseJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceBulkDeleteResponse]
type accountStorageKvNamespaceBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceBulkDeleteResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountStorageKvNamespaceBulkDeleteResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceBulkDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceBulkDeleteResponseError]
type accountStorageKvNamespaceBulkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceBulkDeleteResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountStorageKvNamespaceBulkDeleteResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceBulkDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceBulkDeleteResponseMessage]
type accountStorageKvNamespaceBulkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountStorageKvNamespaceBulkDeleteResponseResultUnknown] or
// [shared.UnionString].
type AccountStorageKvNamespaceBulkDeleteResponseResult interface {
	ImplementsAccountStorageKvNamespaceBulkDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceBulkDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountStorageKvNamespaceBulkDeleteResponseSuccess bool

const (
	AccountStorageKvNamespaceBulkDeleteResponseSuccessTrue AccountStorageKvNamespaceBulkDeleteResponseSuccess = true
)

type AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponse struct {
	Errors   []AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponse]
type accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseError struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseError]
type accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseMessage struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseMessage]
type accountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseResultUnknown]
// or [shared.UnionString].
type AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseResult interface {
	ImplementsAccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseSuccess bool

const (
	AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseSuccessTrue AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsResponseSuccess = true
)

type AccountStorageKvNamespaceBulkDeleteParams struct {
	Body param.Field[[]string] `json:"body,required"`
}

func (r AccountStorageKvNamespaceBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParams struct {
	Body param.Field[[]AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParamsBody] `json:"body,required"`
}

func (r AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParamsBody struct {
	// Whether or not the server should base64 decode the value before storing it.
	// Useful for writing values that wouldn't otherwise be valid JSON strings, such as
	// images.
	Base64 param.Field[bool] `json:"base64"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// should expire.
	Expiration param.Field[float64] `json:"expiration"`
	// The number of seconds for which the key should be visible before it expires. At
	// least 60.
	ExpirationTtl param.Field[float64] `json:"expiration_ttl"`
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key"`
	// Arbitrary JSON that is associated with a key.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value"`
}

func (r AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
