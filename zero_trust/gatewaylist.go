// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	path := fmt.Sprintf("accounts/%s/gateway/lists", params.AccountID)
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
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust lists for an account.
func (r *GatewayListService) List(ctx context.Context, query GatewayListListParams, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayListListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/gateway/lists", query.AccountID)
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

// Fetches all Zero Trust lists for an account.
func (r *GatewayListService) ListAutoPaging(ctx context.Context, query GatewayListListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayListListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Zero Trust list.
func (r *GatewayListService) Delete(ctx context.Context, listID string, params GatewayListDeleteParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
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
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", params.AccountID, listID)
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
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Lists struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string    `json:"value"`
	JSON  listsJSON `json:"-"`
}

// listsJSON contains the JSON metadata for the struct [Lists]
type listsJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Lists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listsJSON) RawJSON() string {
	return r.raw
}

type ListsParam struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r ListsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListNewResponse struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []Lists `json:"items"`
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

// The type of list.
type GatewayListNewResponseType string

const (
	GatewayListNewResponseTypeSerial GatewayListNewResponseType = "SERIAL"
	GatewayListNewResponseTypeURL    GatewayListNewResponseType = "URL"
	GatewayListNewResponseTypeDomain GatewayListNewResponseType = "DOMAIN"
	GatewayListNewResponseTypeEmail  GatewayListNewResponseType = "EMAIL"
	GatewayListNewResponseTypeIP     GatewayListNewResponseType = "IP"
)

func (r GatewayListNewResponseType) IsKnown() bool {
	switch r {
	case GatewayListNewResponseTypeSerial, GatewayListNewResponseTypeURL, GatewayListNewResponseTypeDomain, GatewayListNewResponseTypeEmail, GatewayListNewResponseTypeIP:
		return true
	}
	return false
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

func (r GatewayListUpdateResponseType) IsKnown() bool {
	switch r {
	case GatewayListUpdateResponseTypeSerial, GatewayListUpdateResponseTypeURL, GatewayListUpdateResponseTypeDomain, GatewayListUpdateResponseTypeEmail, GatewayListUpdateResponseTypeIP:
		return true
	}
	return false
}

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

func (r GatewayListListResponseType) IsKnown() bool {
	switch r {
	case GatewayListListResponseTypeSerial, GatewayListListResponseTypeURL, GatewayListListResponseTypeDomain, GatewayListListResponseTypeEmail, GatewayListListResponseTypeIP:
		return true
	}
	return false
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

func (r GatewayListEditResponseType) IsKnown() bool {
	switch r {
	case GatewayListEditResponseTypeSerial, GatewayListEditResponseTypeURL, GatewayListEditResponseTypeDomain, GatewayListEditResponseTypeEmail, GatewayListEditResponseTypeIP:
		return true
	}
	return false
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

func (r GatewayListGetResponseType) IsKnown() bool {
	switch r {
	case GatewayListGetResponseTypeSerial, GatewayListGetResponseTypeURL, GatewayListGetResponseTypeDomain, GatewayListGetResponseTypeEmail, GatewayListGetResponseTypeIP:
		return true
	}
	return false
}

type GatewayListNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[GatewayListNewParamsType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// The items in the list.
	Items param.Field[[]ListsParam] `json:"items"`
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

func (r GatewayListNewParamsType) IsKnown() bool {
	switch r {
	case GatewayListNewParamsTypeSerial, GatewayListNewParamsTypeURL, GatewayListNewParamsTypeDomain, GatewayListNewParamsTypeEmail, GatewayListNewParamsTypeIP:
		return true
	}
	return false
}

type GatewayListNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   GatewayListNewResponse `json:"result,required"`
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

// Whether the API call was successful
type GatewayListNewResponseEnvelopeSuccess bool

const (
	GatewayListNewResponseEnvelopeSuccessTrue GatewayListNewResponseEnvelopeSuccess = true
)

func (r GatewayListNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
}

func (r GatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   GatewayListUpdateResponse `json:"result,required"`
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

// Whether the API call was successful
type GatewayListUpdateResponseEnvelopeSuccess bool

const (
	GatewayListUpdateResponseEnvelopeSuccessTrue GatewayListUpdateResponseEnvelopeSuccess = true
)

func (r GatewayListUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayListDeleteParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r GatewayListDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GatewayListDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                        `json:"errors,required"`
	Messages []shared.ResponseInfo                                        `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
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

// Whether the API call was successful
type GatewayListDeleteResponseEnvelopeSuccess bool

const (
	GatewayListDeleteResponseEnvelopeSuccessTrue GatewayListDeleteResponseEnvelopeSuccess = true
)

func (r GatewayListDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The items in the list.
	Append param.Field[[]ListsParam] `json:"append"`
	// A list of the item values you want to remove.
	Remove param.Field[[]string] `json:"remove"`
}

func (r GatewayListEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   GatewayListEditResponse `json:"result,required"`
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

// Whether the API call was successful
type GatewayListEditResponseEnvelopeSuccess bool

const (
	GatewayListEditResponseEnvelopeSuccessTrue GatewayListEditResponseEnvelopeSuccess = true
)

func (r GatewayListEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayListGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   GatewayListGetResponse `json:"result,required"`
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

// Whether the API call was successful
type GatewayListGetResponseEnvelopeSuccess bool

const (
	GatewayListGetResponseEnvelopeSuccessTrue GatewayListGetResponseEnvelopeSuccess = true
)

func (r GatewayListGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
