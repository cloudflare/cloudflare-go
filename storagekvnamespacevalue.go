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

// StorageKVNamespaceValueService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewStorageKVNamespaceValueService] method instead.
type StorageKVNamespaceValueService struct {
	Options []option.RequestOption
}

// NewStorageKVNamespaceValueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageKVNamespaceValueService(opts ...option.RequestOption) (r *StorageKVNamespaceValueService) {
	r = &StorageKVNamespaceValueService{}
	r.Options = opts
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored along with JSON metadata to be associated with the key/value pair.
// Existing values, expirations, and metadata will be overwritten. If neither
// `expiration` nor `expiration_ttl` is specified, the key-value pair will never
// expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.
func (r *StorageKVNamespaceValueService) Update(ctx context.Context, accountID string, namespaceID string, keyName string, body StorageKVNamespaceValueUpdateParams, opts ...option.RequestOption) (res *StorageKVNamespaceValueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceValueUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a KV pair from the namespace. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name.
func (r *StorageKVNamespaceValueService) Delete(ctx context.Context, accountID string, namespaceID string, keyName string, opts ...option.RequestOption) (res *StorageKVNamespaceValueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceValueDeleteResponseEnvelope
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
func (r *StorageKVNamespaceValueService) Get(ctx context.Context, accountID string, namespaceID string, keyName string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Union satisfied by [StorageKVNamespaceValueUpdateResponseUnknown] or
// [shared.UnionString].
type StorageKVNamespaceValueUpdateResponse interface {
	ImplementsStorageKVNamespaceValueUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKVNamespaceValueUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StorageKVNamespaceValueDeleteResponseUnknown] or
// [shared.UnionString].
type StorageKVNamespaceValueDeleteResponse interface {
	ImplementsStorageKVNamespaceValueDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKVNamespaceValueDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKVNamespaceValueUpdateParams struct {
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
}

func (r StorageKVNamespaceValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKVNamespaceValueUpdateResponseEnvelope struct {
	Errors   []StorageKVNamespaceValueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceValueUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKVNamespaceValueUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceValueUpdateResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceValueUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKVNamespaceValueUpdateResponseEnvelope]
type storageKVNamespaceValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceValueUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    storageKVNamespaceValueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceValueUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKVNamespaceValueUpdateResponseEnvelopeErrors]
type storageKVNamespaceValueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceValueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceValueUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    storageKVNamespaceValueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceValueUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKVNamespaceValueUpdateResponseEnvelopeMessages]
type storageKVNamespaceValueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceValueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceValueUpdateResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceValueUpdateResponseEnvelopeSuccessTrue StorageKVNamespaceValueUpdateResponseEnvelopeSuccess = true
)

type StorageKVNamespaceValueDeleteResponseEnvelope struct {
	Errors   []StorageKVNamespaceValueDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceValueDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKVNamespaceValueDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceValueDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceValueDeleteResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceValueDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKVNamespaceValueDeleteResponseEnvelope]
type storageKVNamespaceValueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceValueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceValueDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    storageKVNamespaceValueDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceValueDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKVNamespaceValueDeleteResponseEnvelopeErrors]
type storageKVNamespaceValueDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceValueDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceValueDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    storageKVNamespaceValueDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceValueDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKVNamespaceValueDeleteResponseEnvelopeMessages]
type storageKVNamespaceValueDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceValueDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceValueDeleteResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceValueDeleteResponseEnvelopeSuccessTrue StorageKVNamespaceValueDeleteResponseEnvelopeSuccess = true
)
