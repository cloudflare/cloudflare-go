// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserTokenPermissionGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserTokenPermissionGroupService] method instead.
type UserTokenPermissionGroupService struct {
	Options []option.RequestOption
}

// NewUserTokenPermissionGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserTokenPermissionGroupService(opts ...option.RequestOption) (r *UserTokenPermissionGroupService) {
	r = &UserTokenPermissionGroupService{}
	r.Options = opts
	return
}

// Find all available permission groups.
func (r *UserTokenPermissionGroupService) PermissionGroupsListPermissionGroups(ctx context.Context, opts ...option.RequestOption) (res *SchemasResponseCollection5gNmnUwj, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens/permission_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasResponseCollection5gNmnUwj struct {
	Errors     []SchemasResponseCollection5gNmnUwjError    `json:"errors"`
	Messages   []SchemasResponseCollection5gNmnUwjMessage  `json:"messages"`
	Result     []interface{}                               `json:"result"`
	ResultInfo SchemasResponseCollection5gNmnUwjResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasResponseCollection5gNmnUwjSuccess `json:"success"`
	JSON    schemasResponseCollection5gNmnUwjJSON    `json:"-"`
}

// schemasResponseCollection5gNmnUwjJSON contains the JSON metadata for the struct
// [SchemasResponseCollection5gNmnUwj]
type schemasResponseCollection5gNmnUwjJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollection5gNmnUwj) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollection5gNmnUwjError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    schemasResponseCollection5gNmnUwjErrorJSON `json:"-"`
}

// schemasResponseCollection5gNmnUwjErrorJSON contains the JSON metadata for the
// struct [SchemasResponseCollection5gNmnUwjError]
type schemasResponseCollection5gNmnUwjErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollection5gNmnUwjError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollection5gNmnUwjMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    schemasResponseCollection5gNmnUwjMessageJSON `json:"-"`
}

// schemasResponseCollection5gNmnUwjMessageJSON contains the JSON metadata for the
// struct [SchemasResponseCollection5gNmnUwjMessage]
type schemasResponseCollection5gNmnUwjMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollection5gNmnUwjMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollection5gNmnUwjResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       schemasResponseCollection5gNmnUwjResultInfoJSON `json:"-"`
}

// schemasResponseCollection5gNmnUwjResultInfoJSON contains the JSON metadata for
// the struct [SchemasResponseCollection5gNmnUwjResultInfo]
type schemasResponseCollection5gNmnUwjResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollection5gNmnUwjResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasResponseCollection5gNmnUwjSuccess bool

const (
	SchemasResponseCollection5gNmnUwjSuccessTrue SchemasResponseCollection5gNmnUwjSuccess = true
)
