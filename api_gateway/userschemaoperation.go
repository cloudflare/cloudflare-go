// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// UserSchemaOperationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSchemaOperationService] method instead.
type UserSchemaOperationService struct {
	Options []option.RequestOption
}

// NewUserSchemaOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSchemaOperationService(opts ...option.RequestOption) (r *UserSchemaOperationService) {
	r = &UserSchemaOperationService{}
	r.Options = opts
	return
}

// Retrieves all operations from the schema. Operations that already exist in API
// Shield Endpoint Management will be returned as full operations.
func (r *UserSchemaOperationService) List(ctx context.Context, schemaID string, params UserSchemaOperationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UserSchemaOperationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s/operations", params.ZoneID, schemaID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieves all operations from the schema. Operations that already exist in API
// Shield Endpoint Management will be returned as full operations.
func (r *UserSchemaOperationService) ListAutoPaging(ctx context.Context, schemaID string, params UserSchemaOperationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UserSchemaOperationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, schemaID, params, opts...))
}

type UserSchemaOperationListResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method      UserSchemaOperationListResponseMethod `json:"method,required"`
	LastUpdated time.Time                             `json:"last_updated" format:"date-time"`
	// UUID
	OperationID string `json:"operation_id"`
	// This field can have the runtime type of [APIShieldFeatures].
	Features interface{}                         `json:"features,required"`
	JSON     userSchemaOperationListResponseJSON `json:"-"`
	union    UserSchemaOperationListResponseUnion
}

// userSchemaOperationListResponseJSON contains the JSON metadata for the struct
// [UserSchemaOperationListResponse]
type userSchemaOperationListResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	LastUpdated apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userSchemaOperationListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *UserSchemaOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = UserSchemaOperationListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSchemaOperationListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [api_gateway.APIShield],
// [api_gateway.UserSchemaOperationListResponseAPIShieldBasicOperation].
func (r UserSchemaOperationListResponse) AsUnion() UserSchemaOperationListResponseUnion {
	return r.union
}

// Union satisfied by [api_gateway.APIShield] or
// [api_gateway.UserSchemaOperationListResponseAPIShieldBasicOperation].
type UserSchemaOperationListResponseUnion interface {
	implementsAPIGatewayUserSchemaOperationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSchemaOperationListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIShield{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldBasicOperation{}),
		},
	)
}

type UserSchemaOperationListResponseAPIShieldBasicOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method UserSchemaOperationListResponseAPIShieldBasicOperationMethod `json:"method,required"`
	JSON   userSchemaOperationListResponseAPIShieldBasicOperationJSON   `json:"-"`
}

// userSchemaOperationListResponseAPIShieldBasicOperationJSON contains the JSON
// metadata for the struct [UserSchemaOperationListResponseAPIShieldBasicOperation]
type userSchemaOperationListResponseAPIShieldBasicOperationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldBasicOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldBasicOperationJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldBasicOperation) implementsAPIGatewayUserSchemaOperationListResponse() {
}

// The HTTP method used to access the endpoint.
type UserSchemaOperationListResponseAPIShieldBasicOperationMethod string

const (
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodGet     UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "GET"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodPost    UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "POST"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodHead    UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "HEAD"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodOptions UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "OPTIONS"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodPut     UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "PUT"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodDelete  UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "DELETE"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodConnect UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "CONNECT"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodPatch   UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "PATCH"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodTrace   UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "TRACE"
)

func (r UserSchemaOperationListResponseAPIShieldBasicOperationMethod) IsKnown() bool {
	switch r {
	case UserSchemaOperationListResponseAPIShieldBasicOperationMethodGet, UserSchemaOperationListResponseAPIShieldBasicOperationMethodPost, UserSchemaOperationListResponseAPIShieldBasicOperationMethodHead, UserSchemaOperationListResponseAPIShieldBasicOperationMethodOptions, UserSchemaOperationListResponseAPIShieldBasicOperationMethodPut, UserSchemaOperationListResponseAPIShieldBasicOperationMethodDelete, UserSchemaOperationListResponseAPIShieldBasicOperationMethodConnect, UserSchemaOperationListResponseAPIShieldBasicOperationMethodPatch, UserSchemaOperationListResponseAPIShieldBasicOperationMethodTrace:
		return true
	}
	return false
}

