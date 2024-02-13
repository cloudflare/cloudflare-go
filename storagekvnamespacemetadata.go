// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// StorageKvNamespaceMetadataService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewStorageKvNamespaceMetadataService] method instead.
type StorageKvNamespaceMetadataService struct {
	Options []option.RequestOption
}

// NewStorageKvNamespaceMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewStorageKvNamespaceMetadataService(opts ...option.RequestOption) (r *StorageKvNamespaceMetadataService) {
	r = &StorageKvNamespaceMetadataService{}
	r.Options = opts
	return
}

// Returns the metadata associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name.
func (r *StorageKvNamespaceMetadataService) Get(ctx context.Context, accountID string, namespaceID string, keyName string, opts ...option.RequestOption) (res *StorageKvNamespaceMetadataGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceMetadataGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/metadata/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StorageKvNamespaceMetadataGetResponse = interface{}

type StorageKvNamespaceMetadataGetResponseEnvelope struct {
	Errors   []StorageKvNamespaceMetadataGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceMetadataGetResponseEnvelopeMessages `json:"messages,required"`
	// Arbitrary JSON that is associated with a key.
	Result StorageKvNamespaceMetadataGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceMetadataGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceMetadataGetResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceMetadataGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKvNamespaceMetadataGetResponseEnvelope]
type storageKvNamespaceMetadataGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceMetadataGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceMetadataGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    storageKvNamespaceMetadataGetResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceMetadataGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKvNamespaceMetadataGetResponseEnvelopeErrors]
type storageKvNamespaceMetadataGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceMetadataGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceMetadataGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    storageKvNamespaceMetadataGetResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceMetadataGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKvNamespaceMetadataGetResponseEnvelopeMessages]
type storageKvNamespaceMetadataGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceMetadataGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceMetadataGetResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceMetadataGetResponseEnvelopeSuccessTrue StorageKvNamespaceMetadataGetResponseEnvelopeSuccess = true
)
