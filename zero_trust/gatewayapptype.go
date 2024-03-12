// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// GatewayAppTypeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayAppTypeService] method
// instead.
type GatewayAppTypeService struct {
	Options []option.RequestOption
}

// NewGatewayAppTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayAppTypeService(opts ...option.RequestOption) (r *GatewayAppTypeService) {
	r = &GatewayAppTypeService{}
	r.Options = opts
	return
}

// Fetches all application and application type mappings.
func (r *GatewayAppTypeService) List(ctx context.Context, query GatewayAppTypeListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayAppTypes, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayAppTypeListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/app_types", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [zero_trust.ZeroTrustGatewayAppTypesZeroTrustGatewayApplication] or
// [zero_trust.ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType].
type ZeroTrustGatewayAppTypes interface {
	implementsZeroTrustZeroTrustGatewayAppTypes()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGatewayAppTypes)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGatewayAppTypesZeroTrustGatewayApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType{}),
		},
	)
}

type ZeroTrustGatewayAppTypesZeroTrustGatewayApplication struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The name of the application or application type.
	Name string                                                  `json:"name"`
	JSON zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON `json:"-"`
}

// zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayAppTypesZeroTrustGatewayApplication]
type zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypesZeroTrustGatewayApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGatewayAppTypesZeroTrustGatewayApplication) implementsZeroTrustZeroTrustGatewayAppTypes() {
}

type ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType struct {
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name string                                                      `json:"name"`
	JSON zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON `json:"-"`
}

// zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType]
type zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType) implementsZeroTrustZeroTrustGatewayAppTypes() {
}

type GatewayAppTypeListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayAppTypeListResponseEnvelope struct {
	Errors   []GatewayAppTypeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayAppTypeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayAppTypes                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayAppTypeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayAppTypeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayAppTypeListResponseEnvelopeJSON       `json:"-"`
}

// gatewayAppTypeListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayAppTypeListResponseEnvelope]
type gatewayAppTypeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAppTypeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAppTypeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayAppTypeListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    gatewayAppTypeListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayAppTypeListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayAppTypeListResponseEnvelopeErrors]
type gatewayAppTypeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAppTypeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAppTypeListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayAppTypeListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    gatewayAppTypeListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayAppTypeListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayAppTypeListResponseEnvelopeMessages]
type gatewayAppTypeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAppTypeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAppTypeListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayAppTypeListResponseEnvelopeSuccess bool

const (
	GatewayAppTypeListResponseEnvelopeSuccessTrue GatewayAppTypeListResponseEnvelopeSuccess = true
)

type GatewayAppTypeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       gatewayAppTypeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayAppTypeListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [GatewayAppTypeListResponseEnvelopeResultInfo]
type gatewayAppTypeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAppTypeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAppTypeListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
