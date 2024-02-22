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

// StorageKVNamespaceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStorageKVNamespaceService] method
// instead.
type StorageKVNamespaceService struct {
	Options  []option.RequestOption
	Bulk     *StorageKVNamespaceBulkService
	Keys     *StorageKVNamespaceKeyService
	Metadata *StorageKVNamespaceMetadataService
	Values   *StorageKVNamespaceValueService
}

// NewStorageKVNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageKVNamespaceService(opts ...option.RequestOption) (r *StorageKVNamespaceService) {
	r = &StorageKVNamespaceService{}
	r.Options = opts
	r.Bulk = NewStorageKVNamespaceBulkService(opts...)
	r.Keys = NewStorageKVNamespaceKeyService(opts...)
	r.Metadata = NewStorageKVNamespaceMetadataService(opts...)
	r.Values = NewStorageKVNamespaceValueService(opts...)
	return
}

// Creates a namespace under the given title. A `400` is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *StorageKVNamespaceService) New(ctx context.Context, accountID string, body StorageKVNamespaceNewParams, opts ...option.RequestOption) (res *StorageKVNamespaceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modifies a namespace's title.
func (r *StorageKVNamespaceService) Update(ctx context.Context, accountID string, namespaceID string, body StorageKVNamespaceUpdateParams, opts ...option.RequestOption) (res *StorageKVNamespaceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the namespaces owned by an account.
func (r *StorageKVNamespaceService) List(ctx context.Context, accountID string, query StorageKVNamespaceListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[StorageKVNamespaceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *StorageKVNamespaceService) ListAutoPaging(ctx context.Context, accountID string, query StorageKVNamespaceListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[StorageKVNamespaceListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Deletes the namespace corresponding to the given ID.
func (r *StorageKVNamespaceService) Delete(ctx context.Context, accountID string, namespaceID string, opts ...option.RequestOption) (res *StorageKVNamespaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKVNamespaceDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StorageKVNamespaceNewResponse struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                              `json:"supports_url_encoding"`
	JSON                storageKVNamespaceNewResponseJSON `json:"-"`
}

// storageKVNamespaceNewResponseJSON contains the JSON metadata for the struct
// [StorageKVNamespaceNewResponse]
type storageKVNamespaceNewResponseJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StorageKVNamespaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [StorageKVNamespaceUpdateResponseUnknown] or
// [shared.UnionString].
type StorageKVNamespaceUpdateResponse interface {
	ImplementsStorageKVNamespaceUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKVNamespaceUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKVNamespaceListResponse struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                               `json:"supports_url_encoding"`
	JSON                storageKVNamespaceListResponseJSON `json:"-"`
}

// storageKVNamespaceListResponseJSON contains the JSON metadata for the struct
// [StorageKVNamespaceListResponse]
type storageKVNamespaceListResponseJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StorageKVNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [StorageKVNamespaceDeleteResponseUnknown] or
// [shared.UnionString].
type StorageKVNamespaceDeleteResponse interface {
	ImplementsStorageKVNamespaceDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKVNamespaceDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKVNamespaceNewParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r StorageKVNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKVNamespaceNewResponseEnvelope struct {
	Errors   []StorageKVNamespaceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKVNamespaceNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceNewResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageKVNamespaceNewResponseEnvelope]
type storageKVNamespaceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    storageKVNamespaceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StorageKVNamespaceNewResponseEnvelopeErrors]
type storageKVNamespaceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    storageKVNamespaceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StorageKVNamespaceNewResponseEnvelopeMessages]
type storageKVNamespaceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceNewResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceNewResponseEnvelopeSuccessTrue StorageKVNamespaceNewResponseEnvelopeSuccess = true
)

type StorageKVNamespaceUpdateParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r StorageKVNamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKVNamespaceUpdateResponseEnvelope struct {
	Errors   []StorageKVNamespaceUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKVNamespaceUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceUpdateResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageKVNamespaceUpdateResponseEnvelope]
type storageKVNamespaceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    storageKVNamespaceUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StorageKVNamespaceUpdateResponseEnvelopeErrors]
type storageKVNamespaceUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    storageKVNamespaceUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StorageKVNamespaceUpdateResponseEnvelopeMessages]
type storageKVNamespaceUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceUpdateResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceUpdateResponseEnvelopeSuccessTrue StorageKVNamespaceUpdateResponseEnvelopeSuccess = true
)

type StorageKVNamespaceListParams struct {
	// Direction to order namespaces.
	Direction param.Field[StorageKVNamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[StorageKVNamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [StorageKVNamespaceListParams]'s query parameters as
// `url.Values`.
func (r StorageKVNamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order namespaces.
type StorageKVNamespaceListParamsDirection string

const (
	StorageKVNamespaceListParamsDirectionAsc  StorageKVNamespaceListParamsDirection = "asc"
	StorageKVNamespaceListParamsDirectionDesc StorageKVNamespaceListParamsDirection = "desc"
)

// Field to order results by.
type StorageKVNamespaceListParamsOrder string

const (
	StorageKVNamespaceListParamsOrderID    StorageKVNamespaceListParamsOrder = "id"
	StorageKVNamespaceListParamsOrderTitle StorageKVNamespaceListParamsOrder = "title"
)

type StorageKVNamespaceDeleteResponseEnvelope struct {
	Errors   []StorageKVNamespaceDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKVNamespaceDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKVNamespaceDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKVNamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKVNamespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// storageKVNamespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageKVNamespaceDeleteResponseEnvelope]
type storageKVNamespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    storageKVNamespaceDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKVNamespaceDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StorageKVNamespaceDeleteResponseEnvelopeErrors]
type storageKVNamespaceDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKVNamespaceDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    storageKVNamespaceDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKVNamespaceDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StorageKVNamespaceDeleteResponseEnvelopeMessages]
type storageKVNamespaceDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKVNamespaceDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKVNamespaceDeleteResponseEnvelopeSuccess bool

const (
	StorageKVNamespaceDeleteResponseEnvelopeSuccessTrue StorageKVNamespaceDeleteResponseEnvelopeSuccess = true
)
