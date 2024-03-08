// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustGatewayAppTypeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustGatewayAppTypeService] method instead.
type ZeroTrustGatewayAppTypeService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayAppTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayAppTypeService(opts ...option.RequestOption) (r *ZeroTrustGatewayAppTypeService) {
	r = &ZeroTrustGatewayAppTypeService{}
	r.Options = opts
	return
}

// Fetches all application and application type mappings.
func (r *ZeroTrustGatewayAppTypeService) List(ctx context.Context, query ZeroTrustGatewayAppTypeListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayAppTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayAppTypeListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/app_types", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplication] or
// [ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationType].
type ZeroTrustGatewayAppTypeListResponse interface {
	implementsZeroTrustGatewayAppTypeListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGatewayAppTypeListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationType{}),
		},
	)
}

type ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplication struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The name of the application or application type.
	Name string                                                             `json:"name"`
	JSON zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationJSON `json:"-"`
}

// zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplication]
type zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplication) implementsZeroTrustGatewayAppTypeListResponse() {
}

type ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationType struct {
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name string                                                                 `json:"name"`
	JSON zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationTypeJSON `json:"-"`
}

// zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationTypeJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationType]
type zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationTypeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationTypeJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGatewayAppTypeListResponseZeroTrustGatewayApplicationType) implementsZeroTrustGatewayAppTypeListResponse() {
}

type ZeroTrustGatewayAppTypeListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustGatewayAppTypeListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayAppTypeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayAppTypeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayAppTypeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustGatewayAppTypeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustGatewayAppTypeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustGatewayAppTypeListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustGatewayAppTypeListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayAppTypeListResponseEnvelope]
type zeroTrustGatewayAppTypeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayAppTypeListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustGatewayAppTypeListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayAppTypeListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayAppTypeListResponseEnvelopeErrors]
type zeroTrustGatewayAppTypeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypeListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayAppTypeListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustGatewayAppTypeListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayAppTypeListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayAppTypeListResponseEnvelopeMessages]
type zeroTrustGatewayAppTypeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypeListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayAppTypeListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayAppTypeListResponseEnvelopeSuccessTrue ZeroTrustGatewayAppTypeListResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayAppTypeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       zeroTrustGatewayAppTypeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustGatewayAppTypeListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayAppTypeListResponseEnvelopeResultInfo]
type zeroTrustGatewayAppTypeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypeListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
