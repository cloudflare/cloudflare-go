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

// KVNamespaceBulkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewKVNamespaceBulkService] method
// instead.
type KVNamespaceBulkService struct {
	Options []option.RequestOption
}

// NewKVNamespaceBulkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKVNamespaceBulkService(opts ...option.RequestOption) (r *KVNamespaceBulkService) {
	r = &KVNamespaceBulkService{}
	r.Options = opts
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *KVNamespaceBulkService) Update(ctx context.Context, namespaceID string, params KVNamespaceBulkUpdateParams, opts ...option.RequestOption) (res *KVNamespaceBulkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceBulkUpdateResponseEnvelope
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
func (r *KVNamespaceBulkService) Delete(ctx context.Context, namespaceID string, params KVNamespaceBulkDeleteParams, opts ...option.RequestOption) (res *KVNamespaceBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceBulkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [KVNamespaceBulkUpdateResponseUnknown] or
// [shared.UnionString].
type KVNamespaceBulkUpdateResponse interface {
	ImplementsKVNamespaceBulkUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KVNamespaceBulkUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [KVNamespaceBulkDeleteResponseUnknown] or
// [shared.UnionString].
type KVNamespaceBulkDeleteResponse interface {
	ImplementsKVNamespaceBulkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KVNamespaceBulkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type KVNamespaceBulkUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                            `path:"account_id,required"`
	Body      param.Field[[]KVNamespaceBulkUpdateParamsBody] `json:"body,required"`
}

func (r KVNamespaceBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type KVNamespaceBulkUpdateParamsBody struct {
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

func (r KVNamespaceBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KVNamespaceBulkUpdateResponseEnvelope struct {
	Errors   []KVNamespaceBulkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceBulkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   KVNamespaceBulkUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [KVNamespaceBulkUpdateResponseEnvelope]
type kvNamespaceBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceBulkUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    kvNamespaceBulkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceBulkUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KVNamespaceBulkUpdateResponseEnvelopeErrors]
type kvNamespaceBulkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceBulkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceBulkUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceBulkUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    kvNamespaceBulkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceBulkUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [KVNamespaceBulkUpdateResponseEnvelopeMessages]
type kvNamespaceBulkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceBulkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceBulkUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KVNamespaceBulkUpdateResponseEnvelopeSuccess bool

const (
	KVNamespaceBulkUpdateResponseEnvelopeSuccessTrue KVNamespaceBulkUpdateResponseEnvelopeSuccess = true
)

type KVNamespaceBulkDeleteParams struct {
	// Identifier
	AccountID param.Field[string]   `path:"account_id,required"`
	Body      param.Field[[]string] `json:"body,required"`
}

func (r KVNamespaceBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type KVNamespaceBulkDeleteResponseEnvelope struct {
	Errors   []KVNamespaceBulkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceBulkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   KVNamespaceBulkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [KVNamespaceBulkDeleteResponseEnvelope]
type kvNamespaceBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceBulkDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    kvNamespaceBulkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceBulkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KVNamespaceBulkDeleteResponseEnvelopeErrors]
type kvNamespaceBulkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceBulkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceBulkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceBulkDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    kvNamespaceBulkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceBulkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [KVNamespaceBulkDeleteResponseEnvelopeMessages]
type kvNamespaceBulkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceBulkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceBulkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KVNamespaceBulkDeleteResponseEnvelopeSuccess bool

const (
	KVNamespaceBulkDeleteResponseEnvelopeSuccessTrue KVNamespaceBulkDeleteResponseEnvelopeSuccess = true
)
