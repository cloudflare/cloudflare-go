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

// StorageKVNamespaceMetadataService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewStorageKVNamespaceMetadataService] method instead.
type StorageKVNamespaceMetadataService struct {
	Options []option.RequestOption
}

// NewStorageKVNamespaceMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewStorageKVNamespaceMetadataService(opts ...option.RequestOption) (r *StorageKVNamespaceMetadataService) {
	r = &StorageKVNamespaceMetadataService{}
	r.Options = opts
	return
}

// Returns the metadata associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name.
func (r *StorageKVNamespaceMetadataService) Get(ctx context.Context, namespaceID string, keyName string, query StorageKVNamespaceMetadataGetParams, opts ...option.RequestOption) (res *StorageKVNamespaceMetadataGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceMetadataGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/metadata/%s", query.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StorageKVNamespaceMetadataGetResponse = interface{}

type StorageKVNamespaceMetadataGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StorageKVNamespaceMetadataGetResponseEnvelope struct {
	Errors   []StorageKVNamespaceMetadataGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceMetadataGetResponseEnvelopeMessages `json:"messages,required"`
	// Arbitrary JSON that is associated with a key.
	Result StorageKVNamespaceMetadataGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceMetadataGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceMetadataGetResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceMetadataGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKVNamespaceMetadataGetResponseEnvelope]
type storageKVNamespaceMetadataGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceMetadataGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceMetadataGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    storageKVNamespaceMetadataGetResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceMetadataGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKVNamespaceMetadataGetResponseEnvelopeErrors]
type storageKVNamespaceMetadataGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceMetadataGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceMetadataGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    storageKVNamespaceMetadataGetResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceMetadataGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKVNamespaceMetadataGetResponseEnvelopeMessages]
type storageKVNamespaceMetadataGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceMetadataGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceMetadataGetResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceMetadataGetResponseEnvelopeSuccessTrue StorageKVNamespaceMetadataGetResponseEnvelopeSuccess = true
)