// The HTTP method used to access the endpoint.
type UserSchemaOperationListResponseMethod string

const (
	UserSchemaOperationListResponseMethodGet     UserSchemaOperationListResponseMethod = "GET"
	UserSchemaOperationListResponseMethodPost    UserSchemaOperationListResponseMethod = "POST"
	UserSchemaOperationListResponseMethodHead    UserSchemaOperationListResponseMethod = "HEAD"
	UserSchemaOperationListResponseMethodOptions UserSchemaOperationListResponseMethod = "OPTIONS"
	UserSchemaOperationListResponseMethodPut     UserSchemaOperationListResponseMethod = "PUT"
	UserSchemaOperationListResponseMethodDelete  UserSchemaOperationListResponseMethod = "DELETE"
	UserSchemaOperationListResponseMethodConnect UserSchemaOperationListResponseMethod = "CONNECT"
	UserSchemaOperationListResponseMethodPatch   UserSchemaOperationListResponseMethod = "PATCH"
	UserSchemaOperationListResponseMethodTrace   UserSchemaOperationListResponseMethod = "TRACE"
)

func (r UserSchemaOperationListResponseMethod) IsKnown() bool {
	switch r {
	case UserSchemaOperationListResponseMethodGet, UserSchemaOperationListResponseMethodPost, UserSchemaOperationListResponseMethodHead, UserSchemaOperationListResponseMethodOptions, UserSchemaOperationListResponseMethodPut, UserSchemaOperationListResponseMethodDelete, UserSchemaOperationListResponseMethodConnect, UserSchemaOperationListResponseMethodPatch, UserSchemaOperationListResponseMethodTrace:
		return true
	}
	return false
}

type UserSchemaOperationListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]UserSchemaOperationListParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Filter results by whether operations exist in API Shield Endpoint Management or
	// not. `new` will just return operations from the schema that do not exist in API
	// Shield Endpoint Management. `existing` will just return operations from the
	// schema that already exist in API Shield Endpoint Management.
	OperationStatus param.Field[UserSchemaOperationListParamsOperationStatus] `query:"operation_status"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [UserSchemaOperationListParams]'s query parameters as
// `url.Values`.
func (r UserSchemaOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserSchemaOperationListParamsFeature string

const (
	UserSchemaOperationListParamsFeatureThresholds       UserSchemaOperationListParamsFeature = "thresholds"
	UserSchemaOperationListParamsFeatureParameterSchemas UserSchemaOperationListParamsFeature = "parameter_schemas"
	UserSchemaOperationListParamsFeatureSchemaInfo       UserSchemaOperationListParamsFeature = "schema_info"
)

func (r UserSchemaOperationListParamsFeature) IsKnown() bool {
	switch r {
	case UserSchemaOperationListParamsFeatureThresholds, UserSchemaOperationListParamsFeatureParameterSchemas, UserSchemaOperationListParamsFeatureSchemaInfo:
		return true
	}
	return false
}

// Filter results by whether operations exist in API Shield Endpoint Management or
// not. `new` will just return operations from the schema that do not exist in API
// Shield Endpoint Management. `existing` will just return operations from the
// schema that already exist in API Shield Endpoint Management.
type UserSchemaOperationListParamsOperationStatus string

const (
	UserSchemaOperationListParamsOperationStatusNew      UserSchemaOperationListParamsOperationStatus = "new"
	UserSchemaOperationListParamsOperationStatusExisting UserSchemaOperationListParamsOperationStatus = "existing"
)

func (r UserSchemaOperationListParamsOperationStatus) IsKnown() bool {
	switch r {
	case UserSchemaOperationListParamsOperationStatusNew, UserSchemaOperationListParamsOperationStatusExisting:
		return true
	}
	return false
}
