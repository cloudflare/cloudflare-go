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

// StorageKVNamespaceBulkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStorageKVNamespaceBulkService]
// method instead.
type StorageKVNamespaceBulkService struct {
	Options []option.RequestOption
}

// NewStorageKVNamespaceBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageKVNamespaceBulkService(opts ...option.RequestOption) (r *StorageKVNamespaceBulkService) {
	r = &StorageKVNamespaceBulkService{}
	r.Options = opts
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *StorageKVNamespaceBulkService) Update(ctx context.Context, namespaceID string, params StorageKVNamespaceBulkUpdateParams, opts ...option.RequestOption) (res *StorageKVNamespaceBulkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceBulkUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *StorageKVNamespaceBulkService) Delete(ctx context.Context, namespaceID string, params StorageKVNamespaceBulkDeleteParams, opts ...option.RequestOption) (res *StorageKVNamespaceBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceBulkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StorageKVNamespaceBulkUpdateResponseUnknown] or
// [shared.UnionString].
type StorageKVNamespaceBulkUpdateResponse interface {
	ImplementsStorageKVNamespaceBulkUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKVNamespaceBulkUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StorageKVNamespaceBulkDeleteResponseUnknown] or
// [shared.UnionString].
type StorageKVNamespaceBulkDeleteResponse interface {
	ImplementsStorageKVNamespaceBulkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKVNamespaceBulkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKVNamespaceBulkUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                                   `path:"account_id,required"`
	Body      param.Field[[]StorageKVNamespaceBulkUpdateParamsBody] `json:"body,required"`
}

func (r StorageKVNamespaceBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type StorageKVNamespaceBulkUpdateParamsBody struct {
	// Whether or not the server should base64 decode the value before storing it.
	// Useful for writing values that wouldn't otherwise be valid JSON strings, such as
	// images.
	Base64 param.Field[bool] `json:"base64"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// should expire.
	Expiration param.Field[float64] `json:"expiration"`
	// The number of seconds for which the key should be visible before it expires. At
	// least 60.
	ExpirationTTL param.Field[float64] `json:"expiration_ttl"`
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key"`
	// Arbitrary JSON that is associated with a key.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value"`
}

func (r StorageKVNamespaceBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKVNamespaceBulkUpdateResponseEnvelope struct {
	Errors   []StorageKVNamespaceBulkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceBulkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKVNamespaceBulkUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceBulkUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKVNamespaceBulkUpdateResponseEnvelope]
type storageKVNamespaceBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceBulkUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    storageKVNamespaceBulkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceBulkUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKVNamespaceBulkUpdateResponseEnvelopeErrors]
type storageKVNamespaceBulkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceBulkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceBulkUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    storageKVNamespaceBulkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceBulkUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKVNamespaceBulkUpdateResponseEnvelopeMessages]
type storageKVNamespaceBulkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceBulkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceBulkUpdateResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceBulkUpdateResponseEnvelopeSuccessTrue StorageKVNamespaceBulkUpdateResponseEnvelopeSuccess = true
)

type StorageKVNamespaceBulkDeleteParams struct {
	// Identifier
	AccountID param.Field[string]   `path:"account_id,required"`
	Body      param.Field[[]string] `json:"body,required"`
}

func (r StorageKVNamespaceBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type StorageKVNamespaceBulkDeleteResponseEnvelope struct {
	Errors   []StorageKVNamespaceBulkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceBulkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKVNamespaceBulkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceBulkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKVNamespaceBulkDeleteResponseEnvelope]
type storageKVNamespaceBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceBulkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    storageKVNamespaceBulkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceBulkDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKVNamespaceBulkDeleteResponseEnvelopeErrors]
type storageKVNamespaceBulkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceBulkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceBulkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    storageKVNamespaceBulkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceBulkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKVNamespaceBulkDeleteResponseEnvelopeMessages]
type storageKVNamespaceBulkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceBulkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceBulkDeleteResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceBulkDeleteResponseEnvelopeSuccessTrue StorageKVNamespaceBulkDeleteResponseEnvelopeSuccess = true
)
