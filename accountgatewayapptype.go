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

// List all Application and Application Type mappings.
func (r *AccountGatewayAppTypeService) ZeroTrustGatewayApplicationAndApplicationTypeMappingsListApplicationAndApplicationTypeMappings(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AppTypesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/gateway/app_types", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AppTypesResponseCollection struct {
	Errors     []AppTypesResponseCollectionError    `json:"errors"`
	Messages   []AppTypesResponseCollectionMessage  `json:"messages"`
	Result     []AppTypesResponseCollectionResult   `json:"result"`
	ResultInfo AppTypesResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AppTypesResponseCollectionSuccess `json:"success"`
	JSON    appTypesResponseCollectionJSON    `json:"-"`
}

// appTypesResponseCollectionJSON contains the JSON metadata for the struct
// [AppTypesResponseCollection]
type appTypesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppTypesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AppTypesResponseCollectionError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    appTypesResponseCollectionErrorJSON `json:"-"`
}

// appTypesResponseCollectionErrorJSON contains the JSON metadata for the struct
// [AppTypesResponseCollectionError]
type appTypesResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppTypesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AppTypesResponseCollectionMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    appTypesResponseCollectionMessageJSON `json:"-"`
}

// appTypesResponseCollectionMessageJSON contains the JSON metadata for the struct
// [AppTypesResponseCollectionMessage]
type appTypesResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppTypesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AppTypesResponseCollectionResultApplication5bPv33sq] or
// [AppTypesResponseCollectionResultApplicationType].
type AppTypesResponseCollectionResult interface {
	implementsAppTypesResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AppTypesResponseCollectionResult)(nil)).Elem(), "")
}

type AppTypesResponseCollectionResultApplication5bPv33sq struct {
	// The identifier for this application. There is only one application per id.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of an Application Type that has been
	// returned.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The name of the application or application type.
	Name string                                                  `json:"name"`
	JSON appTypesResponseCollectionResultApplication5bPv33sqJSON `json:"-"`
}

// appTypesResponseCollectionResultApplication5bPv33sqJSON contains the JSON
// metadata for the struct [AppTypesResponseCollectionResultApplication5bPv33sq]
type appTypesResponseCollectionResultApplication5bPv33sqJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AppTypesResponseCollectionResultApplication5bPv33sq) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AppTypesResponseCollectionResultApplication5bPv33sq) implementsAppTypesResponseCollectionResult() {
}

type AppTypesResponseCollectionResultApplicationType struct {
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of an Application Type that has been
	// returned.
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name string                                              `json:"name"`
	JSON appTypesResponseCollectionResultApplicationTypeJSON `json:"-"`
}

// appTypesResponseCollectionResultApplicationTypeJSON contains the JSON metadata
// for the struct [AppTypesResponseCollectionResultApplicationType]
type appTypesResponseCollectionResultApplicationTypeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppTypesResponseCollectionResultApplicationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AppTypesResponseCollectionResultApplicationType) implementsAppTypesResponseCollectionResult() {
}

type AppTypesResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       appTypesResponseCollectionResultInfoJSON `json:"-"`
}

// appTypesResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [AppTypesResponseCollectionResultInfo]
type appTypesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppTypesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AppTypesResponseCollectionSuccess bool

const (
	AppTypesResponseCollectionSuccessTrue AppTypesResponseCollectionSuccess = true
)
