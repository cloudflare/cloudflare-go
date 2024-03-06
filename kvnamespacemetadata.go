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

// KVNamespaceMetadataService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewKVNamespaceMetadataService]
// method instead.
type KVNamespaceMetadataService struct {
	Options []option.RequestOption
}

// NewKVNamespaceMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewKVNamespaceMetadataService(opts ...option.RequestOption) (r *KVNamespaceMetadataService) {
	r = &KVNamespaceMetadataService{}
	r.Options = opts
	return
}

// Returns the metadata associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name.
func (r *KVNamespaceMetadataService) Get(ctx context.Context, namespaceID string, keyName string, query KVNamespaceMetadataGetParams, opts ...option.RequestOption) (res *KVNamespaceMetadataGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceMetadataGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/metadata/%s", query.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type KVNamespaceMetadataGetResponse = interface{}

type KVNamespaceMetadataGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type KVNamespaceMetadataGetResponseEnvelope struct {
	Errors   []KVNamespaceMetadataGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceMetadataGetResponseEnvelopeMessages `json:"messages,required"`
	// Arbitrary JSON that is associated with a key.
	Result KVNamespaceMetadataGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceMetadataGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceMetadataGetResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceMetadataGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [KVNamespaceMetadataGetResponseEnvelope]
type kvNamespaceMetadataGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceMetadataGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceMetadataGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceMetadataGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    kvNamespaceMetadataGetResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceMetadataGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KVNamespaceMetadataGetResponseEnvelopeErrors]
type kvNamespaceMetadataGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceMetadataGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceMetadataGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceMetadataGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    kvNamespaceMetadataGetResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceMetadataGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KVNamespaceMetadataGetResponseEnvelopeMessages]
type kvNamespaceMetadataGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceMetadataGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceMetadataGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KVNamespaceMetadataGetResponseEnvelopeSuccess bool

const (
	KVNamespaceMetadataGetResponseEnvelopeSuccessTrue KVNamespaceMetadataGetResponseEnvelopeSuccess = true
)
