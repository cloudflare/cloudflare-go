// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// KVNamespaceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewKVNamespaceService] method
// instead.
type KVNamespaceService struct {
	Options  []option.RequestOption
	Bulk     *KVNamespaceBulkService
	Keys     *KVNamespaceKeyService
	Metadata *KVNamespaceMetadataService
	Values   *KVNamespaceValueService
}

// NewKVNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKVNamespaceService(opts ...option.RequestOption) (r *KVNamespaceService) {
	r = &KVNamespaceService{}
	r.Options = opts
	r.Bulk = NewKVNamespaceBulkService(opts...)
	r.Keys = NewKVNamespaceKeyService(opts...)
	r.Metadata = NewKVNamespaceMetadataService(opts...)
	r.Values = NewKVNamespaceValueService(opts...)
	return
}

// Creates a namespace under the given title. A `400` is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *KVNamespaceService) New(ctx context.Context, params KVNamespaceNewParams, opts ...option.RequestOption) (res *WorkersKVNamespace, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modifies a namespace's title.
func (r *KVNamespaceService) Update(ctx context.Context, namespaceID string, params KVNamespaceUpdateParams, opts ...option.RequestOption) (res *KVNamespaceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the namespaces owned by an account.
func (r *KVNamespaceService) List(ctx context.Context, params KVNamespaceListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[WorkersKVNamespace], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", params.AccountID)
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

// Returns the namespaces owned by an account.
func (r *KVNamespaceService) ListAutoPaging(ctx context.Context, params KVNamespaceListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[WorkersKVNamespace] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the namespace corresponding to the given ID.
func (r *KVNamespaceService) Delete(ctx context.Context, namespaceID string, body KVNamespaceDeleteParams, opts ...option.RequestOption) (res *KVNamespaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KVNamespaceDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", body.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersKVNamespace struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                   `json:"supports_url_encoding"`
	JSON                workersKVNamespaceJSON `json:"-"`
}

// workersKVNamespaceJSON contains the JSON metadata for the struct
// [WorkersKVNamespace]
type workersKVNamespaceJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WorkersKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [KVNamespaceUpdateResponseUnknown] or [shared.UnionString].
type KVNamespaceUpdateResponse interface {
	ImplementsKVNamespaceUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KVNamespaceUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [KVNamespaceDeleteResponseUnknown] or [shared.UnionString].
type KVNamespaceDeleteResponse interface {
	ImplementsKVNamespaceDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KVNamespaceDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type KVNamespaceNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r KVNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KVNamespaceNewResponseEnvelope struct {
	Errors   []KVNamespaceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceNewResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersKVNamespace                       `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceNewResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [KVNamespaceNewResponseEnvelope]
type kvNamespaceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KVNamespaceNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    kvNamespaceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [KVNamespaceNewResponseEnvelopeErrors]
type kvNamespaceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KVNamespaceNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    kvNamespaceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [KVNamespaceNewResponseEnvelopeMessages]
type kvNamespaceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KVNamespaceNewResponseEnvelopeSuccess bool

const (
	KVNamespaceNewResponseEnvelopeSuccessTrue KVNamespaceNewResponseEnvelopeSuccess = true
)

type KVNamespaceUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r KVNamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KVNamespaceUpdateResponseEnvelope struct {
	Errors   []KVNamespaceUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   KVNamespaceUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceUpdateResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [KVNamespaceUpdateResponseEnvelope]
type kvNamespaceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KVNamespaceUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    kvNamespaceUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [KVNamespaceUpdateResponseEnvelopeErrors]
type kvNamespaceUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KVNamespaceUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    kvNamespaceUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [KVNamespaceUpdateResponseEnvelopeMessages]
type kvNamespaceUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KVNamespaceUpdateResponseEnvelopeSuccess bool

const (
	KVNamespaceUpdateResponseEnvelopeSuccessTrue KVNamespaceUpdateResponseEnvelopeSuccess = true
)

type KVNamespaceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order namespaces.
	Direction param.Field[KVNamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[KVNamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [KVNamespaceListParams]'s query parameters as `url.Values`.
func (r KVNamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order namespaces.
type KVNamespaceListParamsDirection string

const (
	KVNamespaceListParamsDirectionAsc  KVNamespaceListParamsDirection = "asc"
	KVNamespaceListParamsDirectionDesc KVNamespaceListParamsDirection = "desc"
)

// Field to order results by.
type KVNamespaceListParamsOrder string

const (
	KVNamespaceListParamsOrderID    KVNamespaceListParamsOrder = "id"
	KVNamespaceListParamsOrderTitle KVNamespaceListParamsOrder = "title"
)

type KVNamespaceDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type KVNamespaceDeleteResponseEnvelope struct {
	Errors   []KVNamespaceDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KVNamespaceDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   KVNamespaceDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KVNamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    kvNamespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// kvNamespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [KVNamespaceDeleteResponseEnvelope]
type kvNamespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KVNamespaceDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    kvNamespaceDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// kvNamespaceDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [KVNamespaceDeleteResponseEnvelopeErrors]
type kvNamespaceDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KVNamespaceDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    kvNamespaceDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// kvNamespaceDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [KVNamespaceDeleteResponseEnvelopeMessages]
type kvNamespaceDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KVNamespaceDeleteResponseEnvelopeSuccess bool

const (
	KVNamespaceDeleteResponseEnvelopeSuccessTrue KVNamespaceDeleteResponseEnvelopeSuccess = true
)
