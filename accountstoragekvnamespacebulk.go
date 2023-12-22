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

// Remove multiple KV pairs from the Namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *AccountStorageKvNamespaceBulkService) Delete(ctx context.Context, accountIdentifier string, namespaceIdentifier string, body AccountStorageKvNamespaceBulkDeleteParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither expiration nor
// expiration_ttl is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *AccountStorageKvNamespaceBulkService) WorkersKvNamespaceWriteMultipleKeyValuePairs(ctx context.Context, accountIdentifier string, namespaceIdentifier string, body AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

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
	// A UTF-8 encoded string to be stored, up to 10 MB in length.
	Value param.Field[string] `json:"value"`
}

func (r AccountStorageKvNamespaceBulkWorkersKvNamespaceWriteMultipleKeyValuePairsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
