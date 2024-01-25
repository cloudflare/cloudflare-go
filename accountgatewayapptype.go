// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountGatewayAppTypeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayAppTypeService]
// method instead.
type AccountGatewayAppTypeService struct {
	Options []option.RequestOption
}

// NewAccountGatewayAppTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayAppTypeService(opts ...option.RequestOption) (r *AccountGatewayAppTypeService) {
	r = &AccountGatewayAppTypeService{}
	r.Options = opts
	return
}

// Fetches all application and application type mappings.
func (r *AccountGatewayAppTypeService) ZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappings(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/gateway/app_types", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponse struct {
	Errors     []AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseError    `json:"errors"`
	Messages   []AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseMessage  `json:"messages"`
	Result     []AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResult   `json:"result"`
	ResultInfo AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseSuccess `json:"success"`
	JSON    accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseJSON    `json:"-"`
}

// accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponse]
type accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseError struct {
	Code    int64                                                                                                                                `json:"code,required"`
	Message string                                                                                                                               `json:"message,required"`
	JSON    accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseErrorJSON `json:"-"`
}

// accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseError]
type accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseMessage struct {
	Code    int64                                                                                                                                  `json:"code,required"`
	Message string                                                                                                                                 `json:"message,required"`
	JSON    accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseMessageJSON `json:"-"`
}

// accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseMessage]
type accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplication]
// or
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplicationType].
type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResult interface {
	implementsAccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResult)(nil)).Elem(), "")
}

type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplication struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The name of the application or application type.
	Name string                                                                                                                                                   `json:"name"`
	JSON accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFe7OxcxuApplicationJSON `json:"-"`
}

// accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFe7OxcxuApplicationJSON
// contains the JSON metadata for the struct
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplication]
type accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFe7OxcxuApplicationJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplication) implementsAccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResult() {
}

type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplicationType struct {
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name string                                                                                                                                                       `json:"name"`
	JSON accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFe7OxcxuApplicationTypeJSON `json:"-"`
}

// accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFe7OxcxuApplicationTypeJSON
// contains the JSON metadata for the struct
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplicationType]
type accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFe7OxcxuApplicationTypeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplicationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultFE7OxcxuApplicationType) implementsAccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResult() {
}

type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                                   `json:"total_count"`
	JSON       accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultInfoJSON `json:"-"`
}

// accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultInfo]
type accountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseSuccess bool

const (
	AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseSuccessTrue AccountGatewayAppTypeZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappingsResponseSuccess = true
)
