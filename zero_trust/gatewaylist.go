// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// GatewayListService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayListService] method
// instead.
type GatewayListService struct {
	Options []option.RequestOption
	Items   *GatewayListItemService
}

// NewGatewayListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayListService(opts ...option.RequestOption) (r *GatewayListService) {
	r = &GatewayListService{}
	r.Options = opts
	r.Items = NewGatewayListItemService(opts...)
	return
}

// Creates a new Zero Trust list.
func (r *GatewayListService) New(ctx context.Context, params GatewayListNewParams, opts ...option.RequestOption) (res *GatewayListNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust list.
func (r *GatewayListService) Update(ctx context.Context, listID string, params GatewayListUpdateParams, opts ...option.RequestOption) (res *GatewayListUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust lists for an account.
func (r *GatewayListService) List(ctx context.Context, query GatewayListListParams, opts ...option.RequestOption) (res *[]GatewayListListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Zero Trust list.
func (r *GatewayListService) Delete(ctx context.Context, listID string, body GatewayListDeleteParams, opts ...option.RequestOption) (res *GatewayListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", body.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Appends or removes an item from a configured Zero Trust list.
func (r *GatewayListService) Edit(ctx context.Context, listID string, params GatewayListEditParams, opts ...option.RequestOption) (res *GatewayListEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust list.
func (r *GatewayListService) Get(ctx context.Context, listID string, query GatewayListGetParams, opts ...option.RequestOption) (res *GatewayListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayListNewResponse struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []GatewayListNewResponseItem `json:"items"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      GatewayListNewResponseType `json:"type"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      gatewayListNewResponseJSON `json:"-"`
}

// gatewayListNewResponseJSON contains the JSON metadata for the struct
// [GatewayListNewResponse]
type gatewayListNewResponseJSON struct {
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

func (r *GatewayListNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListNewResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayListNewResponseItem struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                         `json:"value"`
	JSON  gatewayListNewResponseItemJSON `json:"-"`
}

// gatewayListNewResponseItemJSON contains the JSON metadata for the struct
// [GatewayListNewResponseItem]
type gatewayListNewResponseItemJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListNewResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListNewResponseItemJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type GatewayListNewResponseType string

const (
	GatewayListNewResponseTypeSerial GatewayListNewResponseType = "SERIAL"
	GatewayListNewResponseTypeURL    GatewayListNewResponseType = "URL"
	GatewayListNewResponseTypeDomain GatewayListNewResponseType = "DOMAIN"
	GatewayListNewResponseTypeEmail  GatewayListNewResponseType = "EMAIL"
	GatewayListNewResponseTypeIP     GatewayListNewResponseType = "IP"
)

type GatewayListUpdateResponse struct {
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
	Type      GatewayListUpdateResponseType `json:"type"`
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      gatewayListUpdateResponseJSON `json:"-"`
}

// gatewayListUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayListUpdateResponse]
type gatewayListUpdateResponseJSON struct {
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

func (r *GatewayListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type GatewayListUpdateResponseType string

const (
	GatewayListUpdateResponseTypeSerial GatewayListUpdateResponseType = "SERIAL"
	GatewayListUpdateResponseTypeURL    GatewayListUpdateResponseType = "URL"
	GatewayListUpdateResponseTypeDomain GatewayListUpdateResponseType = "DOMAIN"
	GatewayListUpdateResponseTypeEmail  GatewayListUpdateResponseType = "EMAIL"
	GatewayListUpdateResponseTypeIP     GatewayListUpdateResponseType = "IP"
)

type GatewayListListResponse struct {
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
	Type      GatewayListListResponseType `json:"type"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      gatewayListListResponseJSON `json:"-"`
}

// gatewayListListResponseJSON contains the JSON metadata for the struct
// [GatewayListListResponse]
type gatewayListListResponseJSON struct {
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

func (r *GatewayListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListListResponseJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type GatewayListListResponseType string

const (
	GatewayListListResponseTypeSerial GatewayListListResponseType = "SERIAL"
	GatewayListListResponseTypeURL    GatewayListListResponseType = "URL"
	GatewayListListResponseTypeDomain GatewayListListResponseType = "DOMAIN"
	GatewayListListResponseTypeEmail  GatewayListListResponseType = "EMAIL"
	GatewayListListResponseTypeIP     GatewayListListResponseType = "IP"
)

// Union satisfied by [zero_trust.GatewayListDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayListDeleteResponse interface {
	ImplementsZeroTrustGatewayListDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayListDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayListEditResponse struct {
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
	Type      GatewayListEditResponseType `json:"type"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      gatewayListEditResponseJSON `json:"-"`
}

// gatewayListEditResponseJSON contains the JSON metadata for the struct
// [GatewayListEditResponse]
type gatewayListEditResponseJSON struct {
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

func (r *GatewayListEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListEditResponseJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type GatewayListEditResponseType string

const (
	GatewayListEditResponseTypeSerial GatewayListEditResponseType = "SERIAL"
	GatewayListEditResponseTypeURL    GatewayListEditResponseType = "URL"
	GatewayListEditResponseTypeDomain GatewayListEditResponseType = "DOMAIN"
	GatewayListEditResponseTypeEmail  GatewayListEditResponseType = "EMAIL"
	GatewayListEditResponseTypeIP     GatewayListEditResponseType = "IP"
)

type GatewayListGetResponse struct {
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
	Type      GatewayListGetResponseType `json:"type"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      gatewayListGetResponseJSON `json:"-"`
}

// gatewayListGetResponseJSON contains the JSON metadata for the struct
// [GatewayListGetResponse]
type gatewayListGetResponseJSON struct {
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

func (r *GatewayListGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListGetResponseJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type GatewayListGetResponseType string

const (
	GatewayListGetResponseTypeSerial GatewayListGetResponseType = "SERIAL"
	GatewayListGetResponseTypeURL    GatewayListGetResponseType = "URL"
	GatewayListGetResponseTypeDomain GatewayListGetResponseType = "DOMAIN"
	GatewayListGetResponseTypeEmail  GatewayListGetResponseType = "EMAIL"
	GatewayListGetResponseTypeIP     GatewayListGetResponseType = "IP"
)

type GatewayListNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[GatewayListNewParamsType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// The items in the list.
	Items param.Field[[]GatewayListNewParamsItem] `json:"items"`
}

func (r GatewayListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of list.
type GatewayListNewParamsType string

const (
	GatewayListNewParamsTypeSerial GatewayListNewParamsType = "SERIAL"
	GatewayListNewParamsTypeURL    GatewayListNewParamsType = "URL"
	GatewayListNewParamsTypeDomain GatewayListNewParamsType = "DOMAIN"
	GatewayListNewParamsTypeEmail  GatewayListNewParamsType = "EMAIL"
	GatewayListNewParamsTypeIP     GatewayListNewParamsType = "IP"
)

type GatewayListNewParamsItem struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r GatewayListNewParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListNewResponseEnvelope struct {
	Errors   []GatewayListNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListNewResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayListNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListNewResponseEnvelope]
type gatewayListNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayListNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    gatewayListNewResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListNewResponseEnvelopeErrors]
type gatewayListNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayListNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    gatewayListNewResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayListNewResponseEnvelopeMessages]
type gatewayListNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListNewResponseEnvelopeSuccess bool

const (
	GatewayListNewResponseEnvelopeSuccessTrue GatewayListNewResponseEnvelopeSuccess = true
)

type GatewayListUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
}

func (r GatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListUpdateResponseEnvelope struct {
	Errors   []GatewayListUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayListUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListUpdateResponseEnvelope]
type gatewayListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayListUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayListUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListUpdateResponseEnvelopeErrors]
type gatewayListUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayListUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayListUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayListUpdateResponseEnvelopeMessages]
type gatewayListUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListUpdateResponseEnvelopeSuccess bool

const (
	GatewayListUpdateResponseEnvelopeSuccessTrue GatewayListUpdateResponseEnvelopeSuccess = true
)

type GatewayListListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayListListResponseEnvelope struct {
	Errors   []GatewayListListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListListResponseEnvelopeMessages `json:"messages,required"`
	Result   []GatewayListListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayListListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayListListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayListListResponseEnvelopeJSON       `json:"-"`
}

// gatewayListListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListListResponseEnvelope]
type gatewayListListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayListListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    gatewayListListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListListResponseEnvelopeErrors]
type gatewayListListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayListListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayListListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayListListResponseEnvelopeMessages]
type gatewayListListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListListResponseEnvelopeSuccess bool

const (
	GatewayListListResponseEnvelopeSuccessTrue GatewayListListResponseEnvelopeSuccess = true
)

type GatewayListListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       gatewayListListResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayListListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [GatewayListListResponseEnvelopeResultInfo]
type gatewayListListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type GatewayListDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayListDeleteResponseEnvelope struct {
	Errors   []GatewayListDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayListDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListDeleteResponseEnvelope]
type gatewayListDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayListDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayListDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListDeleteResponseEnvelopeErrors]
type gatewayListDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayListDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayListDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayListDeleteResponseEnvelopeMessages]
type gatewayListDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListDeleteResponseEnvelopeSuccess bool

const (
	GatewayListDeleteResponseEnvelopeSuccessTrue GatewayListDeleteResponseEnvelopeSuccess = true
)

type GatewayListEditParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The items in the list.
	Append param.Field[[]GatewayListEditParamsAppend] `json:"append"`
	// A list of the item values you want to remove.
	Remove param.Field[[]string] `json:"remove"`
}

func (r GatewayListEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListEditParamsAppend struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r GatewayListEditParamsAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListEditResponseEnvelope struct {
	Errors   []GatewayListEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListEditResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListEditResponseEnvelopeJSON    `json:"-"`
}

// gatewayListEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListEditResponseEnvelope]
type gatewayListEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayListEditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    gatewayListEditResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListEditResponseEnvelopeErrors]
type gatewayListEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayListEditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayListEditResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayListEditResponseEnvelopeMessages]
type gatewayListEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListEditResponseEnvelopeSuccess bool

const (
	GatewayListEditResponseEnvelopeSuccessTrue GatewayListEditResponseEnvelopeSuccess = true
)

type GatewayListGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayListGetResponseEnvelope struct {
	Errors   []GatewayListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayListGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListGetResponseEnvelope]
type gatewayListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayListGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    gatewayListGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListGetResponseEnvelopeErrors]
type gatewayListGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayListGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    gatewayListGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayListGetResponseEnvelopeMessages]
type gatewayListGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListGetResponseEnvelopeSuccess bool

const (
	GatewayListGetResponseEnvelopeSuccessTrue GatewayListGetResponseEnvelopeSuccess = true
)
