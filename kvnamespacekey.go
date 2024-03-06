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

// KVNamespaceKeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewKVNamespaceKeyService] method
// instead.
type KVNamespaceKeyService struct {
	Options []option.RequestOption
}

// NewKVNamespaceKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKVNamespaceKeyService(opts ...option.RequestOption) (r *KVNamespaceKeyService) {
	r = &KVNamespaceKeyService{}
	r.Options = opts
	return
}

// Lists a namespace's keys.
func (r *KVNamespaceKeyService) List(ctx context.Context, namespaceID string, params KVNamespaceKeyListParams, opts ...option.RequestOption) (res *[]KVNamespaceKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceKeyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/keys", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A name for a value. A value stored under a given key may be retrieved via the
// same key.
type KVNamespaceKeyListResponse struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64 `json:"expiration"`
	// Arbitrary JSON that is associated with a key.
	Metadata interface{}                    `json:"metadata"`
	JSON     kvNamespaceKeyListResponseJSON `json:"-"`
}

// kvNamespaceKeyListResponseJSON contains the JSON metadata for the struct
// [KVNamespaceKeyListResponse]
type kvNamespaceKeyListResponseJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceKeyListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

// URLQuery serializes [KVNamespaceKeyListParams]'s query parameters as
// `url.Values`.
func (r KVNamespaceKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type KVNamespaceKeyListResponseEnvelope struct {
	Errors   []KVNamespaceKeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceKeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []KVNamespaceKeyListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    KVNamespaceKeyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo KVNamespaceKeyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       kvNamespaceKeyListResponseEnvelopeJSON       `json:"-"`
}

// kvNamespaceKeyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [KVNamespaceKeyListResponseEnvelope]
type kvNamespaceKeyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceKeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceKeyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceKeyListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    kvNamespaceKeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceKeyListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [KVNamespaceKeyListResponseEnvelopeErrors]
type kvNamespaceKeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceKeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceKeyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KVNamespaceKeyListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    kvNamespaceKeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceKeyListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [KVNamespaceKeyListResponseEnvelopeMessages]
type kvNamespaceKeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceKeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceKeyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KVNamespaceKeyListResponseEnvelopeSuccess bool

const (
	KVNamespaceKeyListResponseEnvelopeSuccessTrue KVNamespaceKeyListResponseEnvelopeSuccess = true
)

type KVNamespaceKeyListResponseEnvelopeResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the cursors object
	// in the result_info structure.
	Cursor string                                           `json:"cursor"`
	JSON   kvNamespaceKeyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// kvNamespaceKeyListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [KVNamespaceKeyListResponseEnvelopeResultInfo]
type kvNamespaceKeyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceKeyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceKeyListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
