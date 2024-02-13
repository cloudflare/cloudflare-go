// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// StorageKvNamespaceKeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStorageKvNamespaceKeyService]
// method instead.
type StorageKvNamespaceKeyService struct {
	Options []option.RequestOption
}

// NewStorageKvNamespaceKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageKvNamespaceKeyService(opts ...option.RequestOption) (r *StorageKvNamespaceKeyService) {
	r = &StorageKvNamespaceKeyService{}
	r.Options = opts
	return
}

// Lists a namespace's keys.
func (r *StorageKvNamespaceKeyService) List(ctx context.Context, accountID string, namespaceID string, query StorageKvNamespaceKeyListParams, opts ...option.RequestOption) (res *[]StorageKvNamespaceKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceKeyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/keys", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A name for a value. A value stored under a given key may be retrieved via the
// same key.
type StorageKvNamespaceKeyListResponse struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64 `json:"expiration"`
	// Arbitrary JSON that is associated with a key.
	Metadata interface{}                           `json:"metadata"`
	JSON     storageKvNamespaceKeyListResponseJSON `json:"-"`
}

// storageKvNamespaceKeyListResponseJSON contains the JSON metadata for the struct
// [StorageKvNamespaceKeyListResponse]
type storageKvNamespaceKeyListResponseJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceKeyListParams struct {
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the `cursors`
	// object in the `result_info` structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of keys to return. The cursor attribute may be used to iterate over
	// the next batch of keys if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
	// A string prefix used to filter down which keys will be returned. Exact matches
	// and any key names that begin with the prefix will be returned.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [StorageKvNamespaceKeyListParams]'s query parameters as
// `url.Values`.
func (r StorageKvNamespaceKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StorageKvNamespaceKeyListResponseEnvelope struct {
	Errors   []StorageKvNamespaceKeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceKeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []StorageKvNamespaceKeyListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    StorageKvNamespaceKeyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo StorageKvNamespaceKeyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       storageKvNamespaceKeyListResponseEnvelopeJSON       `json:"-"`
}

// storageKvNamespaceKeyListResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageKvNamespaceKeyListResponseEnvelope]
type storageKvNamespaceKeyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceKeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceKeyListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    storageKvNamespaceKeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceKeyListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StorageKvNamespaceKeyListResponseEnvelopeErrors]
type storageKvNamespaceKeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceKeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceKeyListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    storageKvNamespaceKeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceKeyListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StorageKvNamespaceKeyListResponseEnvelopeMessages]
type storageKvNamespaceKeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceKeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceKeyListResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceKeyListResponseEnvelopeSuccessTrue StorageKvNamespaceKeyListResponseEnvelopeSuccess = true
)

type StorageKvNamespaceKeyListResponseEnvelopeResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the cursors object
	// in the result_info structure.
	Cursor string                                                  `json:"cursor"`
	JSON   storageKvNamespaceKeyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// storageKvNamespaceKeyListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [StorageKvNamespaceKeyListResponseEnvelopeResultInfo]
type storageKvNamespaceKeyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceKeyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
