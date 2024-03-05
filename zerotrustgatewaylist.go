// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustGatewayListService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustGatewayListService]
// method instead.
type ZeroTrustGatewayListService struct {
	Options []option.RequestOption
	Items   *ZeroTrustGatewayListItemService
}

// NewZeroTrustGatewayListService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayListService(opts ...option.RequestOption) (r *ZeroTrustGatewayListService) {
	r = &ZeroTrustGatewayListService{}
	r.Options = opts
	r.Items = NewZeroTrustGatewayListItemService(opts...)
	return
}

// Creates a new Zero Trust list.
func (r *ZeroTrustGatewayListService) New(ctx context.Context, params ZeroTrustGatewayListNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayListNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust list.
func (r *ZeroTrustGatewayListService) Update(ctx context.Context, listID string, params ZeroTrustGatewayListUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLists, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust lists for an account.
func (r *ZeroTrustGatewayListService) List(ctx context.Context, query ZeroTrustGatewayListListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayLists, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Zero Trust list.
func (r *ZeroTrustGatewayListService) Delete(ctx context.Context, listID string, body ZeroTrustGatewayListDeleteParams, opts ...option.RequestOption) (res *ZeroTrustGatewayListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", body.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Appends or removes an item from a configured Zero Trust list.
func (r *ZeroTrustGatewayListService) Edit(ctx context.Context, listID string, params ZeroTrustGatewayListEditParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLists, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust list.
func (r *ZeroTrustGatewayListService) Get(ctx context.Context, listID string, query ZeroTrustGatewayListGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLists, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayLists struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      ZeroTrustGatewayListsType `json:"type"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayListsJSON `json:"-"`
}

// zeroTrustGatewayListsJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayLists]
type zeroTrustGatewayListsJSON struct {
	ID          apijson.Field
	Count       apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type ZeroTrustGatewayListsType string

const (
	ZeroTrustGatewayListsTypeSerial ZeroTrustGatewayListsType = "SERIAL"
	ZeroTrustGatewayListsTypeURL    ZeroTrustGatewayListsType = "URL"
	ZeroTrustGatewayListsTypeDomain ZeroTrustGatewayListsType = "DOMAIN"
	ZeroTrustGatewayListsTypeEmail  ZeroTrustGatewayListsType = "EMAIL"
	ZeroTrustGatewayListsTypeIP     ZeroTrustGatewayListsType = "IP"
)

type ZeroTrustGatewayListNewResponse struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []ZeroTrustGatewayListNewResponseItem `json:"items"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      ZeroTrustGatewayListNewResponseType `json:"type"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayListNewResponseJSON `json:"-"`
}

// zeroTrustGatewayListNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayListNewResponse]
type zeroTrustGatewayListNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Items       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListNewResponseItem struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                                  `json:"value"`
	JSON  zeroTrustGatewayListNewResponseItemJSON `json:"-"`
}

// zeroTrustGatewayListNewResponseItemJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayListNewResponseItem]
type zeroTrustGatewayListNewResponseItemJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListNewResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type ZeroTrustGatewayListNewResponseType string

const (
	ZeroTrustGatewayListNewResponseTypeSerial ZeroTrustGatewayListNewResponseType = "SERIAL"
	ZeroTrustGatewayListNewResponseTypeURL    ZeroTrustGatewayListNewResponseType = "URL"
	ZeroTrustGatewayListNewResponseTypeDomain ZeroTrustGatewayListNewResponseType = "DOMAIN"
	ZeroTrustGatewayListNewResponseTypeEmail  ZeroTrustGatewayListNewResponseType = "EMAIL"
	ZeroTrustGatewayListNewResponseTypeIP     ZeroTrustGatewayListNewResponseType = "IP"
)

// Union satisfied by [ZeroTrustGatewayListDeleteResponseUnknown] or
// [shared.UnionString].
type ZeroTrustGatewayListDeleteResponse interface {
	ImplementsZeroTrustGatewayListDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGatewayListDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustGatewayListNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[ZeroTrustGatewayListNewParamsType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// The items in the list.
	Items param.Field[[]ZeroTrustGatewayListNewParamsItem] `json:"items"`
}

func (r ZeroTrustGatewayListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of list.
type ZeroTrustGatewayListNewParamsType string

const (
	ZeroTrustGatewayListNewParamsTypeSerial ZeroTrustGatewayListNewParamsType = "SERIAL"
	ZeroTrustGatewayListNewParamsTypeURL    ZeroTrustGatewayListNewParamsType = "URL"
	ZeroTrustGatewayListNewParamsTypeDomain ZeroTrustGatewayListNewParamsType = "DOMAIN"
	ZeroTrustGatewayListNewParamsTypeEmail  ZeroTrustGatewayListNewParamsType = "EMAIL"
	ZeroTrustGatewayListNewParamsTypeIP     ZeroTrustGatewayListNewParamsType = "IP"
)

type ZeroTrustGatewayListNewParamsItem struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r ZeroTrustGatewayListNewParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayListNewResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayListNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayListNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayListNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayListNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayListNewResponseEnvelope]
type zeroTrustGatewayListNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustGatewayListNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayListNewResponseEnvelopeErrors]
type zeroTrustGatewayListNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustGatewayListNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListNewResponseEnvelopeMessages]
type zeroTrustGatewayListNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayListNewResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListNewResponseEnvelopeSuccessTrue ZeroTrustGatewayListNewResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayListUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
}

func (r ZeroTrustGatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayListUpdateResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLists                                `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayListUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayListUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayListUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayListUpdateResponseEnvelope]
type zeroTrustGatewayListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayListUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListUpdateResponseEnvelopeErrors]
type zeroTrustGatewayListUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayListUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayListUpdateResponseEnvelopeMessages]
type zeroTrustGatewayListUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayListUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListUpdateResponseEnvelopeSuccessTrue ZeroTrustGatewayListUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayListListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayListListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayLists                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustGatewayListListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustGatewayListListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustGatewayListListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustGatewayListListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayListListResponseEnvelope]
type zeroTrustGatewayListListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustGatewayListListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListListResponseEnvelopeErrors]
type zeroTrustGatewayListListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayListListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListListResponseEnvelopeMessages]
type zeroTrustGatewayListListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayListListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListListResponseEnvelopeSuccessTrue ZeroTrustGatewayListListResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayListListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       zeroTrustGatewayListListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustGatewayListListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayListListResponseEnvelopeResultInfo]
type zeroTrustGatewayListListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayListDeleteResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayListDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayListDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayListDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayListDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayListDeleteResponseEnvelope]
type zeroTrustGatewayListDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayListDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListDeleteResponseEnvelopeErrors]
type zeroTrustGatewayListDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayListDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayListDeleteResponseEnvelopeMessages]
type zeroTrustGatewayListDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayListDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListDeleteResponseEnvelopeSuccessTrue ZeroTrustGatewayListDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayListEditParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The items in the list.
	Append param.Field[[]ZeroTrustGatewayListEditParamsAppend] `json:"append"`
	// A list of the item values you want to remove.
	Remove param.Field[[]string] `json:"remove"`
}

func (r ZeroTrustGatewayListEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayListEditParamsAppend struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r ZeroTrustGatewayListEditParamsAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayListEditResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLists                              `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayListEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayListEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayListEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayListEditResponseEnvelope]
type zeroTrustGatewayListEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustGatewayListEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListEditResponseEnvelopeErrors]
type zeroTrustGatewayListEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayListEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListEditResponseEnvelopeMessages]
type zeroTrustGatewayListEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayListEditResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListEditResponseEnvelopeSuccessTrue ZeroTrustGatewayListEditResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayListGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayListGetResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLists                             `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayListGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayListGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayListGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayListGetResponseEnvelope]
type zeroTrustGatewayListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustGatewayListGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayListGetResponseEnvelopeErrors]
type zeroTrustGatewayListGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayListGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustGatewayListGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayListGetResponseEnvelopeMessages]
type zeroTrustGatewayListGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayListGetResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListGetResponseEnvelopeSuccessTrue ZeroTrustGatewayListGetResponseEnvelopeSuccess = true
)
