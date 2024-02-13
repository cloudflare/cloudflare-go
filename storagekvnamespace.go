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

// StorageKvNamespaceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStorageKvNamespaceService] method
// instead.
type StorageKvNamespaceService struct {
	Options  []option.RequestOption
	Bulks    *StorageKvNamespaceBulkService
	Keys     *StorageKvNamespaceKeyService
	Metadata *StorageKvNamespaceMetadataService
	Values   *StorageKvNamespaceValueService
}

// NewStorageKvNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageKvNamespaceService(opts ...option.RequestOption) (r *StorageKvNamespaceService) {
	r = &StorageKvNamespaceService{}
	r.Options = opts
	r.Bulks = NewStorageKvNamespaceBulkService(opts...)
	r.Keys = NewStorageKvNamespaceKeyService(opts...)
	r.Metadata = NewStorageKvNamespaceMetadataService(opts...)
	r.Values = NewStorageKvNamespaceValueService(opts...)
	return
}

// Modifies a namespace's title.
func (r *StorageKvNamespaceService) Update(ctx context.Context, accountID string, namespaceID string, body StorageKvNamespaceUpdateParams, opts ...option.RequestOption) (res *StorageKvNamespaceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the namespaces owned by an account.
func (r *StorageKvNamespaceService) List(ctx context.Context, accountID string, query StorageKvNamespaceListParams, opts ...option.RequestOption) (res *[]StorageKvNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the namespace corresponding to the given ID.
func (r *StorageKvNamespaceService) Delete(ctx context.Context, accountID string, namespaceID string, opts ...option.RequestOption) (res *StorageKvNamespaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a namespace under the given title. A `400` is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *StorageKvNamespaceService) WorkersKvNamespaceNewANamespace(ctx context.Context, accountID string, body StorageKvNamespaceWorkersKvNamespaceNewANamespaceParams, opts ...option.RequestOption) (res *StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StorageKvNamespaceUpdateResponseUnknown] or
// [shared.UnionString].
type StorageKvNamespaceUpdateResponse interface {
	ImplementsStorageKvNamespaceUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKvNamespaceUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKvNamespaceListResponse struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                               `json:"supports_url_encoding"`
	JSON                storageKvNamespaceListResponseJSON `json:"-"`
}

// storageKvNamespaceListResponseJSON contains the JSON metadata for the struct
// [StorageKvNamespaceListResponse]
type storageKvNamespaceListResponseJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StorageKvNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [StorageKvNamespaceDeleteResponseUnknown] or
// [shared.UnionString].
type StorageKvNamespaceDeleteResponse interface {
	ImplementsStorageKvNamespaceDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKvNamespaceDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                                                          `json:"supports_url_encoding"`
	JSON                storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON `json:"-"`
}

// storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON contains the JSON
// metadata for the struct
// [StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse]
type storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceUpdateParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r StorageKvNamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKvNamespaceUpdateResponseEnvelope struct {
	Errors   []StorageKvNamespaceUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceUpdateResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageKvNamespaceUpdateResponseEnvelope]
type storageKvNamespaceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    storageKvNamespaceUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StorageKvNamespaceUpdateResponseEnvelopeErrors]
type storageKvNamespaceUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    storageKvNamespaceUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StorageKvNamespaceUpdateResponseEnvelopeMessages]
type storageKvNamespaceUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceUpdateResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceUpdateResponseEnvelopeSuccessTrue StorageKvNamespaceUpdateResponseEnvelopeSuccess = true
)

type StorageKvNamespaceListParams struct {
	// Direction to order namespaces.
	Direction param.Field[StorageKvNamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[StorageKvNamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [StorageKvNamespaceListParams]'s query parameters as
// `url.Values`.
func (r StorageKvNamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order namespaces.
type StorageKvNamespaceListParamsDirection string

const (
	StorageKvNamespaceListParamsDirectionAsc  StorageKvNamespaceListParamsDirection = "asc"
	StorageKvNamespaceListParamsDirectionDesc StorageKvNamespaceListParamsDirection = "desc"
)

// Field to order results by.
type StorageKvNamespaceListParamsOrder string

const (
	StorageKvNamespaceListParamsOrderID    StorageKvNamespaceListParamsOrder = "id"
	StorageKvNamespaceListParamsOrderTitle StorageKvNamespaceListParamsOrder = "title"
)

type StorageKvNamespaceListResponseEnvelope struct {
	Errors   []StorageKvNamespaceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []StorageKvNamespaceListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    StorageKvNamespaceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo StorageKvNamespaceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       storageKvNamespaceListResponseEnvelopeJSON       `json:"-"`
}

// storageKvNamespaceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageKvNamespaceListResponseEnvelope]
type storageKvNamespaceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    storageKvNamespaceListResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StorageKvNamespaceListResponseEnvelopeErrors]
type storageKvNamespaceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    storageKvNamespaceListResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StorageKvNamespaceListResponseEnvelopeMessages]
type storageKvNamespaceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceListResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceListResponseEnvelopeSuccessTrue StorageKvNamespaceListResponseEnvelopeSuccess = true
)

type StorageKvNamespaceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       storageKvNamespaceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// storageKvNamespaceListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [StorageKvNamespaceListResponseEnvelopeResultInfo]
type storageKvNamespaceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceDeleteResponseEnvelope struct {
	Errors   []StorageKvNamespaceDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageKvNamespaceDeleteResponseEnvelope]
type storageKvNamespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    storageKvNamespaceDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StorageKvNamespaceDeleteResponseEnvelopeErrors]
type storageKvNamespaceDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    storageKvNamespaceDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StorageKvNamespaceDeleteResponseEnvelopeMessages]
type storageKvNamespaceDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceDeleteResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceDeleteResponseEnvelopeSuccessTrue StorageKvNamespaceDeleteResponseEnvelopeSuccess = true
)

type StorageKvNamespaceWorkersKvNamespaceNewANamespaceParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r StorageKvNamespaceWorkersKvNamespaceNewANamespaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelope struct {
	Errors   []StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelope]
type storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeErrors]
type storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeMessages]
type storageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeSuccessTrue StorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseEnvelopeSuccess = true
)
