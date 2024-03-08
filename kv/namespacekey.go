// File generated from our OpenAPI spec by Stainless.

package kv

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// NamespaceKeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNamespaceKeyService] method
// instead.
type NamespaceKeyService struct {
	Options []option.RequestOption
}

// NewNamespaceKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceKeyService(opts ...option.RequestOption) (r *NamespaceKeyService) {
	r = &NamespaceKeyService{}
	r.Options = opts
	return
}

// Lists a namespace's keys.
func (r *NamespaceKeyService) List(ctx context.Context, namespaceID string, params NamespaceKeyListParams, opts ...option.RequestOption) (res *[]WorkersKVKey, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceKeyListResponseEnvelope
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
type WorkersKVKey struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64 `json:"expiration"`
	// Arbitrary JSON that is associated with a key.
	Metadata interface{}      `json:"metadata"`
	JSON     workersKVKeyJSON `json:"-"`
}

// workersKVKeyJSON contains the JSON metadata for the struct [WorkersKVKey]
type workersKVKeyJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersKVKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVKeyJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyListParams struct {
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

// URLQuery serializes [NamespaceKeyListParams]'s query parameters as `url.Values`.
func (r NamespaceKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NamespaceKeyListResponseEnvelope struct {
	Errors   []NamespaceKeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceKeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkersKVKey                             `json:"result,required"`
	// Whether the API call was successful
	Success    NamespaceKeyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo NamespaceKeyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       namespaceKeyListResponseEnvelopeJSON       `json:"-"`
}

// namespaceKeyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceKeyListResponseEnvelope]
type namespaceKeyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    namespaceKeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceKeyListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceKeyListResponseEnvelopeErrors]
type namespaceKeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    namespaceKeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceKeyListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [NamespaceKeyListResponseEnvelopeMessages]
type namespaceKeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceKeyListResponseEnvelopeSuccess bool

const (
	NamespaceKeyListResponseEnvelopeSuccessTrue NamespaceKeyListResponseEnvelopeSuccess = true
)

type NamespaceKeyListResponseEnvelopeResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the cursors object
	// in the result_info structure.
	Cursor string                                         `json:"cursor"`
	JSON   namespaceKeyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// namespaceKeyListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [NamespaceKeyListResponseEnvelopeResultInfo]
type namespaceKeyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
