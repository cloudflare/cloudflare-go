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

// AccountStorageKvNamespaceMetadataService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountStorageKvNamespaceMetadataService] method instead.
type AccountStorageKvNamespaceMetadataService struct {
	Options []option.RequestOption
}

// NewAccountStorageKvNamespaceMetadataService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceMetadataService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceMetadataService) {
	r = &AccountStorageKvNamespaceMetadataService{}
	r.Options = opts
	return
}

// Returns the metadata associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name.
func (r *AccountStorageKvNamespaceMetadataService) Get(ctx context.Context, accountIdentifier string, namespaceIdentifier string, keyName string, opts ...option.RequestOption) (res *AccountStorageKvNamespaceMetadataGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/metadata/%s", accountIdentifier, namespaceIdentifier, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStorageKvNamespaceMetadataGetResponse struct {
	Errors   []AccountStorageKvNamespaceMetadataGetResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceMetadataGetResponseMessage `json:"messages"`
	// Arbitrary JSON that is associated with a key.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceMetadataGetResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceMetadataGetResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceMetadataGetResponseJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceMetadataGetResponse]
type accountStorageKvNamespaceMetadataGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceMetadataGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceMetadataGetResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountStorageKvNamespaceMetadataGetResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceMetadataGetResponseErrorJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceMetadataGetResponseError]
type accountStorageKvNamespaceMetadataGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceMetadataGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceMetadataGetResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountStorageKvNamespaceMetadataGetResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceMetadataGetResponseMessageJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceMetadataGetResponseMessage]
type accountStorageKvNamespaceMetadataGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceMetadataGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStorageKvNamespaceMetadataGetResponseSuccess bool

const (
	AccountStorageKvNamespaceMetadataGetResponseSuccessTrue AccountStorageKvNamespaceMetadataGetResponseSuccess = true
)
