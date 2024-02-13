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

// Updates a configured Zero Trust list.
func (r *GatewayListService) Update(ctx context.Context, accountID interface{}, listID string, body GatewayListUpdateParams, opts ...option.RequestOption) (res *GatewayListUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// Creates a new Zero Trust list.
func (r *GatewayListService) ZeroTrustListsNewZeroTrustList(ctx context.Context, accountID interface{}, body GatewayListZeroTrustListsNewZeroTrustListParams, opts ...option.RequestOption) (res *GatewayListZeroTrustListsNewZeroTrustListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListZeroTrustListsNewZeroTrustListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust lists for an account.
func (r *GatewayListService) ZeroTrustListsListZeroTrustLists(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]GatewayListZeroTrustListsListZeroTrustListsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListZeroTrustListsListZeroTrustListsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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

// The type of list.
type GatewayListUpdateResponseType string

const (
	GatewayListUpdateResponseTypeSerial GatewayListUpdateResponseType = "SERIAL"
	GatewayListUpdateResponseTypeURL    GatewayListUpdateResponseType = "URL"
	GatewayListUpdateResponseTypeDomain GatewayListUpdateResponseType = "DOMAIN"
	GatewayListUpdateResponseTypeEmail  GatewayListUpdateResponseType = "EMAIL"
	GatewayListUpdateResponseTypeIP     GatewayListUpdateResponseType = "IP"
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

type GatewayListZeroTrustListsNewZeroTrustListResponse struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []GatewayListZeroTrustListsNewZeroTrustListResponseItem `json:"items"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      GatewayListZeroTrustListsNewZeroTrustListResponseType `json:"type"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      gatewayListZeroTrustListsNewZeroTrustListResponseJSON `json:"-"`
}

// gatewayListZeroTrustListsNewZeroTrustListResponseJSON contains the JSON metadata
// for the struct [GatewayListZeroTrustListsNewZeroTrustListResponse]
type gatewayListZeroTrustListsNewZeroTrustListResponseJSON struct {
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

func (r *GatewayListZeroTrustListsNewZeroTrustListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListZeroTrustListsNewZeroTrustListResponseItem struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                                                    `json:"value"`
	JSON  gatewayListZeroTrustListsNewZeroTrustListResponseItemJSON `json:"-"`
}

// gatewayListZeroTrustListsNewZeroTrustListResponseItemJSON contains the JSON
// metadata for the struct [GatewayListZeroTrustListsNewZeroTrustListResponseItem]
type gatewayListZeroTrustListsNewZeroTrustListResponseItemJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsNewZeroTrustListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type GatewayListZeroTrustListsNewZeroTrustListResponseType string

const (
	GatewayListZeroTrustListsNewZeroTrustListResponseTypeSerial GatewayListZeroTrustListsNewZeroTrustListResponseType = "SERIAL"
	GatewayListZeroTrustListsNewZeroTrustListResponseTypeURL    GatewayListZeroTrustListsNewZeroTrustListResponseType = "URL"
	GatewayListZeroTrustListsNewZeroTrustListResponseTypeDomain GatewayListZeroTrustListsNewZeroTrustListResponseType = "DOMAIN"
	GatewayListZeroTrustListsNewZeroTrustListResponseTypeEmail  GatewayListZeroTrustListsNewZeroTrustListResponseType = "EMAIL"
	GatewayListZeroTrustListsNewZeroTrustListResponseTypeIP     GatewayListZeroTrustListsNewZeroTrustListResponseType = "IP"
)

type GatewayListZeroTrustListsListZeroTrustListsResponse struct {
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
	Type      GatewayListZeroTrustListsListZeroTrustListsResponseType `json:"type"`
	UpdatedAt time.Time                                               `json:"updated_at" format:"date-time"`
	JSON      gatewayListZeroTrustListsListZeroTrustListsResponseJSON `json:"-"`
}

// gatewayListZeroTrustListsListZeroTrustListsResponseJSON contains the JSON
// metadata for the struct [GatewayListZeroTrustListsListZeroTrustListsResponse]
type gatewayListZeroTrustListsListZeroTrustListsResponseJSON struct {
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

func (r *GatewayListZeroTrustListsListZeroTrustListsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type GatewayListZeroTrustListsListZeroTrustListsResponseType string

const (
	GatewayListZeroTrustListsListZeroTrustListsResponseTypeSerial GatewayListZeroTrustListsListZeroTrustListsResponseType = "SERIAL"
	GatewayListZeroTrustListsListZeroTrustListsResponseTypeURL    GatewayListZeroTrustListsListZeroTrustListsResponseType = "URL"
	GatewayListZeroTrustListsListZeroTrustListsResponseTypeDomain GatewayListZeroTrustListsListZeroTrustListsResponseType = "DOMAIN"
	GatewayListZeroTrustListsListZeroTrustListsResponseTypeEmail  GatewayListZeroTrustListsListZeroTrustListsResponseType = "EMAIL"
	GatewayListZeroTrustListsListZeroTrustListsResponseTypeIP     GatewayListZeroTrustListsListZeroTrustListsResponseType = "IP"
)

type GatewayListUpdateParams struct {
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

// Whether the API call was successful
type GatewayListUpdateResponseEnvelopeSuccess bool

const (
	GatewayListUpdateResponseEnvelopeSuccessTrue GatewayListUpdateResponseEnvelopeSuccess = true
)

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

type GatewayListZeroTrustListsNewZeroTrustListParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[GatewayListZeroTrustListsNewZeroTrustListParamsType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// The items in the list.
	Items param.Field[[]GatewayListZeroTrustListsNewZeroTrustListParamsItem] `json:"items"`
}

func (r GatewayListZeroTrustListsNewZeroTrustListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of list.
type GatewayListZeroTrustListsNewZeroTrustListParamsType string

const (
	GatewayListZeroTrustListsNewZeroTrustListParamsTypeSerial GatewayListZeroTrustListsNewZeroTrustListParamsType = "SERIAL"
	GatewayListZeroTrustListsNewZeroTrustListParamsTypeURL    GatewayListZeroTrustListsNewZeroTrustListParamsType = "URL"
	GatewayListZeroTrustListsNewZeroTrustListParamsTypeDomain GatewayListZeroTrustListsNewZeroTrustListParamsType = "DOMAIN"
	GatewayListZeroTrustListsNewZeroTrustListParamsTypeEmail  GatewayListZeroTrustListsNewZeroTrustListParamsType = "EMAIL"
	GatewayListZeroTrustListsNewZeroTrustListParamsTypeIP     GatewayListZeroTrustListsNewZeroTrustListParamsType = "IP"
)

type GatewayListZeroTrustListsNewZeroTrustListParamsItem struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r GatewayListZeroTrustListsNewZeroTrustListParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListZeroTrustListsNewZeroTrustListResponseEnvelope struct {
	Errors   []GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListZeroTrustListsNewZeroTrustListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeJSON    `json:"-"`
}

// gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [GatewayListZeroTrustListsNewZeroTrustListResponseEnvelope]
type gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsNewZeroTrustListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeErrors]
type gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeMessages]
type gatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeSuccess bool

const (
	GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeSuccessTrue GatewayListZeroTrustListsNewZeroTrustListResponseEnvelopeSuccess = true
)

type GatewayListZeroTrustListsListZeroTrustListsResponseEnvelope struct {
	Errors   []GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeMessages `json:"messages,required"`
	Result   []GatewayListZeroTrustListsListZeroTrustListsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeJSON       `json:"-"`
}

// gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [GatewayListZeroTrustListsListZeroTrustListsResponseEnvelope]
type gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsListZeroTrustListsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeErrors]
type gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeMessages]
type gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeSuccess bool

const (
	GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeSuccessTrue GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeSuccess = true
)

type GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                   `json:"total_count"`
	JSON       gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeResultInfo]
type gatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListZeroTrustListsListZeroTrustListsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
