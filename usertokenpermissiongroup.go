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
func (r *UserTokenPermissionGroupService) PermissionGroupsListPermissionGroups(ctx context.Context, opts ...option.RequestOption) (res *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens/permission_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponse struct {
	Errors     []UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseError    `json:"errors"`
	Messages   []UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseMessage  `json:"messages"`
	Result     []interface{}                                                                  `json:"result"`
	ResultInfo UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseSuccess `json:"success"`
	JSON    userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseJSON    `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponse]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseErrorJSON `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseErrorJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseError]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseMessageJSON `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseMessageJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseMessage]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                            `json:"total_count"`
	JSON       userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseResultInfoJSON `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseResultInfo]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseSuccess bool

const (
	UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseSuccessTrue UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseSuccess = true
)
