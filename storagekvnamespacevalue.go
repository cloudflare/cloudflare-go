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

// StorageKvNamespaceValueService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewStorageKvNamespaceValueService] method instead.
type StorageKvNamespaceValueService struct {
	Options []option.RequestOption
}

// NewStorageKvNamespaceValueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageKvNamespaceValueService(opts ...option.RequestOption) (r *StorageKvNamespaceValueService) {
	r = &StorageKvNamespaceValueService{}
	r.Options = opts
	return
}

// Remove a KV pair from the namespace. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name.
func (r *StorageKvNamespaceValueService) Delete(ctx context.Context, accountID string, namespaceID string, keyName string, opts ...option.RequestOption) (res *StorageKvNamespaceValueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceValueDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
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
func (r *StorageKvNamespaceValueService) Get(ctx context.Context, accountID string, namespaceID string, keyName string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored along with JSON metadata to be associated with the key/value pair.
// Existing values, expirations, and metadata will be overwritten. If neither
// `expiration` nor `expiration_ttl` is specified, the key-value pair will never
// expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.
func (r *StorageKvNamespaceValueService) Replace(ctx context.Context, accountID string, namespaceID string, keyName string, body StorageKvNamespaceValueReplaceParams, opts ...option.RequestOption) (res *StorageKvNamespaceValueReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceValueReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StorageKvNamespaceValueDeleteResponseUnknown] or
// [shared.UnionString].
type StorageKvNamespaceValueDeleteResponse interface {
	ImplementsStorageKvNamespaceValueDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKvNamespaceValueDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StorageKvNamespaceValueReplaceResponseUnknown] or
// [shared.UnionString].
type StorageKvNamespaceValueReplaceResponse interface {
	ImplementsStorageKvNamespaceValueReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKvNamespaceValueReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKvNamespaceValueDeleteResponseEnvelope struct {
	Errors   []StorageKvNamespaceValueDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceValueDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceValueDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceValueDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceValueDeleteResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceValueDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKvNamespaceValueDeleteResponseEnvelope]
type storageKvNamespaceValueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceValueDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    storageKvNamespaceValueDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceValueDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKvNamespaceValueDeleteResponseEnvelopeErrors]
type storageKvNamespaceValueDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceValueDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    storageKvNamespaceValueDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceValueDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKvNamespaceValueDeleteResponseEnvelopeMessages]
type storageKvNamespaceValueDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceValueDeleteResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceValueDeleteResponseEnvelopeSuccessTrue StorageKvNamespaceValueDeleteResponseEnvelopeSuccess = true
)

type StorageKvNamespaceValueReplaceParams struct {
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
}

func (r StorageKvNamespaceValueReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKvNamespaceValueReplaceResponseEnvelope struct {
	Errors   []StorageKvNamespaceValueReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceValueReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceValueReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceValueReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceValueReplaceResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceValueReplaceResponseEnvelopeJSON contains the JSON metadata
// for the struct [StorageKvNamespaceValueReplaceResponseEnvelope]
type storageKvNamespaceValueReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceValueReplaceResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    storageKvNamespaceValueReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceValueReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKvNamespaceValueReplaceResponseEnvelopeErrors]
type storageKvNamespaceValueReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceValueReplaceResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    storageKvNamespaceValueReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceValueReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKvNamespaceValueReplaceResponseEnvelopeMessages]
type storageKvNamespaceValueReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceValueReplaceResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceValueReplaceResponseEnvelopeSuccessTrue StorageKvNamespaceValueReplaceResponseEnvelopeSuccess = true
)
