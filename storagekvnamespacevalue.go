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

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored along with JSON metadata to be associated with the key/value pair.
// Existing values, expirations, and metadata will be overwritten. If neither
// `expiration` nor `expiration_ttl` is specified, the key-value pair will never
// expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.
func (r *StorageKvNamespaceValueService) Update(ctx context.Context, accountID string, namespaceID string, keyName string, body StorageKvNamespaceValueUpdateParams, opts ...option.RequestOption) (res *StorageKvNamespaceValueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceValueUpdateResponseEnvelope
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

// Union satisfied by [StorageKvNamespaceValueUpdateResponseUnknown] or
// [shared.UnionString].
type StorageKvNamespaceValueUpdateResponse interface {
	ImplementsStorageKvNamespaceValueUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKvNamespaceValueUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
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

type StorageKvNamespaceValueUpdateParams struct {
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
}

func (r StorageKvNamespaceValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKvNamespaceValueUpdateResponseEnvelope struct {
	Errors   []StorageKvNamespaceValueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceValueUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceValueUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceValueUpdateResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceValueUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKvNamespaceValueUpdateResponseEnvelope]
type storageKvNamespaceValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceValueUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    storageKvNamespaceValueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceValueUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKvNamespaceValueUpdateResponseEnvelopeErrors]
type storageKvNamespaceValueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceValueUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    storageKvNamespaceValueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceValueUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKvNamespaceValueUpdateResponseEnvelopeMessages]
type storageKvNamespaceValueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceValueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceValueUpdateResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceValueUpdateResponseEnvelopeSuccessTrue StorageKvNamespaceValueUpdateResponseEnvelopeSuccess = true
)

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
