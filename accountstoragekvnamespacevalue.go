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

// AccountStorageKvNamespaceValueService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountStorageKvNamespaceValueService] method instead.
type AccountStorageKvNamespaceValueService struct {
	Options []option.RequestOption
}

// NewAccountStorageKvNamespaceValueService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceValueService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceValueService) {
	r = &AccountStorageKvNamespaceValueService{}
	r.Options = opts
	return
}

// Returns the value associated with the given key in the given namespace. Use
// URL-encoding to use special characters (e.g. :, !, %) in the key name. If the
// KV-pair is set to expire at some point, the expiration time as measured in
// seconds since the UNIX epoch will be returned in the "Expiration" response
// header.
func (r *AccountStorageKvNamespaceValueService) Get(ctx context.Context, accountIdentifier string, namespaceIdentifier string, keyName string, opts ...option.RequestOption) (res *KvValue, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountIdentifier, namespaceIdentifier, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (e.g. :, !, %) in the key name. Body should be the value to be stored along with
// json metadata to be associated with the key/value pair. Existing values,
// expirations and metadata will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored.
func (r *AccountStorageKvNamespaceValueService) Update(ctx context.Context, accountIdentifier string, namespaceIdentifier string, keyName string, body AccountStorageKvNamespaceValueUpdateParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountIdentifier, namespaceIdentifier, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove a KV pair from the Namespace. Use URL-encoding to use special characters
// (e.g. :, !, %) in the key name.
func (r *AccountStorageKvNamespaceValueService) Delete(ctx context.Context, accountIdentifier string, namespaceIdentifier string, keyName string, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountIdentifier, namespaceIdentifier, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type KvValue = string

type KvValueParam = string

type AccountStorageKvNamespaceValueUpdateParams struct {
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 10 MB in length.
	Value param.Field[KvValueParam] `json:"value,required"`
}

func (r AccountStorageKvNamespaceValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
