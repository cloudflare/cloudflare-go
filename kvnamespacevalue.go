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

// KVNamespaceValueService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewKVNamespaceValueService] method
// instead.
type KVNamespaceValueService struct {
	Options []option.RequestOption
}

// NewKVNamespaceValueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewKVNamespaceValueService(opts ...option.RequestOption) (r *KVNamespaceValueService) {
	r = &KVNamespaceValueService{}
	r.Options = opts
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored along with JSON metadata to be associated with the key/value pair.
// Existing values, expirations, and metadata will be overwritten. If neither
// `expiration` nor `expiration_ttl` is specified, the key-value pair will never
// expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.
func (r *KVNamespaceValueService) Update(ctx context.Context, namespaceID string, keyName string, params KVNamespaceValueUpdateParams, opts ...option.RequestOption) (res *KVNamespaceValueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceValueUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", params.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a KV pair from the namespace. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name.
func (r *KVNamespaceValueService) Delete(ctx context.Context, namespaceID string, keyName string, body KVNamespaceValueDeleteParams, opts ...option.RequestOption) (res *KVNamespaceValueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceValueDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", body.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the value associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name. If the KV-pair is set to expire at some point, the expiration time as
// measured in seconds since the UNIX epoch will be returned in the `expiration`
// response header.
func (r *KVNamespaceValueService) Get(ctx context.Context, namespaceID string, keyName string, query KVNamespaceValueGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", query.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [KVNamespaceValueUpdateResponseUnknown] or
// [shared.UnionString].
type KVNamespaceValueUpdateResponse interface {
	ImplementsKVNamespaceValueUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KVNamespaceValueUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [KVNamespaceValueDeleteResponseUnknown] or
// [shared.UnionString].
type KVNamespaceValueDeleteResponse interface {
	ImplementsKVNamespaceValueDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KVNamespaceValueDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type KVNamespaceValueUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
}

func (r KVNamespaceValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KVNamespaceValueUpdateResponseEnvelope struct {
	Errors   []KVNamespaceValueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceValueUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   KVNamespaceValueUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceValueUpdateResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceValueUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [KVNamespaceValueUpdateResponseEnvelope]
type kvNamespaceValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceValueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceValueUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    kvNamespaceValueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceValueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KVNamespaceValueUpdateResponseEnvelopeErrors]
type kvNamespaceValueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceValueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceValueUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceValueUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    kvNamespaceValueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceValueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KVNamespaceValueUpdateResponseEnvelopeMessages]
type kvNamespaceValueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceValueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceValueUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KVNamespaceValueUpdateResponseEnvelopeSuccess bool

const (
	KVNamespaceValueUpdateResponseEnvelopeSuccessTrue KVNamespaceValueUpdateResponseEnvelopeSuccess = true
)

type KVNamespaceValueDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type KVNamespaceValueDeleteResponseEnvelope struct {
	Errors   []KVNamespaceValueDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceValueDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   KVNamespaceValueDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceValueDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceValueDeleteResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceValueDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [KVNamespaceValueDeleteResponseEnvelope]
type kvNamespaceValueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceValueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceValueDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceValueDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    kvNamespaceValueDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceValueDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KVNamespaceValueDeleteResponseEnvelopeErrors]
type kvNamespaceValueDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceValueDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceValueDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceValueDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    kvNamespaceValueDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceValueDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KVNamespaceValueDeleteResponseEnvelopeMessages]
type kvNamespaceValueDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceValueDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceValueDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KVNamespaceValueDeleteResponseEnvelopeSuccess bool

const (
	KVNamespaceValueDeleteResponseEnvelopeSuccessTrue KVNamespaceValueDeleteResponseEnvelopeSuccess = true
)

type KVNamespaceValueGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
