// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// NamespaceKeyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceKeyService] method instead.
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
func (r *NamespaceKeyService) List(ctx context.Context, namespaceID string, params NamespaceKeyListParams, opts ...option.RequestOption) (res *pagination.CursorLimitPagination[Key], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/keys", params.AccountID, namespaceID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists a namespace's keys.
func (r *NamespaceKeyService) ListAutoPaging(ctx context.Context, namespaceID string, params NamespaceKeyListParams, opts ...option.RequestOption) *pagination.CursorLimitPaginationAutoPager[Key] {
	return pagination.NewCursorLimitPaginationAutoPager(r.List(ctx, namespaceID, params, opts...))
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
//
// Deprecated: Please use kv.namespaces.bulk_delete instead
func (r *NamespaceKeyService) BulkDelete(ctx context.Context, namespaceID string, params NamespaceKeyBulkDeleteParams, opts ...option.RequestOption) (res *NamespaceKeyBulkDeleteResponse, err error) {
	var env NamespaceKeyBulkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk/delete", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
//
// Deprecated: Please use kv.namespaces.bulk_update instead
func (r *NamespaceKeyService) BulkUpdate(ctx context.Context, namespaceID string, params NamespaceKeyBulkUpdateParams, opts ...option.RequestOption) (res *NamespaceKeyBulkUpdateResponse, err error) {
	var env NamespaceKeyBulkUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A name for a value. A value stored under a given key may be retrieved via the
// same key.
type Key struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64 `json:"expiration"`
	// Arbitrary JSON that is associated with a key.
	Metadata map[string]interface{} `json:"metadata"`
	JSON     keyJSON                `json:"-"`
}

// keyJSON contains the JSON metadata for the struct [Key]
type keyJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Key) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyBulkDeleteResponse struct {
	// Number of keys successfully updated
	SuccessfulKeyCount float64 `json:"successful_key_count"`
	// Name of the keys that failed to be fully updated. They should be retried.
	UnsuccessfulKeys []string                           `json:"unsuccessful_keys"`
	JSON             namespaceKeyBulkDeleteResponseJSON `json:"-"`
}

// namespaceKeyBulkDeleteResponseJSON contains the JSON metadata for the struct
// [NamespaceKeyBulkDeleteResponse]
type namespaceKeyBulkDeleteResponseJSON struct {
	SuccessfulKeyCount apijson.Field
	UnsuccessfulKeys   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *NamespaceKeyBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyBulkUpdateResponse struct {
	// Number of keys successfully updated
	SuccessfulKeyCount float64 `json:"successful_key_count"`
	// Name of the keys that failed to be fully updated. They should be retried.
	UnsuccessfulKeys []string                           `json:"unsuccessful_keys"`
	JSON             namespaceKeyBulkUpdateResponseJSON `json:"-"`
}

// namespaceKeyBulkUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceKeyBulkUpdateResponse]
type namespaceKeyBulkUpdateResponseJSON struct {
	SuccessfulKeyCount apijson.Field
	UnsuccessfulKeys   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *NamespaceKeyBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkUpdateResponseJSON) RawJSON() string {
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceKeyBulkDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      []string            `json:"body,required"`
}

func (r NamespaceKeyBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceKeyBulkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success NamespaceKeyBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceKeyBulkDeleteResponse                `json:"result,nullable"`
	JSON    namespaceKeyBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceKeyBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceKeyBulkDeleteResponseEnvelope]
type namespaceKeyBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceKeyBulkDeleteResponseEnvelopeSuccess bool

const (
	NamespaceKeyBulkDeleteResponseEnvelopeSuccessTrue NamespaceKeyBulkDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceKeyBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceKeyBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceKeyBulkUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                `path:"account_id,required"`
	Body      []NamespaceKeyBulkUpdateParamsBody `json:"body,required"`
}

func (r NamespaceKeyBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceKeyBulkUpdateParamsBody struct {
	// Whether or not the server should base64 decode the value before storing it.
	// Useful for writing values that wouldn't otherwise be valid JSON strings, such as
	// images.
	Base64 param.Field[bool] `json:"base64"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// should expire.
	Expiration param.Field[float64] `json:"expiration"`
	// The number of seconds for which the key should be visible before it expires. At
	// least 60.
	ExpirationTTL param.Field[float64] `json:"expiration_ttl"`
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key"`
	// Arbitrary JSON that is associated with a key.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value"`
}

func (r NamespaceKeyBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceKeyBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success NamespaceKeyBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceKeyBulkUpdateResponse                `json:"result,nullable"`
	JSON    namespaceKeyBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceKeyBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceKeyBulkUpdateResponseEnvelope]
type namespaceKeyBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceKeyBulkUpdateResponseEnvelopeSuccess bool

const (
	NamespaceKeyBulkUpdateResponseEnvelopeSuccessTrue NamespaceKeyBulkUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceKeyBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceKeyBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
