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
func (r *GatewayListService) New(ctx context.Context, accountID interface{}, body GatewayListNewParams, opts ...option.RequestOption) (res *GatewayListNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust lists for an account.
func (r *GatewayListService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]GatewayListListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Zero Trust list.
func (r *GatewayListService) Delete(ctx context.Context, accountID interface{}, listID string, opts ...option.RequestOption) (res *GatewayListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust list.
func (r *GatewayListService) Get(ctx context.Context, accountID interface{}, listID string, opts ...option.RequestOption) (res *GatewayListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust list.
func (r *GatewayListService) Replace(ctx context.Context, accountID interface{}, listID string, body GatewayListReplaceParams, opts ...option.RequestOption) (res *GatewayListReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// The type of list.
type GatewayListNewResponseType string

const (
	GatewayListNewResponseTypeSerial GatewayListNewResponseType = "SERIAL"
	GatewayListNewResponseTypeURL    GatewayListNewResponseType = "URL"
	GatewayListNewResponseTypeDomain GatewayListNewResponseType = "DOMAIN"
	GatewayListNewResponseTypeEmail  GatewayListNewResponseType = "EMAIL"
	GatewayListNewResponseTypeIP     GatewayListNewResponseType = "IP"
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

// The type of list.
type GatewayListListResponseType string

const (
	GatewayListListResponseTypeSerial GatewayListListResponseType = "SERIAL"
	GatewayListListResponseTypeURL    GatewayListListResponseType = "URL"
	GatewayListListResponseTypeDomain GatewayListListResponseType = "DOMAIN"
	GatewayListListResponseTypeEmail  GatewayListListResponseType = "EMAIL"
	GatewayListListResponseTypeIP     GatewayListListResponseType = "IP"
)

// Union satisfied by [GatewayListDeleteResponseUnknown] or [shared.UnionString].
type GatewayListDeleteResponse interface {
	ImplementsGatewayListDeleteResponse()
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

// The type of list.
type GatewayListGetResponseType string

const (
	GatewayListGetResponseTypeSerial GatewayListGetResponseType = "SERIAL"
	GatewayListGetResponseTypeURL    GatewayListGetResponseType = "URL"
	GatewayListGetResponseTypeDomain GatewayListGetResponseType = "DOMAIN"
	GatewayListGetResponseTypeEmail  GatewayListGetResponseType = "EMAIL"
	GatewayListGetResponseTypeIP     GatewayListGetResponseType = "IP"
)

type GatewayListReplaceResponse struct {
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
	Type      GatewayListReplaceResponseType `json:"type"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      gatewayListReplaceResponseJSON `json:"-"`
}

// gatewayListReplaceResponseJSON contains the JSON metadata for the struct
// [GatewayListReplaceResponse]
type gatewayListReplaceResponseJSON struct {
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

func (r *GatewayListReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type GatewayListReplaceResponseType string

const (
	GatewayListReplaceResponseTypeSerial GatewayListReplaceResponseType = "SERIAL"
	GatewayListReplaceResponseTypeURL    GatewayListReplaceResponseType = "URL"
	GatewayListReplaceResponseTypeDomain GatewayListReplaceResponseType = "DOMAIN"
	GatewayListReplaceResponseTypeEmail  GatewayListReplaceResponseType = "EMAIL"
	GatewayListReplaceResponseTypeIP     GatewayListReplaceResponseType = "IP"
)

type GatewayListNewParams struct {
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

// Whether the API call was successful
type GatewayListNewResponseEnvelopeSuccess bool

const (
	GatewayListNewResponseEnvelopeSuccessTrue GatewayListNewResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type GatewayListDeleteResponseEnvelopeSuccess bool

const (
	GatewayListDeleteResponseEnvelopeSuccessTrue GatewayListDeleteResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type GatewayListGetResponseEnvelopeSuccess bool

const (
	GatewayListGetResponseEnvelopeSuccessTrue GatewayListGetResponseEnvelopeSuccess = true
)

type GatewayListReplaceParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
}

func (r GatewayListReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListReplaceResponseEnvelope struct {
	Errors   []GatewayListReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListReplaceResponseEnvelopeJSON    `json:"-"`
}

// gatewayListReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListReplaceResponseEnvelope]
type gatewayListReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListReplaceResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    gatewayListReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayListReplaceResponseEnvelopeErrors]
type gatewayListReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListReplaceResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    gatewayListReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayListReplaceResponseEnvelopeMessages]
type gatewayListReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayListReplaceResponseEnvelopeSuccess bool

const (
	GatewayListReplaceResponseEnvelopeSuccessTrue GatewayListReplaceResponseEnvelopeSuccess = true
)
