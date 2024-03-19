// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// NamespaceValueService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNamespaceValueService] method
// instead.
type NamespaceValueService struct {
	Options []option.RequestOption
}

// NewNamespaceValueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceValueService(opts ...option.RequestOption) (r *NamespaceValueService) {
	r = &NamespaceValueService{}
	r.Options = opts
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored along with JSON metadata to be associated with the key/value pair.
// Existing values, expirations, and metadata will be overwritten. If neither
// `expiration` nor `expiration_ttl` is specified, the key-value pair will never
// expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.
func (r *NamespaceValueService) Update(ctx context.Context, namespaceID string, keyName string, params NamespaceValueUpdateParams, opts ...option.RequestOption) (res *NamespaceValueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceValueUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", params.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a KV pair from the namespace. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name.
func (r *NamespaceValueService) Delete(ctx context.Context, namespaceID string, keyName string, body NamespaceValueDeleteParams, opts ...option.RequestOption) (res *NamespaceValueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceValueDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", body.AccountID, namespaceID, keyName)
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
func (r *NamespaceValueService) Get(ctx context.Context, namespaceID string, keyName string, query NamespaceValueGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", query.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Union satisfied by [kv.NamespaceValueUpdateResponseUnknown] or
// [shared.UnionString].
type NamespaceValueUpdateResponse interface {
	ImplementsKVNamespaceValueUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceValueUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [kv.NamespaceValueDeleteResponseUnknown] or
// [shared.UnionString].
type NamespaceValueDeleteResponse interface {
	ImplementsKVNamespaceValueDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceValueDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NamespaceValueUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
}

func (r NamespaceValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceValueUpdateResponseEnvelope struct {
	Errors   []NamespaceValueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceValueUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   NamespaceValueUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceValueUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceValueUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceValueUpdateResponseEnvelope]
type namespaceValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceValueUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    namespaceValueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceValueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [NamespaceValueUpdateResponseEnvelopeErrors]
type namespaceValueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceValueUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    namespaceValueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceValueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [NamespaceValueUpdateResponseEnvelopeMessages]
type namespaceValueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceValueUpdateResponseEnvelopeSuccess bool

const (
	NamespaceValueUpdateResponseEnvelopeSuccessTrue NamespaceValueUpdateResponseEnvelopeSuccess = true
)

type NamespaceValueDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceValueDeleteResponseEnvelope struct {
	Errors   []NamespaceValueDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceValueDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   NamespaceValueDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceValueDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceValueDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceValueDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceValueDeleteResponseEnvelope]
type namespaceValueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceValueDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    namespaceValueDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceValueDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [NamespaceValueDeleteResponseEnvelopeErrors]
type namespaceValueDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceValueDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    namespaceValueDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceValueDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [NamespaceValueDeleteResponseEnvelopeMessages]
type namespaceValueDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceValueDeleteResponseEnvelopeSuccess bool

const (
	NamespaceValueDeleteResponseEnvelopeSuccessTrue NamespaceValueDeleteResponseEnvelopeSuccess = true
)

type NamespaceValueGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
