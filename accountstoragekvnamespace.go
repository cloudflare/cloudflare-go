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

// AccountStorageKvNamespaceService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountStorageKvNamespaceService] method instead.
type AccountStorageKvNamespaceService struct {
	Options  []option.RequestOption
	Bulks    *AccountStorageKvNamespaceBulkService
	Keys     *AccountStorageKvNamespaceKeyService
	Metadata *AccountStorageKvNamespaceMetadataService
	Values   *AccountStorageKvNamespaceValueService
}

// NewAccountStorageKvNamespaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceService) {
	r = &AccountStorageKvNamespaceService{}
	r.Options = opts
	r.Bulks = NewAccountStorageKvNamespaceBulkService(opts...)
	r.Keys = NewAccountStorageKvNamespaceKeyService(opts...)
	r.Metadata = NewAccountStorageKvNamespaceMetadataService(opts...)
	r.Values = NewAccountStorageKvNamespaceValueService(opts...)
	return
}

// Modifies a namespace's title.
func (r *AccountStorageKvNamespaceService) Update(ctx context.Context, accountIdentifier string, namespaceIdentifier string, body AccountStorageKvNamespaceUpdateParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the namespaces owned by an account.
func (r *AccountStorageKvNamespaceService) List(ctx context.Context, accountIdentifier string, query AccountStorageKvNamespaceListParams, opts ...option.RequestOption) (res *shared.Page[AccountStorageKvNamespaceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountIdentifier)
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

// Deletes the namespace corresponding to the given ID.
func (r *AccountStorageKvNamespaceService) Delete(ctx context.Context, accountIdentifier string, namespaceIdentifier string, opts ...option.RequestOption) (res *AccountStorageKvNamespaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountIdentifier, namespaceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a namespace under the given title. A `400` is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *AccountStorageKvNamespaceService) WorkersKvNamespaceNewANamespace(ctx context.Context, accountIdentifier string, body AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountStorageKvNamespaceUpdateResponse struct {
	Errors   []AccountStorageKvNamespaceUpdateResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceUpdateResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceUpdateResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceUpdateResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceUpdateResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceUpdateResponse]
type accountStorageKvNamespaceUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceUpdateResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountStorageKvNamespaceUpdateResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceUpdateResponseError]
type accountStorageKvNamespaceUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceUpdateResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountStorageKvNamespaceUpdateResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceUpdateResponseMessage]
type accountStorageKvNamespaceUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountStorageKvNamespaceUpdateResponseResultUnknown] or
// [shared.UnionString].
type AccountStorageKvNamespaceUpdateResponseResult interface {
	ImplementsAccountStorageKvNamespaceUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceUpdateResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountStorageKvNamespaceUpdateResponseSuccess bool

const (
	AccountStorageKvNamespaceUpdateResponseSuccessTrue AccountStorageKvNamespaceUpdateResponseSuccess = true
)

type AccountStorageKvNamespaceListResponse struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                                      `json:"supports_url_encoding"`
	JSON                accountStorageKvNamespaceListResponseJSON `json:"-"`
}

// accountStorageKvNamespaceListResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceListResponse]
type accountStorageKvNamespaceListResponseJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceDeleteResponse struct {
	Errors   []AccountStorageKvNamespaceDeleteResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceDeleteResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceDeleteResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceDeleteResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceDeleteResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceDeleteResponse]
type accountStorageKvNamespaceDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceDeleteResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountStorageKvNamespaceDeleteResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceDeleteResponseError]
type accountStorageKvNamespaceDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceDeleteResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountStorageKvNamespaceDeleteResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceDeleteResponseMessage]
type accountStorageKvNamespaceDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountStorageKvNamespaceDeleteResponseResultUnknown] or
// [shared.UnionString].
type AccountStorageKvNamespaceDeleteResponseResult interface {
	ImplementsAccountStorageKvNamespaceDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountStorageKvNamespaceDeleteResponseSuccess bool

const (
	AccountStorageKvNamespaceDeleteResponseSuccessTrue AccountStorageKvNamespaceDeleteResponseSuccess = true
)

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse struct {
	Errors   []AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON contains
// the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                                                                       `json:"supports_url_encoding"`
	JSON                accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResultJSON `json:"-"`
}

// accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResultJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult]
type accountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResultJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccess bool

const (
	AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccessTrue AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceResponseSuccess = true
)

type AccountStorageKvNamespaceUpdateParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r AccountStorageKvNamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStorageKvNamespaceListParams struct {
	// Direction to order namespaces.
	Direction param.Field[AccountStorageKvNamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[AccountStorageKvNamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountStorageKvNamespaceListParams]'s query parameters as
// `url.Values`.
func (r AccountStorageKvNamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order namespaces.
type AccountStorageKvNamespaceListParamsDirection string

const (
	AccountStorageKvNamespaceListParamsDirectionAsc  AccountStorageKvNamespaceListParamsDirection = "asc"
	AccountStorageKvNamespaceListParamsDirectionDesc AccountStorageKvNamespaceListParamsDirection = "desc"
)

// Field to order results by.
type AccountStorageKvNamespaceListParamsOrder string

const (
	AccountStorageKvNamespaceListParamsOrderID    AccountStorageKvNamespaceListParamsOrder = "id"
	AccountStorageKvNamespaceListParamsOrderTitle AccountStorageKvNamespaceListParamsOrder = "title"
)

type AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceParams struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r AccountStorageKvNamespaceWorkersKvNamespaceNewANamespaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
